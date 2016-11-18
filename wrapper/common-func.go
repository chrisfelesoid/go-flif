// Copyright 2016 chrisfelesoid. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package wrapper

/*
#include <stdlib.h>
#include <flif.h>
*/
import "C"
import "unsafe"

type flifImage struct {
	img *C.FLIF_IMAGE
}

var flifCreateImage = func(width, height int) *flifImage {
	p := C.flif_create_image(C.uint32_t(width), C.uint32_t(height))
	if p == nil {
		return nil
	}
	return &flifImage{p}
}

var flifCreateImageHDR = func(width, height int) *flifImage {
	p := C.flif_create_image_HDR(C.uint32_t(width), C.uint32_t(height))
	if p == nil {
		return nil
	}
	return &flifImage{p}
}

var flifImportImageRGBA = func(width, height int, data []byte, stride int) *flifImage {
	p := C.flif_import_image_RGBA(C.uint32_t(width), C.uint32_t(height), unsafe.Pointer(&data[0]), C.uint32_t(stride))
	if p == nil {
		return nil
	}
	return &flifImage{p}
}

var flifImportImageRGB = func(width, height int, data []byte, stride int) *flifImage {
	p := C.flif_import_image_RGB(C.uint32_t(width), C.uint32_t(height), unsafe.Pointer(&data[0]), C.uint32_t(stride))
	if p == nil {
		return nil
	}
	return &flifImage{p}
}

var flifImportImageGRAY = func(width, height int, data []byte, stride int) *flifImage {
	p := C.flif_import_image_GRAY(C.uint32_t(width), C.uint32_t(height), unsafe.Pointer(&data[0]), C.uint32_t(stride))
	if p == nil {
		return nil
	}
	return &flifImage{p}
}

var flifDestroyImage = func(image *flifImage) {
	C.flif_destroy_image(image.img)
}

var flifImageGetWidth = func(image *flifImage) int {
	return int(C.flif_image_get_width(image.img))
}

var flifImageGetHeight = func(image *flifImage) int {
	return int(C.flif_image_get_height(image.img))
}

var flifImageGetNbChannels = func(image *flifImage) int {
	return int(C.flif_image_get_nb_channels(image.img))
}

var flifImageGetDepth = func(image *flifImage) int {
	return int(C.flif_image_get_depth(image.img))
}

var flifImageGetFrameDelay = func(image *flifImage) int {
	return int(C.flif_image_get_frame_delay(image.img))
}

var flifImageSetFrameDelay = func(image *flifImage, delay int) {
	C.flif_image_set_frame_delay(image.img, C.uint32_t(delay))
}

var flifImageSetMetadata = func(image *flifImage, chunkname string, data []byte) {
	cchunk := C.CString(chunkname)
	defer C.free(unsafe.Pointer(cchunk))
	C.flif_image_set_metadata(image.img, cchunk, (*C.uchar)(unsafe.Pointer(&data[0])), (C.size_t)(len(data)))
}

var flifImageGetMetadata = func(image *flifImage, chunkname string, data *[]byte) {
	*data = nil

	cchunk := C.CString(chunkname)
	defer C.free(unsafe.Pointer(cchunk))
	var p unsafe.Pointer
	var num C.size_t

	C.flif_image_get_metadata(image.img, cchunk, (**C.uchar)(unsafe.Pointer(&p)), (*C.size_t)(unsafe.Pointer(&num)))
	if p == nil {
		return
	}
	defer flifImageFreeMetadata(image, p)
	*data = C.GoBytes(p, C.int(num))
}

var flifImageFreeMetadata = func(image *flifImage, data unsafe.Pointer) {
	C.flif_image_free_metadata(image.img, (*C.uchar)(unsafe.Pointer(data)))
}

var flifImageWriteRowRGBA8 = func(image *flifImage, row int, data []byte) {
	C.flif_image_write_row_RGBA8(image.img, C.uint32_t(row), unsafe.Pointer(&data[0]), C.size_t(len(data)))
}

var flifImageReadRowRGBA8 = func(image *flifImage, row int) []byte {
	w := flifImageGetWidth(image)
	b := make([]byte, C.sizeof_uint8_t*4*w)
	C.flif_image_read_row_RGBA8(image.img, C.uint32_t(row), unsafe.Pointer(&b[0]), C.size_t(len(b)))
	return b
}

var flifImageWriteRowRGBA16 = func(image *flifImage, row int, data []byte) {
	C.flif_image_write_row_RGBA16(image.img, C.uint32_t(row), unsafe.Pointer(&data[0]), C.size_t(len(data)))
}

var flifImageReadRowRGBA16 = func(image *flifImage, row int) []byte {
	w := flifImageGetWidth(image)
	b := make([]byte, C.sizeof_uint16_t*4*w)
	C.flif_image_read_row_RGBA16(image.img, C.uint32_t(row), unsafe.Pointer(&b[0]), C.size_t(len(b)))
	return b
}

var flifFreeMemory = func(p unsafe.Pointer) {
	C.flif_free_memory(p)
}
