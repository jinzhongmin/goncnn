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
	ncnnLib, err = c.NewLib(path, mod)
	if err != nil {
		panic(err)
	}
}

func Version() string {
	return ncnnLib.Call(_func_ncnn_version_, nil).StrFree()
}

type Allocator struct{ c unsafe.Pointer }

func CreatePoolAllocator() *Allocator {
	return &Allocator{c: ncnnLib.Call(_func_ncnn_allocator_create_pool_allocator_, nil).PtrFree()}
}
func CreateUnlockedPoolAllocator() *Allocator {
	return &Allocator{c: ncnnLib.Call(_func_ncnn_allocator_create_unlocked_pool_allocator_, nil).PtrFree()}
}
func (alc *Allocator) Destroy() {
	ncnnLib.Call(_func_ncnn_allocator_destroy_, []interface{}{&alc.c})
}

type Option struct{ c unsafe.Pointer }

func CreateOption() *Option {
	return &Option{c: ncnnLib.Call(_func_ncnn_option_create_, nil).PtrFree()}
}
func (opt *Option) Destroy() {
	ncnnLib.Call(_func_ncnn_option_destroy_, []interface{}{&opt.c})
}
func (opt *Option) GetNumThreads() int32 {
	return ncnnLib.Call(_func_ncnn_option_get_num_threads_, []interface{}{&opt.c}).I32Free()
}
func (opt *Option) SetNumThreads(num int32) {
	ncnnLib.Call(_func_ncnn_option_set_num_threads_, []interface{}{&opt.c, &num})
}
func (opt *Option) GetUseLocalPoolAllocator() int32 {
	return ncnnLib.Call(_func_ncnn_option_get_use_local_pool_allocator_, []interface{}{&opt.c}).I32Free()
}
func (opt *Option) SetUseLocalPoolAllocator(n int32) {
	ncnnLib.Call(_func_ncnn_option_set_use_local_pool_allocator_, []interface{}{&opt.c, &n})
}
func (opt *Option) SetBlobAllocator(alc *Allocator) {
	ncnnLib.Call(_func_ncnn_option_set_blob_allocator_, []interface{}{&opt.c, &alc.c})
}
func (opt *Option) SetWorkspaceAllocator(alc *Allocator) {
	ncnnLib.Call(_func_ncnn_option_set_workspace_allocator_, []interface{}{&opt.c, &alc.c})
}
func (opt *Option) GetUseVulkanCompute() bool {
	return ncnnLib.Call(_func_ncnn_option_get_use_vulkan_compute_, []interface{}{&opt.c}).BoolFree()
}
func (opt *Option) SetUseVulkanCompute(enable bool) {
	n := c.CBool(enable)
	ncnnLib.Call(_func_ncnn_option_set_use_vulkan_compute_, []interface{}{&opt.c, &n})
}

type Mat struct{ c unsafe.Pointer }

func (mat *Mat) Destroy() {
	ncnnLib.Call(_func_ncnn_mat_destroy_, []interface{}{&mat.c})
}
func CreateMat() *Mat {
	return &Mat{c: ncnnLib.Call(_func_ncnn_mat_create_, nil).PtrFree()}
}
func CreateMat1D(w int32, alc *Allocator) *Mat {
	return &Mat{c: ncnnLib.Call(_func_ncnn_mat_create_1d_, []interface{}{&w, &alc.c}).PtrFree()}
}
func CreateMat2D(w, h int32, alc *Allocator) *Mat {
	return &Mat{c: ncnnLib.Call(_func_ncnn_mat_create_2d_, []interface{}{&w, &h, &alc.c}).PtrFree()}
}
func CreateMat3D(w, h, ch int32, alc *Allocator) *Mat {
	return &Mat{c: ncnnLib.Call(_func_ncnn_mat_create_3d_, []interface{}{&w, &h, &ch, &alc.c}).PtrFree()}
}
func CreateMat4D(w, h, d, ch int32, alc *Allocator) *Mat {
	return &Mat{c: ncnnLib.Call(_func_ncnn_mat_create_4d_, []interface{}{&w, &h, &d, &ch, &alc.c}).PtrFree()}
}
func CreateMatExternal1D(w int32, data interface{}, alc *Allocator) *Mat {
	dt := usf.AddrOf(data)
	return &Mat{c: ncnnLib.Call(_func_ncnn_mat_create_external_1d_, []interface{}{&w, &dt, &alc.c}).PtrFree()}
}
func CreateMatExternal2D(w, h int32, data interface{}, alc *Allocator) *Mat {
	dt := usf.AddrOf(data)
	return &Mat{c: ncnnLib.Call(_func_ncnn_mat_create_external_2d_, []interface{}{&w, &h, &dt, &alc.c}).PtrFree()}
}
func CreateMatExternal3D(w, h, ch int32, data interface{}, alc *Allocator) *Mat {
	dt := usf.AddrOf(data)
	return &Mat{c: ncnnLib.Call(_func_ncnn_mat_create_external_3d_, []interface{}{&w, &h, &ch, &dt, &alc.c}).PtrFree()}
}
func CreateMatExternal4D(w, h, d, ch int32, data interface{}, alc *Allocator) *Mat {
	dt := usf.AddrOf(data)
	return &Mat{c: ncnnLib.Call(_func_ncnn_mat_create_external_4d_, []interface{}{&w, &h, &d, &ch, &dt, &alc.c}).PtrFree()}
}
func CreateMat1DElm(w int32, elmSize uint64, elmPack int32, alc *Allocator) *Mat {
	return &Mat{c: ncnnLib.Call(_func_ncnn_mat_create_1d_elem_, []interface{}{&w, &elmSize, &elmPack, &alc.c}).PtrFree()}
}
func CreateMat2DElm(w, h int32, elmSize uint64, elmPack int32, alc *Allocator) *Mat {
	return &Mat{c: ncnnLib.Call(_func_ncnn_mat_create_2d_elem_, []interface{}{&w, &h, &elmSize, &elmPack, &alc.c}).PtrFree()}
}
func CreateMat3DElm(w, h, ch int32, elmSize uint64, elmPack int32, alc *Allocator) *Mat {
	return &Mat{c: ncnnLib.Call(_func_ncnn_mat_create_3d_elem_, []interface{}{&w, &h, &ch, &elmSize, &elmPack, &alc.c}).PtrFree()}
}
func CreateMat4DElm(w, h, d, ch int32, elmSize uint64, elmPack int32, alc *Allocator) *Mat {
	return &Mat{c: ncnnLib.Call(_func_ncnn_mat_create_4d_elem_, []interface{}{&w, &h, &d, &ch, &elmSize, &elmPack, &alc.c}).PtrFree()}
}
func CreateMatExternal1DElm(w int32, data interface{}, elmSize uint64, elmPack int32, alc *Allocator) *Mat {
	dt := usf.AddrOf(data)
	return &Mat{c: ncnnLib.Call(_func_ncnn_mat_create_external_1d_elem_, []interface{}{&w, &dt, &elmSize, &elmPack, &alc.c}).PtrFree()}
}
func CreateMatExternal2DElm(w, h int32, data interface{},
	elmSize uint64, elmPack int32, alc *Allocator) *Mat {
	dt := usf.AddrOf(data)
	return &Mat{c: ncnnLib.Call(_func_ncnn_mat_create_external_2d_elem_, []interface{}{&w, &h, &dt, &elmSize, &elmPack, &alc.c}).PtrFree()}
}
func CreateMatExternal3DElm(w, h, ch int32, data interface{},
	elmSize uint64, elmPack int32, alc *Allocator) *Mat {
	dt := usf.AddrOf(data)
	return &Mat{c: ncnnLib.Call(_func_ncnn_mat_create_external_3d_elem_, []interface{}{&w, &h, &ch, &dt, &elmSize, &elmPack, &alc.c}).PtrFree()}
}
func CreateMatExternal4DElm(w, h, d, ch int32, data interface{},
	elmSize uint64, elmPack int32, alc *Allocator) *Mat {
	dt := usf.AddrOf(data)
	return &Mat{c: ncnnLib.Call(_func_ncnn_mat_create_external_4d_elem_, []interface{}{&w, &h, &dt, &ch, &d, &elmSize, &elmPack, &alc.c}).PtrFree()}
}
func (mat *Mat) FillFloat(f float32) {
	ncnnLib.Call(_func_ncnn_mat_fill_float_, []interface{}{&mat.c, &f})
}
func (mat *Mat) Clone(alc *Allocator) *Mat {
	return &Mat{c: ncnnLib.Call(_func_ncnn_mat_clone_, []interface{}{&mat.c, &alc.c}).PtrFree()}
}
func (mat *Mat) Reshape1D(w int32, alc *Allocator) *Mat {
	return &Mat{c: ncnnLib.Call(_func_ncnn_mat_reshape_1d_, []interface{}{&mat.c, &w, &alc.c}).PtrFree()}
}
func (mat *Mat) Reshape2D(w, h int32, alc *Allocator) *Mat {
	return &Mat{c: ncnnLib.Call(_func_ncnn_mat_reshape_2d_, []interface{}{&mat.c, &w, &h, &alc.c}).PtrFree()}
}
func (mat *Mat) Reshape3D(w, h, ch int32, alc *Allocator) *Mat {
	return &Mat{c: ncnnLib.Call(_func_ncnn_mat_reshape_3d_, []interface{}{&mat.c, &w, &h, &ch, &alc.c}).PtrFree()}
}
func (mat *Mat) Reshape4D(w, h, d, ch int32, alc *Allocator) *Mat {
	return &Mat{c: ncnnLib.Call(_func_ncnn_mat_reshape_4d_, []interface{}{&mat.c, &w, &h, &d, &ch, &alc.c}).PtrFree()}
}
func (mat *Mat) GetDims() int32 {
	return ncnnLib.Call(_func_ncnn_mat_get_dims_, []interface{}{&mat.c}).I32Free()
}
func (mat *Mat) GetW() int32 {
	return ncnnLib.Call(_func_ncnn_mat_get_w_, []interface{}{&mat.c}).I32Free()
}
func (mat *Mat) GetH() int32 {
	return ncnnLib.Call(_func_ncnn_mat_get_h_, []interface{}{&mat.c}).I32Free()
}
func (mat *Mat) GetD() int32 {
	return ncnnLib.Call(_func_ncnn_mat_get_d_, []interface{}{&mat.c}).I32Free()
}
func (mat *Mat) GetC() int32 {
	return ncnnLib.Call(_func_ncnn_mat_get_c_, []interface{}{&mat.c}).I32Free()
}
func (mat *Mat) GetElemsize() uint64 {
	return ncnnLib.Call(_func_ncnn_mat_get_elemsize_, []interface{}{&mat.c}).U64Free()
}
func (mat *Mat) GetElempack() int32 {
	return ncnnLib.Call(_func_ncnn_mat_get_elempack_, []interface{}{&mat.c}).I32Free()
}
func (mat *Mat) GetCstep() uint64 {
	return ncnnLib.Call(_func_ncnn_mat_get_cstep_, []interface{}{&mat.c}).U64Free()
}
func (mat *Mat) GetData() unsafe.Pointer {
	return ncnnLib.Call(_func_ncnn_mat_get_data_, []interface{}{&mat.c}).PtrFree()
}
func (mat *Mat) GetChannelData(ch int32) unsafe.Pointer {
	return ncnnLib.Call(_func_ncnn_mat_get_channel_data_, []interface{}{&mat.c, &ch}).PtrFree()
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
	return &Mat{c: ncnnLib.Call(_func_ncnn_mat_from_pixels_, []interface{}{&d, &typ, &w, &h, &stride, &alc}).PtrFree()}
}
func CreateMatFromPixelsResize(pixels []byte, typ PixelType, w, h, stride int32, targetW, targetH int32, alc *Allocator) *Mat {
	d := &pixels[0]
	return &Mat{c: ncnnLib.Call(_func_ncnn_mat_from_pixels_resize_, []interface{}{&d, &typ, &w, &h, &stride, &targetW, &targetH, &alc}).PtrFree()}
}
func CreateMatFromPixelsRoi(pixels []byte, typ PixelType, w, h, stride int32, roix, roiy, roiw, roih int32, alc *Allocator) *Mat {
	d := &pixels[0]
	return &Mat{c: ncnnLib.Call(_func_ncnn_mat_from_pixels_roi_, []interface{}{&d, &typ, &w, &h, &stride, &roix, &roiy, &roiw, &roih, &alc}).PtrFree()}
}
func CreateMatFromPixelsRoiResize(pixels []byte, typ PixelType, w, h, stride int32, roix, roiy, roiw, roih int32, targetW, targetH int32, alc *Allocator) *Mat {
	d := &pixels[0]
	return &Mat{c: ncnnLib.Call(_func_ncnn_mat_from_pixels_roi_resize_, []interface{}{&d, &typ, &w, &h, &stride, &roix, &roiy, &roiw, &roih, &targetW, &targetH, &alc}).PtrFree()}
}

