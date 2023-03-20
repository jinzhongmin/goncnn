package ncnn

//
import "C"
import (
	"runtime"
	"unsafe"

	"github.com/jinzhongmin/dlfcn/pkg/dlfcn"
	"github.com/jinzhongmin/goffi/pkg/ffi"
	"github.com/jinzhongmin/usf"
)

var dll *dlfcn.Handle

func Loadlib(dllPath string, mod dlfcn.Mode) {
	var err error
	dll, err = dlfcn.DlOpen(dllPath, mod)
	if err != nil {
		panic(err)
	}
}

func cbool(b bool) C.int {
	if b {
		return 1
	}
	return 0
}
func cstr(s string) (*C.char, unsafe.Pointer) {
	p := C.CString(s)
	return p, unsafe.Pointer(p)
}
func gostr(p unsafe.Pointer) string {
	return C.GoString((*C.char)(p))
}
func call(fn string, outTyp ffi.Type, inTyp []ffi.Type, args []interface{}) unsafe.Pointer {
	fnp, err := dll.Symbol(fn)
	if err != nil {
		panic(err)
	}

	cif, err := ffi.NewCif(ffi.AbiDefault, outTyp, inTyp)
	if err != nil {
		panic(err)
	}
	defer cif.Free()

	r := cif.Call(unsafe.Pointer(fnp), args)
	p := usf.Malloc(1, 8)
	usf.Push(p, (*(*[1]unsafe.Pointer)(r))[0])
	runtime.SetFinalizer(&p, func(_p *unsafe.Pointer) {
		usf.Free(*_p)
	})
	return p
}

func Version() string {
	v := call("ncnn_version", ffi.Pointer, []ffi.Type{}, []interface{}{})
	return gostr(usf.Pop(v))
}

type Allocator struct{ c unsafe.Pointer }

func CreatePoolAllocator() *Allocator {
	p := call("ncnn_allocator_create_pool_allocator",
		ffi.Pointer, []ffi.Type{}, []interface{}{})
	return &Allocator{c: usf.Pop(p)}
}
func CreateUnlockedPoolAllocator() *Allocator {
	p := call("ncnn_allocator_create_unlocked_pool_allocator",
		ffi.Pointer, []ffi.Type{}, []interface{}{})
	return &Allocator{c: usf.Pop(p)}
}
func (alc *Allocator) Destroy() {
	call("ncnn_allocator_destroy", ffi.Void,
		[]ffi.Type{ffi.Pointer}, []interface{}{&alc.c})
}

type Option struct{ c unsafe.Pointer }

func CreateOption() *Option {
	p := call("ncnn_option_create", ffi.Pointer, []ffi.Type{}, []interface{}{})
	return &Option{c: usf.Pop(p)}
}

func (opt *Option) Destroy() {
	call("ncnn_option_destroy", ffi.Void,
		[]ffi.Type{ffi.Pointer}, []interface{}{&opt.c})
}
func (opt *Option) GetNumThreads() int32 {
	n := call("ncnn_option_get_num_threads", ffi.Int32,
		[]ffi.Type{ffi.Pointer}, []interface{}{&opt.c})
	return (*(*[1]int32)(n))[0]
}
func (opt *Option) SetNumThreads(num int32) {
	call("ncnn_option_set_num_threads", ffi.Void,
		[]ffi.Type{ffi.Pointer, ffi.Int32},
		[]interface{}{&opt.c, &num})
}
func (opt *Option) GetUseLocalPoolAllocator() int32 {
	n := call("ncnn_option_get_use_local_pool_allocator", ffi.Int32,
		[]ffi.Type{ffi.Pointer},
		[]interface{}{&opt.c})
	return (*(*[1]int32)(n))[0]
}
func (opt *Option) SetUseLocalPoolAllocator(n int32) {
	call("ncnn_option_set_use_local_pool_allocator", ffi.Void,
		[]ffi.Type{ffi.Pointer, ffi.Int32},
		[]interface{}{&opt.c, &n})
}
func (opt *Option) SetBlobAllocator(alc *Allocator) {
	call("ncnn_option_set_blob_allocator", ffi.Void,
		[]ffi.Type{ffi.Pointer, ffi.Pointer},
		[]interface{}{&opt.c, &alc.c})
}
func (opt *Option) SetWorkspaceAllocator(alc *Allocator) {
	call("ncnn_option_set_workspace_allocator", ffi.Void,
		[]ffi.Type{ffi.Pointer, ffi.Pointer},
		[]interface{}{&opt.c, &alc.c})
}
func (opt *Option) GetUseVulkanCompute() bool {
	n := call("ncnn_option_get_use_vulkan_compute", ffi.Int32,
		[]ffi.Type{ffi.Pointer},
		[]interface{}{&opt.c})
	return (*(*[1]int32)(n))[0] == 1
}
func (opt *Option) SetUseVulkanCompute(enable bool) {
	n := cbool(enable)
	call("ncnn_option_set_use_vulkan_compute", ffi.Void,
		[]ffi.Type{ffi.Pointer, ffi.Int32},
		[]interface{}{&opt.c, &n})
}

type Mat struct {
	c unsafe.Pointer
}

