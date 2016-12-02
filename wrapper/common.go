// Copyright 2016 chrisfelesoid. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package wrapper

import (
	"sync"
	"github.com/chrisfelesoid/go-flif/internal/wrapper"
)

// FlifImage FLIF_IMAGE
type FlifImage struct {
	images []*wrapper.CflifImage
	once   sync.Once
}

func NewFlifImage(width, height int) *FlifImage {
	img := wrapper.CflifCreateImage(width, height)
	if img == nil {
		return nil
	}
	return &FlifImage{
		images: []*wrapper.CflifImage{img},
	}
}

func NewFlifImageHDR(width, height int) *FlifImage {
	img := wrapper.CflifCreateImageHDR(width, height)
	if img == nil {
		return nil
	}
	return &FlifImage{
		images: []*wrapper.CflifImage{img},
	}
}

func NewFlifImageFromRGBA(width, height int, data []byte) *FlifImage {
	stride := width * 4 // R,G,B,A = 4
	img := wrapper.CflifImportImageRGBA(width, height, data, stride)
	if img == nil {
		return nil
	}
	return &FlifImage{
		images: []*wrapper.CflifImage{img},
	}
}

func NewFlifImageFromRGB(width, height int, data []byte) *FlifImage {
	stride := width * 3 // R,G,B = 3
	img := wrapper.CflifImportImageRGB(width, height, data, stride)
	if img == nil {
		return nil
	}
	return &FlifImage{
		images: []*wrapper.CflifImage{img},
	}
}

func NewFlifImageFromGray(width, height int, data []byte) *FlifImage {
	stride := width * 1 // GrayScale = 1
	img := wrapper.CflifImportImageGRAY(width, height, data, stride)
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
			wrapper.CflifDestroyImage(img)
		}
		f.images = nil
	})
}

func (f *FlifImage) GetImageCount() int {
	return len(f.images)
}

func (f *FlifImage) GetWidth() int {
	return wrapper.CflifImageGetWidth(f.images[0])
}

func (f *FlifImage) GetHeight() int {
	return wrapper.CflifImageGetHeight(f.images[0])
}

func (f *FlifImage) GetChannel() int {
	return wrapper.CflifImageGetNbChannels(f.images[0])
}

func (f *FlifImage) GetDepth() int {
	return wrapper.CflifImageGetDepth(f.images[0])
}

func (f *FlifImage) GetFrameDelay() int {
	return wrapper.CflifImageGetFrameDelay(f.images[0])
}

func (f *FlifImage) SetFrameDelay(delay int) {
	wrapper.CflifImageSetFrameDelay(f.images[0], delay)
}

func (f *FlifImage) SetMetadata(name string, data []byte) {
	wrapper.CflifImageSetMetadata(f.images[0], name, data)
}

func (f *FlifImage) GetMetadata(name string) []byte {
	var data []byte
	wrapper.CflifImageGetMetadata(f.images[0], name, &data)
	return data
}

func (f *FlifImage) WriteRowRGBA8(row int, data []byte, index int) {
	if row >= f.GetHeight() {
		return
	}
	wrapper.CflifImageWriteRowRGBA8(f.images[index], row, data)
}

func (f *FlifImage) ReadRowRGBA8(row, index int) []byte {
	if row >= f.GetHeight() {
		return nil
	}
	return wrapper.CflifImageReadRowRGBA8(f.images[index], row)
}

func (f *FlifImage) WriteRowRGBA16(row int, data []byte, index int) {
	if row >= f.GetHeight() {
		return
	}
	wrapper.CflifImageWriteRowRGBA16(f.images[index], row, data)
}

func (f *FlifImage) ReadRowRGBA16(row, index int) []byte {
	if row >= f.GetHeight() {
		return nil
	}
	return wrapper.CflifImageReadRowRGBA16(f.images[index], row)
}
