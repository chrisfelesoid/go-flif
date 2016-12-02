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

func CflifCreateDecoder() *CflifDecoder {
	p := C.flif_create_decoder()
	if p == nil {
		return nil
	}
	return &CflifDecoder{p}
}

func CflifDecoderDecodeFile(decoder *CflifDecoder, filename string) int {
	cname := C.CString(filename)
	defer C.free(unsafe.Pointer(cname))
	return int(C.flif_decoder_decode_file(decoder.dec, cname))
}

func CflifDecoderDecodeMemory(decoder *CflifDecoder, data []byte) int {
	return int(C.flif_decoder_decode_memory(decoder.dec, unsafe.Pointer(&data[0]), C.size_t(len(data))))
}

func CflifDecoderNumImages(decoder *CflifDecoder) int {
	return int(C.flif_decoder_num_images(decoder.dec))
}

func CflifDecoderGetNumLoops(decoder *CflifDecoder) int {
	return int(C.flif_decoder_num_loops(decoder.dec))
}

func CflifDecoderGetImage(decoder *CflifDecoder, index int) *CflifImage {
	p := C.flif_decoder_get_image(decoder.dec, C.size_t(index))
	if p == nil {
		return nil
	}
	return &CflifImage{p}
}

func CflifDestroyDecoder(decoder *CflifDecoder) {
	C.flif_destroy_decoder(decoder.dec)
}

func CflifAbortDecoder(decoder *CflifDecoder) int {
	return int(C.flif_abort_decoder(decoder.dec))
}

func CflifDecoderSetCrcCheck(decoder *CflifDecoder, crcCheck int) {
	C.flif_decoder_set_crc_check(decoder.dec, C.int32_t(crcCheck))
}

func CflifDecoderSetQuality(decoder *CflifDecoder, quality int) {
	C.flif_decoder_set_quality(decoder.dec, C.int32_t(quality))
}

func CflifDecoderSetScale(decoder *CflifDecoder, scale int) {
	C.flif_decoder_set_scale(decoder.dec, C.uint32_t(scale))
}

func CflifDecoderSetResize(decoder *CflifDecoder, width, height int) {
	C.flif_decoder_set_resize(decoder.dec, C.uint32_t(width), C.uint32_t(height))
}

func CflifDecoderSetFit(decoder *CflifDecoder, width, height int) {
	C.flif_decoder_set_fit(decoder.dec, C.uint32_t(width), C.uint32_t(height))
}

func CflifDecoderSetCallback() {
	//C.flif_decoder_set_callback(FLIF_DECODER* decoder, uint32_t (*callback)(int32_t quality, int64_t bytes_read))
}

func CflifDecoderSetFirstCallbackQuality() {
	//C.flif_decoder_set_first_callback_quality(FLIF_DECODER* decoder, int32_t quality) // valid quality: 0-10000
}

func CflifReadInfoFromMemory(data []byte) *CflifInfo {
	p := C.flif_read_info_from_memory(unsafe.Pointer(&data[0]), C.size_t(len(data)))
	if p == nil {
		return nil
	}
	return &CflifInfo{p}
}

func CflifDestroyInfo(info *CflifInfo) {
	C.flif_destroy_info(info.info)
}

func CflifInfoGetWidth(info *CflifInfo) int {
	return int(C.flif_info_get_width(info.info))
}

func CflifInfoGetHeight(info *CflifInfo) int {
	return int(C.flif_info_get_height(info.info))
}

func CflifInfoGetNbChannels(info *CflifInfo) int {
	return int(C.flif_info_get_nb_channels(info.info))
}

func CflifInfoGetDepth(info *CflifInfo) int {
	return int(C.flif_info_get_depth(info.info))
}

func CflifInfoNumImages(info *CflifInfo) int {
	return int(C.flif_info_num_images(info.info))
}