func (mat *Mat) Destroy() {
	call("ncnn_mat_destroy", ffi.Void,
		[]ffi.Type{ffi.Pointer},
		[]interface{}{&mat.c})
}
func CreateMat() *Mat {
	m := call("ncnn_mat_create", ffi.Pointer,
		[]ffi.Type{}, []interface{}{})
	return &Mat{c: usf.Pop(m)}
}
func CreateMat1D(w int32, alc *Allocator) *Mat {
	m := call("ncnn_mat_create_1d", ffi.Pointer,
		[]ffi.Type{ffi.Int32, ffi.Pointer},
		[]interface{}{&w, &alc.c})
	return &Mat{c: usf.Pop(m)}
}
func CreateMat2D(w, h int32, alc *Allocator) *Mat {
	m := call("ncnn_mat_create_2d", ffi.Pointer,
		[]ffi.Type{ffi.Int32, ffi.Int32, ffi.Pointer},
		[]interface{}{&w, &h, &alc.c})
	return &Mat{c: usf.Pop(m)}
}
func CreateMat3D(w, h, c int32, alc *Allocator) *Mat {
	m := call("ncnn_mat_create_3d", ffi.Pointer,
		[]ffi.Type{ffi.Int32, ffi.Int32, ffi.Int32, ffi.Pointer},
		[]interface{}{&w, &h, &c, &alc.c})
	return &Mat{c: usf.Pop(m)}
}
func CreateMat4D(w, h, d, c int32, alc *Allocator) *Mat {
	m := call("ncnn_mat_create_4d", ffi.Pointer,
		[]ffi.Type{ffi.Int32, ffi.Int32, ffi.Int32, ffi.Int32, ffi.Pointer},
		[]interface{}{&w, &h, &d, &c, &alc.c})
	return &Mat{c: usf.Pop(m)}
}
func CreateMatExternal1D(w int32, data unsafe.Pointer, alc *Allocator) *Mat {
	m := call("ncnn_mat_create_external_1d", ffi.Pointer,
		[]ffi.Type{ffi.Int32, ffi.Pointer, ffi.Pointer},
		[]interface{}{&w, &data, &alc.c})
	return &Mat{c: usf.Pop(m)}
}
func CreateMatExternal2D(w, h int32, data unsafe.Pointer, alc *Allocator) *Mat {
	m := call("ncnn_mat_create_external_2d", ffi.Pointer,
		[]ffi.Type{ffi.Int32, ffi.Int32, ffi.Pointer, ffi.Pointer},
		[]interface{}{&w, &h, &data, &alc.c})
	return &Mat{c: usf.Pop(m)}
}

func CreateMatExternal3D(w, h, c int32, data unsafe.Pointer, alc *Allocator) *Mat {
	m := call("ncnn_mat_create_external_3d", ffi.Pointer,
		[]ffi.Type{ffi.Int32, ffi.Int32, ffi.Int32, ffi.Pointer, ffi.Pointer},
		[]interface{}{&w, &h, &c, &data, &alc.c})
	return &Mat{c: usf.Pop(m)}
}
func CreateMatExternal4D(w, h, d, c int32, data unsafe.Pointer, alc *Allocator) *Mat {
	m := call("ncnn_mat_create_external_4d", ffi.Pointer,
		[]ffi.Type{ffi.Int32, ffi.Int32, ffi.Int32, ffi.Int32, ffi.Pointer, ffi.Pointer},
		[]interface{}{&w, &h, &d, &c, &data, &alc.c})
	return &Mat{c: usf.Pop(m)}
}
func CreateMat1DElm(w int32, elmSize uint64, elmPack int32, alc *Allocator) *Mat {
	m := call("ncnn_mat_create_1d_elem", ffi.Pointer,
		[]ffi.Type{ffi.Int32, ffi.Uint64, ffi.Int32, ffi.Pointer},
		[]interface{}{&w, &elmSize, &elmPack, &alc.c})
	return &Mat{c: usf.Pop(m)}
}
func CreateMat2DElm(w, h int32, elmSize uint64, elmPack int32, alc *Allocator) *Mat {
	m := call("ncnn_mat_create_2d_elem", ffi.Pointer,
		[]ffi.Type{ffi.Int32, ffi.Int32, ffi.Uint64, ffi.Int32, ffi.Pointer},
		[]interface{}{&w, &h, &elmSize, &elmPack, &alc.c})
	return &Mat{c: usf.Pop(m)}
}
func CreateMat3DElm(w, h, c int32, elmSize uint64, elmPack int32, alc *Allocator) *Mat {
	m := call("ncnn_mat_create_3d_elem", ffi.Pointer,
		[]ffi.Type{ffi.Int32, ffi.Int32, ffi.Int32, ffi.Uint64, ffi.Int32, ffi.Pointer},
		[]interface{}{&w, &h, &c, &elmSize, &elmPack, &alc.c})
	return &Mat{c: usf.Pop(m)}
}
func CreateMat4DElm(w, h, d, c int32, elmSize uint64, elmPack int32, alc *Allocator) *Mat {
	m := call("ncnn_mat_create_4d_elem", ffi.Pointer,
		[]ffi.Type{ffi.Int32, ffi.Int32, ffi.Int32, ffi.Int32, ffi.Uint64, ffi.Int32, ffi.Pointer},
		[]interface{}{&w, &h, &d, &c, &elmSize, &elmPack, &alc.c})
	return &Mat{c: usf.Pop(m)}
}
func CreateMatExternal1DElm(w int32, data unsafe.Pointer,
	elmSize uint64, elmPack int32, alc *Allocator) *Mat {
	m := call("ncnn_mat_create_external_1d_elem", ffi.Pointer,
		[]ffi.Type{ffi.Int32, ffi.Pointer, ffi.Uint64, ffi.Int32, ffi.Pointer},
		[]interface{}{&w, &data, &elmSize, &elmPack, &alc.c})
	return &Mat{c: usf.Pop(m)}
}
func CreateMatExternal2DElm(w, h int32, data unsafe.Pointer,
	elmSize uint64, elmPack int32, alc *Allocator) *Mat {
	m := call("ncnn_mat_create_external_2d_elem", ffi.Pointer,
		[]ffi.Type{ffi.Int32, ffi.Int32, ffi.Pointer, ffi.Uint64, ffi.Int32, ffi.Pointer},
		[]interface{}{&w, &h, &data, &elmSize, &elmPack, &alc.c})
	return &Mat{c: usf.Pop(m)}
}
func CreateMatExternal3DElm(w, h, c int32, data unsafe.Pointer,
	elmSize uint64, elmPack int32, alc *Allocator) *Mat {
	m := call("ncnn_mat_create_external_3d_elem", ffi.Pointer,
		[]ffi.Type{ffi.Int32, ffi.Int32, ffi.Int32, ffi.Pointer, ffi.Uint64, ffi.Int32, ffi.Pointer},
		[]interface{}{&w, &h, &c, &data, &elmSize, &elmPack, &alc.c})
	return &Mat{c: usf.Pop(m)}
}

