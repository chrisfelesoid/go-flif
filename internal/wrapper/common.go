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

type CflifImage struct {
	img *C.FLIF_IMAGE
}

func CflifCreateImage(width, height int) *CflifImage {
	p := C.flif_create_image(C.uint32_t(width), C.uint32_t(height))
	if p == nil {
		return nil
	}
	return &CflifImage{p}
}

func CflifCreateImageHDR(width, height int) *CflifImage {
	p := C.flif_create_image_HDR(C.uint32_t(width), C.uint32_t(height))
	if p == nil {
		return nil
	}
	return &CflifImage{p}
}

func CflifImportImageRGBA(width, height int, data []byte, stride int) *CflifImage {
	p := C.flif_import_image_RGBA(C.uint32_t(width), C.uint32_t(height), unsafe.Pointer(&data[0]), C.uint32_t(stride))
	if p == nil {
		return nil
	}
	return &CflifImage{p}
}

func CflifImportImageRGB(width, height int, data []byte, stride int) *CflifImage {
	p := C.flif_import_image_RGB(C.uint32_t(width), C.uint32_t(height), unsafe.Pointer(&data[0]), C.uint32_t(stride))
	if p == nil {
		return nil
	}
	return &CflifImage{p}
}

func CflifImportImageGRAY(width, height int, data []byte, stride int) *CflifImage {
	p := C.flif_import_image_GRAY(C.uint32_t(width), C.uint32_t(height), unsafe.Pointer(&data[0]), C.uint32_t(stride))
	if p == nil {
		return nil
	}
	return &CflifImage{p}
}

func CflifDestroyImage(image *CflifImage) {
	C.flif_destroy_image(image.img)
}

func CflifImageGetWidth(image *CflifImage) int {
	return int(C.flif_image_get_width(image.img))
}

func CflifImageGetHeight(image *CflifImage) int {
	return int(C.flif_image_get_height(image.img))
}

func CflifImageGetNbChannels(image *CflifImage) int {
	return int(C.flif_image_get_nb_channels(image.img))
}

func CflifImageGetDepth(image *CflifImage) int {
	return int(C.flif_image_get_depth(image.img))
}

func CflifImageGetFrameDelay(image *CflifImage) int {
	return int(C.flif_image_get_frame_delay(image.img))
}

func CflifImageSetFrameDelay(image *CflifImage, delay int) {
	C.flif_image_set_frame_delay(image.img, C.uint32_t(delay))
}

func CflifImageSetMetadata(image *CflifImage, chunkname string, data []byte) {
	cchunk := C.CString(chunkname)
	defer C.free(unsafe.Pointer(cchunk))
	C.flif_image_set_metadata(image.img, cchunk, (*C.uchar)(unsafe.Pointer(&data[0])), (C.size_t)(len(data)))
}

func CflifImageGetMetadata(image *CflifImage, chunkname string, data *[]byte) {
	*data = nil

	cchunk := C.CString(chunkname)
	defer C.free(unsafe.Pointer(cchunk))
	var p unsafe.Pointer
	var num C.size_t

	C.flif_image_get_metadata(image.img, cchunk, (**C.uchar)(unsafe.Pointer(&p)), (*C.size_t)(unsafe.Pointer(&num)))
	if p == nil {
		return
	}
	defer CflifImageFreeMetadata(image, p)
	*data = C.GoBytes(p, C.int(num))
}

func CflifImageFreeMetadata(image *CflifImage, data unsafe.Pointer) {
	C.flif_image_free_metadata(image.img, (*C.uchar)(unsafe.Pointer(data)))
}

func CflifImageWriteRowRGBA8(image *CflifImage, row int, data []byte) {
	C.flif_image_write_row_RGBA8(image.img, C.uint32_t(row), unsafe.Pointer(&data[0]), C.size_t(len(data)))
}

func CflifImageReadRowRGBA8(image *CflifImage, row int) []byte {
	w := CflifImageGetWidth(image)
	b := make([]byte, C.sizeof_uint8_t*4*w)
	C.flif_image_read_row_RGBA8(image.img, C.uint32_t(row), unsafe.Pointer(&b[0]), C.size_t(len(b)))
	return b
}

func CflifImageWriteRowRGBA16(image *CflifImage, row int, data []byte) {
	C.flif_image_write_row_RGBA16(image.img, C.uint32_t(row), unsafe.Pointer(&data[0]), C.size_t(len(data)))
}

func CflifImageReadRowRGBA16(image *CflifImage, row int) []byte {
	w := CflifImageGetWidth(image)
	b := make([]byte, C.sizeof_uint16_t*4*w)
	C.flif_image_read_row_RGBA16(image.img, C.uint32_t(row), unsafe.Pointer(&b[0]), C.size_t(len(b)))
	return b
}

func CflifFreeMemory(p unsafe.Pointer) {
	C.flif_free_memory(p)
}
