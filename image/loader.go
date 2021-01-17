package image

// #include "../stb_image/config.h"
// #include "../stb_image/stb_image.h"
// #cgo LDFLAGS: -lm
import "C"
import (
	"errors"
	"unsafe"
)

// Load is to load an image file from a specified path
func Load(path string) (*uint8, int32, int32, int32, error) {
	var x, y, comp int32
	dataPtr := (*uint8)(C.stbi_load(C.CString(path),
		(*C.int)(unsafe.Pointer(&x)),
		(*C.int)(unsafe.Pointer(&y)),
		(*C.int)(unsafe.Pointer(&comp)),
		0))
	if dataPtr == nil {
		return nil, 0, 0, 0, errors.New("Load image failed: " + path)
	}
	return dataPtr, x, y, comp, nil
}

// Free is to free an image data
func Free(dataPtr *uint8) {
	if dataPtr != nil {
		C.stbi_image_free(unsafe.Pointer(dataPtr))
	}
}