func CreateMatExternal4DElm(w, h, d, c int32, data unsafe.Pointer,
	elmSize uint64, elmPack int32, alc *Allocator) *Mat {
	m := call("ncnn_mat_create_external_4d_elem", ffi.Pointer,
		[]ffi.Type{ffi.Int32, ffi.Int32, ffi.Int32, ffi.Int32, ffi.Pointer, ffi.Uint64, ffi.Int32, ffi.Pointer},
		[]interface{}{&w, &h, &d, &c, &data, &elmSize, &elmPack, &alc.c})
	return &Mat{c: usf.Pop(m)}
}
func (mat *Mat) FillFloat(f float32) {
	call("ncnn_mat_fill_float", ffi.Void,
		[]ffi.Type{ffi.Pointer, ffi.Float},
		[]interface{}{&mat.c, &f})
}
func (mat *Mat) Clone(alc *Allocator) *Mat {
	m := call("ncnn_mat_clone", ffi.Pointer,
		[]ffi.Type{ffi.Pointer, ffi.Pointer},
		[]interface{}{&mat.c, &alc.c})
	return &Mat{c: usf.Pop(m)}
}
func (mat *Mat) Reshape1D(w int32, alc *Allocator) *Mat {
	m := call("ncnn_mat_reshape_1d", ffi.Pointer,
		[]ffi.Type{ffi.Pointer, ffi.Int32, ffi.Pointer},
		[]interface{}{&mat.c, &w, &alc.c})
	return &Mat{c: usf.Pop(m)}
}
func (mat *Mat) Reshape2D(w, h int32, alc *Allocator) *Mat {
	m := call("ncnn_mat_reshape_2d", ffi.Pointer,
		[]ffi.Type{ffi.Pointer, ffi.Int32, ffi.Int32, ffi.Pointer},
		[]interface{}{&mat.c, &w, &h, &alc.c})
	return &Mat{c: usf.Pop(m)}
}
func (mat *Mat) Reshape3D(w, h, c int32, alc *Allocator) *Mat {
	m := call("ncnn_mat_reshape_3d", ffi.Pointer,
		[]ffi.Type{ffi.Pointer, ffi.Int32, ffi.Int32, ffi.Int32, ffi.Pointer},
		[]interface{}{&mat.c, &w, &h, &c, &alc.c})
	return &Mat{c: usf.Pop(m)}
}
func (mat *Mat) Reshape4D(w, h, d, c int32, alc *Allocator) *Mat {
	m := call("ncnn_mat_reshape_4d", ffi.Pointer,
		[]ffi.Type{ffi.Pointer, ffi.Int32, ffi.Int32, ffi.Int32, ffi.Int32, ffi.Pointer},
		[]interface{}{&mat.c, &w, &h, &d, &c, &alc.c})
	return &Mat{c: usf.Pop(m)}
}
func (mat *Mat) GetDims() int32 {
	n := call("ncnn_mat_get_dims", ffi.Int32,
		[]ffi.Type{ffi.Pointer},
		[]interface{}{&mat.c})
	return (*(*[1]int32)(n))[0]
}
func (mat *Mat) GetW() int32 {
	n := call("ncnn_mat_get_w", ffi.Int32,
		[]ffi.Type{ffi.Pointer},
		[]interface{}{&mat.c})
	return (*(*[1]int32)(n))[0]
}
func (mat *Mat) GetH() int32 {
	n := call("ncnn_mat_get_h", ffi.Int32,
		[]ffi.Type{ffi.Pointer},
		[]interface{}{&mat.c})
	return (*(*[1]int32)(n))[0]
}
func (mat *Mat) GetD() int32 {
	n := call("ncnn_mat_get_d", ffi.Int32,
		[]ffi.Type{ffi.Pointer},
		[]interface{}{&mat.c})
	return (*(*[1]int32)(n))[0]
}
func (mat *Mat) GetC() int32 {
	n := call("ncnn_mat_get_c", ffi.Int32,
		[]ffi.Type{ffi.Pointer},
		[]interface{}{&mat.c})
	return (*(*[1]int32)(n))[0]
}
func (mat *Mat) GetElemsize() uint64 {
	n := call("ncnn_mat_get_elemsize", ffi.Uint64,
		[]ffi.Type{ffi.Pointer},
		[]interface{}{&mat.c})
	return (*(*[1]uint64)(n))[0]
}
func (mat *Mat) GetElempack() int32 {
	n := call("ncnn_mat_get_elempack", ffi.Int32,
		[]ffi.Type{ffi.Pointer},
		[]interface{}{&mat.c})
	return (*(*[1]int32)(n))[0]
}
func (mat *Mat) GetCstep() uint64 {
	n := call("ncnn_mat_get_cstep", ffi.Uint64,
		[]ffi.Type{ffi.Pointer},
		[]interface{}{&mat.c})
	return (*(*[1]uint64)(n))[0]
}

