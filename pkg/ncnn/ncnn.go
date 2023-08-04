package ncnn

//
import "C"
import (
	"unsafe"

	"github.com/jinzhongmin/goffi/pkg/c"
	"github.com/jinzhongmin/usf"
)

var ncnnLib *c.Lib

// init library by shared library
func InitLib(path string, mod c.LibMode) {
	var err error
	ncnnLib, err = c.NewLib(path, mod, true)
	if err != nil {
		panic(err)
	}
}

// var dll *dlfcn.Handle

// func Loadlib(dllPath string, mod dlfcn.Mode) {
// 	var err error
// 	dll, err = dlfcn.DlOpen(dllPath, mod)
// 	if err != nil {
// 		panic(err)
// 	}
// }

// func c.CBool(b bool) C.int {
// 	if b {
// 		return 1
// 	}
// 	return 0
// }
// func  c.CStr(s string) unsafe.Pointer {
// 	p := C.CString(s)
// 	return unsafe.Pointer(p)
// }
// func gostr(p unsafe.Pointer) string {
// 	return C.GoString((*C.char)(p))
// }
// func Call(fn string, outTyp ffi.Type, inTyp []ffi.Type, args []interface{}) unsafe.Pointer {
// 	fnp, err := dll.Symbol(fn)
// 	if err != nil {
// 		panic(err)
// 	}

// 	cif, err := ffi.NewCif(ffi.AbiDefault, outTyp, inTyp)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer cif.Free()

// 	r := cif.Call(unsafe.Pointer(fnp), args)
// 	p := usf.Malloc(1, 8)
// 	usf.Push(p, (*(*[1]unsafe.Pointer)(r))[0])
// 	runtime.SetFinalizer(&p, func(_p *unsafe.Pointer) {
// 		usf.Free(*_p)
// 	})
// 	return p
// }

func Version() string {
	s := ncnnLib.Call("ncnn_version", c.Pointer, nil, nil)
	defer s.Free()
	return s.Str()
}

type Allocator struct{ c unsafe.Pointer }

func CreatePoolAllocator() *Allocator {
	p := ncnnLib.Call("ncnn_allocator_create_pool_allocator", c.Pointer, nil, nil)
	return &Allocator{c: p.PtrFree()}
}
func CreateUnlockedPoolAllocator() *Allocator {
	p := ncnnLib.Call("ncnn_allocator_create_unlocked_pool_allocator", c.Pointer, nil, nil)
	return &Allocator{c: p.PtrFree()}
}
func (alc *Allocator) Destroy() {
	ncnnLib.Call("ncnn_allocator_destroy", c.Void, []c.Type{c.Pointer}, []interface{}{&alc.c})
}

type Option struct{ c unsafe.Pointer }

func CreateOption() *Option {
	p := ncnnLib.Call("ncnn_option_create", c.Pointer, nil, nil)
	return &Option{c: p.PtrFree()}
}

func (opt *Option) Destroy() {
	ncnnLib.Call("ncnn_option_destroy", c.Void, []c.Type{c.Pointer}, []interface{}{&opt.c})
}
func (opt *Option) GetNumThreads() int32 {
	n := ncnnLib.Call("ncnn_option_get_num_threads", c.I32, []c.Type{c.Pointer}, []interface{}{&opt.c})
	return n.I32Free()
}
func (opt *Option) SetNumThreads(num int32) {
	ncnnLib.Call("ncnn_option_set_num_threads", c.Void, []c.Type{c.Pointer, c.I32}, []interface{}{&opt.c, &num})
}
func (opt *Option) GetUseLocalPoolAllocator() int32 {
	n := ncnnLib.Call("ncnn_option_get_use_local_pool_allocator", c.I32,
		[]c.Type{c.Pointer}, []interface{}{&opt.c})
	return n.I32Free()
}
func (opt *Option) SetUseLocalPoolAllocator(n int32) {
	ncnnLib.Call("ncnn_option_set_use_local_pool_allocator", c.Void,
		[]c.Type{c.Pointer, c.I32}, []interface{}{&opt.c, &n})
}
func (opt *Option) SetBlobAllocator(alc *Allocator) {
	ncnnLib.Call("ncnn_option_set_blob_allocator", c.Void,
		[]c.Type{c.Pointer, c.Pointer}, []interface{}{&opt.c, &alc.c})
}
func (opt *Option) SetWorkspaceAllocator(alc *Allocator) {
	ncnnLib.Call("ncnn_option_set_workspace_allocator", c.Void,
		[]c.Type{c.Pointer, c.Pointer}, []interface{}{&opt.c, &alc.c})
}
func (opt *Option) GetUseVulkanCompute() bool {
	n := ncnnLib.Call("ncnn_option_get_use_vulkan_compute", c.I32, []c.Type{c.Pointer}, []interface{}{&opt.c})
	return n.BoolFree()
}
func (opt *Option) SetUseVulkanCompute(enable bool) {
	n := c.CBool(enable)
	ncnnLib.Call("ncnn_option_set_use_vulkan_compute", c.Void,
		[]c.Type{c.Pointer, c.I32}, []interface{}{&opt.c, &n})
}

type Mat struct{ c unsafe.Pointer }

