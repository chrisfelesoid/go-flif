// Copyright 2016 chrisfelesoid. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package wrapper

import "github.com/chrisfelesoid/go-flif/internal/wrapper"

type mockEncoderWrapper struct {
	flifCreateEncoder               func() *wrapper.CflifEncoder
	flifEncoderAddImage             func(encoder *wrapper.CflifEncoder, image *wrapper.CflifImage)
	flifEncoderEncodeFile           func(encoder *wrapper.CflifEncoder, filename string) int
	flifEncoderEncodeMemory         func(encoder *wrapper.CflifEncoder, data *[]byte) int
	flifDestroyEncoder              func(encoder *wrapper.CflifEncoder)
	flifEncoderSetInterlaced        func(encoder *wrapper.CflifEncoder, interlaced int)
	flifEncoderSetLearnRepeat       func(encoder *wrapper.CflifEncoder, learnRepeat int)
	flifEncoderSetAutoColorBuckets  func(encoder *wrapper.CflifEncoder, autoColorBuckets int)
	flifEncoderSetPaletteSize       func(encoder *wrapper.CflifEncoder, paletteSize int)
	flifEncoderSetLookback          func(encoder *wrapper.CflifEncoder, lookback int)
	flifEncoderSetDivisor           func(encoder *wrapper.CflifEncoder, divisor int)
	flifEncoderSetMinSize           func(encoder *wrapper.CflifEncoder, minSize int)
	flifEncoderSetSplitThreshold    func(encoder *wrapper.CflifEncoder, splitThreshold int)
	flifEncoderSetAlphaZeroLossless func(encoder *wrapper.CflifEncoder)
	flifEncoderSetChanceCutoff      func(encoder *wrapper.CflifEncoder, chanceCutoff int)
	flifEncoderSetChanceAlpha       func(encoder *wrapper.CflifEncoder, chanceAlpha int)
	flifEncoderSetCrcCheck          func(encoder *wrapper.CflifEncoder, crcCheck int)
	flifEncoderSetChannelCompact    func(encoder *wrapper.CflifEncoder, channelCompact int)
	flifEncoderSetYcocg             func(encoder *wrapper.CflifEncoder, ycocg int)
	flifEncoderSetFrameShape        func(encoder *wrapper.CflifEncoder, frameShape int)
	flifEncoderSetLossy             func(encoder *wrapper.CflifEncoder, lossy int)
}

func (m *mockEncoderWrapper) FlifCreateEncoder() *wrapper.CflifEncoder {
	return m.flifCreateEncoder()
}

func (m *mockEncoderWrapper) FlifEncoderAddImage(encoder *wrapper.CflifEncoder, image *wrapper.CflifImage) {
	m.flifEncoderAddImage(encoder, image)
}

func (m *mockEncoderWrapper) FlifEncoderEncodeFile(encoder *wrapper.CflifEncoder, filename string) int {
	return m.flifEncoderEncodeFile(encoder, filename)
}

func (m *mockEncoderWrapper) FlifEncoderEncodeMemory(encoder *wrapper.CflifEncoder, data *[]byte) int {
	return m.flifEncoderEncodeMemory(encoder, data)
}

func (m *mockEncoderWrapper) FlifDestroyEncoder(encoder *wrapper.CflifEncoder) {
	m.flifDestroyEncoder(encoder)
}

func (m *mockEncoderWrapper) FlifEncoderSetInterlaced(encoder *wrapper.CflifEncoder, interlaced int) {
	m.flifEncoderSetInterlaced(encoder, interlaced)
}

func (m *mockEncoderWrapper) FlifEncoderSetLearnRepeat(encoder *wrapper.CflifEncoder, learnRepeat int) {
	m.flifEncoderSetLearnRepeat(encoder, learnRepeat)
}

func (m *mockEncoderWrapper) FlifEncoderSetAutoColorBuckets(encoder *wrapper.CflifEncoder, autoColorBuckets int) {
	m.flifEncoderSetAutoColorBuckets(encoder, autoColorBuckets)
}

func (m *mockEncoderWrapper) FlifEncoderSetPaletteSize(encoder *wrapper.CflifEncoder, paletteSize int) {
	m.flifEncoderSetPaletteSize(encoder, paletteSize)
}

func (m *mockEncoderWrapper) FlifEncoderSetLookback(encoder *wrapper.CflifEncoder, lookback int) {
	m.flifEncoderSetLookback(encoder, lookback)
}

func (m *mockEncoderWrapper) FlifEncoderSetDivisor(encoder *wrapper.CflifEncoder, divisor int) {
	m.flifEncoderSetDivisor(encoder, divisor)
}

func (m *mockEncoderWrapper) FlifEncoderSetMinSize(encoder *wrapper.CflifEncoder, minSize int) {
	m.flifEncoderSetMinSize(encoder, minSize)
}

func (m *mockEncoderWrapper) FlifEncoderSetSplitThreshold(encoder *wrapper.CflifEncoder, splitThreshold int) {
	m.flifEncoderSetSplitThreshold(encoder, splitThreshold)
}

func (m *mockEncoderWrapper) FlifEncoderSetAlphaZeroLossless(encoder *wrapper.CflifEncoder) {
	m.flifEncoderSetAlphaZeroLossless(encoder)
}

func (m *mockEncoderWrapper) FlifEncoderSetChanceCutoff(encoder *wrapper.CflifEncoder, chanceCutoff int) {
	m.flifEncoderSetChanceCutoff(encoder, chanceCutoff)
}

func (m *mockEncoderWrapper) FlifEncoderSetChanceAlpha(encoder *wrapper.CflifEncoder, chanceAlpha int) {
	m.flifEncoderSetChanceAlpha(encoder, chanceAlpha)
}

func (m *mockEncoderWrapper) FlifEncoderSetCrcCheck(encoder *wrapper.CflifEncoder, crcCheck int) {
	m.flifEncoderSetCrcCheck(encoder, crcCheck)
}

func (m *mockEncoderWrapper) FlifEncoderSetChannelCompact(encoder *wrapper.CflifEncoder, channelCompact int) {
	m.flifEncoderSetChannelCompact(encoder, channelCompact)
}

func (m *mockEncoderWrapper) FlifEncoderSetYcocg(encoder *wrapper.CflifEncoder, ycocg int) {
	m.flifEncoderSetYcocg(encoder, ycocg)
}

func (m *mockEncoderWrapper) FlifEncoderSetFrameShape(encoder *wrapper.CflifEncoder, frameShape int) {
	m.flifEncoderSetFrameShape(encoder, frameShape)
}

func (m *mockEncoderWrapper) FlifEncoderSetLossy(encoder *wrapper.CflifEncoder, lossy int) {
	m.flifEncoderSetLossy(encoder, lossy)
}