func (mat *Mat) GetData() unsafe.Pointer {
	n := call("ncnn_mat_get_data", ffi.Pointer,
		[]ffi.Type{ffi.Pointer},
		[]interface{}{&mat.c})
	return usf.Pop(n)
}

func (mat *Mat) GetChannelData(c int32) unsafe.Pointer {
	n := call("ncnn_mat_get_channel_data", ffi.Pointer,
		[]ffi.Type{ffi.Pointer, ffi.Int32},
		[]interface{}{&mat.c, &c})
	return usf.Pop(n)
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
	m := call("ncnn_mat_from_pixels", ffi.Pointer,
		[]ffi.Type{ffi.Pointer, ffi.Int32, ffi.Int32, ffi.Int32, ffi.Int32, ffi.Pointer},
		[]interface{}{&d, &typ, &w, &h, &stride, &alc})
	return &Mat{c: usf.Pop(m)}
}
func CreateMatFromPixelsResize(pixels []byte, typ PixelType, w, h, stride int32, targetW, targetH int32, alc *Allocator) *Mat {
	d := &pixels[0]
	m := call("ncnn_mat_from_pixels_resize", ffi.Pointer,
		[]ffi.Type{ffi.Pointer, ffi.Int32, ffi.Int32, ffi.Int32, ffi.Int32, ffi.Int32, ffi.Int32, ffi.Pointer},
		[]interface{}{&d, &typ, &w, &h, &stride, &targetW, &targetH, &alc})
	return &Mat{c: usf.Pop(m)}
}
func CreateMatFromPixelsRoi(pixels []byte, typ PixelType, w, h, stride int32, roix, roiy, roiw, roih int32, alc *Allocator) *Mat {
	d := &pixels[0]
	m := call("ncnn_mat_from_pixels_roi", ffi.Pointer,
		[]ffi.Type{ffi.Pointer, ffi.Int32, ffi.Int32, ffi.Int32, ffi.Int32, ffi.Int32, ffi.Int32, ffi.Int32, ffi.Int32, ffi.Pointer},
		[]interface{}{&d, &typ, &w, &h, &stride, &roix, &roiy, &roiw, &roih, &alc})
	return &Mat{c: usf.Pop(m)}
}

func CreateMatFromPixelsRoiResize(pixels []byte, typ PixelType, w, h, stride int32, roix, roiy, roiw, roih int32, targetW, targetH int32, alc *Allocator) *Mat {
	d := &pixels[0]
	m := call("ncnn_mat_from_pixels_roi_resize", ffi.Pointer,
		[]ffi.Type{ffi.Pointer, ffi.Int32, ffi.Int32, ffi.Int32, ffi.Int32, ffi.Int32, ffi.Int32, ffi.Int32, ffi.Int32, ffi.Int32, ffi.Int32, ffi.Pointer},
		[]interface{}{&d, &typ, &w, &h, &stride, &roix, &roiy, &roiw, &roih, &targetW, &targetH, &alc})
	return &Mat{c: usf.Pop(m)}
}

// type slice struct {
// 	ptr unsafe.Pointer
// 	len int
// 	cap int
// }

// func (mat *Mat) ToPixels(typ PixelType, stride int32, outData interface{}) {
// 	p := (*(*[2]unsafe.Pointer)(unsafe.Pointer(&outData)))[1]
// 	p = (*(*slice)((unsafe.Pointer)((*unsafe.Pointer)(p)))).ptr
// 	call("ncnn_mat_to_pixels", ffi.Void,
// 		[]ffi.Type{ffi.Pointer, ffi.Pointer, ffi.Int32, ffi.Int32},
// 		[]interface{}{mat.c, p, typ, stride})
// }
// func (mat *Mat) ToPixelsResize(typ PixelType, targetW, targetH, targetStride int32, outData interface{}) {
// 	p := (*(*[2]unsafe.Pointer)(unsafe.Pointer(&outData)))[1]
// 	p = (*(*slice)((unsafe.Pointer)((*unsafe.Pointer)(p)))).ptr
// 	call("ncnn_mat_to_pixels", ffi.Void,
// 		[]ffi.Type{ffi.Pointer, ffi.Pointer, ffi.Int32, ffi.Int32, ffi.Int32, ffi.Int32},
// 		[]interface{}{mat.c, p, typ, targetW, targetH, targetStride})
// }

func (mat *Mat) SubstractMeanNormalize(mean, norm []float32) {
	m := &mean[0]
	n := &norm[0]
	call("ncnn_mat_substract_mean_normalize", ffi.Void,
		[]ffi.Type{ffi.Pointer, ffi.Pointer, ffi.Pointer},
		[]interface{}{&mat.c, &m, &n})
}
func (mat *Mat) ConvertPacking(elempack int32, opt *Option) *Mat {
	m := usf.Malloc(1, 8)
	defer usf.Free(m)
	call("ncnn_convert_packing", ffi.Void,
		[]ffi.Type{ffi.Pointer, ffi.Pointer, ffi.Int32, ffi.Pointer},
		[]interface{}{&mat.c, &m, &elempack, &opt.c})

	return &Mat{c: usf.Pop(m)}
}