func (mat *Mat) Destroy() {
	ncnnLib.Call("ncnn_mat_destroy", c.Void, []c.Type{c.Pointer}, []interface{}{&mat.c})
}
func CreateMat() *Mat {
	m := ncnnLib.Call("ncnn_mat_create", c.Pointer, nil, nil)
	return &Mat{c: m.PtrFree()}
}
func CreateMat1D(w int32, alc *Allocator) *Mat {
	m := ncnnLib.Call("ncnn_mat_create_1d", c.Pointer,
		[]c.Type{c.I32, c.Pointer}, []interface{}{&w, &alc.c})
	return &Mat{c: m.PtrFree()}
}
func CreateMat2D(w, h int32, alc *Allocator) *Mat {
	m := ncnnLib.Call("ncnn_mat_create_2d", c.Pointer,
		[]c.Type{c.I32, c.I32, c.Pointer}, []interface{}{&w, &h, &alc.c})
	return &Mat{c: m.PtrFree()}
}
func CreateMat3D(w, h, ch int32, alc *Allocator) *Mat {
	m := ncnnLib.Call("ncnn_mat_create_3d", c.Pointer,
		[]c.Type{c.I32, c.I32, c.I32, c.Pointer}, []interface{}{&w, &h, &ch, &alc.c})
	return &Mat{c: m.PtrFree()}
}
func CreateMat4D(w, h, d, ch int32, alc *Allocator) *Mat {
	m := ncnnLib.Call("ncnn_mat_create_4d", c.Pointer,
		[]c.Type{c.I32, c.I32, c.I32, c.I32, c.Pointer}, []interface{}{&w, &h, &d, &ch, &alc.c})
	return &Mat{c: m.PtrFree()}
}
func CreateMatExternal1D(w int32, data unsafe.Pointer, alc *Allocator) *Mat {
	m := ncnnLib.Call("ncnn_mat_create_external_1d", c.Pointer,
		[]c.Type{c.I32, c.Pointer, c.Pointer}, []interface{}{&w, &data, &alc.c})
	return &Mat{c: m.PtrFree()}
}
func CreateMatExternal2D(w, h int32, data unsafe.Pointer, alc *Allocator) *Mat {
	m := ncnnLib.Call("ncnn_mat_create_external_2d", c.Pointer,
		[]c.Type{c.I32, c.I32, c.Pointer, c.Pointer}, []interface{}{&w, &h, &data, &alc.c})
	return &Mat{c: m.PtrFree()}
}

func CreateMatExternal3D(w, h, ch int32, data unsafe.Pointer, alc *Allocator) *Mat {
	m := ncnnLib.Call("ncnn_mat_create_external_3d", c.Pointer,
		[]c.Type{c.I32, c.I32, c.I32, c.Pointer, c.Pointer}, []interface{}{&w, &h, &ch, &data, &alc.c})
	return &Mat{c: m.PtrFree()}
}
func CreateMatExternal4D(w, h, d, ch int32, data unsafe.Pointer, alc *Allocator) *Mat {
	m := ncnnLib.Call("ncnn_mat_create_external_4d", c.Pointer,
		[]c.Type{c.I32, c.I32, c.I32, c.I32, c.Pointer, c.Pointer},
		[]interface{}{&w, &h, &d, &ch, &data, &alc.c})
	return &Mat{c: m.PtrFree()}
}
func CreateMat1DElm(w int32, elmSize uint64, elmPack int32, alc *Allocator) *Mat {
	m := ncnnLib.Call("ncnn_mat_create_1d_elem", c.Pointer,
		[]c.Type{c.I32, c.U64, c.I32, c.Pointer}, []interface{}{&w, &elmSize, &elmPack, &alc.c})
	return &Mat{c: m.PtrFree()}
}
func CreateMat2DElm(w, h int32, elmSize uint64, elmPack int32, alc *Allocator) *Mat {
	m := ncnnLib.Call("ncnn_mat_create_2d_elem", c.Pointer,
		[]c.Type{c.I32, c.I32, c.U64, c.I32, c.Pointer},
		[]interface{}{&w, &h, &elmSize, &elmPack, &alc.c})
	return &Mat{c: m.PtrFree()}
}
func CreateMat3DElm(w, h, ch int32, elmSize uint64, elmPack int32, alc *Allocator) *Mat {
	m := ncnnLib.Call("ncnn_mat_create_3d_elem", c.Pointer,
		[]c.Type{c.I32, c.I32, c.I32, c.U64, c.I32, c.Pointer},
		[]interface{}{&w, &h, &ch, &elmSize, &elmPack, &alc.c})
	return &Mat{c: m.PtrFree()}
}
func CreateMat4DElm(w, h, d, ch int32, elmSize uint64, elmPack int32, alc *Allocator) *Mat {
	m := ncnnLib.Call("ncnn_mat_create_4d_elem", c.Pointer,
		[]c.Type{c.I32, c.I32, c.I32, c.I32, c.U64, c.I32, c.Pointer},
		[]interface{}{&w, &h, &d, &ch, &elmSize, &elmPack, &alc.c})
	return &Mat{c: m.PtrFree()}
}
func CreateMatExternal1DElm(w int32, data unsafe.Pointer, elmSize uint64, elmPack int32, alc *Allocator) *Mat {
	m := ncnnLib.Call("ncnn_mat_create_external_1d_elem", c.Pointer,
		[]c.Type{c.I32, c.Pointer, c.U64, c.I32, c.Pointer},
		[]interface{}{&w, &data, &elmSize, &elmPack, &alc.c})
	return &Mat{c: m.PtrFree()}
}
func CreateMatExternal2DElm(w, h int32, data unsafe.Pointer,
	elmSize uint64, elmPack int32, alc *Allocator) *Mat {
	m := ncnnLib.Call("ncnn_mat_create_external_2d_elem", c.Pointer,
		[]c.Type{c.I32, c.I32, c.Pointer, c.U64, c.I32, c.Pointer},
		[]interface{}{&w, &h, &data, &elmSize, &elmPack, &alc.c})
	return &Mat{c: m.PtrFree()}
}
func CreateMatExternal3DElm(w, h, ch int32, data unsafe.Pointer,
	elmSize uint64, elmPack int32, alc *Allocator) *Mat {
	m := ncnnLib.Call("ncnn_mat_create_external_3d_elem", c.Pointer,
		[]c.Type{c.I32, c.I32, c.I32, c.Pointer, c.U64, c.I32, c.Pointer},
		[]interface{}{&w, &h, &ch, &data, &elmSize, &elmPack, &alc.c})
	return &Mat{c: m.PtrFree()}
}

