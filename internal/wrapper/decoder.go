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

type CflifDecoder struct {
	dec *C.FLIF_DECODER
}

type CflifInfo struct {
	info *C.FLIF_INFO
}

type CdecoderFunctionWrapper interface {
	FlifCreateDecoder() *CflifDecoder
	FlifDecoderDecodeFile(decoder *CflifDecoder, filename string) int
	FlifDecoderDecodeMemory(decoder *CflifDecoder, data []byte) int
	FlifDecoderNumImages(decoder *CflifDecoder) int
	FlifDecoderGetNumLoops(decoder *CflifDecoder) int
	FlifDecoderGetImage(decoder *CflifDecoder, index int) *CflifImage
	FlifDestroyDecoder(decoder *CflifDecoder)
	FlifAbortDecoder(decoder *CflifDecoder) int
	FlifDecoderSetCrcCheck(decoder *CflifDecoder, crcCheck int)
	FlifDecoderSetQuality(decoder *CflifDecoder, quality int)
	FlifDecoderSetScale(decoder *CflifDecoder, scale int)
	FlifDecoderSetResize(decoder *CflifDecoder, width, height int)
	FlifDecoderSetFit(decoder *CflifDecoder, width, height int)
	// FlifDecoderSetCallback()
	// FlifDecoderSetFirstCallbackQuality()
	FlifReadInfoFromMemory(data []byte) *CflifInfo
	FlifDestroyInfo(info *CflifInfo)
	FlifInfoGetWidth(info *CflifInfo) int
	FlifInfoGetHeight(info *CflifInfo) int
	FlifInfoGetNbChannels(info *CflifInfo) int
	FlifInfoGetDepth(info *CflifInfo) int
	FlifInfoNumImages(info *CflifInfo) int
}

type CdecoderWrapper struct {
}

func (d *CdecoderWrapper) FlifCreateDecoder() *CflifDecoder {
	p := C.flif_create_decoder()
	if p == nil {
		return nil
	}
	return &CflifDecoder{p}
}

func (d *CdecoderWrapper) FlifDecoderDecodeFile(decoder *CflifDecoder, filename string) int {
	cname := C.CString(filename)
	defer C.free(unsafe.Pointer(cname))
	return int(C.flif_decoder_decode_file(decoder.dec, cname))
}

func (d *CdecoderWrapper) FlifDecoderDecodeMemory(decoder *CflifDecoder, data []byte) int {
	return int(C.flif_decoder_decode_memory(decoder.dec, unsafe.Pointer(&data[0]), C.size_t(len(data))))
}

func (d *CdecoderWrapper) FlifDecoderNumImages(decoder *CflifDecoder) int {
	return int(C.flif_decoder_num_images(decoder.dec))
}

func (d *CdecoderWrapper) FlifDecoderGetNumLoops(decoder *CflifDecoder) int {
	return int(C.flif_decoder_num_loops(decoder.dec))
}

func (d *CdecoderWrapper) FlifDecoderGetImage(decoder *CflifDecoder, index int) *CflifImage {
	p := C.flif_decoder_get_image(decoder.dec, C.size_t(index))
	if p == nil {
		return nil
	}
	return &CflifImage{p}
}

func (d *CdecoderWrapper) FlifDestroyDecoder(decoder *CflifDecoder) {
	C.flif_destroy_decoder(decoder.dec)
}

func (d *CdecoderWrapper) FlifAbortDecoder(decoder *CflifDecoder) int {
	return int(C.flif_abort_decoder(decoder.dec))
}

func (d *CdecoderWrapper) FlifDecoderSetCrcCheck(decoder *CflifDecoder, crcCheck int) {
	C.flif_decoder_set_crc_check(decoder.dec, C.int32_t(crcCheck))
}

func (d *CdecoderWrapper) FlifDecoderSetQuality(decoder *CflifDecoder, quality int) {
	C.flif_decoder_set_quality(decoder.dec, C.int32_t(quality))
}

func (d *CdecoderWrapper) FlifDecoderSetScale(decoder *CflifDecoder, scale int) {
	C.flif_decoder_set_scale(decoder.dec, C.uint32_t(scale))
}

func (d *CdecoderWrapper) FlifDecoderSetResize(decoder *CflifDecoder, width, height int) {
	C.flif_decoder_set_resize(decoder.dec, C.uint32_t(width), C.uint32_t(height))
}

func (d *CdecoderWrapper) FlifDecoderSetFit(decoder *CflifDecoder, width, height int) {
	C.flif_decoder_set_fit(decoder.dec, C.uint32_t(width), C.uint32_t(height))
}

// func (d *CdecoderWrapper) FlifDecoderSetCallback() {
// 	//C.flif_decoder_set_callback(FLIF_DECODER* decoder, uint32_t (*callback)(int32_t quality, int64_t bytes_read))
// }

// func (d *CdecoderWrapper) FlifDecoderSetFirstCallbackQuality() {
// 	//C.flif_decoder_set_first_callback_quality(FLIF_DECODER* decoder, int32_t quality) // valid quality: 0-10000
// }

func (d *CdecoderWrapper) FlifReadInfoFromMemory(data []byte) *CflifInfo {
	p := C.flif_read_info_from_memory(unsafe.Pointer(&data[0]), C.size_t(len(data)))
	if p == nil {
		return nil
	}
	return &CflifInfo{p}
}

func (d *CdecoderWrapper) FlifDestroyInfo(info *CflifInfo) {
	C.flif_destroy_info(info.info)
}

func (d *CdecoderWrapper) FlifInfoGetWidth(info *CflifInfo) int {
	return int(C.flif_info_get_width(info.info))
}

func (d *CdecoderWrapper) FlifInfoGetHeight(info *CflifInfo) int {
	return int(C.flif_info_get_height(info.info))
}

func (d *CdecoderWrapper) FlifInfoGetNbChannels(info *CflifInfo) int {
	return int(C.flif_info_get_nb_channels(info.info))
}

func (d *CdecoderWrapper) FlifInfoGetDepth(info *CflifInfo) int {
	return int(C.flif_info_get_depth(info.info))
}

func (d *CdecoderWrapper) FlifInfoNumImages(info *CflifInfo) int {
	return int(C.flif_info_num_images(info.info))
}