func (mat *Mat) Flatten(opt *Option) *Mat {
	m := usf.Malloc(1, 8)
	defer usf.Free(m)
	call("ncnn_flatten", ffi.Void,
		[]ffi.Type{ffi.Pointer, ffi.Pointer, ffi.Pointer},
		[]interface{}{&mat.c, &m, &opt.c})

	return &Mat{c: usf.Pop(m)}
}

type Blob struct{ c unsafe.Pointer }

func (blob *Blob) GetName() string {
	v := call("ncnn_blob_get_name", ffi.Pointer,
		[]ffi.Type{ffi.Pointer},
		[]interface{}{&blob.c})
	return C.GoString((*C.char)(usf.Pop(v)))
}
func (blob *Blob) GetProducer() int32 {
	v := call("ncnn_blob_get_producer", ffi.Pointer,
		[]ffi.Type{ffi.Pointer},
		[]interface{}{&blob.c})
	return (*(*[1]int32)(v))[0]
}
func (blob *Blob) GetConsumer() int32 {
	v := call("ncnn_blob_get_consumer", ffi.Pointer,
		[]ffi.Type{ffi.Pointer},
		[]interface{}{&blob.c})
	return (*(*[1]int32)(v))[0]
}

type Paramdict struct{ c unsafe.Pointer }

func CreateParamdict() *Paramdict {
	pd := call("ncnn_paramdict_create", ffi.Pointer, []ffi.Type{}, []interface{}{})
	return &Paramdict{c: usf.Pop(pd)}
}
func (pd *Paramdict) Destroy() {
	call("ncnn_paramdict_destroy", ffi.Void,
		[]ffi.Type{ffi.Pointer},
		[]interface{}{&pd.c})
}
func (pd *Paramdict) GetType(id int32) int32 {
	i := call("ncnn_paramdict_get_type", ffi.Int32,
		[]ffi.Type{ffi.Pointer, ffi.Int32},
		[]interface{}{&pd.c, &id})
	return (*(*[1]int32)(i))[0]
}
func (pd *Paramdict) GetInt32(id int32, def int32) int32 {
	i := call("ncnn_paramdict_get_int", ffi.Int32,
		[]ffi.Type{ffi.Pointer, ffi.Int32, ffi.Int32},
		[]interface{}{&pd.c, &id, &def})
	return (*(*[1]int32)(i))[0]
}
func (pd *Paramdict) GetFloat32(id int32, def float32) float32 {
	f := call("ncnn_paramdict_get_float", ffi.Float,
		[]ffi.Type{ffi.Pointer, ffi.Int32, ffi.Float},
		[]interface{}{&pd.c, &id, &def})
	return (*(*[1]float32)(f))[0]
}
func (pd *Paramdict) GetArray(id int32, def *Mat) *Mat {
	m := call("ncnn_paramdict_get_array", ffi.Pointer,
		[]ffi.Type{ffi.Pointer, ffi.Int32, ffi.Pointer},
		[]interface{}{&pd.c, &id, &def})
	return &Mat{c: usf.Pop(m)}
}
func (pd *Paramdict) SetInt32(id int32, i int32) {
	call("ncnn_paramdict_set_int", ffi.Void,
		[]ffi.Type{ffi.Pointer, ffi.Int32, ffi.Int32},
		[]interface{}{&pd.c, &id, &i})
}
func (pd *Paramdict) SetFloat32(id int32, f float32) {
	call("ncnn_paramdict_set_float", ffi.Void,
		[]ffi.Type{ffi.Pointer, ffi.Int32, ffi.Float},
		[]interface{}{&pd.c, &id, &f})
}
func (pd *Paramdict) SetArray(id int32, m *Mat) {
	call("ncnn_paramdict_set_array", ffi.Void,
		[]ffi.Type{ffi.Pointer, ffi.Int32, ffi.Pointer},
		[]interface{}{&pd.c, &id, &m})
}

type DataReader struct{ c unsafe.Pointer }

func CreateDataReader() *DataReader {
	d := call("ncnn_datareader_create", ffi.Pointer, []ffi.Type{}, []interface{}{})
	return &DataReader{c: usf.Pop(d)}
}
func CreateDataReaderFromMemory(mem []byte) *DataReader {
	p := unsafe.Pointer(&mem[0])
	pp := usf.Malloc(1, 8)
	usf.Push(pp, p)
	defer usf.Free(pp)
	d := call("ncnn_datareader_create_from_memory", ffi.Pointer,
		[]ffi.Type{ffi.Pointer},
		[]interface{}{&pp})
	return &DataReader{c: usf.Pop(d)}
}
func (d *DataReader) Destroy() {
	call("ncnn_datareader_destroy", ffi.Void,
		[]ffi.Type{ffi.Pointer},
		[]interface{}{&d.c})
}

type ModelBin struct{ c unsafe.Pointer }

