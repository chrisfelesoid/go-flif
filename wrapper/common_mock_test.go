// Copyright 2016 chrisfelesoid. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package wrapper

import (
	"unsafe"

	"github.com/chrisfelesoid/go-flif/internal/wrapper"
)

type mockCommonWrapper struct {
	flifCreateImage         func(width, height int) *wrapper.CflifImage
	flifCreateImageHDR      func(width, height int) *wrapper.CflifImage
	flifImportImageRGBA     func(width, height int, data []byte, stride int) *wrapper.CflifImage
	flifImportImageRGB      func(width, height int, data []byte, stride int) *wrapper.CflifImage
	flifImportImageGRAY     func(width, height int, data []byte, stride int) *wrapper.CflifImage
	flifDestroyImage        func(image *wrapper.CflifImage)
	flifImageGetWidth       func(image *wrapper.CflifImage) int
	flifImageGetHeight      func(image *wrapper.CflifImage) int
	flifImageGetNbChannels  func(image *wrapper.CflifImage) int
	flifImageGetDepth       func(image *wrapper.CflifImage) int
	flifImageGetFrameDelay  func(image *wrapper.CflifImage) int
	flifImageSetFrameDelay  func(image *wrapper.CflifImage, delay int)
	flifImageSetMetadata    func(image *wrapper.CflifImage, chunkname string, data []byte)
	flifImageGetMetadata    func(image *wrapper.CflifImage, chunkname string, data *[]byte)
	flifImageFreeMetadata   func(image *wrapper.CflifImage, data unsafe.Pointer)
	flifImageWriteRowRGBA8  func(image *wrapper.CflifImage, row int, data []byte)
	flifImageReadRowRGBA8   func(image *wrapper.CflifImage, row int) []byte
	flifImageWriteRowRGBA16 func(image *wrapper.CflifImage, row int, data []byte)
	flifImageReadRowRGBA16  func(image *wrapper.CflifImage, row int) []byte
	flifFreeMemory          func(p unsafe.Pointer)
}

func (m *mockCommonWrapper) FlifCreateImage(width, height int) *wrapper.CflifImage {
	return m.flifCreateImage(width, height)
}

func (m *mockCommonWrapper) FlifCreateImageHDR(width, height int) *wrapper.CflifImage {
	return m.flifCreateImageHDR(width, height)
}
func (m *mockCommonWrapper) FlifImportImageRGBA(width, height int, data []byte, stride int) *wrapper.CflifImage {
	return m.flifImportImageRGBA(width, height, data, stride)
}

func (m *mockCommonWrapper) FlifImportImageRGB(width, height int, data []byte, stride int) *wrapper.CflifImage {
	return m.flifImportImageRGB(width, height, data, stride)
}

func (m *mockCommonWrapper) FlifImportImageGRAY(width, height int, data []byte, stride int) *wrapper.CflifImage {
	return m.flifImportImageGRAY(width, height, data, stride)
}

func (m *mockCommonWrapper) FlifDestroyImage(image *wrapper.CflifImage) {
	m.flifDestroyImage(image)
}

func (m *mockCommonWrapper) FlifImageGetWidth(image *wrapper.CflifImage) int {
	return m.flifImageGetWidth(image)
}

func (m *mockCommonWrapper) FlifImageGetHeight(image *wrapper.CflifImage) int {
	return m.flifImageGetHeight(image)
}

func (m *mockCommonWrapper) FlifImageGetNbChannels(image *wrapper.CflifImage) int {
	return m.flifImageGetNbChannels(image)
}

func (m *mockCommonWrapper) FlifImageGetDepth(image *wrapper.CflifImage) int {
	return m.flifImageGetDepth(image)
}

func (m *mockCommonWrapper) FlifImageGetFrameDelay(image *wrapper.CflifImage) int {
	return m.flifImageGetFrameDelay(image)
}

func (m *mockCommonWrapper) FlifImageSetFrameDelay(image *wrapper.CflifImage, delay int) {
	m.flifImageSetFrameDelay(image, delay)
}

func (m *mockCommonWrapper) FlifImageSetMetadata(image *wrapper.CflifImage, chunkname string, data []byte) {
	m.flifImageSetMetadata(image, chunkname, data)
}

func (m *mockCommonWrapper) FlifImageGetMetadata(image *wrapper.CflifImage, chunkname string, data *[]byte) {
	m.flifImageGetMetadata(image, chunkname, data)
}

func (m *mockCommonWrapper) FlifImageFreeMetadata(image *wrapper.CflifImage, data unsafe.Pointer) {
	m.flifImageFreeMetadata(image, data)
}

func (m *mockCommonWrapper) FlifImageWriteRowRGBA8(image *wrapper.CflifImage, row int, data []byte) {
	m.flifImageWriteRowRGBA8(image, row, data)
}

func (m *mockCommonWrapper) FlifImageReadRowRGBA8(image *wrapper.CflifImage, row int) []byte {
	return m.flifImageReadRowRGBA8(image, row)
}

func (m *mockCommonWrapper) FlifImageWriteRowRGBA16(image *wrapper.CflifImage, row int, data []byte) {
	m.flifImageWriteRowRGBA16(image, row, data)
}

func (m *mockCommonWrapper) FlifImageReadRowRGBA16(image *wrapper.CflifImage, row int) []byte {
	return m.flifImageReadRowRGBA16(image, row)
}

func (m *mockCommonWrapper) FlifFreeMemory(p unsafe.Pointer) {
	m.flifFreeMemory(p)
}
