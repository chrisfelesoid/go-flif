// Copyright 2016 chrisfelesoid. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package wrapper

import (
	"sync"

	"github.com/chrisfelesoid/go-flif/internal/wrapper"
)

// FlifDecoder FLIF_DECODER
type FlifDecoder struct {
	dec      *wrapper.CflifDecoder
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

var decoder wrapper.CdecoderFunctionWrapper

// NewFlifDecoder flif_create_decoder
func NewFlifDecoder() *FlifDecoder {
	p := decoder.FlifCreateDecoder()
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
		decoder.FlifDestroyDecoder(d.dec)
		d.dec = nil
	})
}

// Abort flif_abort_decoder
func (d *FlifDecoder) Abort() error {
	return checkResult(decoder.FlifAbortDecoder(d.dec))
}

// DecodeFile flif_decoder_decode_file
func (d *FlifDecoder) DecodeFile(name string) error {
	setDecoderOptions(d)
	return checkResult(decoder.FlifDecoderDecodeFile(d.dec, name))
}

// DecodeMemory flif_decoder_decode_memory
func (d *FlifDecoder) DecodeMemory(data []byte) error {
	setDecoderOptions(d)
	return checkResult(decoder.FlifDecoderDecodeMemory(d.dec, data))
}

// GetImageCount flif_decoder_num_images
func (d *FlifDecoder) GetImageCount() int {
	return decoder.FlifDecoderNumImages(d.dec)
}

// GetLoopCount flif_decoder_num_loops
func (d *FlifDecoder) GetLoopCount() int {
	return decoder.FlifDecoderGetNumLoops(d.dec)
}

// GetImage flif_decoder_get_image
// TODO: this FlifImage must not destroy. should destroy from FlifDecoder.
func (d *FlifDecoder) GetImage() *FlifImage {
	// get all images because flif_decoder_get_image is invalidate internal image.
	num := d.GetImageCount()
	fi := &FlifImage{
		images: make([]*wrapper.CflifImage, num),
	}

	for i := 0; i < num; i++ {
		p := decoder.FlifDecoderGetImage(d.dec, i)
		fi.images[i] = p
	}

	return fi
}

// GetInfo flif_info_xxx
func GetInfo(data []byte) *FlifInfo {
	info := decoder.FlifReadInfoFromMemory(data)
	if info == nil {
		return nil
	}
	defer decoder.FlifDestroyInfo(info)

	return &FlifInfo{
		Width:      decoder.FlifInfoGetWidth(info),
		Height:     decoder.FlifInfoGetHeight(info),
		Channel:    decoder.FlifInfoGetNbChannels(info),
		Depth:      decoder.FlifInfoGetDepth(info),
		ImageCount: decoder.FlifInfoNumImages(info),
	}
}

var setDecoderOptions = func(d *FlifDecoder) {
	crc := 0
	if d.CrcCheck {
		crc = 1
	}
	decoder.FlifDecoderSetCrcCheck(d.dec, crc)
	decoder.FlifDecoderSetQuality(d.dec, d.Quality)
	decoder.FlifDecoderSetScale(d.dec, d.Scale)
	if !d.Fit {
		decoder.FlifDecoderSetResize(d.dec, d.Width, d.Height)
	} else {
		decoder.FlifDecoderSetFit(d.dec, d.Width, d.Height)
	}
}

func init() {
	decoder = &wrapper.CdecoderWrapper{}
}