func CreateMatExternal4DElm(w, h, d, ch int32, data unsafe.Pointer,
	elmSize uint64, elmPack int32, alc *Allocator) *Mat {
	m := ncnnLib.Call("ncnn_mat_create_external_4d_elem", c.Pointer,
		[]c.Type{c.I32, c.I32, c.I32, c.I32, c.Pointer, c.U64, c.I32, c.Pointer},
		[]interface{}{&w, &h, &d, &ch, &data, &elmSize, &elmPack, &alc.c})
	return &Mat{c: m.PtrFree()}
}
func (mat *Mat) FillFloat(f float32) {
	ncnnLib.Call("ncnn_mat_fill_float", c.Void, []c.Type{c.Pointer, c.F32}, []interface{}{&mat.c, &f})
}
func (mat *Mat) Clone(alc *Allocator) *Mat {
	m := ncnnLib.Call("ncnn_mat_clone", c.Pointer,
		[]c.Type{c.Pointer, c.Pointer}, []interface{}{&mat.c, &alc.c})
	return &Mat{c: m.PtrFree()}
}
func (mat *Mat) Reshape1D(w int32, alc *Allocator) *Mat {
	m := ncnnLib.Call("ncnn_mat_reshape_1d", c.Pointer,
		[]c.Type{c.Pointer, c.I32, c.Pointer}, []interface{}{&mat.c, &w, &alc.c})
	return &Mat{c: m.PtrFree()}
}
func (mat *Mat) Reshape2D(w, h int32, alc *Allocator) *Mat {
	m := ncnnLib.Call("ncnn_mat_reshape_2d", c.Pointer,
		[]c.Type{c.Pointer, c.I32, c.I32, c.Pointer}, []interface{}{&mat.c, &w, &h, &alc.c})
	return &Mat{c: m.PtrFree()}
}
func (mat *Mat) Reshape3D(w, h, ch int32, alc *Allocator) *Mat {
	m := ncnnLib.Call("ncnn_mat_reshape_3d", c.Pointer,
		[]c.Type{c.Pointer, c.I32, c.I32, c.I32, c.Pointer}, []interface{}{&mat.c, &w, &h, &ch, &alc.c})
	return &Mat{c: m.PtrFree()}
}
func (mat *Mat) Reshape4D(w, h, d, ch int32, alc *Allocator) *Mat {
	m := ncnnLib.Call("ncnn_mat_reshape_4d", c.Pointer,
		[]c.Type{c.Pointer, c.I32, c.I32, c.I32, c.I32, c.Pointer},
		[]interface{}{&mat.c, &w, &h, &d, &ch, &alc.c})
	return &Mat{c: m.PtrFree()}
}
func (mat *Mat) GetDims() int32 {
	return ncnnLib.Call("ncnn_mat_get_dims", c.I32,
		[]c.Type{c.Pointer}, []interface{}{&mat.c}).I32Free()
}
func (mat *Mat) GetW() int32 {
	return ncnnLib.Call("ncnn_mat_get_w", c.I32,
		[]c.Type{c.Pointer}, []interface{}{&mat.c}).I32Free()

}
func (mat *Mat) GetH() int32 {
	return ncnnLib.Call("ncnn_mat_get_h", c.I32,
		[]c.Type{c.Pointer}, []interface{}{&mat.c}).I32Free()
}
func (mat *Mat) GetD() int32 {
	return ncnnLib.Call("ncnn_mat_get_d", c.I32,
		[]c.Type{c.Pointer}, []interface{}{&mat.c}).I32Free()
}
func (mat *Mat) GetC() int32 {
	return ncnnLib.Call("ncnn_mat_get_c", c.I32,
		[]c.Type{c.Pointer}, []interface{}{&mat.c}).I32Free()
}
func (mat *Mat) GetElemsize() uint64 {
	return ncnnLib.Call("ncnn_mat_get_elemsize", c.U64,
		[]c.Type{c.Pointer}, []interface{}{&mat.c}).U64()
}
func (mat *Mat) GetElempack() int32 {
	return ncnnLib.Call("ncnn_mat_get_elempack", c.I32,
		[]c.Type{c.Pointer}, []interface{}{&mat.c}).I32Free()
}
func (mat *Mat) GetCstep() uint64 {
	return ncnnLib.Call("ncnn_mat_get_cstep", c.U64,
		[]c.Type{c.Pointer}, []interface{}{&mat.c}).U64()
}

func (mat *Mat) GetData() unsafe.Pointer {
	return ncnnLib.Call("ncnn_mat_get_data", c.Pointer,
		[]c.Type{c.Pointer}, []interface{}{&mat.c}).Ptr()
}

func (mat *Mat) GetChannelData(ch int32) unsafe.Pointer {
	return ncnnLib.Call("ncnn_mat_get_channel_data", c.Pointer,
		[]c.Type{c.Pointer, c.I32}, []interface{}{&mat.c, &ch}).Ptr()
}

type PixelType int32

const (
	PixelTypeRGB  PixelType = 1
	PixelTypeBGR  PixelType = 2
	PixelTypeGRAY PixelType = 3
	PixelTypeRGBA PixelType = 4
	PixelTypeBGRA PixelType = 5
)

func PixelTypeX2Y(x, y int32) PixelType { return PixelType(x | (y << 16)) }

func CreateMatFromPixels(pixels []byte, typ PixelType, w, h, stride int32, alc *Allocator) *Mat {
	d := &pixels[0]
	m := ncnnLib.Call("ncnn_mat_from_pixels", c.Pointer,
		[]c.Type{c.Pointer, c.I32, c.I32, c.I32, c.I32, c.Pointer},
		[]interface{}{&d, &typ, &w, &h, &stride, &alc})
	return &Mat{c: m.PtrFree()}
}
func CreateMatFromPixelsResize(pixels []byte, typ PixelType, w, h, stride int32, targetW, targetH int32, alc *Allocator) *Mat {
	d := &pixels[0]
	m := ncnnLib.Call("ncnn_mat_from_pixels_resize", c.Pointer,
		[]c.Type{c.Pointer, c.I32, c.I32, c.I32, c.I32, c.I32, c.I32, c.Pointer},
		[]interface{}{&d, &typ, &w, &h, &stride, &targetW, &targetH, &alc})
	return &Mat{c: m.PtrFree()}
}
func CreateMatFromPixelsRoi(pixels []byte, typ PixelType, w, h, stride int32, roix, roiy, roiw, roih int32, alc *Allocator) *Mat {
	d := &pixels[0]
	m := ncnnLib.Call("ncnn_mat_from_pixels_roi", c.Pointer,
		[]c.Type{c.Pointer, c.I32, c.I32, c.I32, c.I32, c.I32, c.I32, c.I32, c.I32, c.Pointer},
		[]interface{}{&d, &typ, &w, &h, &stride, &roix, &roiy, &roiw, &roih, &alc})
	return &Mat{c: m.PtrFree()}
}

