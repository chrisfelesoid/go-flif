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

type CencoderFunctionWrapper interface {
	FlifCreateEncoder() *CflifEncoder
	FlifEncoderAddImage(encoder *CflifEncoder, image *CflifImage)
	FlifEncoderEncodeFile(encoder *CflifEncoder, filename string) int
	FlifEncoderEncodeMemory(encoder *CflifEncoder, data *[]byte) int
	FlifDestroyEncoder(encoder *CflifEncoder)
	FlifEncoderSetInterlaced(encoder *CflifEncoder, interlaced int)
	FlifEncoderSetLearnRepeat(encoder *CflifEncoder, learnRepeat int)
	FlifEncoderSetAutoColorBuckets(encoder *CflifEncoder, autoColorBuckets int)
	FlifEncoderSetPaletteSize(encoder *CflifEncoder, paletteSize int)
	FlifEncoderSetLookback(encoder *CflifEncoder, lookback int)
	FlifEncoderSetDivisor(encoder *CflifEncoder, divisor int)
	FlifEncoderSetMinSize(encoder *CflifEncoder, minSize int)
	FlifEncoderSetSplitThreshold(encoder *CflifEncoder, splitThreshold int)
	FlifEncoderSetAlphaZeroLossless(encoder *CflifEncoder)
	FlifEncoderSetChanceCutoff(encoder *CflifEncoder, chanceCutoff int)
	FlifEncoderSetChanceAlpha(encoder *CflifEncoder, chanceAlpha int)
	FlifEncoderSetCrcCheck(encoder *CflifEncoder, crcCheck int)
	FlifEncoderSetChannelCompact(encoder *CflifEncoder, channelCompact int)
	FlifEncoderSetYcocg(encoder *CflifEncoder, ycocg int)
	FlifEncoderSetFrameShape(encoder *CflifEncoder, frameShape int)
	FlifEncoderSetLossy(encoder *CflifEncoder, lossy int)
}

type CencoderWrapper struct {
}

func (e *CencoderWrapper) FlifCreateEncoder() *CflifEncoder {
	p := C.flif_create_encoder()
	if p == nil {
		return nil
	}
	return &CflifEncoder{p}
}

func (e *CencoderWrapper) FlifEncoderAddImage(encoder *CflifEncoder, image *CflifImage) {
	C.flif_encoder_add_image(encoder.enc, image.img)
}

func (e *CencoderWrapper) FlifEncoderEncodeFile(encoder *CflifEncoder, filename string) int {
	cname := C.CString(filename)
	defer C.free(unsafe.Pointer(cname))
	return int(C.flif_encoder_encode_file(encoder.enc, cname))
}

func (e *CencoderWrapper) FlifEncoderEncodeMemory(encoder *CflifEncoder, data *[]byte) int {
	*data = nil
	var p unsafe.Pointer
	var num C.size_t

	stat := int(C.flif_encoder_encode_memory(encoder.enc, &p, &num))
	if stat == 0 {
		return stat
	}
	defer common.FlifFreeMemory(p)

	*data = C.GoBytes(p, C.int(num))
	return stat
}

func (e *CencoderWrapper) FlifDestroyEncoder(encoder *CflifEncoder) {
	C.flif_destroy_encoder(encoder.enc)
}

func (e *CencoderWrapper) FlifEncoderSetInterlaced(encoder *CflifEncoder, interlaced int) {
	C.flif_encoder_set_interlaced(encoder.enc, C.uint32_t(interlaced))
}

func (e *CencoderWrapper) FlifEncoderSetLearnRepeat(encoder *CflifEncoder, learnRepeat int) {
	C.flif_encoder_set_learn_repeat(encoder.enc, C.uint32_t(learnRepeat))
}

func (e *CencoderWrapper) FlifEncoderSetAutoColorBuckets(encoder *CflifEncoder, autoColorBuckets int) {
	C.flif_encoder_set_auto_color_buckets(encoder.enc, C.uint32_t(autoColorBuckets))
}

func (e *CencoderWrapper) FlifEncoderSetPaletteSize(encoder *CflifEncoder, paletteSize int) {
	C.flif_encoder_set_palette_size(encoder.enc, C.int32_t(paletteSize))
}

func (e *CencoderWrapper) FlifEncoderSetLookback(encoder *CflifEncoder, lookback int) {
	C.flif_encoder_set_lookback(encoder.enc, C.int32_t(lookback))
}

func (e *CencoderWrapper) FlifEncoderSetDivisor(encoder *CflifEncoder, divisor int) {
	C.flif_encoder_set_divisor(encoder.enc, C.int32_t(divisor))
}

func (e *CencoderWrapper) FlifEncoderSetMinSize(encoder *CflifEncoder, minSize int) {
	C.flif_encoder_set_min_size(encoder.enc, C.int32_t(minSize))
}

func (e *CencoderWrapper) FlifEncoderSetSplitThreshold(encoder *CflifEncoder, splitThreshold int) {
	C.flif_encoder_set_split_threshold(encoder.enc, C.int32_t(splitThreshold))
}

func (e *CencoderWrapper) FlifEncoderSetAlphaZeroLossless(encoder *CflifEncoder) {
	C.flif_encoder_set_alpha_zero_lossless(encoder.enc)
}

func (e *CencoderWrapper) FlifEncoderSetChanceCutoff(encoder *CflifEncoder, chanceCutoff int) {
	C.flif_encoder_set_chance_cutoff(encoder.enc, C.int32_t(chanceCutoff))
}

func (e *CencoderWrapper) FlifEncoderSetChanceAlpha(encoder *CflifEncoder, chanceAlpha int) {
	C.flif_encoder_set_chance_alpha(encoder.enc, C.int32_t(chanceAlpha))
}

func (e *CencoderWrapper) FlifEncoderSetCrcCheck(encoder *CflifEncoder, crcCheck int) {
	C.flif_encoder_set_crc_check(encoder.enc, C.uint32_t(crcCheck))
}

func (e *CencoderWrapper) FlifEncoderSetChannelCompact(encoder *CflifEncoder, channelCompact int) {
	C.flif_encoder_set_channel_compact(encoder.enc, C.uint32_t(channelCompact))
}

func (e *CencoderWrapper) FlifEncoderSetYcocg(encoder *CflifEncoder, ycocg int) {
	C.flif_encoder_set_ycocg(encoder.enc, C.uint32_t(ycocg))
}

func (e *CencoderWrapper) FlifEncoderSetFrameShape(encoder *CflifEncoder, frameShape int) {
	C.flif_encoder_set_frame_shape(encoder.enc, C.uint32_t(frameShape))
}

func (e *CencoderWrapper) FlifEncoderSetLossy(encoder *CflifEncoder, lossy int) {
	C.flif_encoder_set_lossy(encoder.enc, C.int32_t(lossy))
}
