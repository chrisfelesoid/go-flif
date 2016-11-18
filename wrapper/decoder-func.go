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

type flifDecoder struct {
	dec *C.FLIF_DECODER
}

type flifInfo struct {
	info *C.FLIF_INFO
}

var flifCreateDecoder = func() *flifDecoder {
	p := C.flif_create_decoder()
	if p == nil {
		return nil
	}
	return &flifDecoder{p}
}

var flifDecoderDecodeFile = func(decoder *flifDecoder, filename string) int {
	cname := C.CString(filename)
	defer C.free(unsafe.Pointer(cname))
	return int(C.flif_decoder_decode_file(decoder.dec, cname))
}

var flifDecoderDecodeMemory = func(decoder *flifDecoder, data []byte) int {
	return int(C.flif_decoder_decode_memory(decoder.dec, unsafe.Pointer(&data[0]), C.size_t(len(data))))
}

var flifDecoderNumImages = func(decoder *flifDecoder) int {
	return int(C.flif_decoder_num_images(decoder.dec))
}

var flifDecoderGetNumLoops = func(decoder *flifDecoder) int {
	return int(C.flif_decoder_num_loops(decoder.dec))
}

var flifDecoderGetImage = func(decoder *flifDecoder, index int) *flifImage {
	p := C.flif_decoder_get_image(decoder.dec, C.size_t(index))
	if p == nil {
		return nil
	}
	return &flifImage{p}
}

var flifDestroyDecoder = func(decoder *flifDecoder) {
	C.flif_destroy_decoder(decoder.dec)
}

var flifAbortDecoder = func(decoder *flifDecoder) int {
	return int(C.flif_abort_decoder(decoder.dec))
}

var flifDecoderSetCrcCheck = func(decoder *flifDecoder, crcCheck int) {
	C.flif_decoder_set_crc_check(decoder.dec, C.int32_t(crcCheck))
}

var flifDecoderSetQuality = func(decoder *flifDecoder, quality int) {
	C.flif_decoder_set_quality(decoder.dec, C.int32_t(quality))
}

var flifDecoderSetScale = func(decoder *flifDecoder, scale int) {
	C.flif_decoder_set_scale(decoder.dec, C.uint32_t(scale))
}

var flifDecoderSetResize = func(decoder *flifDecoder, width, height int) {
	C.flif_decoder_set_resize(decoder.dec, C.uint32_t(width), C.uint32_t(height))
}

var flifDecoderSetFit = func(decoder *flifDecoder, width, height int) {
	C.flif_decoder_set_fit(decoder.dec, C.uint32_t(width), C.uint32_t(height))
}

var flifDecoderSetCallback = func() {
	//C.flif_decoder_set_callback(FLIF_DECODER* decoder, uint32_t (*callback)(int32_t quality, int64_t bytes_read))
}

var flifDecoderSetFirstCallbackQuality = func() {
	//C.flif_decoder_set_first_callback_quality(FLIF_DECODER* decoder, int32_t quality) // valid quality: 0-10000
}

var flifReadInfoFromMemory = func(data []byte) *flifInfo {
	p := C.flif_read_info_from_memory(unsafe.Pointer(&data[0]), C.size_t(len(data)))
	if p == nil {
		return nil
	}
	return &flifInfo{p}
}

var flifDestroyInfo = func(info *flifInfo) {
	C.flif_destroy_info(info.info)
}

var flifInfoGetWidth = func(info *flifInfo) int {
	return int(C.flif_info_get_width(info.info))
}

var flifInfoGetHeight = func(info *flifInfo) int {
	return int(C.flif_info_get_height(info.info))
}

var flifInfoGetNbChannels = func(info *flifInfo) int {
	return int(C.flif_info_get_nb_channels(info.info))
}

var flifInfoGetDepth = func(info *flifInfo) int {
	return int(C.flif_info_get_depth(info.info))
}

var flifInfoNumImages = func(info *flifInfo) int {
	return int(C.flif_info_num_images(info.info))
}