func CreateMatFromPixelsRoiResize(pixels []byte, typ PixelType, w, h, stride int32, roix, roiy, roiw, roih int32, targetW, targetH int32, alc *Allocator) *Mat {
	d := &pixels[0]
	m := ncnnLib.Call("ncnn_mat_from_pixels_roi_resize", c.Pointer,
		[]c.Type{c.Pointer, c.I32, c.I32, c.I32, c.I32, c.I32, c.I32, c.I32, c.I32, c.I32, c.I32, c.Pointer},
		[]interface{}{&d, &typ, &w, &h, &stride, &roix, &roiy, &roiw, &roih, &targetW, &targetH, &alc})
	return &Mat{c: m.PtrFree()}
}

// type slice struct {
// 	ptr unsafe.Pointer
// 	len int
// 	cap int
// }

// func (mat *Mat) ToPixels(typ PixelType, stride int32, outData interface{}) {
// 	p := (*(*[2]unsafe.Pointer)(unsafe.Pointer(&outData)))[1]
// 	p = (*(*slice)((unsafe.Pointer)((*unsafe.Pointer)(p)))).ptr
// 	ncnnLib.Call("ncnn_mat_to_pixels", c.Void,
// 		[]c.Type{c.Pointer, c.Pointer, c.I32, c.I32},
// 		[]interface{}{mat.c, p, typ, stride})
// }
// func (mat *Mat) ToPixelsResize(typ PixelType, targetW, targetH, targetStride int32, outData interface{}) {
// 	p := (*(*[2]unsafe.Pointer)(unsafe.Pointer(&outData)))[1]
// 	p = (*(*slice)((unsafe.Pointer)((*unsafe.Pointer)(p)))).ptr
// 	ncnnLib.Call("ncnn_mat_to_pixels", c.Void,
// 		[]c.Type{c.Pointer, c.Pointer, c.I32, c.I32, c.I32, c.I32},
// 		[]interface{}{mat.c, p, typ, targetW, targetH, targetStride})
// }

func (mat *Mat) SubstractMeanNormalize(mean, norm []float32) {
	m := &mean[0]
	n := &norm[0]
	ncnnLib.Call("ncnn_mat_substract_mean_normalize", c.Void,
		[]c.Type{c.Pointer, c.Pointer, c.Pointer},
		[]interface{}{&mat.c, &m, &n})
}
func (mat *Mat) ConvertPacking(elempack int32, opt *Option) *Mat {
	m := usf.Malloc(1, 8)
	defer usf.Free(m)
	ncnnLib.Call("ncnn_convert_packing", c.Void,
		[]c.Type{c.Pointer, c.Pointer, c.I32, c.Pointer},
		[]interface{}{&mat.c, &m, &elempack, &opt.c})
	return &Mat{c: usf.Pop(m)}
}

func (mat *Mat) Flatten(opt *Option) *Mat {
	m := usf.Malloc(1, 8)
	defer usf.Free(m)
	ncnnLib.Call("ncnn_flatten", c.Void,
		[]c.Type{c.Pointer, c.Pointer, c.Pointer},
		[]interface{}{&mat.c, &m, &opt.c})
	return &Mat{c: usf.Pop(m)}
}

type Blob struct{ c unsafe.Pointer }

func (blob *Blob) GetName() string {
	return ncnnLib.Call("ncnn_blob_get_name", c.Pointer,
		[]c.Type{c.Pointer}, []interface{}{&blob.c}).Str()
}
func (blob *Blob) GetProducer() int32 {
	return ncnnLib.Call("ncnn_blob_get_producer", c.I32,
		[]c.Type{c.Pointer}, []interface{}{&blob.c}).I32Free()
}
func (blob *Blob) GetConsumer() int32 {
	return ncnnLib.Call("ncnn_blob_get_consumer", c.I32,
		[]c.Type{c.Pointer}, []interface{}{&blob.c}).I32Free()
}

type Paramdict struct{ c unsafe.Pointer }

func CreateParamdict() *Paramdict {
	pd := ncnnLib.Call("ncnn_paramdict_create", c.Pointer, nil, nil)
	return &Paramdict{c: pd.PtrFree()}
}
func (pd *Paramdict) Destroy() {
	ncnnLib.Call("ncnn_paramdict_destroy", c.Void,
		[]c.Type{c.Pointer}, []interface{}{&pd.c})
}
func (pd *Paramdict) GetType(id int32) int32 {
	return ncnnLib.Call("ncnn_paramdict_get_type", c.I32,
		[]c.Type{c.Pointer, c.I32}, []interface{}{&pd.c, &id}).I32Free()
}
func (pd *Paramdict) GetInt32(id int32, def int32) int32 {
	return ncnnLib.Call("ncnn_paramdict_get_int", c.I32,
		[]c.Type{c.Pointer, c.I32, c.I32}, []interface{}{&pd.c, &id, &def}).I32Free()
}
func (pd *Paramdict) GetFloat32(id int32, def float32) float32 {
	return ncnnLib.Call("ncnn_paramdict_get_float", c.F32,
		[]c.Type{c.Pointer, c.I32, c.F32}, []interface{}{&pd.c, &id, &def}).F32Free()
}
func (pd *Paramdict) GetArray(id int32, def *Mat) *Mat {
	m := ncnnLib.Call("ncnn_paramdict_get_array", c.Pointer,
		[]c.Type{c.Pointer, c.I32, c.Pointer},
		[]interface{}{&pd.c, &id, &def})
	return &Mat{c: m.PtrFree()}
}
func (pd *Paramdict) SetInt32(id int32, i int32) {
	ncnnLib.Call("ncnn_paramdict_set_int", c.Void,
		[]c.Type{c.Pointer, c.I32, c.I32},
		[]interface{}{&pd.c, &id, &i})
}
func (pd *Paramdict) SetFloat32(id int32, f float32) {
	ncnnLib.Call("ncnn_paramdict_set_float", c.Void,
		[]c.Type{c.Pointer, c.I32, c.F32},
		[]interface{}{&pd.c, &id, &f})
}
func (pd *Paramdict) SetArray(id int32, m *Mat) {
	ncnnLib.Call("ncnn_paramdict_set_array", c.Void,
		[]c.Type{c.Pointer, c.I32, c.Pointer},
		[]interface{}{&pd.c, &id, &m})
}

