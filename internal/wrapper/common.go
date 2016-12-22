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

var common = &CcommonWrapper{}

type CflifImage struct {
	img *C.FLIF_IMAGE
}

type CcommonFunctionWrapper interface {
	FlifCreateImage(width, height int) *CflifImage
	FlifCreateImageHDR(width, height int) *CflifImage
	FlifImportImageRGBA(width, height int, data []byte, stride int) *CflifImage
	FlifImportImageRGB(width, height int, data []byte, stride int) *CflifImage
	FlifImportImageGRAY(width, height int, data []byte, stride int) *CflifImage
	FlifDestroyImage(image *CflifImage)
	FlifImageGetWidth(image *CflifImage) int
	FlifImageGetHeight(image *CflifImage) int
	FlifImageGetNbChannels(image *CflifImage) int
	FlifImageGetDepth(image *CflifImage) int
	FlifImageGetFrameDelay(image *CflifImage) int
	FlifImageSetFrameDelay(image *CflifImage, delay int)
	FlifImageSetMetadata(image *CflifImage, chunkname string, data []byte)
	FlifImageGetMetadata(image *CflifImage, chunkname string, data *[]byte)
	FlifImageFreeMetadata(image *CflifImage, data unsafe.Pointer)
	FlifImageWriteRowRGBA8(image *CflifImage, row int, data []byte)
	FlifImageReadRowRGBA8(image *CflifImage, row int) []byte
	FlifImageWriteRowRGBA16(image *CflifImage, row int, data []byte)
	FlifImageReadRowRGBA16(image *CflifImage, row int) []byte
	FlifFreeMemory(p unsafe.Pointer)
}

type CcommonWrapper struct {
}

func (c *CcommonWrapper) FlifCreateImage(width, height int) *CflifImage {
	p := C.flif_create_image(C.uint32_t(width), C.uint32_t(height))
	if p == nil {
		return nil
	}
	return &CflifImage{p}
}

func (c *CcommonWrapper) FlifCreateImageHDR(width, height int) *CflifImage {
	p := C.flif_create_image_HDR(C.uint32_t(width), C.uint32_t(height))
	if p == nil {
		return nil
	}
	return &CflifImage{p}
}

func (c *CcommonWrapper) FlifImportImageRGBA(width, height int, data []byte, stride int) *CflifImage {
	p := C.flif_import_image_RGBA(C.uint32_t(width), C.uint32_t(height), unsafe.Pointer(&data[0]), C.uint32_t(stride))
	if p == nil {
		return nil
	}
	return &CflifImage{p}
}

func (c *CcommonWrapper) FlifImportImageRGB(width, height int, data []byte, stride int) *CflifImage {
	p := C.flif_import_image_RGB(C.uint32_t(width), C.uint32_t(height), unsafe.Pointer(&data[0]), C.uint32_t(stride))
	if p == nil {
		return nil
	}
	return &CflifImage{p}
}

func (c *CcommonWrapper) FlifImportImageGRAY(width, height int, data []byte, stride int) *CflifImage {
	p := C.flif_import_image_GRAY(C.uint32_t(width), C.uint32_t(height), unsafe.Pointer(&data[0]), C.uint32_t(stride))
	if p == nil {
		return nil
	}
	return &CflifImage{p}
}

func (c *CcommonWrapper) FlifDestroyImage(image *CflifImage) {
	C.flif_destroy_image(image.img)
}

func (c *CcommonWrapper) FlifImageGetWidth(image *CflifImage) int {
	return int(C.flif_image_get_width(image.img))
}

func (c *CcommonWrapper) FlifImageGetHeight(image *CflifImage) int {
	return int(C.flif_image_get_height(image.img))
}

func (c *CcommonWrapper) FlifImageGetNbChannels(image *CflifImage) int {
	return int(C.flif_image_get_nb_channels(image.img))
}

func (c *CcommonWrapper) FlifImageGetDepth(image *CflifImage) int {
	return int(C.flif_image_get_depth(image.img))
}

func (c *CcommonWrapper) FlifImageGetFrameDelay(image *CflifImage) int {
	return int(C.flif_image_get_frame_delay(image.img))
}

func (c *CcommonWrapper) FlifImageSetFrameDelay(image *CflifImage, delay int) {
	C.flif_image_set_frame_delay(image.img, C.uint32_t(delay))
}

func (c *CcommonWrapper) FlifImageSetMetadata(image *CflifImage, chunkname string, data []byte) {
	cchunk := C.CString(chunkname)
	defer C.free(unsafe.Pointer(cchunk))
	C.flif_image_set_metadata(image.img, cchunk, (*C.uchar)(unsafe.Pointer(&data[0])), (C.size_t)(len(data)))
}

func (c *CcommonWrapper) FlifImageGetMetadata(image *CflifImage, chunkname string, data *[]byte) {
	*data = nil

	cchunk := C.CString(chunkname)
	defer C.free(unsafe.Pointer(cchunk))
	var p unsafe.Pointer
	var num C.size_t

	C.flif_image_get_metadata(image.img, cchunk, (**C.uchar)(unsafe.Pointer(&p)), (*C.size_t)(unsafe.Pointer(&num)))
	if p == nil {
		return
	}
	defer c.FlifImageFreeMetadata(image, p)
	*data = C.GoBytes(p, C.int(num))
}

func (c *CcommonWrapper) FlifImageFreeMetadata(image *CflifImage, data unsafe.Pointer) {
	C.flif_image_free_metadata(image.img, (*C.uchar)(unsafe.Pointer(data)))
}

func (c *CcommonWrapper) FlifImageWriteRowRGBA8(image *CflifImage, row int, data []byte) {
	C.flif_image_write_row_RGBA8(image.img, C.uint32_t(row), unsafe.Pointer(&data[0]), C.size_t(len(data)))
}

func (c *CcommonWrapper) FlifImageReadRowRGBA8(image *CflifImage, row int) []byte {
	w := c.FlifImageGetWidth(image)
	b := make([]byte, C.sizeof_uint8_t*4*w)
	C.flif_image_read_row_RGBA8(image.img, C.uint32_t(row), unsafe.Pointer(&b[0]), C.size_t(len(b)))
	return b
}

func (c *CcommonWrapper) FlifImageWriteRowRGBA16(image *CflifImage, row int, data []byte) {
	C.flif_image_write_row_RGBA16(image.img, C.uint32_t(row), unsafe.Pointer(&data[0]), C.size_t(len(data)))
}

func (c *CcommonWrapper) FlifImageReadRowRGBA16(image *CflifImage, row int) []byte {
	w := c.FlifImageGetWidth(image)
	b := make([]byte, C.sizeof_uint16_t*4*w)
	C.flif_image_read_row_RGBA16(image.img, C.uint32_t(row), unsafe.Pointer(&b[0]), C.size_t(len(b)))
	return b
}

func (c *CcommonWrapper) FlifFreeMemory(p unsafe.Pointer) {
	C.flif_free_memory(p)
}