func CreateModelBinFromDataReader(dr *DataReader) *ModelBin {
	m := call("ncnn_modelbin_create_from_datareader", ffi.Pointer,
		[]ffi.Type{ffi.Pointer},
		[]interface{}{&dr.c})
	return &ModelBin{c: usf.Pop(m)}
}
func CreateModelBinFromMatArray(mats []*Mat) *ModelBin {
	l := uint32(len(mats))
	ms := usf.Malloc(uint64(l), 8)
	defer usf.Free(ms)
	m := call("ncnn_modelbin_create_from_mat_array", ffi.Pointer,
		[]ffi.Type{ffi.Pointer, ffi.Int32},
		[]interface{}{&ms, &l})
	return &ModelBin{c: usf.Pop(m)}
}
func (mb *ModelBin) Destroy() {
	call("ncnn_modelbin_destroy", ffi.Void,
		[]ffi.Type{ffi.Pointer},
		[]interface{}{&mb.c})
}

type Layer struct{ c unsafe.Pointer }

func CreateLayer() *Layer {
	l := call("ncnn_layer_create", ffi.Pointer, []ffi.Type{}, []interface{}{})
	return &Layer{c: usf.Pop(l)}
}
func CreateLayerByTypeindex(typIdx int32) *Layer {
	l := call("ncnn_layer_create_by_typeindex", ffi.Pointer,
		[]ffi.Type{ffi.Int32},
		[]interface{}{&typIdx})
	return &Layer{c: usf.Pop(l)}
}
func CreateLayerByType(typ string) *Layer {
	t, p := cstr(typ)
	defer usf.Free(p)
	l := call("ncnn_layer_create_by_typeindex", ffi.Pointer,
		[]ffi.Type{ffi.Pointer},
		[]interface{}{&t})
	return &Layer{c: usf.Pop(l)}
}
func (l *Layer) Destroy() {
	call("ncnn_layer_destroy", ffi.Void,
		[]ffi.Type{ffi.Pointer},
		[]interface{}{&l.c})
}
func (l *Layer) GetName() string {
	n := call("ncnn_layer_get_name", ffi.Pointer,
		[]ffi.Type{ffi.Pointer},
		[]interface{}{&l.c})
	return gostr(usf.Pop(n))
}
func (l *Layer) GetTypeIndex() int32 {
	n := call("ncnn_layer_get_typeindex", ffi.Int32,
		[]ffi.Type{ffi.Pointer},
		[]interface{}{&l.c})
	return (*(*[1]int32)(n))[0]
}
func (l *Layer) GetType() string {
	n := call("ncnn_layer_get_type", ffi.Pointer,
		[]ffi.Type{ffi.Pointer},
		[]interface{}{&l.c})
	return gostr(usf.Pop(n))
}

func (l *Layer) GetOneBlobOnly() bool {
	n := call("ncnn_layer_get_one_blob_only", ffi.Int32,
		[]ffi.Type{ffi.Pointer},
		[]interface{}{&l.c})
	return (*(*[1]int32)(n))[0] == 1
}
func (l *Layer) GetSupportInplace() bool {
	n := call("ncnn_layer_get_support_inplace", ffi.Int32,
		[]ffi.Type{ffi.Pointer},
		[]interface{}{&l.c})
	return (*(*[1]int32)(n))[0] == 1
}
func (l *Layer) GetSupportVulkan() bool {
	n := call("ncnn_layer_get_support_vulkan", ffi.Int32,
		[]ffi.Type{ffi.Pointer},
		[]interface{}{&l.c})
	return (*(*[1]int32)(n))[0] == 1
}
func (l *Layer) GetSupportPacking() bool {
	n := call("ncnn_layer_get_support_packing", ffi.Int32,
		[]ffi.Type{ffi.Pointer},
		[]interface{}{&l.c})
	return (*(*[1]int32)(n))[0] == 1
}
func (l *Layer) GetSupportBf16Storage() bool {
	n := call("ncnn_layer_get_support_bf16_storage", ffi.Int32,
		[]ffi.Type{ffi.Pointer},
		[]interface{}{&l.c})
	return (*(*[1]int32)(n))[0] == 1
}
func (l *Layer) GetSupportFp16Storage() bool {
	n := call("ncnn_layer_get_support_fp16_storage", ffi.Int32,
		[]ffi.Type{ffi.Pointer},
		[]interface{}{&l.c})
	return (*(*[1]int32)(n))[0] == 1
}
func (l *Layer) GetSupportImageStorage() bool {
	n := call("ncnn_layer_get_support_image_storage", ffi.Int32,
		[]ffi.Type{ffi.Pointer},
		[]interface{}{&l.c})
	return (*(*[1]int32)(n))[0] == 1
}