type DataReader struct{ c unsafe.Pointer }

func CreateDataReader() *DataReader {
	d := ncnnLib.Call("ncnn_datareader_create", c.Pointer, nil, nil)
	return &DataReader{c: d.PtrFree()}
}
func CreateDataReaderFromMemory(mem []byte) *DataReader {
	p := unsafe.Pointer(&mem[0])
	pp := usf.Malloc(1, 8)
	usf.Push(pp, p)
	defer usf.Free(pp)
	d := ncnnLib.Call("ncnn_datareader_create_from_memory", c.Pointer,
		[]c.Type{c.Pointer}, []interface{}{&pp})
	return &DataReader{c: d.PtrFree()}
}
func (d *DataReader) Destroy() {
	ncnnLib.Call("ncnn_datareader_destroy", c.Void, []c.Type{c.Pointer}, []interface{}{&d.c})
}

type ModelBin struct{ c unsafe.Pointer }

func CreateModelBinFromDataReader(dr *DataReader) *ModelBin {
	m := ncnnLib.Call("ncnn_modelbin_create_from_datareader", c.Pointer,
		[]c.Type{c.Pointer}, []interface{}{&dr.c})
	return &ModelBin{c: m.PtrFree()}
}
func CreateModelBinFromMatArray(mats []*Mat) *ModelBin {
	l := uint32(len(mats))
	ms := usf.Malloc(uint64(l), 8)
	defer usf.Free(ms)
	m := ncnnLib.Call("ncnn_modelbin_create_from_mat_array", c.Pointer,
		[]c.Type{c.Pointer, c.I32}, []interface{}{&ms, &l})
	return &ModelBin{c: m.PtrFree()}
}
func (mb *ModelBin) Destroy() {
	ncnnLib.Call("ncnn_modelbin_destroy", c.Void,
		[]c.Type{c.Pointer}, []interface{}{&mb.c})
}

type Layer struct {
	c          unsafe.Pointer
	load_param *c.Fn
	load_model *c.Fn

	create_pipeline  *c.Fn
	destroy_pipeline *c.Fn

	forward_1 *c.Fn
	forward_n *c.Fn

	forward_inplace_1 *c.Fn
	forward_inplace_n *c.Fn
}

type _c_layer struct {
	this unsafe.Pointer

	load_param unsafe.Pointer
	load_model unsafe.Pointer

	create_pipeline  unsafe.Pointer
	destroy_pipeline unsafe.Pointer

	forward_1 unsafe.Pointer
	forward_n unsafe.Pointer

	forward_inplace_1 unsafe.Pointer
	forward_inplace_n unsafe.Pointer
}

