// Copyright 2016 chrisfelesoid. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package wrapper

import "sync"

// FlifImage FLIF_IMAGE
type FlifImage struct {
	images []*flifImage
	once   sync.Once
}

func NewFlifImage(width, height int) *FlifImage {
	img := flifCreateImage(width, height)
	if img == nil {
		return nil
	}
	return &FlifImage{
		images: []*flifImage{img},
	}
}

func NewFlifImageHDR(width, height int) *FlifImage {
	img := flifCreateImageHDR(width, height)
	if img == nil {
		return nil
	}
	return &FlifImage{
		images: []*flifImage{img},
	}
}

func NewFlifImageFromRGBA(width, height int, data []byte) *FlifImage {
	stride := width * 4 // R,G,B,A = 4
	img := flifImportImageRGBA(width, height, data, stride)
	if img == nil {
		return nil
	}
	return &FlifImage{
		images: []*flifImage{img},
	}
}

func NewFlifImageFromRGB(width, height int, data []byte) *FlifImage {
	stride := width * 3 // R,G,B = 3
	img := flifImportImageRGB(width, height, data, stride)
	if img == nil {
		return nil
	}
	return &FlifImage{
		images: []*flifImage{img},
	}
}

func NewFlifImageFromGray(width, height int, data []byte) *FlifImage {
	stride := width * 1 // GrayScale = 1
	img := flifImportImageGRAY(width, height, data, stride)
	if img == nil {
		return nil
	}
	return &FlifImage{
		images: []*flifImage{img},
	}
}

func (f *FlifImage) Destroy() {
	if f.images == nil {
		return
	}

	f.once.Do(func() {
		for _, img := range f.images {
			flifDestroyImage(img)
		}
		f.images = nil
	})
}

func (f *FlifImage) GetWidth() int {
	return flifImageGetWidth(f.images[0])
}

func (f *FlifImage) GetHeight() int {
	return flifImageGetHeight(f.images[0])
}

func (f *FlifImage) GetChannel() int {
	return flifImageGetNbChannels(f.images[0])
}

func (f *FlifImage) GetDepth() int {
	return flifImageGetDepth(f.images[0])
}

func (f *FlifImage) GetFrameDelay() int {
	return flifImageGetFrameDelay(f.images[0])
}

func (f *FlifImage) SetFrameDelay(delay int) {
	flifImageSetFrameDelay(f.images[0], delay)
}

func (f *FlifImage) SetMetadata(name string, data []byte) {
	flifImageSetMetadata(f.images[0], name, data)
}

func (f *FlifImage) GetMetadata(name string) []byte {
	var data []byte
	flifImageGetMetadata(f.images[0], name, &data)
	return data
}

func (f *FlifImage) WriteRowRGBA8(row int, data []byte, index int) {
	if row >= f.GetHeight() {
		return
	}
	flifImageWriteRowRGBA8(f.images[index], row, data)
}

func (f *FlifImage) ReadRowRGBA8(row, index int) []byte {
	if row >= f.GetHeight() {
		return nil
	}
	return flifImageReadRowRGBA8(f.images[index], row)
}

func (f *FlifImage) WriteRowRGBA16(row int, data []byte, index int) {
	if row >= f.GetHeight() {
		return
	}
	flifImageWriteRowRGBA16(f.images[index], row, data)
}

func (f *FlifImage) ReadRowRGBA16(row, index int) []byte {
	if row >= f.GetHeight() {
		return nil
	}
	return flifImageReadRowRGBA16(f.images[index], row)
}
