package main

import (
	"fmt"
	"image"
	"os"
	"sort"
	"unsafe"

	"github.com/anthonynsimon/bild/imgio"
	"github.com/anthonynsimon/bild/parallel"
	"github.com/jinzhongmin/goncnn/pkg/ncnn"
	"github.com/jinzhongmin/usf"
)

// 1、download ncnn lib from https://github.com/Tencent/ncnn
var lib = "./ncnn.dll"

// 2、transform model yolov5 model using pnnx https://github.com/pnnx/pnnx https://github.com/ultralytics/yolov5
var model_param = "./yolov5s.ncnn.param"
var model_bin = "./yolov5s.ncnn.bin"

// 3、img
var img_path = "./bird.png"

func main() {
	ncnn.Loadlib(lib) //download from ncnn

	//CreateOption
	opt := ncnn.CreateOption()
	opt.SetUseVulkanCompute(true)

	//init net
	net := ncnn.CreateNet()
	net.SetOption(opt)
	p, _ := os.ReadFile(model_param)
	net.LoadParamMemory(p)
	b, _ := os.ReadFile(model_bin)
	net.LoadModelMemory(b)

	alc := ncnn.CreatePoolAllocator()
	ex := ncnn.CreateExtractor(net)
	ex.SetOption(opt)

	//load img
	srcImg, _ := imgio.Open(img_path)
	imgBlob := blob(&srcImg, [3]float32{0, 0, 0}, [3]float32{1.0 / 255, 1.0 / 255, 1.0 / 255})

	inMat := ncnn.CreateMatExternal3D(640, 640, 3, unsafe.Pointer(&imgBlob[0]), alc)
	defer inMat.Destroy()

	ex.Input("in0", inMat)
	outMat := ex.Extract("out0")
	defer outMat.Destroy()

	w := outMat.GetW()
	h := outMat.GetH()
	l := w * h
	outData := *(*[]float32)(usf.Slice(outMat.GetData(), uint64(l)))

	results := make([]*result, 25200)
	parallel.Line(25200, func(start, end int) {
		for r := start; r < end; r++ {
			p := outData[r*85 : r*85+85]
			if p[4] < 0.4 {
				continue
			}
			cx := float64(p[0])
			cy := float64(p[1])
			w := float64(p[2])
			h := float64(p[3])
			p = outData[r*85+5 : r*85+85]
			var bestVal float32 = 0.0
			bestIdx := 0
			for i := range p {
				if p[i] <= bestVal {
					continue
				}
				bestIdx = i
				bestVal = p[i]
			}
			if bestVal <= 0.98 {
				continue
			}

			x0 := int(cx - w/2)
			y0 := int(cy - h/2)
			x1 := x0 + int(w)
			y1 := y0 + int(h)

			results[r] = &result{
				rect:  image.Rect(x0, y0, x1, y1),
				score: bestVal,
				class: bestIdx,
			}
		}
	})

	_results := make([]*result, 0)
	for i := 0; i < 25200; i++ {
		if results[i] == nil {
			continue
		}
		_results = append(_results, results[i])
	}
	rs := nms(_results, 0.6)
	fmt.Println(rs)
}

func blob(im *image.Image, mean [3]float32, scale [3]float32) []float32 {
	rows := (*im).Bounds().Dy()
	cols := (*im).Bounds().Dx()
	frame := rows * cols
	frame2 := frame * 2

	rgb := make([]float32, rows*cols*3)
	s0 := 255 * scale[0]
	s1 := 255 * scale[1]
	s2 := 255 * scale[2]
	parallel.Line(rows, func(start, end int) {
		idx := start * rows

		for row := start; row < end; row++ {
			for col := 0; col < cols; col++ {
				r, g, b, _ := (*im).At(col, row).RGBA()
				rgb[idx] = (float32(r>>8) - mean[0]) / s0
				rgb[idx+frame] = (float32(g>>8) - mean[1]) / s1
				rgb[idx+frame2] = (float32(b>>8) - mean[2]) / s2
				idx += 1
			}
		}
	})
	return rgb
}

type result struct {
	rect  image.Rectangle
	class int
	score float32
}

type results []*result

func (r results) Len() int           { return len(r) }
func (r results) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
func (r results) Less(i, j int) bool { return r[i].score < r[j].score }

func IOUFloat32(r1, r2 image.Rectangle) float32 {
	intersection := r1.Intersect(r2)
	interArea := intersection.Dx() * intersection.Dy()
	r1Area := r1.Dx() * r1.Dy()
	r2Area := r2.Dx() * r2.Dy()
	return float32(interArea) / float32(r1Area+r2Area-interArea)
}
func nms(rs []*result, thre float32) []*result {
	sort.Sort(results(rs))
	rnms := make([]*result, 0)
	if len(rs) == 0 {
		return rnms
	}
	rnms = append(rnms, rs[0])

	for i := 1; i < len(rs); i++ {
		tocheck, del := len(rnms), false
		for j := 0; j < tocheck; j++ {
			currIOU := IOUFloat32(rs[i].rect, rs[j].rect)
			if currIOU > thre && rs[i].class == rnms[j].class {
				del = true
				break
			}
		}
		if !del {
			rnms = append(rnms, rs[i])
		}
	}
	return rnms
}