func CreateLayer() *Layer {
	l := ncnnLib.Call("ncnn_layer_create", c.Pointer, nil, nil)
	return &Layer{c: l.PtrFree()}
}
func CreateLayerByTypeindex(typIdx int32) *Layer {
	l := ncnnLib.Call("ncnn_layer_create_by_typeindex", c.Pointer,
		[]c.Type{c.I32}, []interface{}{&typIdx})
	return &Layer{c: l.PtrFree()}
}
func CreateLayerByType(typ string) *Layer {
	t := c.CStr(typ)
	defer usf.Free(t)
	l := ncnnLib.Call("ncnn_layer_create_by_typeindex", c.Pointer,
		[]c.Type{c.Pointer}, []interface{}{&t})
	return &Layer{c: l.PtrFree()}
}
func (l *Layer) SetLoadParam(fn func(l *Layer, pd *Paramdict) int32) {
	cls := c.NewFn(c.AbiDefault, c.I32, []c.Type{c.Pointer, c.Pointer}, func(args []c.Value, ret c.Value) {
		lc := args[0].Ptr()
		pc := args[1].Ptr()
		r := fn(&Layer{c: lc}, &Paramdict{c: pc})
		ret.SetI32(r)
	})

	((*_c_layer)(l.c)).load_param = cls.Cptr()
	l.load_param = cls
}
func (l *Layer) SetLoadModel(fn func(l *Layer, mb *ModelBin) int32) {
	cls := c.NewFn(c.AbiDefault, c.I32, []c.Type{c.Pointer, c.Pointer}, func(args []c.Value, ret c.Value) {
		lc := args[0].Ptr()
		mc := args[1].Ptr()
		r := fn(&Layer{c: lc}, &ModelBin{c: mc})
		ret.SetI32(r)
	})

	((*_c_layer)(l.c)).load_model = cls.Cptr()
	l.load_model = cls
}
func (l *Layer) SetCreatePipeline(fn func(l *Layer, opt *Option) int32) {
	cls := c.NewFn(c.AbiDefault, c.I32, []c.Type{c.Pointer, c.Pointer}, func(args []c.Value, ret c.Value) {
		lc := args[0].Ptr()
		oc := args[1].Ptr()
		r := fn(&Layer{c: lc}, &Option{c: oc})
		ret.SetI32(r)
	})

	((*_c_layer)(l.c)).create_pipeline = cls.Cptr()
	l.create_pipeline = cls
}
func (l *Layer) SetDestroyPipeline(fn func(l *Layer, opt *Option) int32) {
	cls := c.NewFn(c.AbiDefault, c.I32, []c.Type{c.Pointer, c.Pointer}, func(args []c.Value, ret c.Value) {
		lc, oc := args[0].Ptr(), args[1].Ptr()
		r := fn(&Layer{c: lc}, &Option{c: oc})
		ret.SetI32(r)
	})

	((*_c_layer)(l.c)).destroy_pipeline = cls.Cptr()
	l.destroy_pipeline = cls
}
func (l *Layer) SetForward1(fn func(l *Layer, bottomBlob *Mat, topBlob *Mat, opt *Option) int32) {
	cls := c.NewFn(c.AbiDefault, c.I32, []c.Type{c.Pointer, c.Pointer, c.Pointer, c.Pointer}, func(args []c.Value, ret c.Value) {
		lc, bbc, tbc, oc := args[0].Ptr(), args[1].Ptr(), args[2].Ptr(), args[3].Ptr()
		tbc = usf.Pop(tbc)
		r := fn(&Layer{c: lc}, &Mat{c: bbc}, &Mat{c: tbc}, &Option{c: oc})
		ret.SetI32(r)
	})

	((*_c_layer)(l.c)).forward_1 = cls.Cptr()
	l.forward_1 = cls
}
func (l *Layer) SetForwardN(fn func(l *Layer, bottomBlob []*Mat, topBlob []*Mat, opt *Option) int32) {
	cls := c.NewFn(c.AbiDefault, c.I32,
		[]c.Type{c.Pointer, c.Pointer, c.I32, c.Pointer, c.I32, c.Pointer},
		func(args []c.Value, ret c.Value) {
			lc, oc := args[0].Ptr(), args[5].Ptr()
			bp, bn := args[1].Ptr(), args[2].I32()
			tp, tn := args[3].Ptr(), args[4].I32()

			bbgo := make([]*Mat, 0)
			bs := *(*[]unsafe.Pointer)(usf.Slice(bp, uint64(bn)))
			for i := range bs {
				bbgo = append(bbgo, &Mat{c: bs[i]})
			}

			ttgo := make([]*Mat, 0)
			ts := *(*[]unsafe.Pointer)(usf.Slice(tp, uint64(tn)))
			for i := range ts {
				ttgo = append(ttgo, &Mat{c: ts[i]})
			}

			r := fn(&Layer{c: lc}, bbgo, ttgo, &Option{c: oc})
			ret.SetI32(r)
		})

	((*_c_layer)(l.c)).forward_n = cls.Cptr()
	l.forward_n = cls
}
func (l *Layer) SetForwardInplace1(fn func(l *Layer, bottomTopBlob *Blob, opt *Option) int32) {
	cls := c.NewFn(c.AbiDefault, c.I32, []c.Type{c.Pointer, c.Pointer, c.Pointer},
		func(args []c.Value, ret c.Value) {
			lc, btb, oc := args[0].Ptr(), args[1].Ptr(), args[2].Ptr()
			r := fn(&Layer{c: lc}, &Blob{c: btb}, &Option{c: oc})
			ret.SetI32(r)
		})

	((*_c_layer)(l.c)).forward_inplace_1 = cls.Cptr()
	l.forward_inplace_1 = cls
}
func (l *Layer) SetForwardInplaceN(fn func(l *Layer, bottomTopBlob []*Blob, opt *Option) int32) {
	cls := c.NewFn(c.AbiDefault, c.I32, []c.Type{c.Pointer, c.Pointer, c.I32, c.Pointer},
		func(args []c.Value, ret c.Value) {
			lc, btb, btbn, oc := args[0].Ptr(), args[1].Ptr(), args[2].I32(), args[2].Ptr()
			btbgo := make([]*Blob, 0)
			btbs := *(*[]unsafe.Pointer)(usf.Slice(btb, uint64(btbn)))
			for i := range btbs {
				btbgo = append(btbgo, &Blob{c: btbs[i]})
			}

			r := fn(&Layer{c: lc}, btbgo, &Option{c: oc})
			ret.SetI32(r)
		})

	((*_c_layer)(l.c)).forward_inplace_n = cls.Cptr()
	l.forward_inplace_n = cls
}
func (l *Layer) Destroy() {
	ncnnLib.Call("ncnn_layer_destroy", c.Void, []c.Type{c.Pointer}, []interface{}{&l.c})
}
func (l *Layer) GetName() string {
	return ncnnLib.Call("ncnn_layer_get_name", c.Pointer,
		[]c.Type{c.Pointer}, []interface{}{&l.c}).StrFree()
}
func (l *Layer) GetTypeIndex() int32 {
	return ncnnLib.Call("ncnn_layer_get_typeindex", c.I32,
		[]c.Type{c.Pointer}, []interface{}{&l.c}).I32Free()
}
func (l *Layer) GetType() string {
	return ncnnLib.Call("ncnn_layer_get_type", c.Pointer,
		[]c.Type{c.Pointer}, []interface{}{&l.c}).StrFree()
}

func (l *Layer) GetOneBlobOnly() bool {
	return ncnnLib.Call("ncnn_layer_get_one_blob_only", c.I32,
		[]c.Type{c.Pointer}, []interface{}{&l.c}).BoolFree()
}
func (l *Layer) GetSupportInplace() bool {
	return ncnnLib.Call("ncnn_layer_get_support_inplace", c.I32,
		[]c.Type{c.Pointer}, []interface{}{&l.c}).BoolFree()
}
func (l *Layer) GetSupportVulkan() bool {
	return ncnnLib.Call("ncnn_layer_get_support_vulkan", c.I32,
		[]c.Type{c.Pointer}, []interface{}{&l.c}).BoolFree()
}
func (l *Layer) GetSupportPacking() bool {
	return ncnnLib.Call("ncnn_layer_get_support_packing", c.I32,
		[]c.Type{c.Pointer}, []interface{}{&l.c}).BoolFree()
}
func (l *Layer) GetSupportBf16Storage() bool {
	return ncnnLib.Call("ncnn_layer_get_support_bf16_storage", c.I32,
		[]c.Type{c.Pointer}, []interface{}{&l.c}).BoolFree()
}
func (l *Layer) GetSupportFp16Storage() bool {
	return ncnnLib.Call("ncnn_layer_get_support_fp16_storage", c.I32,
		[]c.Type{c.Pointer}, []interface{}{&l.c}).BoolFree()
}
func (l *Layer) GetSupportImageStorage() bool {
	return ncnnLib.Call("ncnn_layer_get_support_image_storage", c.I32,
		[]c.Type{c.Pointer}, []interface{}{&l.c}).BoolFree()
}