//	type slice struct {
//	    ptr unsafe.Pointer
//	    len int
//	    cap int
//	}
//
//	func (mat *Mat) ToPixels(typ PixelType, stride int32, outData interface{}) {
//	    p := (*(*[2]unsafe.Pointer)(unsafe.Pointer(&outData)))[1]
//	    p = (*(*slice)((unsafe.Pointer)((*unsafe.Pointer)(p)))).ptr
//	    ncnnLib.Call("ncnn_mat_to_pixels", c.Void,
//	        []c.Type{c.Pointer, c.Pointer, c.I32, c.I32},
//	        []interface{}{mat.c, p, typ, stride})
//	}
//
//	func (mat *Mat) ToPixelsResize(typ PixelType, targetW, targetH, targetStride int32, outData interface{}) {
//	    p := (*(*[2]unsafe.Pointer)(unsafe.Pointer(&outData)))[1]
//	    p = (*(*slice)((unsafe.Pointer)((*unsafe.Pointer)(p)))).ptr
//	    ncnnLib.Call("ncnn_mat_to_pixels", c.Void,
//	        []c.Type{c.Pointer, c.Pointer, c.I32, c.I32, c.I32, c.I32},
//	        []interface{}{mat.c, p, typ, targetW, targetH, targetStride})
//	}
func (mat *Mat) SubstractMeanNormalize(mean, norm []float32) {
	m := &mean[0]
	n := &norm[0]
	ncnnLib.Call(_func_ncnn_mat_substract_mean_normalize_, []interface{}{&mat.c, &m, &n})
}
func (mat *Mat) ConvertPacking(elempack int32, opt *Option) *Mat {
	m := usf.Malloc(8)
	defer usf.Free(m)
	ncnnLib.Call(_func_ncnn_convert_packing_, []interface{}{&mat.c, &m, &elempack, &opt.c})
	return &Mat{c: usf.Pop(m)}
}
func (mat *Mat) Flatten(opt *Option) *Mat {
	m := usf.MallocN(1, 8)
	defer usf.Free(m)
	ncnnLib.Call(_func_ncnn_flatten_, []interface{}{&mat.c, &m, &opt.c})
	return &Mat{c: usf.Pop(m)}
}

type Blob struct{ c unsafe.Pointer }

func (blob *Blob) GetName() string {
	return ncnnLib.Call(_func_ncnn_blob_get_name_, []interface{}{&blob.c}).StrFree()
}
func (blob *Blob) GetProducer() int32 {
	return ncnnLib.Call(_func_ncnn_blob_get_producer_, []interface{}{&blob.c}).I32Free()
}
func (blob *Blob) GetConsumer() int32 {
	return ncnnLib.Call(_func_ncnn_blob_get_consumer_, []interface{}{&blob.c}).I32Free()
}

type Paramdict struct{ c unsafe.Pointer }

func CreateParamdict() *Paramdict {
	return &Paramdict{c: ncnnLib.Call(_func_ncnn_paramdict_create_, nil).PtrFree()}
}
func (pd *Paramdict) Destroy() {
	ncnnLib.Call(_func_ncnn_paramdict_destroy_, []interface{}{&pd.c})
}
func (pd *Paramdict) GetType(id int32) int32 {
	return ncnnLib.Call(_func_ncnn_paramdict_get_type_, []interface{}{&pd.c, &id}).I32Free()
}
func (pd *Paramdict) GetInt32(id int32, def int32) int32 {
	return ncnnLib.Call(_func_ncnn_paramdict_get_int_, []interface{}{&pd.c, &id, &def}).I32Free()
}
func (pd *Paramdict) GetFloat32(id int32, def float32) float32 {
	return ncnnLib.Call(_func_ncnn_paramdict_get_float_, []interface{}{&pd.c, &id, &def}).F32Free()
}
func (pd *Paramdict) GetArray(id int32, def *Mat) *Mat {
	return &Mat{c: ncnnLib.Call(_func_ncnn_paramdict_get_array_, []interface{}{&pd.c, &id, &def}).PtrFree()}
}
func (pd *Paramdict) SetInt32(id int32, i int32) {
	ncnnLib.Call(_func_ncnn_paramdict_set_int_, []interface{}{&pd.c, &id, &i})
}
func (pd *Paramdict) SetFloat32(id int32, f float32) {
	ncnnLib.Call(_func_ncnn_paramdict_set_float_, []interface{}{&pd.c, &id, &f})
}
func (pd *Paramdict) SetArray(id int32, m *Mat) {
	ncnnLib.Call(_func_ncnn_paramdict_set_array_, []interface{}{&pd.c, &id, &m})
}

