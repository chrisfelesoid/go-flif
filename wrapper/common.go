// Copyright 2016 chrisfelesoid. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package wrapper

import (
	"errors"
	"sync"

	"github.com/chrisfelesoid/go-flif/internal/wrapper"
)

// FlifImage FLIF_IMAGE
type FlifImage struct {
	images []*wrapper.CflifImage
	once   sync.Once
}

var (
	ErrOutOfRange = errors.New("index out of range")
)

var common wrapper.CcommonFunctionWrapper

func NewFlifImage(width, height int) *FlifImage {
	img := common.FlifCreateImage(width, height)
	if img == nil {
		return nil
	}
	return &FlifImage{
		images: []*wrapper.CflifImage{img},
	}
}

func NewFlifImageHDR(width, height int) *FlifImage {
	img := common.FlifCreateImageHDR(width, height)
	if img == nil {
		return nil
	}
	return &FlifImage{
		images: []*wrapper.CflifImage{img},
	}
}

func NewFlifImageFromRGBA(width, height int, data []byte) *FlifImage {
	stride := width * 4 // R,G,B,A = 4
	img := common.FlifImportImageRGBA(width, height, data, stride)
	if img == nil {
		return nil
	}
	return &FlifImage{
		images: []*wrapper.CflifImage{img},
	}
}

func NewFlifImageFromRGB(width, height int, data []byte) *FlifImage {
	stride := width * 3 // R,G,B = 3
	img := common.FlifImportImageRGB(width, height, data, stride)
	if img == nil {
		return nil
	}
	return &FlifImage{
		images: []*wrapper.CflifImage{img},
	}
}

func NewFlifImageFromGray(width, height int, data []byte) *FlifImage {
	stride := width * 1 // GrayScale = 1
	img := common.FlifImportImageGRAY(width, height, data, stride)
	if img == nil {
		return nil
	}
	return &FlifImage{
		images: []*wrapper.CflifImage{img},
	}
}

func (f *FlifImage) Destroy() {
	if f.images == nil {
		return
	}

	f.once.Do(func() {
		for _, img := range f.images {
			common.FlifDestroyImage(img)
		}
		f.images = nil
	})
}

func (f *FlifImage) GetImageCount() int {
	return len(f.images)
}

func (f *FlifImage) GetWidth() int {
	if f.images == nil {
		return 0
	}
	return common.FlifImageGetWidth(f.images[0])
}

func (f *FlifImage) GetHeight() int {
	if f.images == nil {
		return 0
	}
	return common.FlifImageGetHeight(f.images[0])
}

func (f *FlifImage) GetChannel() int {
	if f.images == nil {
		return 0
	}
	return common.FlifImageGetNbChannels(f.images[0])
}

func (f *FlifImage) GetDepth() int {
	if f.images == nil {
		return 0
	}
	return common.FlifImageGetDepth(f.images[0])
}

func (f *FlifImage) GetFrameDelay() int {
	if f.images == nil {
		return 0
	}
	return common.FlifImageGetFrameDelay(f.images[0])
}

func (f *FlifImage) SetFrameDelay(delay int) {
	if f.images == nil {
		return
	}
	common.FlifImageSetFrameDelay(f.images[0], delay)
}

func (f *FlifImage) SetMetadata(name string, data []byte) {
	if f.images == nil {
		return
	}
	common.FlifImageSetMetadata(f.images[0], name, data)
}

func (f *FlifImage) GetMetadata(name string) []byte {
	if f.images == nil {
		return nil
	}
	var data []byte
	common.FlifImageGetMetadata(f.images[0], name, &data)
	return data
}

func (f *FlifImage) WriteRowRGBA8(row int, data []byte, index int) error {
	w := f.GetWidth()
	h := f.GetHeight()
	if h == 0 || row >= h || row < 0 {
		return ErrOutOfRange
	}
	if w == 0 || len(data) < w*4 {
		return ErrOutOfRange
	}
	if index < 0 || len(f.images) <= index {
		return ErrOutOfRange
	}
	common.FlifImageWriteRowRGBA8(f.images[index], row, data)
	return nil
}

func (f *FlifImage) ReadRowRGBA8(row, index int) ([]byte, error) {
	h := f.GetHeight()
	if h == 0 || row >= h || row < 0 {
		return nil, ErrOutOfRange
	}
	if index < 0 || len(f.images) <= index {
		return nil, ErrOutOfRange
	}
	data := common.FlifImageReadRowRGBA8(f.images[index], row)
	if data == nil {
		return nil, ErrUnknown
	}
	return data, nil
}

func (f *FlifImage) WriteRowRGBA16(row int, data []byte, index int) error {
	w := f.GetWidth()
	h := f.GetHeight()
	if h == 0 || row >= h {
		return ErrOutOfRange
	}
	if w == 0 || len(data) < w*4*2 {
		return ErrOutOfRange
	}
	if index < 0 || len(f.images) <= index {
		return ErrOutOfRange
	}
	common.FlifImageWriteRowRGBA16(f.images[index], row, data)
	return nil
}

func (f *FlifImage) ReadRowRGBA16(row, index int) ([]byte, error) {
	h := f.GetHeight()
	if h == 0 || row >= h || row < 0 {
		return nil, ErrOutOfRange
	}
	if index < 0 || len(f.images) <= index {
		return nil, ErrOutOfRange
	}
	data := common.FlifImageReadRowRGBA16(f.images[index], row)
	if data == nil {
		return nil, ErrUnknown
	}
	return data, nil
}

func init() {
	common = &wrapper.CcommonWrapper{}
}