func (l *Layer) SetOneBlobOnly(enable bool) {
	b := c.CBool(enable)
	ncnnLib.Call("ncnn_layer_set_one_blob_only", c.Void,
		[]c.Type{c.Pointer, c.I32}, []interface{}{&l.c, &b})
}
func (l *Layer) SetSupportInplace(enable bool) {
	b := c.CBool(enable)
	ncnnLib.Call("ncnn_layer_set_support_inplace", c.Void,
		[]c.Type{c.Pointer, c.I32}, []interface{}{&l.c, &b})
}
func (l *Layer) SetSupportVulkan(enable bool) {
	b := c.CBool(enable)
	ncnnLib.Call("ncnn_layer_set_support_vulkan", c.Void,
		[]c.Type{c.Pointer, c.I32}, []interface{}{&l.c, &b})
}
func (l *Layer) SetSupportPacking(enable bool) {
	b := c.CBool(enable)
	ncnnLib.Call("ncnn_layer_set_support_packing", c.Void,
		[]c.Type{c.Pointer, c.I32}, []interface{}{&l.c, &b})
}
func (l *Layer) SetSupportBf16Storage(enable bool) {
	b := c.CBool(enable)
	ncnnLib.Call("ncnn_layer_set_support_bf16_storage", c.Void,
		[]c.Type{c.Pointer, c.I32}, []interface{}{&l.c, &b})
}
func (l *Layer) SetSupportFp16Storage(enable bool) {
	b := c.CBool(enable)
	ncnnLib.Call("ncnn_layer_set_support_fp16_storage", c.Void,
		[]c.Type{c.Pointer, c.I32}, []interface{}{&l.c, &b})
}
func (l *Layer) SetSupportImageStorage(enable bool) {
	b := c.CBool(enable)
	ncnnLib.Call("ncnn_layer_set_support_image_storage", c.Void,
		[]c.Type{c.Pointer, c.I32}, []interface{}{&l.c, &b})
}
func (l *Layer) GetBottomCount() int32 {
	return ncnnLib.Call("ncnn_layer_get_bottom_count", c.I32,
		[]c.Type{c.Pointer}, []interface{}{&l.c}).I32Free()
}
func (l *Layer) GetBottom(i int32) int32 {
	return ncnnLib.Call("ncnn_layer_get_bottom_count", c.I32,
		[]c.Type{c.Pointer, c.I32}, []interface{}{&l.c, &i}).I32Free()
}
func (l *Layer) GetTopCount() int32 {
	return ncnnLib.Call("ncnn_layer_get_top_count", c.I32,
		[]c.Type{c.Pointer}, []interface{}{&l.c}).I32Free()
}
func (l *Layer) GetTop(i int32) int32 {
	return ncnnLib.Call("ncnn_layer_get_top", c.I32,
		[]c.Type{c.Pointer, c.I32}, []interface{}{&l.c, &i}).I32Free()
}
func (l *Layer) GetBottomShape(i int32) (dims int32, w int32, h int32, ch int32) {
	dims, w, h, ch = int32(0), int32(1), int32(1), int32(1)
	_dims, _w, _h, _ch := &dims, &w, &h, &ch

	ncnnLib.Call("ncnn_blob_get_bottom_shape", c.Void,
		[]c.Type{c.Pointer, c.I32}, []interface{}{&l.c, &i, &_dims, &_w, &_h, &_ch})
	return
}
func (l *Layer) GetTopShape(i int32) (dims int32, w int32, h int32, ch int32) {
	dims, w, h, ch = int32(0), int32(1), int32(1), int32(1)
	_dims, _w, _h, _ch := &dims, &w, &h, &ch

	ncnnLib.Call("ncnn_blob_get_top_shape", c.Void,
		[]c.Type{c.Pointer, c.I32}, []interface{}{&l.c, &i, &_dims, &_w, &_h, &_ch})
	return
}

type _net_custom_layer struct {
	name      unsafe.Pointer
	creator   *c.Fn
	destroyer *c.Fn
}

func (l *_net_custom_layer) free() {
	if l == nil {
		return
	}
	if l.name != nil {
		usf.Free(l.name)
	}
	if l.creator != nil {
		l.creator.Free()
		return
	}
	if l.destroyer != nil {
		l.destroyer.Free()
	}
}

type Net struct {
	c             unsafe.Pointer
	_custom_layer []*_net_custom_layer
}