type DataReader struct{ c unsafe.Pointer }

func CreateDataReader() *DataReader {
	return &DataReader{c: ncnnLib.Call(_func_ncnn_datareader_create_, nil).PtrFree()}
}
func CreateDataReaderFromMemory(mem []byte) *DataReader {
	p := unsafe.Pointer(&mem[0])
	pp := usf.Malloc(8)
	usf.Push(pp, p)
	defer usf.Free(pp)
	return &DataReader{c: ncnnLib.Call(_func_ncnn_datareader_create_from_memory_, []interface{}{&pp}).PtrFree()}
}
func (d *DataReader) Destroy() {
	ncnnLib.Call(_func_ncnn_datareader_destroy_, []interface{}{&d.c})
}

type ModelBin struct{ c unsafe.Pointer }

func CreateModelBinFromDataReader(dr *DataReader) *ModelBin {
	return &ModelBin{c: ncnnLib.Call(_func_ncnn_modelbin_create_from_datareader_, []interface{}{&dr.c}).PtrFree()}
}

//	func CreateModelBinFromMatArray(mats []*Mat) *ModelBin {
//		l := uint32(len(mats))
//		ms := usf.MallocN(uint64(l), 8)
//		defer usf.Free(ms)
//		m := ncnnLib.Call(_func_ncnn_modelbin_create_from_mat_array_, []interface{}{&ms, &l})
//		return &ModelBin{c: m.PtrFree()}
//	}
func (mb *ModelBin) Destroy() {
	ncnnLib.Call(_func_ncnn_modelbin_destroy_, []interface{}{&mb.c})
}

type Layer struct {
	c                 unsafe.Pointer
	load_param        *c.Callback
	load_model        *c.Callback
	create_pipeline   *c.Callback
	destroy_pipeline  *c.Callback
	forward_1         *c.Callback
	forward_n         *c.Callback
	forward_inplace_1 *c.Callback
	forward_inplace_n *c.Callback
}
type _c_layer struct {
	this              unsafe.Pointer
	load_param        unsafe.Pointer
	load_model        unsafe.Pointer
	create_pipeline   unsafe.Pointer
	destroy_pipeline  unsafe.Pointer
	forward_1         unsafe.Pointer
	forward_n         unsafe.Pointer
	forward_inplace_1 unsafe.Pointer
	forward_inplace_n unsafe.Pointer
}

func CreateLayer() *Layer {
	return &Layer{c: ncnnLib.Call(_func_ncnn_layer_create_, nil).PtrFree()}
}
func CreateLayerByTypeindex(typIdx int32) *Layer {
	return &Layer{c: ncnnLib.Call(_func_ncnn_layer_create_by_typeindex_, []interface{}{&typIdx}).PtrFree()}
}
func CreateLayerByType(typ string) *Layer {
	t := c.CStr(typ)
	defer usf.Free(t)
	return &Layer{c: ncnnLib.Call(_func_ncnn_layer_create_by_type_, []interface{}{&t}).PtrFree()}
}

var loadParamCallbackPrototype = c.DefineCallbackPrototype(c.AbiDefault, c.I32, []c.Type{c.Pointer, c.Pointer})

func (l *Layer) SetLoadParam(fn func(l *Layer, pd *Paramdict) int32) {
	if l.load_param != nil {
		l.load_param.Free()
	}

	cls := loadParamCallbackPrototype.CreateCallback(func(args []*c.Value, ret *c.Value) {
		lc := args[0].Ptr()
		pc := args[1].Ptr()
		r := fn(&Layer{c: lc}, &Paramdict{c: pc})
		ret.SetI32(r)
	})
	((*_c_layer)(l.c)).load_param = cls.CFuncPtr()
	l.load_param = cls
}

var loadModelCallbackPrototype = c.DefineCallbackPrototype(c.AbiDefault, c.I32, []c.Type{c.Pointer, c.Pointer})

func (l *Layer) SetLoadModel(fn func(l *Layer, mb *ModelBin) int32) {
	if l.load_model != nil {
		l.load_model.Free()
	}

	cls := loadModelCallbackPrototype.CreateCallback(func(args []*c.Value, ret *c.Value) {
		lc := args[0].Ptr()
		mc := args[1].Ptr()
		r := fn(&Layer{c: lc}, &ModelBin{c: mc})
		ret.SetI32(r)
	})
	((*_c_layer)(l.c)).load_model = cls.CFuncPtr()
	l.load_model = cls
}

var createPipelineCallbackPrototype = c.DefineCallbackPrototype(c.AbiDefault, c.I32, []c.Type{c.Pointer, c.Pointer})

func (l *Layer) SetCreatePipeline(fn func(l *Layer, opt *Option) int32) {
	if l.create_pipeline != nil {
		l.create_pipeline.Free()
	}

	cls := createPipelineCallbackPrototype.CreateCallback(func(args []*c.Value, ret *c.Value) {
		lc := args[0].Ptr()
		oc := args[1].Ptr()
		r := fn(&Layer{c: lc}, &Option{c: oc})
		ret.SetI32(r)
	})
	((*_c_layer)(l.c)).create_pipeline = cls.CFuncPtr()
	l.create_pipeline = cls
}

var destroyPipelineCallbackPrototype = c.DefineCallbackPrototype(c.AbiDefault, c.I32, []c.Type{c.Pointer, c.Pointer})

func (l *Layer) SetDestroyPipeline(fn func(l *Layer, opt *Option) int32) {
	if l.destroy_pipeline != nil {
		l.destroy_pipeline.Free()
	}

	cls := destroyPipelineCallbackPrototype.CreateCallback(func(args []*c.Value, ret *c.Value) {
		lc, oc := args[0].Ptr(), args[1].Ptr()
		r := fn(&Layer{c: lc}, &Option{c: oc})
		ret.SetI32(r)
	})
	((*_c_layer)(l.c)).destroy_pipeline = cls.CFuncPtr()
	l.destroy_pipeline = cls
}

var forward1CallbackPrototype = c.DefineCallbackPrototype(c.AbiDefault, c.I32, []c.Type{c.Pointer, c.Pointer, c.Pointer, c.Pointer})

func (l *Layer) SetForward1(fn func(l *Layer, bottomBlob *Mat, topBlob *Mat, opt *Option) int32) {
	if l.forward_1 != nil {
		l.forward_1.Free()
	}

	cls := forward1CallbackPrototype.CreateCallback(func(args []*c.Value, ret *c.Value) {
		lc, bbc, tbc, oc := args[0].Ptr(), args[1].Ptr(), args[2].Ptr(), args[3].Ptr()
		tbc = usf.Pop(tbc)
		r := fn(&Layer{c: lc}, &Mat{c: bbc}, &Mat{c: tbc}, &Option{c: oc})
		ret.SetI32(r)
	})
	((*_c_layer)(l.c)).forward_1 = cls.CFuncPtr()
	l.forward_1 = cls

}

var forwardNCallbackPrototype = c.DefineCallbackPrototype(c.AbiDefault, c.I32, []c.Type{c.Pointer, c.Pointer, c.I32, c.Pointer, c.I32, c.Pointer})

