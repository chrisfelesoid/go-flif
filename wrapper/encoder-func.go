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

type flifEncoder struct {
	enc *C.FLIF_ENCODER
}

var flifCreateEncoder = func() *flifEncoder {
	p := C.flif_create_encoder()
	if p == nil {
		return nil
	}
	return &flifEncoder{p}
}

var flifEncoderAddImage = func(encoder *flifEncoder, image *flifImage) {
	C.flif_encoder_add_image(encoder.enc, image.img)
}

var flifEncoderEncodeFile = func(encoder *flifEncoder, filename string) int {
	cname := C.CString(filename)
	defer C.free(unsafe.Pointer(cname))
	return int(C.flif_encoder_encode_file(encoder.enc, cname))
}

var flifEncoderEncodeMemory = func(encoder *flifEncoder, data *[]byte) int {
	*data = nil
	var p unsafe.Pointer
	var num C.size_t

	stat := int(C.flif_encoder_encode_memory(encoder.enc, &p, &num))
	if stat == 0 {
		return stat
	}
	defer flifFreeMemory(p)

	*data = C.GoBytes(p, C.int(num))
	return stat
}

var flifDestroyEncoder = func(encoder *flifEncoder) {
	C.flif_destroy_encoder(encoder.enc)
}

var flifEncoderSetInterlaced = func(encoder *flifEncoder, interlaced int) {
	C.flif_encoder_set_interlaced(encoder.enc, C.uint32_t(interlaced))
}

var flifEncoderSetLearnRepeat = func(encoder *flifEncoder, learnRepeat int) {
	C.flif_encoder_set_learn_repeat(encoder.enc, C.uint32_t(learnRepeat))
}

var flifEncoderSetAutoColorBuckets = func(encoder *flifEncoder, autoColorBuckets int) {
	C.flif_encoder_set_auto_color_buckets(encoder.enc, C.uint32_t(autoColorBuckets))
}

var flifEncoderSetPaletteSize = func(encoder *flifEncoder, paletteSize int) {
	C.flif_encoder_set_palette_size(encoder.enc, C.int32_t(paletteSize))
}

var flifEncoderSetLookback = func(encoder *flifEncoder, lookback int) {
	C.flif_encoder_set_lookback(encoder.enc, C.int32_t(lookback))
}

var flifEncoderSetDivisor = func(encoder *flifEncoder, divisor int) {
	C.flif_encoder_set_divisor(encoder.enc, C.int32_t(divisor))
}

var flifEncoderSetMinSize = func(encoder *flifEncoder, minSize int) {
	C.flif_encoder_set_min_size(encoder.enc, C.int32_t(minSize))
}

var flifEncoderSetSplitThreshold = func(encoder *flifEncoder, splitThreshold int) {
	C.flif_encoder_set_split_threshold(encoder.enc, C.int32_t(splitThreshold))
}

var flifEncoderSetAlphaZeroLossless = func(encoder *flifEncoder) {
	C.flif_encoder_set_alpha_zero_lossless(encoder.enc)
}

var flifEncoderSetChanceCutoff = func(encoder *flifEncoder, chanceCutoff int) {
	C.flif_encoder_set_chance_cutoff(encoder.enc, C.int32_t(chanceCutoff))
}

var flifEncoderSetChanceAlpha = func(encoder *flifEncoder, chanceAlpha int) {
	C.flif_encoder_set_chance_alpha(encoder.enc, C.int32_t(chanceAlpha))
}

var flifEncoderSetCrcCheck = func(encoder *flifEncoder, crcCheck int) {
	C.flif_encoder_set_crc_check(encoder.enc, C.uint32_t(crcCheck))
}

var flifEncoderSetChannelCompact = func(encoder *flifEncoder, channelCompact int) {
	C.flif_encoder_set_channel_compact(encoder.enc, C.uint32_t(channelCompact))
}

var flifEncoderSetYcocg = func(encoder *flifEncoder, ycocg int) {
	C.flif_encoder_set_ycocg(encoder.enc, C.uint32_t(ycocg))
}

var flifEncoderSetFrameShape = func(encoder *flifEncoder, frameShape int) {
	C.flif_encoder_set_frame_shape(encoder.enc, C.uint32_t(frameShape))
}

var flifEncoderSetLossy = func(encoder *flifEncoder, lossy int) {
	C.flif_encoder_set_lossy(encoder.enc, C.int32_t(lossy))
}
