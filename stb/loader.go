package stb

// #include "../stb_image/config.h"
// #include "../stb_image/stb_image.h"
// #cgo LDFLAGS: -lm -O3
import "C"
import (
	"errors"
	"fmt"
	"io/ioutil"
	"unsafe"
)

// Load is to load image file from a specified path
func Load(path string) (*uint8, int, int, int, error) {
	byts, e := ioutil.ReadFile(path)
	if e != nil {
		return nil, 0, 0, 0, fmt.Errorf("read image failed %s, %w", path, e)
	}

	return LoadBytes(byts)
}

// LoadBytes is to load image from []byte
func LoadBytes(data []byte) (*uint8, int, int, int, error) {
	var x, y, comp int32
	dataPtr := (*uint8)(C.stbi_load_from_memory(
		(*C.uint8_t)(unsafe.Pointer(&data[0])),
		C.int(len(data)),
		(*C.int)(unsafe.Pointer(&x)),
		(*C.int)(unsafe.Pointer(&y)),
		(*C.int)(unsafe.Pointer(&comp)),
		0))
	if dataPtr == nil {
		return nil, 0, 0, 0, errors.New("decode image failed")
	}
	return dataPtr, int(x), int(y), int(comp), nil
}

// Free is to free an image data
func Free(dataPtr *uint8) {
	if dataPtr != nil {
		C.stbi_image_free(unsafe.Pointer(dataPtr))
	}
}