func (l *Layer) SetForwardN(fn func(l *Layer, bottomBlob []*Mat, topBlob []*Mat, opt *Option) int32) {
	if l.forward_n != nil {
		l.forward_n.Free()
	}

	cls := forwardNCallbackPrototype.CreateCallback(func(args []*c.Value, ret *c.Value) {
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
	((*_c_layer)(l.c)).forward_n = cls.CFuncPtr()
	l.forward_n = cls
}

var forwardInplace1CallbackPrototype = c.DefineCallbackPrototype(c.AbiDefault, c.I32, []c.Type{c.Pointer, c.Pointer, c.Pointer})

func (l *Layer) SetForwardInplace1(fn func(l *Layer, bottomTopBlob *Blob, opt *Option) int32) {
	if l.forward_inplace_1 != nil {
		l.forward_inplace_1.Free()
	}

	cls := forwardInplace1CallbackPrototype.CreateCallback(func(args []*c.Value, ret *c.Value) {
		lc, btb, oc := args[0].Ptr(), args[1].Ptr(), args[2].Ptr()
		r := fn(&Layer{c: lc}, &Blob{c: btb}, &Option{c: oc})
		ret.SetI32(r)
	})
	((*_c_layer)(l.c)).forward_inplace_1 = cls.CFuncPtr()
	l.forward_inplace_1 = cls
}

var forwardInplaceNCallbackPrototype = c.DefineCallbackPrototype(c.AbiDefault, c.I32, []c.Type{c.Pointer, c.Pointer, c.I32, c.Pointer})

func (l *Layer) SetForwardInplaceN(fn func(l *Layer, bottomTopBlob []*Blob, opt *Option) int32) {
	if l.forward_inplace_n != nil {
		l.forward_inplace_n.Free()
	}

	cls := forwardInplaceNCallbackPrototype.CreateCallback(func(args []*c.Value, ret *c.Value) {
		lc, btb, btbn, oc := args[0].Ptr(), args[1].Ptr(), args[2].I32(), args[2].Ptr()
		btbgo := make([]*Blob, 0)
		btbs := *(*[]unsafe.Pointer)(usf.Slice(btb, uint64(btbn)))
		for i := range btbs {
			btbgo = append(btbgo, &Blob{c: btbs[i]})
		}
		r := fn(&Layer{c: lc}, btbgo, &Option{c: oc})
		ret.SetI32(r)
	})
	((*_c_layer)(l.c)).forward_inplace_n = cls.CFuncPtr()
	l.forward_inplace_n = cls
}
func (l *Layer) Destroy() {
	ncnnLib.Call(_func_ncnn_layer_destroy_, []interface{}{&l.c})
}
func (l *Layer) GetName() string {
	return ncnnLib.Call(_func_ncnn_layer_get_name_, []interface{}{&l.c}).StrFree()
}
func (l *Layer) GetTypeIndex() int32 {
	return ncnnLib.Call(_func_ncnn_layer_get_typeindex_, []interface{}{&l.c}).I32Free()
}
func (l *Layer) GetType() string {
	return ncnnLib.Call(_func_ncnn_layer_get_type_, []interface{}{&l.c}).StrFree()
}
func (l *Layer) GetOneBlobOnly() bool {
	return ncnnLib.Call(_func_ncnn_layer_get_one_blob_only_, []interface{}{&l.c}).BoolFree()
}
func (l *Layer) GetSupportInplace() bool {
	return ncnnLib.Call(_func_ncnn_layer_get_support_inplace_, []interface{}{&l.c}).BoolFree()
}
func (l *Layer) GetSupportVulkan() bool {
	return ncnnLib.Call(_func_ncnn_layer_get_support_vulkan_, []interface{}{&l.c}).BoolFree()
}
func (l *Layer) GetSupportPacking() bool {
	return ncnnLib.Call(_func_ncnn_layer_get_support_packing_, []interface{}{&l.c}).BoolFree()
}
func (l *Layer) GetSupportBf16Storage() bool {
	return ncnnLib.Call(_func_ncnn_layer_get_support_bf16_storage_, []interface{}{&l.c}).BoolFree()
}
func (l *Layer) GetSupportFp16Storage() bool {
	return ncnnLib.Call(_func_ncnn_layer_get_support_fp16_storage_, []interface{}{&l.c}).BoolFree()
}
func (l *Layer) GetSupportImageStorage() bool {
	return ncnnLib.Call(_func_ncnn_layer_get_support_image_storage_, []interface{}{&l.c}).BoolFree()
}
func (l *Layer) SetOneBlobOnly(enable bool) {
	b := c.CBool(enable)
	ncnnLib.Call(_func_ncnn_layer_set_one_blob_only_, []interface{}{&l.c, &b})
}
func (l *Layer) SetSupportInplace(enable bool) {
	b := c.CBool(enable)
	ncnnLib.Call(_func_ncnn_layer_set_support_inplace_, []interface{}{&l.c, &b})
}
func (l *Layer) SetSupportVulkan(enable bool) {
	b := c.CBool(enable)
	ncnnLib.Call(_func_ncnn_layer_set_support_vulkan_, []interface{}{&l.c, &b})
}
func (l *Layer) SetSupportPacking(enable bool) {
	b := c.CBool(enable)
	ncnnLib.Call(_func_ncnn_layer_set_support_packing_, []interface{}{&l.c, &b})
}
func (l *Layer) SetSupportBf16Storage(enable bool) {
	b := c.CBool(enable)
	ncnnLib.Call(_func_ncnn_layer_set_support_bf16_storage_, []interface{}{&l.c, &b})
}
func (l *Layer) SetSupportFp16Storage(enable bool) {
	b := c.CBool(enable)
	ncnnLib.Call(_func_ncnn_layer_set_support_fp16_storage_, []interface{}{&l.c, &b})
}
func (l *Layer) SetSupportImageStorage(enable bool) {
	b := c.CBool(enable)
	ncnnLib.Call(_func_ncnn_layer_set_support_image_storage_, []interface{}{&l.c, &b})
}
func (l *Layer) GetBottomCount() int32 {
	return ncnnLib.Call(_func_ncnn_layer_get_bottom_count_, []interface{}{&l.c}).I32Free()
}
func (l *Layer) GetBottom(i int32) int32 {
	return ncnnLib.Call(_func_ncnn_layer_get_bottom_, []interface{}{&l.c, &i}).I32Free()
}
func (l *Layer) GetTopCount() int32 {
	return ncnnLib.Call(_func_ncnn_layer_get_top_count_, []interface{}{&l.c}).I32Free()
}
func (l *Layer) GetTop(i int32) int32 {
	return ncnnLib.Call(_func_ncnn_layer_get_top_, []interface{}{&l.c, &i}).I32Free()
}
func (l *Layer) GetBottomShape(i int32) (dims int32, w int32, h int32, ch int32) {
	dims, w, h, ch = int32(0), int32(1), int32(1), int32(1)
	_dims, _w, _h, _ch := &dims, &w, &h, &ch
	ncnnLib.Call(_func_ncnn_blob_get_bottom_shape_, []interface{}{&l.c, &i, &_dims, &_w, &_h, &_ch})
	return
}
func (l *Layer) GetTopShape(i int32) (dims int32, w int32, h int32, ch int32) {
	dims, w, h, ch = int32(0), int32(1), int32(1), int32(1)
	_dims, _w, _h, _ch := &dims, &w, &h, &ch
	ncnnLib.Call(_func_ncnn_blob_get_top_shape_, []interface{}{&l.c, &i, &_dims, &_w, &_h, &_ch})
	return
}

type _net_custom_layer struct {
	name      unsafe.Pointer
	creator   *c.Callback
	destroyer *c.Callback
	data      unsafe.Pointer
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
	c                 unsafe.Pointer
	_custom_layer     map[string]*_net_custom_layer
	_custom_layer_idx map[int32]*_net_custom_layer
}

func CreateNet() *Net {
	return &Net{c: ncnnLib.Call(_func_ncnn_net_create_, nil).PtrFree()}
}
func (net *Net) Destroy() {
	ncnnLib.Call(_func_ncnn_net_destroy_, []interface{}{&net.c})
	for i := range net._custom_layer {
		net._custom_layer[i].free()
	}
	for i := range net._custom_layer_idx {
		net._custom_layer_idx[i].free()
	}
}
func (net *Net) GetOption() *Option {
	return &Option{c: ncnnLib.Call(_func_ncnn_net_get_option_, []interface{}{&net.c}).PtrFree()}
}
func (net *Net) SetOption(opt *Option) {
	ncnnLib.Call(_func_ncnn_net_set_option_, []interface{}{&net.c, &opt.c})
}

var creatorCallbackPrototype = c.DefineCallbackPrototype(c.AbiDefault, c.Pointer, []c.Type{c.Pointer})
var destroyerCallbackPrototype = c.DefineCallbackPrototype(c.AbiDefault, c.Void, []c.Type{c.Pointer, c.Pointer})

func (net *Net) RegisterCustomLayerByType(typ string, creator func() *Layer, destroyer func(*Layer)) {
	if net._custom_layer == nil {
		net._custom_layer = make(map[string]*_net_custom_layer)
	}

	_, ok := net._custom_layer[typ]
	if ok {
		net._custom_layer[typ].free()
	}
	t := c.CStr(typ)
	create := creatorCallbackPrototype.CreateCallback(func(args []*c.Value, ret *c.Value) {
		l := creator()
		ret.SetPtr(l.c)
	})

	dest := destroyerCallbackPrototype.CreateCallback(func(args []*c.Value, ret *c.Value) {
		destroyer(&Layer{c: args[0].Ptr()})
	})

	cc, d, a := create.CFuncPtr(), dest.CFuncPtr(), usf.Malloc(8)
	usf.Memset(a, 0, 8)
	ncl := &_net_custom_layer{
		name:      t,
		creator:   create,
		destroyer: dest,
		data:      a,
	}
	net._custom_layer[typ] = ncl
	ncnnLib.Call(_func_ncnn_net_register_custom_layer_by_type_, []interface{}{&net.c, &t, &cc, &d, a})
	net._custom_layer[typ].creator = create
	net._custom_layer[typ].destroyer = dest
}
func (net *Net) RegisterCustomLayerByTypeIndex(typIdx int32, creator func() *Layer, destroyer func(*Layer)) {
	if net._custom_layer_idx == nil {
		net._custom_layer_idx = make(map[int32]*_net_custom_layer)
	}
	_, ok := net._custom_layer_idx[typIdx]
	if ok {
		net._custom_layer_idx[typIdx].free()
	}

	create := creatorCallbackPrototype.CreateCallback(func(args []*c.Value, ret *c.Value) {
		l := creator()
		ret.SetPtr(l.c)
	})

	dest := destroyerCallbackPrototype.CreateCallback(func(args []*c.Value, ret *c.Value) {
		destroyer(&Layer{c: args[0].Ptr()})
	})

	cc, d, a := create.CFuncPtr(), dest.CFuncPtr(), usf.Malloc(8)
	usf.Memset(a, 0, 8)
	ncnnLib.Call(_func_ncnn_net_register_custom_layer_by_typeindex_, []interface{}{&net.c, &typIdx, &cc, &d, &a})
	usf.Memset(a, 0, 8)
	ncl := &_net_custom_layer{
		name:      nil,
		creator:   create,
		destroyer: dest,
		data:      a,
	}
	net._custom_layer_idx[typIdx] = ncl
	ncnnLib.Call(_func_ncnn_net_register_custom_layer_by_typeindex_, []interface{}{&net.c, &typIdx, &cc, &d, &a})
	net._custom_layer_idx[typIdx].creator = create
	net._custom_layer_idx[typIdx].destroyer = dest
}
func (net *Net) LoadParam(path string) int32 {
	p := c.CStr(path)
	defer usf.Free(p)
	return ncnnLib.Call(_func_ncnn_net_load_param_, []interface{}{&net.c, &p}).I32Free()
}
func (net *Net) LoadParamBin(path string) int32 {
	p := c.CStr(path)
	defer usf.Free(p)
	return ncnnLib.Call(_func_ncnn_net_load_param_bin_, []interface{}{&net.c, &p}).I32Free()
}
func (net *Net) LoadModel(path string) int32 {
	p := c.CStr(path)
	defer usf.Free(p)
	return ncnnLib.Call(_func_ncnn_net_load_model_, []interface{}{&net.c, &p}).I32Free()
}
func (net *Net) LoadParamMemory(data []byte) int32 {
	d := &data[0]
	return ncnnLib.Call(_func_ncnn_net_load_param_memory_, []interface{}{&net.c, &d}).I32Free()
}
func (net *Net) LoadParamBinMemory(data []byte) int32 {
	d := &data[0]
	return ncnnLib.Call(_func_ncnn_net_load_param_bin_memory_, []interface{}{&net.c, &d}).I32Free()
}
func (net *Net) LoadModelMemory(data []byte) int32 {
	d := &data[0]
	return ncnnLib.Call(_func_ncnn_net_load_model_memory_, []interface{}{&net.c, &d}).I32Free()
}
func (net *Net) Clear() {
	ncnnLib.Call(_func_ncnn_net_clear_, []interface{}{&net.c})
}
func (net *Net) GetInputCount() int32 {
	return ncnnLib.Call(_func_ncnn_net_get_input_count_, []interface{}{&net.c}).I32Free()
}
func (net *Net) GetOutputCount() int32 {
	return ncnnLib.Call(_func_ncnn_net_get_output_count_, []interface{}{&net.c}).I32Free()
}
func (net *Net) GetInputName(i int32) string {
	return ncnnLib.Call(_func_ncnn_net_get_input_name_, []interface{}{&net.c, &i}).StrFree()
}
func (net *Net) GetOutputName(i int32) string {
	return ncnnLib.Call(_func_ncnn_net_get_output_name_, []interface{}{&net.c, &i}).StrFree()
}

type Extractor struct{ c unsafe.Pointer }

func CreateExtractor(net *Net) *Extractor {
	return &Extractor{c: ncnnLib.Call(_func_ncnn_extractor_create_, []interface{}{&net.c}).PtrFree()}
}
func (ex *Extractor) Destroy() {
	ncnnLib.Call(_func_ncnn_extractor_destroy_, []interface{}{&ex.c})
}
func (ex *Extractor) SetOption(opt *Option) {
	ncnnLib.Call(_func_ncnn_extractor_set_option_, []interface{}{&ex.c, &opt.c})
}
func (ex *Extractor) Input(name string, mat *Mat) int32 {
	m := c.CStr(name)
	defer usf.Free(m)
	return ncnnLib.Call(_func_ncnn_extractor_input_, []interface{}{&ex.c, &m, &mat.c}).I32Free()
}
func (ex *Extractor) Extract(name string) *Mat {
	m, n := usf.MallocN(1, 8), c.CStr(name)
	defer usf.Free(m)
	defer usf.Free(n)
	ncnnLib.Call(_func_ncnn_extractor_extract_, []interface{}{&ex.c, &n, &m})
	return &Mat{c: usf.Pop(m)}
}
func (ex *Extractor) InputIndex(idx int32, mat *Mat) int32 {
	return ncnnLib.Call(_func_ncnn_extractor_input_index_, []interface{}{&ex.c, &idx, &mat.c}).I32Free()
}
func (ex *Extractor) ExtractIndex(idx int32) *Mat {
	m := usf.MallocN(1, 8)
	defer usf.Free(m)
	ncnnLib.Call(_func_ncnn_extractor_extract_index_, []interface{}{&ex.c, &idx, &m})
	return &Mat{c: usf.Pop(m)}
}

var (
	_func_ncnn_version_                                  = &c.FuncPrototype{Name: "ncnn_version", OutType: c.Pointer, InTypes: nil}
	_func_ncnn_paramdict_get_type_                       = &c.FuncPrototype{Name: "ncnn_paramdict_get_type", OutType: c.I32, InTypes: []c.Type{c.Pointer, c.I32}}
	_func_ncnn_paramdict_get_int_                        = &c.FuncPrototype{Name: "ncnn_paramdict_get_int", OutType: c.I32, InTypes: []c.Type{c.Pointer, c.I32, c.I32}}
	_func_ncnn_paramdict_get_float_                      = &c.FuncPrototype{Name: "ncnn_paramdict_get_float", OutType: c.F32, InTypes: []c.Type{c.Pointer, c.I32, c.F32}}
	_func_ncnn_net_load_param_memory_                    = &c.FuncPrototype{Name: "ncnn_net_load_param_memory", OutType: c.I32, InTypes: []c.Type{c.Pointer, c.Pointer}}
	_func_ncnn_net_load_param_bin_memory_                = &c.FuncPrototype{Name: "ncnn_net_load_param_bin_memory", OutType: c.I32, InTypes: []c.Type{c.Pointer, c.Pointer}}
	_func_ncnn_net_load_param_bin_                       = &c.FuncPrototype{Name: "ncnn_net_load_param_bin", OutType: c.I32, InTypes: []c.Type{c.Pointer, c.Pointer}}
	_func_ncnn_net_load_param_                           = &c.FuncPrototype{Name: "ncnn_net_load_param", OutType: c.I32, InTypes: []c.Type{c.Pointer, c.Pointer}}
	_func_ncnn_net_load_model_memory_                    = &c.FuncPrototype{Name: "ncnn_net_load_model_memory", OutType: c.I32, InTypes: []c.Type{c.Pointer, c.Pointer}}
	_func_ncnn_net_load_model_                           = &c.FuncPrototype{Name: "ncnn_net_load_model", OutType: c.I32, InTypes: []c.Type{c.Pointer, c.Pointer}}
	_func_ncnn_net_get_output_name_                      = &c.FuncPrototype{Name: "ncnn_net_get_output_name", OutType: c.Pointer, InTypes: []c.Type{c.Pointer, c.I32}}
	_func_ncnn_net_get_output_count_                     = &c.FuncPrototype{Name: "ncnn_net_get_output_count", OutType: c.I32, InTypes: []c.Type{c.Pointer}}
	_func_ncnn_net_get_input_name_                       = &c.FuncPrototype{Name: "ncnn_net_get_input_name", OutType: c.Pointer, InTypes: []c.Type{c.Pointer, c.I32}}
	_func_ncnn_net_get_input_count_                      = &c.FuncPrototype{Name: "ncnn_net_get_input_count", OutType: c.I32, InTypes: []c.Type{c.Pointer}}
	_func_ncnn_mat_get_w_                                = &c.FuncPrototype{Name: "ncnn_mat_get_w", OutType: c.I32, InTypes: []c.Type{c.Pointer}}
	_func_ncnn_mat_get_h_                                = &c.FuncPrototype{Name: "ncnn_mat_get_h", OutType: c.I32, InTypes: []c.Type{c.Pointer}}
	_func_ncnn_mat_get_elemsize_                         = &c.FuncPrototype{Name: "ncnn_mat_get_elemsize", OutType: c.U64, InTypes: []c.Type{c.Pointer}}
	_func_ncnn_mat_get_elempack_                         = &c.FuncPrototype{Name: "ncnn_mat_get_elempack", OutType: c.I32, InTypes: []c.Type{c.Pointer}}
	_func_ncnn_mat_get_dims_                             = &c.FuncPrototype{Name: "ncnn_mat_get_dims", OutType: c.I32, InTypes: []c.Type{c.Pointer}}
	_func_ncnn_mat_get_data_                             = &c.FuncPrototype{Name: "ncnn_mat_get_data", OutType: c.Pointer, InTypes: []c.Type{c.Pointer}}
	_func_ncnn_mat_get_d_                                = &c.FuncPrototype{Name: "ncnn_mat_get_d", OutType: c.I32, InTypes: []c.Type{c.Pointer}}
	_func_ncnn_mat_get_cstep_                            = &c.FuncPrototype{Name: "ncnn_mat_get_cstep", OutType: c.U64, InTypes: []c.Type{c.Pointer}}
	_func_ncnn_mat_get_channel_data_                     = &c.FuncPrototype{Name: "ncnn_mat_get_channel_data", OutType: c.Pointer, InTypes: []c.Type{c.Pointer, c.I32}}
	_func_ncnn_mat_get_c_                                = &c.FuncPrototype{Name: "ncnn_mat_get_c", OutType: c.I32, InTypes: []c.Type{c.Pointer}}
	_func_ncnn_layer_get_typeindex_                      = &c.FuncPrototype{Name: "ncnn_layer_get_typeindex", OutType: c.I32, InTypes: []c.Type{c.Pointer}}
	_func_ncnn_layer_get_type_                           = &c.FuncPrototype{Name: "ncnn_layer_get_type", OutType: c.Pointer, InTypes: []c.Type{c.Pointer}}
	_func_ncnn_layer_get_top_count_                      = &c.FuncPrototype{Name: "ncnn_layer_get_top_count", OutType: c.I32, InTypes: []c.Type{c.Pointer}}
	_func_ncnn_layer_get_top_                            = &c.FuncPrototype{Name: "ncnn_layer_get_top", OutType: c.I32, InTypes: []c.Type{c.Pointer, c.I32}}
	_func_ncnn_layer_get_support_vulkan_                 = &c.FuncPrototype{Name: "ncnn_layer_get_support_vulkan", OutType: c.I32, InTypes: []c.Type{c.Pointer}}
	_func_ncnn_layer_get_support_packing_                = &c.FuncPrototype{Name: "ncnn_layer_get_support_packing", OutType: c.I32, InTypes: []c.Type{c.Pointer}}
	_func_ncnn_layer_get_support_inplace_                = &c.FuncPrototype{Name: "ncnn_layer_get_support_inplace", OutType: c.I32, InTypes: []c.Type{c.Pointer}}
	_func_ncnn_layer_get_support_image_storage_          = &c.FuncPrototype{Name: "ncnn_layer_get_support_image_storage", OutType: c.I32, InTypes: []c.Type{c.Pointer}}
	_func_ncnn_layer_get_support_fp16_storage_           = &c.FuncPrototype{Name: "ncnn_layer_get_support_fp16_storage", OutType: c.I32, InTypes: []c.Type{c.Pointer}}
	_func_ncnn_layer_get_support_bf16_storage_           = &c.FuncPrototype{Name: "ncnn_layer_get_support_bf16_storage", OutType: c.I32, InTypes: []c.Type{c.Pointer}}
	_func_ncnn_layer_get_one_blob_only_                  = &c.FuncPrototype{Name: "ncnn_layer_get_one_blob_only", OutType: c.I32, InTypes: []c.Type{c.Pointer}}
	_func_ncnn_layer_get_name_                           = &c.FuncPrototype{Name: "ncnn_layer_get_name", OutType: c.Pointer, InTypes: []c.Type{c.Pointer}}
	_func_ncnn_layer_get_bottom_count_                   = &c.FuncPrototype{Name: "ncnn_layer_get_bottom_count", OutType: c.I32, InTypes: []c.Type{c.Pointer}}
	_func_ncnn_layer_get_bottom_                         = &c.FuncPrototype{Name: "ncnn_layer_get_bottom_count", OutType: c.I32, InTypes: []c.Type{c.Pointer, c.I32}}
	_func_ncnn_extractor_input_index_                    = &c.FuncPrototype{Name: "ncnn_extractor_input_index", OutType: c.I32, InTypes: []c.Type{c.Pointer, c.I32, c.Pointer}}
	_func_ncnn_extractor_input_                          = &c.FuncPrototype{Name: "ncnn_extractor_input", OutType: c.I32, InTypes: []c.Type{c.Pointer, c.Pointer, c.Pointer}}
	_func_ncnn_blob_get_producer_                        = &c.FuncPrototype{Name: "ncnn_blob_get_producer", OutType: c.I32, InTypes: []c.Type{c.Pointer}}
	_func_ncnn_blob_get_name_                            = &c.FuncPrototype{Name: "ncnn_blob_get_name", OutType: c.Pointer, InTypes: []c.Type{c.Pointer}}
	_func_ncnn_blob_get_consumer_                        = &c.FuncPrototype{Name: "ncnn_blob_get_consumer", OutType: c.I32, InTypes: []c.Type{c.Pointer}}
	_func_ncnn_paramdict_create_                         = &c.FuncPrototype{Name: "ncnn_paramdict_create", OutType: c.Pointer, InTypes: nil}
	_func_ncnn_option_create_                            = &c.FuncPrototype{Name: "ncnn_option_create", OutType: c.Pointer, InTypes: nil}
	_func_ncnn_allocator_create_unlocked_pool_allocator_ = &c.FuncPrototype{Name: "ncnn_allocator_create_unlocked_pool_allocator", OutType: c.Pointer, InTypes: nil}
	_func_ncnn_allocator_create_pool_allocator_          = &c.FuncPrototype{Name: "ncnn_allocator_create_pool_allocator", OutType: c.Pointer, InTypes: nil}
	_func_ncnn_net_get_option_                           = &c.FuncPrototype{Name: "ncnn_net_get_option", OutType: c.Pointer, InTypes: []c.Type{c.Pointer}}
	_func_ncnn_paramdict_set_int_                        = &c.FuncPrototype{Name: "ncnn_paramdict_set_int", OutType: c.Void, InTypes: []c.Type{c.Pointer, c.I32, c.I32}}
	_func_ncnn_paramdict_set_float_                      = &c.FuncPrototype{Name: "ncnn_paramdict_set_float", OutType: c.Void, InTypes: []c.Type{c.Pointer, c.I32, c.F32}}
	_func_ncnn_paramdict_set_array_                      = &c.FuncPrototype{Name: "ncnn_paramdict_set_array", OutType: c.Void, InTypes: []c.Type{c.Pointer, c.I32, c.Pointer}}
	_func_ncnn_paramdict_destroy_                        = &c.FuncPrototype{Name: "ncnn_paramdict_destroy", OutType: c.Void, InTypes: []c.Type{c.Pointer}}
	_func_ncnn_option_set_workspace_allocator_           = &c.FuncPrototype{Name: "ncnn_option_set_workspace_allocator", OutType: c.Void, InTypes: []c.Type{c.Pointer, c.Pointer}}
	_func_ncnn_option_set_use_vulkan_compute_            = &c.FuncPrototype{Name: "ncnn_option_set_use_vulkan_compute", OutType: c.Void, InTypes: []c.Type{c.Pointer, c.I32}}
	_func_ncnn_option_set_use_local_pool_allocator_      = &c.FuncPrototype{Name: "ncnn_option_set_use_local_pool_allocator", OutType: c.Void, InTypes: []c.Type{c.Pointer, c.I32}}
	_func_ncnn_option_set_num_threads_                   = &c.FuncPrototype{Name: "ncnn_option_set_num_threads", OutType: c.Void, InTypes: []c.Type{c.Pointer, c.I32}}
	_func_ncnn_option_set_blob_allocator_                = &c.FuncPrototype{Name: "ncnn_option_set_blob_allocator", OutType: c.Void, InTypes: []c.Type{c.Pointer, c.Pointer}}
	_func_ncnn_option_destroy_                           = &c.FuncPrototype{Name: "ncnn_option_destroy", OutType: c.Void, InTypes: []c.Type{c.Pointer}}
	_func_ncnn_net_set_option_                           = &c.FuncPrototype{Name: "ncnn_net_set_option", OutType: c.Void, InTypes: []c.Type{c.Pointer, c.Pointer}}
	_func_ncnn_net_register_custom_layer_by_typeindex_   = &c.FuncPrototype{Name: "ncnn_net_register_custom_layer_by_typeindex", OutType: c.Void, InTypes: []c.Type{c.Pointer, c.I32, c.Pointer, c.Pointer, c.Pointer}}
	_func_ncnn_net_register_custom_layer_by_type_        = &c.FuncPrototype{Name: "ncnn_net_register_custom_layer_by_type", OutType: c.Void, InTypes: []c.Type{c.Pointer, c.Pointer, c.Pointer, c.Pointer, c.Pointer}}
	_func_ncnn_net_destroy_                              = &c.FuncPrototype{Name: "ncnn_net_destroy", OutType: c.Void, InTypes: []c.Type{c.Pointer}}
	_func_ncnn_net_clear_                                = &c.FuncPrototype{Name: "ncnn_net_clear", OutType: c.Void, InTypes: []c.Type{c.Pointer}}
	_func_ncnn_modelbin_destroy_                         = &c.FuncPrototype{Name: "ncnn_modelbin_destroy", OutType: c.Void, InTypes: []c.Type{c.Pointer}}
	_func_ncnn_mat_substract_mean_normalize_             = &c.FuncPrototype{Name: "ncnn_mat_substract_mean_normalize", OutType: c.Void, InTypes: []c.Type{c.Pointer, c.Pointer, c.Pointer}}
	_func_ncnn_mat_fill_float_                           = &c.FuncPrototype{Name: "ncnn_mat_fill_float", OutType: c.Void, InTypes: []c.Type{c.Pointer, c.F32}}
	_func_ncnn_mat_destroy_                              = &c.FuncPrototype{Name: "ncnn_mat_destroy", OutType: c.Void, InTypes: []c.Type{c.Pointer}}
	_func_ncnn_layer_set_support_vulkan_                 = &c.FuncPrototype{Name: "ncnn_layer_set_support_vulkan", OutType: c.Void, InTypes: []c.Type{c.Pointer, c.I32}}
	_func_ncnn_layer_set_support_packing_                = &c.FuncPrototype{Name: "ncnn_layer_set_support_packing", OutType: c.Void, InTypes: []c.Type{c.Pointer, c.I32}}
	_func_ncnn_layer_set_support_inplace_                = &c.FuncPrototype{Name: "ncnn_layer_set_support_inplace", OutType: c.Void, InTypes: []c.Type{c.Pointer, c.I32}}
	_func_ncnn_layer_set_support_image_storage_          = &c.FuncPrototype{Name: "ncnn_layer_set_support_image_storage", OutType: c.Void, InTypes: []c.Type{c.Pointer, c.I32}}
	_func_ncnn_layer_set_support_fp16_storage_           = &c.FuncPrototype{Name: "ncnn_layer_set_support_fp16_storage", OutType: c.Void, InTypes: []c.Type{c.Pointer, c.I32}}
	_func_ncnn_layer_set_support_bf16_storage_           = &c.FuncPrototype{Name: "ncnn_layer_set_support_bf16_storage", OutType: c.Void, InTypes: []c.Type{c.Pointer, c.I32}}
	_func_ncnn_layer_set_one_blob_only_                  = &c.FuncPrototype{Name: "ncnn_layer_set_one_blob_only", OutType: c.Void, InTypes: []c.Type{c.Pointer, c.I32}}
	_func_ncnn_layer_destroy_                            = &c.FuncPrototype{Name: "ncnn_layer_destroy", OutType: c.Void, InTypes: []c.Type{c.Pointer}}
	_func_ncnn_flatten_                                  = &c.FuncPrototype{Name: "ncnn_flatten", OutType: c.Void, InTypes: []c.Type{c.Pointer, c.Pointer, c.Pointer}}
	_func_ncnn_extractor_set_option_                     = &c.FuncPrototype{Name: "ncnn_extractor_set_option", OutType: c.Void, InTypes: []c.Type{c.Pointer, c.Pointer}}
	_func_ncnn_extractor_extract_index_                  = &c.FuncPrototype{Name: "ncnn_extractor_extract_index", OutType: c.I32, InTypes: []c.Type{c.Pointer, c.I32, c.Pointer}}
	_func_ncnn_extractor_extract_                        = &c.FuncPrototype{Name: "ncnn_extractor_extract", OutType: c.I32, InTypes: []c.Type{c.Pointer, c.Pointer, c.Pointer}}
	_func_ncnn_extractor_destroy_                        = &c.FuncPrototype{Name: "ncnn_extractor_destroy", OutType: c.Void, InTypes: []c.Type{c.Pointer}}
	_func_ncnn_datareader_destroy_                       = &c.FuncPrototype{Name: "ncnn_datareader_destroy", OutType: c.Void, InTypes: []c.Type{c.Pointer}}
	_func_ncnn_convert_packing_                          = &c.FuncPrototype{Name: "ncnn_convert_packing", OutType: c.Void, InTypes: []c.Type{c.Pointer, c.Pointer, c.I32, c.Pointer}}
	_func_ncnn_blob_get_top_shape_                       = &c.FuncPrototype{Name: "ncnn_blob_get_top_shape", OutType: c.Void, InTypes: []c.Type{c.Pointer, c.I32}}
	_func_ncnn_blob_get_bottom_shape_                    = &c.FuncPrototype{Name: "ncnn_blob_get_bottom_shape", OutType: c.Void, InTypes: []c.Type{c.Pointer, c.I32}}
	_func_ncnn_allocator_destroy_                        = &c.FuncPrototype{Name: "ncnn_allocator_destroy", OutType: c.Void, InTypes: []c.Type{c.Pointer}}
	_func_ncnn_option_get_use_vulkan_compute_            = &c.FuncPrototype{Name: "ncnn_option_get_use_vulkan_compute", OutType: c.I32, InTypes: []c.Type{c.Pointer}}
	_func_ncnn_option_get_use_local_pool_allocator_      = &c.FuncPrototype{Name: "ncnn_option_get_use_local_pool_allocator", OutType: c.I32, InTypes: []c.Type{c.Pointer}}
	_func_ncnn_option_get_num_threads_                   = &c.FuncPrototype{Name: "ncnn_option_get_num_threads", OutType: c.I32, InTypes: []c.Type{c.Pointer}}
	_func_ncnn_net_create_                               = &c.FuncPrototype{Name: "ncnn_net_create", OutType: c.Pointer, InTypes: nil}
	_func_ncnn_paramdict_get_array_                      = &c.FuncPrototype{Name: "ncnn_paramdict_get_array", OutType: c.Pointer, InTypes: []c.Type{c.Pointer, c.I32, c.Pointer}}
	_func_ncnn_modelbin_create_from_mat_array_           = &c.FuncPrototype{Name: "ncnn_modelbin_create_from_mat_array", OutType: c.Pointer, InTypes: []c.Type{c.Pointer, c.I32}}
	_func_ncnn_modelbin_create_from_datareader_          = &c.FuncPrototype{Name: "ncnn_modelbin_create_from_datareader", OutType: c.Pointer, InTypes: []c.Type{c.Pointer}}
	_func_ncnn_mat_reshape_4d_                           = &c.FuncPrototype{Name: "ncnn_mat_reshape_4d", OutType: c.Pointer, InTypes: []c.Type{c.Pointer, c.I32, c.I32, c.I32, c.I32, c.Pointer}}
	_func_ncnn_mat_reshape_3d_                           = &c.FuncPrototype{Name: "ncnn_mat_reshape_3d", OutType: c.Pointer, InTypes: []c.Type{c.Pointer, c.I32, c.I32, c.I32, c.Pointer}}
	_func_ncnn_mat_reshape_2d_                           = &c.FuncPrototype{Name: "ncnn_mat_reshape_2d", OutType: c.Pointer, InTypes: []c.Type{c.Pointer, c.I32, c.I32, c.Pointer}}
	_func_ncnn_mat_reshape_1d_                           = &c.FuncPrototype{Name: "ncnn_mat_reshape_1d", OutType: c.Pointer, InTypes: []c.Type{c.Pointer, c.I32, c.Pointer}}
	_func_ncnn_mat_from_pixels_roi_resize_               = &c.FuncPrototype{Name: "ncnn_mat_from_pixels_roi_resize", OutType: c.Pointer, InTypes: []c.Type{c.Pointer, c.I32, c.I32, c.I32, c.I32, c.I32, c.I32, c.I32, c.I32, c.I32, c.I32, c.Pointer}}
	_func_ncnn_mat_from_pixels_roi_                      = &c.FuncPrototype{Name: "ncnn_mat_from_pixels_roi", OutType: c.Pointer, InTypes: []c.Type{c.Pointer, c.I32, c.I32, c.I32, c.I32, c.I32, c.I32, c.I32, c.I32, c.Pointer}}
	_func_ncnn_mat_from_pixels_resize_                   = &c.FuncPrototype{Name: "ncnn_mat_from_pixels_resize", OutType: c.Pointer, InTypes: []c.Type{c.Pointer, c.I32, c.I32, c.I32, c.I32, c.I32, c.I32, c.Pointer}}
	_func_ncnn_mat_from_pixels_                          = &c.FuncPrototype{Name: "ncnn_mat_from_pixels", OutType: c.Pointer, InTypes: []c.Type{c.Pointer, c.I32, c.I32, c.I32, c.I32, c.Pointer}}
	_func_ncnn_mat_create_external_4d_elem_              = &c.FuncPrototype{Name: "ncnn_mat_create_external_4d_elem", OutType: c.Pointer, InTypes: []c.Type{c.I32, c.I32, c.I32, c.I32, c.Pointer, c.U64, c.I32, c.Pointer}}
	_func_ncnn_mat_create_external_4d_                   = &c.FuncPrototype{Name: "ncnn_mat_create_external_4d", OutType: c.Pointer, InTypes: []c.Type{c.I32, c.I32, c.I32, c.I32, c.Pointer, c.Pointer}}
	_func_ncnn_mat_create_external_3d_elem_              = &c.FuncPrototype{Name: "ncnn_mat_create_external_3d_elem", OutType: c.Pointer, InTypes: []c.Type{c.I32, c.I32, c.I32, c.Pointer, c.U64, c.I32, c.Pointer}}
	_func_ncnn_mat_create_external_3d_                   = &c.FuncPrototype{Name: "ncnn_mat_create_external_3d", OutType: c.Pointer, InTypes: []c.Type{c.I32, c.I32, c.I32, c.Pointer, c.Pointer}}
	_func_ncnn_mat_create_external_2d_elem_              = &c.FuncPrototype{Name: "ncnn_mat_create_external_2d_elem", OutType: c.Pointer, InTypes: []c.Type{c.I32, c.I32, c.Pointer, c.U64, c.I32, c.Pointer}}
	_func_ncnn_mat_create_external_2d_                   = &c.FuncPrototype{Name: "ncnn_mat_create_external_2d", OutType: c.Pointer, InTypes: []c.Type{c.I32, c.I32, c.Pointer, c.Pointer}}
	_func_ncnn_mat_create_external_1d_elem_              = &c.FuncPrototype{Name: "ncnn_mat_create_external_1d_elem", OutType: c.Pointer, InTypes: []c.Type{c.I32, c.Pointer, c.U64, c.I32, c.Pointer}}
	_func_ncnn_mat_create_external_1d_                   = &c.FuncPrototype{Name: "ncnn_mat_create_external_1d", OutType: c.Pointer, InTypes: []c.Type{c.I32, c.Pointer, c.Pointer}}
	_func_ncnn_mat_create_4d_elem_                       = &c.FuncPrototype{Name: "ncnn_mat_create_4d_elem", OutType: c.Pointer, InTypes: []c.Type{c.I32, c.I32, c.I32, c.I32, c.U64, c.I32, c.Pointer}}
	_func_ncnn_mat_create_4d_                            = &c.FuncPrototype{Name: "ncnn_mat_create_4d", OutType: c.Pointer, InTypes: []c.Type{c.I32, c.I32, c.I32, c.I32, c.Pointer}}
	_func_ncnn_mat_create_3d_elem_                       = &c.FuncPrototype{Name: "ncnn_mat_create_3d_elem", OutType: c.Pointer, InTypes: []c.Type{c.I32, c.I32, c.I32, c.U64, c.I32, c.Pointer}}
	_func_ncnn_mat_create_3d_                            = &c.FuncPrototype{Name: "ncnn_mat_create_3d", OutType: c.Pointer, InTypes: []c.Type{c.I32, c.I32, c.I32, c.Pointer}}
	_func_ncnn_mat_create_2d_elem_                       = &c.FuncPrototype{Name: "ncnn_mat_create_2d_elem", OutType: c.Pointer, InTypes: []c.Type{c.I32, c.I32, c.U64, c.I32, c.Pointer}}
	_func_ncnn_mat_create_2d_                            = &c.FuncPrototype{Name: "ncnn_mat_create_2d", OutType: c.Pointer, InTypes: []c.Type{c.I32, c.I32, c.Pointer}}
	_func_ncnn_mat_create_1d_elem_                       = &c.FuncPrototype{Name: "ncnn_mat_create_1d_elem", OutType: c.Pointer, InTypes: []c.Type{c.I32, c.U64, c.I32, c.Pointer}}
	_func_ncnn_mat_create_1d_                            = &c.FuncPrototype{Name: "ncnn_mat_create_1d", OutType: c.Pointer, InTypes: []c.Type{c.I32, c.Pointer}}
	_func_ncnn_mat_create_                               = &c.FuncPrototype{Name: "ncnn_mat_create", OutType: c.Pointer, InTypes: nil}
	_func_ncnn_mat_clone_                                = &c.FuncPrototype{Name: "ncnn_mat_clone", OutType: c.Pointer, InTypes: []c.Type{c.Pointer, c.Pointer}}
	_func_ncnn_layer_create_by_type_                     = &c.FuncPrototype{Name: "ncnn_layer_create_by_type", OutType: c.Pointer, InTypes: []c.Type{c.Pointer}}
	_func_ncnn_layer_create_by_typeindex_                = &c.FuncPrototype{Name: "ncnn_layer_create_by_typeindex", OutType: c.Pointer, InTypes: []c.Type{c.I32}}
	_func_ncnn_layer_create_                             = &c.FuncPrototype{Name: "ncnn_layer_create", OutType: c.Pointer, InTypes: nil}
	_func_ncnn_extractor_create_                         = &c.FuncPrototype{Name: "ncnn_extractor_create", OutType: c.Pointer, InTypes: []c.Type{c.Pointer}}
	_func_ncnn_datareader_create_from_memory_            = &c.FuncPrototype{Name: "ncnn_datareader_create_from_memory", OutType: c.Pointer, InTypes: []c.Type{c.Pointer}}
	_func_ncnn_datareader_create_                        = &c.FuncPrototype{Name: "ncnn_datareader_create", OutType: c.Pointer, InTypes: nil}
)
