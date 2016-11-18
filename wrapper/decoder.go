// Copyright 2016 chrisfelesoid. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package wrapper

import "sync"

// FlifDecoder FLIF_DECODER
type FlifDecoder struct {
	dec      *flifDecoder
	once     sync.Once
	CrcCheck bool
	Quality  int
	Scale    int
	Width    int
	Height   int
	Fit      bool
}

// FlifInfo FLIF_INFO
type FlifInfo struct {
	Width      int
	Height     int
	Channel    int
	Depth      int
	ImageCount int
}

// NewFlifDecoder flif_create_decoder
func NewFlifDecoder() *FlifDecoder {
	p := flifCreateDecoder()
	if p == nil {
		return nil
	}
	return &FlifDecoder{
		dec:      p,
		CrcCheck: false,
		Quality:  100,
		Scale:    1,
		Width:    0,
		Height:   0,
		Fit:      false,
	}
}

// Destroy flif_destroy_decoder
func (d *FlifDecoder) Destroy() {
	if d.dec == nil {
		return
	}

	d.once.Do(func() {
		flifDestroyDecoder(d.dec)
		d.dec = nil
	})
}

// Abort flif_abort_decoder
func (d *FlifDecoder) Abort() int {
	return flifAbortDecoder(d.dec)
}

func (d *FlifDecoder) setOptions() {
	crc := 0
	if d.CrcCheck {
		crc = 1
	}
	flifDecoderSetCrcCheck(d.dec, crc)
	flifDecoderSetQuality(d.dec, d.Quality)
	flifDecoderSetScale(d.dec, d.Scale)
	if !d.Fit {
		flifDecoderSetResize(d.dec, d.Width, d.Height)
	} else {
		flifDecoderSetFit(d.dec, d.Width, d.Height)
	}
}

// DecodeFile flif_decoder_decode_file
func (d *FlifDecoder) DecodeFile(name string) int {
	d.setOptions()
	return flifDecoderDecodeFile(d.dec, name)
}

// DecodeMemory flif_decoder_decode_memory
func (d *FlifDecoder) DecodeMemory(data []byte) int {
	d.setOptions()
	return flifDecoderDecodeMemory(d.dec, data)
}

// GetImageCount flif_decoder_num_images
func (d *FlifDecoder) GetImageCount() int {
	return flifDecoderNumImages(d.dec)
}

// GetLoopCount flif_decoder_num_loops
func (d *FlifDecoder) GetLoopCount() int {
	return flifDecoderGetNumLoops(d.dec)
}

// GetImage flif_decoder_get_image
// TODO: this FlifImage must not destroy. should destroy from FlifDecoder.
func (d *FlifDecoder) GetImage() *FlifImage {
	// get all images because flif_decoder_get_image is invalidate internal image.
	num := d.GetImageCount()
	fi := &FlifImage{
		images: make([]*flifImage, num),
	}

	for i := 0; i < num; i++ {
		p := flifDecoderGetImage(d.dec, i)
		fi.images[i] = p
	}

	return fi
}

// GetInfo flif_info_xxx
func GetInfo(data []byte) *FlifInfo {
	info := flifReadInfoFromMemory(data)
	if info == nil {
		return nil
	}
	defer flifDestroyInfo(info)

	return &FlifInfo{
		Width:      flifInfoGetWidth(info),
		Height:     flifInfoGetHeight(info),
		Channel:    flifInfoGetNbChannels(info),
		Depth:      flifInfoGetDepth(info),
		ImageCount: flifInfoNumImages(info),
	}
}