func (l *Layer) SetOneBlobOnly(enable bool) {
	b := cbool(enable)
	call("ncnn_layer_set_one_blob_only", ffi.Void,
		[]ffi.Type{ffi.Pointer, ffi.Int32},
		[]interface{}{&l.c, &b})
}
func (l *Layer) SetSupportInplace(enable bool) {
	b := cbool(enable)
	call("ncnn_layer_set_support_inplace", ffi.Void,
		[]ffi.Type{ffi.Pointer, ffi.Int32},
		[]interface{}{&l.c, &b})
}
func (l *Layer) SetSupportVulkan(enable bool) {
	b := cbool(enable)
	call("ncnn_layer_set_support_vulkan", ffi.Void,
		[]ffi.Type{ffi.Pointer, ffi.Int32},
		[]interface{}{&l.c, &b})
}
func (l *Layer) SetSupportPacking(enable bool) {
	b := cbool(enable)
	call("ncnn_layer_set_support_packing", ffi.Void,
		[]ffi.Type{ffi.Pointer, ffi.Int32},
		[]interface{}{&l.c, &b})
}
func (l *Layer) SetSupportBf16Storage(enable bool) {
	b := cbool(enable)
	call("ncnn_layer_set_support_bf16_storage", ffi.Void,
		[]ffi.Type{ffi.Pointer, ffi.Int32},
		[]interface{}{&l.c, &b})
}
func (l *Layer) SetSupportFp16Storage(enable bool) {
	b := cbool(enable)
	call("ncnn_layer_set_support_fp16_storage", ffi.Void,
		[]ffi.Type{ffi.Pointer, ffi.Int32},
		[]interface{}{&l.c, &b})
}
func (l *Layer) SetSupportImageStorage(enable bool) {
	b := cbool(enable)
	call("ncnn_layer_set_support_image_storage", ffi.Void,
		[]ffi.Type{ffi.Pointer, ffi.Int32},
		[]interface{}{&l.c, &b})
}
func (l *Layer) GetBottomCount() int32 {
	n := call("ncnn_layer_get_bottom_count", ffi.Int32,
		[]ffi.Type{ffi.Pointer},
		[]interface{}{&l.c})
	return (*(*[1]int32)(n))[0]
}
func (l *Layer) GetBottom(i int32) int32 {
	n := call("ncnn_layer_get_bottom_count", ffi.Int32,
		[]ffi.Type{ffi.Pointer, ffi.Int32},
		[]interface{}{&l.c, &i})
	return (*(*[1]int32)(n))[0]
}
func (l *Layer) GetTopCount() int32 {
	n := call("ncnn_layer_get_top_count", ffi.Int32,
		[]ffi.Type{ffi.Pointer},
		[]interface{}{&l.c})
	return (*(*[1]int32)(n))[0]
}
func (l *Layer) GetTop(i int32) int32 {
	n := call("ncnn_layer_get_top", ffi.Int32,
		[]ffi.Type{ffi.Pointer, ffi.Int32},
		[]interface{}{&l.c, &i})
	return (*(*[1]int32)(n))[0]
}
func (l *Layer) GetBottomShape(i int32) (dims int32, w int32, h int32, c int32) {
	dims = int32(0)
	w = int32(1)
	h = int32(1)
	c = int32(1)

	_dims := &dims
	_w := &w
	_h := &h
	_c := &c
	call("ncnn_blob_get_bottom_shape", ffi.Void,
		[]ffi.Type{ffi.Pointer, ffi.Int32},
		[]interface{}{&l.c, &i, &_dims, &_w, &_h, &_c})
	return
}
func (l *Layer) GetTopShape(i int32) (dims int32, w int32, h int32, c int32) {
	dims = int32(0)
	w = int32(1)
	h = int32(1)
	c = int32(1)

	_dims := &dims
	_w := &w
	_h := &h
	_c := &c
	call("ncnn_blob_get_top_shape", ffi.Void,
		[]ffi.Type{ffi.Pointer, ffi.Int32},
		[]interface{}{&l.c, &i, &_dims, &_w, &_h, &_c})
	return
}

type Net struct{ c unsafe.Pointer }

func CreateNet() *Net {
	n := call("ncnn_net_create", ffi.Pointer, []ffi.Type{}, []interface{}{})
	return &Net{c: usf.Pop(n)}
}
func (net *Net) Destroy() {
	call("ncnn_net_destroy", ffi.Void,
		[]ffi.Type{ffi.Pointer},
		[]interface{}{&net.c})
}
func (net *Net) GetOption() *Option {
	o := call("ncnn_net_get_option", ffi.Pointer,
		[]ffi.Type{ffi.Pointer},
		[]interface{}{&net.c})
	return &Option{c: usf.Pop(o)}
}
func (net *Net) SetOption(opt *Option) {
	call("ncnn_net_set_option", ffi.Void,
		[]ffi.Type{ffi.Pointer, ffi.Pointer},
		[]interface{}{&net.c, &opt.c})
}

