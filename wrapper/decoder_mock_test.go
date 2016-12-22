// Copyright 2016 chrisfelesoid. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package wrapper

import "github.com/chrisfelesoid/go-flif/internal/wrapper"

type mockDecoderWrapper struct {
	flifCreateDecoder       func() *wrapper.CflifDecoder
	flifDecoderDecodeFile   func(decoder *wrapper.CflifDecoder, filename string) int
	flifDecoderDecodeMemory func(decoder *wrapper.CflifDecoder, data []byte) int
	flifDecoderNumImages    func(decoder *wrapper.CflifDecoder) int
	flifDecoderGetNumLoops  func(decoder *wrapper.CflifDecoder) int
	flifDecoderGetImage     func(decoder *wrapper.CflifDecoder, index int) *wrapper.CflifImage
	flifDestroyDecoder      func(decoder *wrapper.CflifDecoder)
	flifAbortDecoder        func(decoder *wrapper.CflifDecoder) int
	flifDecoderSetCrcCheck  func(decoder *wrapper.CflifDecoder, crcCheck int)
	flifDecoderSetQuality   func(decoder *wrapper.CflifDecoder, quality int)
	flifDecoderSetScale     func(decoder *wrapper.CflifDecoder, scale int)
	flifDecoderSetResize    func(decoder *wrapper.CflifDecoder, width, height int)
	flifDecoderSetFit       func(decoder *wrapper.CflifDecoder, width, height int)
	// flifDecoderSetCallback func()
	// flifDecoderSetFirstCallbackQuality func()
	flifReadInfoFromMemory func(data []byte) *wrapper.CflifInfo
	flifDestroyInfo        func(info *wrapper.CflifInfo)
	flifInfoGetWidth       func(info *wrapper.CflifInfo) int
	flifInfoGetHeight      func(info *wrapper.CflifInfo) int
	flifInfoGetNbChannels  func(info *wrapper.CflifInfo) int
	flifInfoGetDepth       func(info *wrapper.CflifInfo) int
	flifInfoNumImages      func(info *wrapper.CflifInfo) int
}

func (m *mockDecoderWrapper) FlifCreateDecoder() *wrapper.CflifDecoder {
	return m.flifCreateDecoder()
}

func (m *mockDecoderWrapper) FlifDecoderDecodeFile(decoder *wrapper.CflifDecoder, filename string) int {
	return m.flifDecoderDecodeFile(decoder, filename)
}

func (m *mockDecoderWrapper) FlifDecoderDecodeMemory(decoder *wrapper.CflifDecoder, data []byte) int {
	return m.flifDecoderDecodeMemory(decoder, data)
}

func (m *mockDecoderWrapper) FlifDecoderNumImages(decoder *wrapper.CflifDecoder) int {
	return m.flifDecoderNumImages(decoder)
}

func (m *mockDecoderWrapper) FlifDecoderGetNumLoops(decoder *wrapper.CflifDecoder) int {
	return m.flifDecoderGetNumLoops(decoder)
}

func (m *mockDecoderWrapper) FlifDecoderGetImage(decoder *wrapper.CflifDecoder, index int) *wrapper.CflifImage {
	return m.flifDecoderGetImage(decoder, index)
}

func (m *mockDecoderWrapper) FlifDestroyDecoder(decoder *wrapper.CflifDecoder) {
	m.flifDestroyDecoder(decoder)
}

func (m *mockDecoderWrapper) FlifAbortDecoder(decoder *wrapper.CflifDecoder) int {
	return m.flifAbortDecoder(decoder)
}

func (m *mockDecoderWrapper) FlifDecoderSetCrcCheck(decoder *wrapper.CflifDecoder, crcCheck int) {
	m.flifDecoderSetCrcCheck(decoder, crcCheck)
}

func (m *mockDecoderWrapper) FlifDecoderSetQuality(decoder *wrapper.CflifDecoder, quality int) {
	m.flifDecoderSetQuality(decoder, quality)
}

func (m *mockDecoderWrapper) FlifDecoderSetScale(decoder *wrapper.CflifDecoder, scale int) {
	m.flifDecoderSetScale(decoder, scale)
}

func (m *mockDecoderWrapper) FlifDecoderSetResize(decoder *wrapper.CflifDecoder, width, height int) {
	m.flifDecoderSetResize(decoder, width, height)
}

func (m *mockDecoderWrapper) FlifDecoderSetFit(decoder *wrapper.CflifDecoder, width, height int) {
	m.flifDecoderSetFit(decoder, width, height)
}

// func (m *mockDecoderWrapper) FlifDecoderSetCallback() {
//     m.flifDecoderSetCallback()
// }

// func (m *mockDecoderWrapper) FlifDecoderSetFirstCallbackQuality() {
//     m.flifDecoderSetFirstCallbackQuality()
// }

func (m *mockDecoderWrapper) FlifReadInfoFromMemory(data []byte) *wrapper.CflifInfo {
	return m.flifReadInfoFromMemory(data)
}

func (m *mockDecoderWrapper) FlifDestroyInfo(info *wrapper.CflifInfo) {
	m.flifDestroyInfo(info)
}

func (m *mockDecoderWrapper) FlifInfoGetWidth(info *wrapper.CflifInfo) int {
	return m.flifInfoGetWidth(info)
}

func (m *mockDecoderWrapper) FlifInfoGetHeight(info *wrapper.CflifInfo) int {
	return m.flifInfoGetHeight(info)
}

func (m *mockDecoderWrapper) FlifInfoGetNbChannels(info *wrapper.CflifInfo) int {
	return m.flifInfoGetNbChannels(info)
}

func (m *mockDecoderWrapper) FlifInfoGetDepth(info *wrapper.CflifInfo) int {
	return m.flifInfoGetDepth(info)
}

func (m *mockDecoderWrapper) FlifInfoNumImages(info *wrapper.CflifInfo) int {
	return m.flifInfoNumImages(info)
}