func CreateNet() *Net {
	n := ncnnLib.Call("ncnn_net_create", c.Pointer, nil, nil)
	return &Net{c: n.PtrFree()}
}
func (net *Net) Destroy() {
	ncnnLib.Call("ncnn_net_destroy", c.Void, []c.Type{c.Pointer}, []interface{}{&net.c})
	for i := range net._custom_layer {
		net._custom_layer[i].free()
	}
}
func (net *Net) GetOption() *Option {
	o := ncnnLib.Call("ncnn_net_get_option", c.Pointer,
		[]c.Type{c.Pointer}, []interface{}{&net.c})
	return &Option{c: o.PtrFree()}
}
func (net *Net) SetOption(opt *Option) {
	ncnnLib.Call("ncnn_net_set_option", c.Void,
		[]c.Type{c.Pointer, c.Pointer}, []interface{}{&net.c, &opt.c})
}
func (net *Net) RegisterCustomLayerByType(typ string, creator func() *Layer, destroyer func(*Layer)) {
	t := c.CStr(typ)

	create := c.NewFn(c.AbiDefault, c.Pointer, []c.Type{c.Pointer},
		func(args []c.Value, ret c.Value) {
			l := creator()
			ret.SetPtr(l.c)
		})

	dest := c.NewFn(c.AbiDefault, c.Void, []c.Type{c.Pointer, c.Pointer},
		func(args []c.Value, ret c.Value) {
			destroyer(&Layer{c: args[0].Ptr()})
		})

	cc := create.Cptr()
	d := dest.Cptr()
	a := usf.MallocOf(1, 8)
	usf.Memset(a, 0, 8)

	ncnnLib.Call("ncnn_net_register_custom_layer_by_type", c.Void,
		[]c.Type{c.Pointer, c.Pointer, c.Pointer, c.Pointer, c.Pointer},
		[]interface{}{&net.c, &t, &cc, &d, a})
	net._custom_layer = append(net._custom_layer, &_net_custom_layer{
		name:      t,
		creator:   create,
		destroyer: dest,
	})
}
func (net *Net) RegisterCustomLayerByTypeIndex(typIdx int32, creator func() *Layer, destroyer func(*Layer)) {
	create := c.NewFn(c.AbiDefault, c.Pointer, []c.Type{c.Pointer},
		func(args []c.Value, ret c.Value) {
			l := creator()
			ret.SetPtr(l.c)
		})
	dest := c.NewFn(c.AbiDefault, c.Void, []c.Type{c.Pointer, c.Pointer},
		func(args []c.Value, ret c.Value) {
			destroyer(&Layer{c: args[0].Ptr()})
		})

	cc, d, a := create.Cptr(), dest.Cptr(), usf.MallocOf(1, 8)
	usf.Memset(a, 0, 8)
	defer usf.Free(a)

	ncnnLib.Call("ncnn_net_register_custom_layer_by_typeindex", c.Void,
		[]c.Type{c.Pointer, c.I32, c.Pointer, c.Pointer, c.Pointer},
		[]interface{}{&net.c, &typIdx, &cc, &d, &a})
	net._custom_layer = append(net._custom_layer, &_net_custom_layer{
		creator:   create,
		destroyer: dest,
	})
}
func (net *Net) LoadParam(path string) int32 {
	p := c.CStr(path)
	defer usf.Free(p)
	return ncnnLib.Call("ncnn_net_load_param", c.I32,
		[]c.Type{c.Pointer, c.Pointer}, []interface{}{&net.c, &p}).I32Free()
}
func (net *Net) LoadParamBin(path string) int32 {
	p := c.CStr(path)
	defer usf.Free(p)
	return ncnnLib.Call("ncnn_net_load_param_bin", c.I32,
		[]c.Type{c.Pointer, c.Pointer}, []interface{}{&net.c, &p}).I32Free()
}
func (net *Net) LoadModel(path string) int32 {
	p := c.CStr(path)
	defer usf.Free(p)
	return ncnnLib.Call("ncnn_net_load_model", c.I32,
		[]c.Type{c.Pointer, c.Pointer}, []interface{}{&net.c, &p}).I32Free()
}
func (net *Net) LoadParamMemory(data []byte) int32 {
	d := &data[0]
	return ncnnLib.Call("ncnn_net_load_param_memory", c.I32,
		[]c.Type{c.Pointer, c.Pointer}, []interface{}{&net.c, &d}).I32Free()
}
func (net *Net) LoadParamBinMemory(data []byte) int32 {
	d := &data[0]
	return ncnnLib.Call("ncnn_net_load_param_bin_memory", c.I32,
		[]c.Type{c.Pointer, c.Pointer}, []interface{}{&net.c, &d}).I32Free()
}
func (net *Net) LoadModelMemory(data []byte) int32 {
	d := &data[0]
	return ncnnLib.Call("ncnn_net_load_model_memory", c.I32,
		[]c.Type{c.Pointer, c.Pointer}, []interface{}{&net.c, &d}).I32Free()
}
func (net *Net) Clear() {
	ncnnLib.Call("ncnn_net_clear", c.Void, []c.Type{c.Pointer}, []interface{}{&net.c})
}
func (net *Net) GetInputCount() int32 {
	return ncnnLib.Call("ncnn_net_get_input_count", c.I32,
		[]c.Type{c.Pointer}, []interface{}{&net.c}).I32Free()
}
func (net *Net) GetOutputCount() int32 {
	return ncnnLib.Call("ncnn_net_get_output_count", c.I32,
		[]c.Type{c.Pointer}, []interface{}{&net.c}).I32Free()
}
func (net *Net) GetInputName(i int32) string {
	return ncnnLib.Call("ncnn_net_get_input_name", c.Pointer,
		[]c.Type{c.Pointer, c.I32}, []interface{}{&net.c, &i}).StrFree()
}
func (net *Net) GetOutputName(i int32) string {
	return ncnnLib.Call("ncnn_net_get_output_name", c.Pointer,
		[]c.Type{c.Pointer, c.I32}, []interface{}{&net.c, &i}).StrFree()
}

type Extractor struct{ c unsafe.Pointer }

func CreateExtractor(net *Net) *Extractor {
	e := ncnnLib.Call("ncnn_extractor_create", c.Pointer, []c.Type{c.Pointer}, []interface{}{&net.c})
	return &Extractor{c: e.PtrFree()}
}
func (ex *Extractor) Destroy() {
	ncnnLib.Call("ncnn_extractor_destroy", c.Void, []c.Type{c.Pointer}, []interface{}{&ex.c})
}
func (ex *Extractor) SetOption(opt *Option) {
	ncnnLib.Call("ncnn_extractor_set_option", c.Void,
		[]c.Type{c.Pointer, c.Pointer}, []interface{}{&ex.c, &opt.c})
}
func (ex *Extractor) Input(name string, mat *Mat) int32 {
	m := c.CStr(name)
	defer usf.Free(m)
	return ncnnLib.Call("ncnn_extractor_input", c.I32,
		[]c.Type{c.Pointer, c.Pointer, c.Pointer}, []interface{}{&ex.c, &m, &mat.c}).I32Free()
}
func (ex *Extractor) Extract(name string) *Mat {
	m, n := usf.Malloc(1, 8), c.CStr(name)
	defer usf.Free(m)
	defer usf.Free(n)

	ncnnLib.Call("ncnn_extractor_extract", c.I32,
		[]c.Type{c.Pointer, c.Pointer, c.Pointer}, []interface{}{&ex.c, &n, &m})
	return &Mat{c: usf.Pop(m)}
}

func (ex *Extractor) InputIndex(idx int32, mat *Mat) int32 {
	return ncnnLib.Call("ncnn_extractor_input_index", c.I32,
		[]c.Type{c.Pointer, c.I32, c.Pointer}, []interface{}{&ex.c, &idx, &mat.c}).I32Free()
}

func (ex *Extractor) ExtractIndex(idx int32) *Mat {
	m := usf.Malloc(1, 8)
	defer usf.Free(m)

	ncnnLib.Call("ncnn_extractor_extract_index", c.I32,
		[]c.Type{c.Pointer, c.I32, c.Pointer}, []interface{}{&ex.c, &idx, &m})
	return &Mat{c: usf.Pop(m)}
}