// 不知道怎么使用
func (net *Net) RegisterCustomLayerByType(typ string, creator func() *Layer, destroyer func(*Layer)) {
	t, p := cstr(typ)
	defer usf.Free(p)
	creat := ffi.NewClosure(ffi.ClosureConf{
		Abi:    ffi.AbiDefault,
		Output: ffi.Pointer,
		Inputs: []ffi.Type{},
	}, func(cp *ffi.ClosureParams) {
		l := creator()
		usf.Push(cp.Return, l.c)
	}, []interface{}{})

	dest := ffi.NewClosure(ffi.ClosureConf{
		Abi:    ffi.AbiDefault,
		Output: ffi.Void,
		Inputs: []ffi.Type{ffi.Pointer, ffi.Pointer},
	}, func(cp *ffi.ClosureParams) {
		l := usf.Pop(cp.Args[0])
		destroyer(&Layer{c: l})
	}, []interface{}{})

	c := creat.Cfn()
	d := dest.Cfn()
	a := usf.MallocOf(1, 8)
	usf.Memset(a, 0, 8)
	call("ncnn_net_register_custom_layer_by_type", ffi.Void,
		[]ffi.Type{ffi.Pointer, ffi.Pointer, ffi.Pointer, ffi.Pointer, ffi.Pointer},
		[]interface{}{&net.c, &t, &c, &d, &a})
}
func (net *Net) LoadParam(path string) int32 {
	c, p := cstr(path)
	defer usf.Free(p)
	n := call("ncnn_net_load_param", ffi.Int32,
		[]ffi.Type{ffi.Pointer, ffi.Pointer},
		[]interface{}{&net.c, &c})
	return (*(*[1]int32)(n))[0]
}
func (net *Net) LoadParamBin(path string) int32 {
	c, p := cstr(path)
	defer usf.Free(p)
	n := call("ncnn_net_load_param_bin", ffi.Int32,
		[]ffi.Type{ffi.Pointer, ffi.Pointer},
		[]interface{}{&net.c, &c})
	return (*(*[1]int32)(n))[0]
}
func (net *Net) LoadModel(path string) int32 {
	c, p := cstr(path)
	defer usf.Free(p)
	n := call("ncnn_net_load_model", ffi.Int32,
		[]ffi.Type{ffi.Pointer, ffi.Pointer},
		[]interface{}{&net.c, &c})
	return (*(*[1]int32)(n))[0]
}
func (net *Net) LoadParamMemory(data []byte) int32 {
	d := &data[0]
	n := call("ncnn_net_load_param_memory", ffi.Int32,
		[]ffi.Type{ffi.Pointer, ffi.Pointer},
		[]interface{}{&net.c, &d})
	return (*(*[1]int32)(n))[0]
}
func (net *Net) LoadParamBinMemory(data []byte) int32 {
	d := &data[0]
	n := call("ncnn_net_load_param_bin_memory", ffi.Int32,
		[]ffi.Type{ffi.Pointer, ffi.Pointer},
		[]interface{}{&net.c, &d})
	return (*(*[1]int32)(n))[0]
}
func (net *Net) LoadModelMemory(data []byte) int32 {
	d := &data[0]
	n := call("ncnn_net_load_model_memory", ffi.Int32,
		[]ffi.Type{ffi.Pointer, ffi.Pointer},
		[]interface{}{&net.c, &d})
	return (*(*[1]int32)(n))[0]
}
func (net *Net) Clear() {
	call("ncnn_net_clear", ffi.Void,
		[]ffi.Type{ffi.Pointer},
		[]interface{}{&net.c})
}
func (net *Net) GetInputCount() int32 {
	n := call("ncnn_net_get_input_count", ffi.Int32,
		[]ffi.Type{ffi.Pointer},
		[]interface{}{&net.c})
	return (*(*[1]int32)(n))[0]
}
func (net *Net) GetOutputCount() int32 {
	n := call("ncnn_net_get_output_count", ffi.Int32,
		[]ffi.Type{ffi.Pointer},
		[]interface{}{&net.c})
	return (*(*[1]int32)(n))[0]
}
func (net *Net) GetInputName(i int32) string {
	n := call("ncnn_net_get_input_name", ffi.Pointer,
		[]ffi.Type{ffi.Pointer, ffi.Int32},
		[]interface{}{&net.c, &i})
	return gostr(usf.Pop(n))
}
func (net *Net) GetOutputName(i int32) string {
	n := call("ncnn_net_get_output_name", ffi.Pointer,
		[]ffi.Type{ffi.Pointer, ffi.Int32},
		[]interface{}{&net.c, &i})
	return gostr(usf.Pop(n))
}

type Extractor struct{ c unsafe.Pointer }

func CreateExtractor(net *Net) *Extractor {
	e := call("ncnn_extractor_create", ffi.Pointer,
		[]ffi.Type{ffi.Pointer},
		[]interface{}{&net.c})
	return &Extractor{c: usf.Pop(e)}
}
func (ex *Extractor) Destroy() {
	call("ncnn_extractor_destroy", ffi.Void,
		[]ffi.Type{ffi.Pointer},
		[]interface{}{&ex.c})
}
func (ex *Extractor) SetOption(opt *Option) {
	call("ncnn_extractor_set_option", ffi.Void,
		[]ffi.Type{ffi.Pointer, ffi.Pointer},
		[]interface{}{&ex.c, &opt.c})
}
func (ex *Extractor) Input(name string, mat *Mat) int32 {
	m, p := cstr(name)
	defer usf.Free(p)
	i := call("ncnn_extractor_input", ffi.Int32,
		[]ffi.Type{ffi.Pointer, ffi.Pointer, ffi.Pointer},
		[]interface{}{&ex.c, &m, &mat.c})
	return (*(*[1]int32)(i))[0]
}
func (ex *Extractor) Extract(name string) *Mat {
	m := usf.Malloc(1, 8)
	defer usf.Free(m)

	n, p := cstr(name)
	defer usf.Free(p)

	call("ncnn_extractor_extract", ffi.Int32,
		[]ffi.Type{ffi.Pointer, ffi.Pointer, ffi.Pointer},
		[]interface{}{&ex.c, &n, &m})
	return &Mat{c: usf.Pop(m)}
}

// 不知道怎么使用
func (ex *Extractor) InputIndex(idx int32, mat *Mat) int32 {
	i := call("ncnn_extractor_input_index", ffi.Int32,
		[]ffi.Type{ffi.Pointer, ffi.Int32, ffi.Pointer},
		[]interface{}{&ex.c, &idx, &mat.c})
	return (*(*[1]int32)(i))[0]
}

// 不知道怎么使用
func (ex *Extractor) ExtractIndex(idx int32) *Mat {
	m := usf.Malloc(1, 8)
	defer usf.Free(m)

	call("ncnn_extractor_extract_index", ffi.Int32,
		[]ffi.Type{ffi.Pointer, ffi.Int32, ffi.Pointer},
		[]interface{}{&ex.c, &idx, &m})
	return &Mat{c: usf.Pop(m)}
}
