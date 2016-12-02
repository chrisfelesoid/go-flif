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

type CflifEncoder struct {
	enc *C.FLIF_ENCODER
}

func CflifCreateEncoder() *CflifEncoder {
	p := C.flif_create_encoder()
	if p == nil {
		return nil
	}
	return &CflifEncoder{p}
}

func CflifEncoderAddImage(encoder *CflifEncoder, image *CflifImage) {
	C.flif_encoder_add_image(encoder.enc, image.img)
}

func CflifEncoderEncodeFile(encoder *CflifEncoder, filename string) int {
	cname := C.CString(filename)
	defer C.free(unsafe.Pointer(cname))
	return int(C.flif_encoder_encode_file(encoder.enc, cname))
}

func CflifEncoderEncodeMemory(encoder *CflifEncoder, data *[]byte) int {
	*data = nil
	var p unsafe.Pointer
	var num C.size_t

	stat := int(C.flif_encoder_encode_memory(encoder.enc, &p, &num))
	if stat == 0 {
		return stat
	}
	defer CflifFreeMemory(p)

	*data = C.GoBytes(p, C.int(num))
	return stat
}

func CflifDestroyEncoder(encoder *CflifEncoder) {
	C.flif_destroy_encoder(encoder.enc)
}

func CflifEncoderSetInterlaced(encoder *CflifEncoder, interlaced int) {
	C.flif_encoder_set_interlaced(encoder.enc, C.uint32_t(interlaced))
}

func CflifEncoderSetLearnRepeat(encoder *CflifEncoder, learnRepeat int) {
	C.flif_encoder_set_learn_repeat(encoder.enc, C.uint32_t(learnRepeat))
}

func CflifEncoderSetAutoColorBuckets(encoder *CflifEncoder, autoColorBuckets int) {
	C.flif_encoder_set_auto_color_buckets(encoder.enc, C.uint32_t(autoColorBuckets))
}

func CflifEncoderSetPaletteSize(encoder *CflifEncoder, paletteSize int) {
	C.flif_encoder_set_palette_size(encoder.enc, C.int32_t(paletteSize))
}

func CflifEncoderSetLookback(encoder *CflifEncoder, lookback int) {
	C.flif_encoder_set_lookback(encoder.enc, C.int32_t(lookback))
}

func CflifEncoderSetDivisor(encoder *CflifEncoder, divisor int) {
	C.flif_encoder_set_divisor(encoder.enc, C.int32_t(divisor))
}

func CflifEncoderSetMinSize(encoder *CflifEncoder, minSize int) {
	C.flif_encoder_set_min_size(encoder.enc, C.int32_t(minSize))
}

func CflifEncoderSetSplitThreshold(encoder *CflifEncoder, splitThreshold int) {
	C.flif_encoder_set_split_threshold(encoder.enc, C.int32_t(splitThreshold))
}

func CflifEncoderSetAlphaZeroLossless(encoder *CflifEncoder) {
	C.flif_encoder_set_alpha_zero_lossless(encoder.enc)
}

func CflifEncoderSetChanceCutoff(encoder *CflifEncoder, chanceCutoff int) {
	C.flif_encoder_set_chance_cutoff(encoder.enc, C.int32_t(chanceCutoff))
}

func CflifEncoderSetChanceAlpha(encoder *CflifEncoder, chanceAlpha int) {
	C.flif_encoder_set_chance_alpha(encoder.enc, C.int32_t(chanceAlpha))
}

func CflifEncoderSetCrcCheck(encoder *CflifEncoder, crCheck int) {
	C.flif_encoder_set_crc_check(encoder.enc, C.uint32_t(crCheck))
}

func CflifEncoderSetChannelCompact(encoder *CflifEncoder, channelCompact int) {
	C.flif_encoder_set_channel_compact(encoder.enc, C.uint32_t(channelCompact))
}

func CflifEncoderSetYcocg(encoder *CflifEncoder, ycocg int) {
	C.flif_encoder_set_ycocg(encoder.enc, C.uint32_t(ycocg))
}

func CflifEncoderSetFrameShape(encoder *CflifEncoder, frameShape int) {
	C.flif_encoder_set_frame_shape(encoder.enc, C.uint32_t(frameShape))
}

func CflifEncoderSetLossy(encoder *CflifEncoder, lossy int) {
	C.flif_encoder_set_lossy(encoder.enc, C.int32_t(lossy))
}
