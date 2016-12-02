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

// NewFlifDecoder flif_create_decoder
func NewFlifDecoder() *FlifDecoder {
	p := wrapper.CflifCreateDecoder()
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
		wrapper.CflifDestroyDecoder(d.dec)
		d.dec = nil
	})
}

// Abort flif_abort_decoder
func (d *FlifDecoder) Abort() int {
	return wrapper.CflifAbortDecoder(d.dec)
}

func (d *FlifDecoder) setOptions() {
	crc := 0
	if d.CrcCheck {
		crc = 1
	}
	wrapper.CflifDecoderSetCrcCheck(d.dec, crc)
	wrapper.CflifDecoderSetQuality(d.dec, d.Quality)
	wrapper.CflifDecoderSetScale(d.dec, d.Scale)
	if !d.Fit {
		wrapper.CflifDecoderSetResize(d.dec, d.Width, d.Height)
	} else {
		wrapper.CflifDecoderSetFit(d.dec, d.Width, d.Height)
	}
}

// DecodeFile flif_decoder_decode_file
func (d *FlifDecoder) DecodeFile(name string) int {
	d.setOptions()
	return wrapper.CflifDecoderDecodeFile(d.dec, name)
}

// DecodeMemory flif_decoder_decode_memory
func (d *FlifDecoder) DecodeMemory(data []byte) int {
	d.setOptions()
	return wrapper.CflifDecoderDecodeMemory(d.dec, data)
}

// GetImageCount flif_decoder_num_images
func (d *FlifDecoder) GetImageCount() int {
	return wrapper.CflifDecoderNumImages(d.dec)
}

// GetLoopCount flif_decoder_num_loops
func (d *FlifDecoder) GetLoopCount() int {
	return wrapper.CflifDecoderGetNumLoops(d.dec)
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
		p := wrapper.CflifDecoderGetImage(d.dec, i)
		fi.images[i] = p
	}

	return fi
}

// GetInfo flif_info_xxx
func GetInfo(data []byte) *FlifInfo {
	info := wrapper.CflifReadInfoFromMemory(data)
	if info == nil {
		return nil
	}
	defer wrapper.CflifDestroyInfo(info)

	return &FlifInfo{
		Width:      wrapper.CflifInfoGetWidth(info),
		Height:     wrapper.CflifInfoGetHeight(info),
		Channel:    wrapper.CflifInfoGetNbChannels(info),
		Depth:      wrapper.CflifInfoGetDepth(info),
		ImageCount: wrapper.CflifInfoNumImages(info),
	}
}
