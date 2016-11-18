// Copyright 2016 chrisfelesoid. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package wrapper

import "sync"

type FlifEncoder struct {
	enc               *flifEncoder
	once              sync.Once
	Interlaced        bool
	LearnRepeat       int
	AutoColorBuckets  int
	PaletteSize       int
	Lookback          int
	Divisor           int
	MinSize           int
	SplitThreashold   int
	AlphaZeroLossless bool
	ChanceCutoff      int
	ChanceAlpha       int
	CrcCheck          bool
	ChannelCompact    int
	YCoCg             int
	FrameShape        int
	Lossy             int
}

func NewFlifEncoder() *FlifEncoder {
	p := flifCreateEncoder()
	if p == nil {
		return nil
	}
	return &FlifEncoder{
		enc:               p,
		Interlaced:        true,
		LearnRepeat:       2, //TREE_LEARN_REPEATS
		AutoColorBuckets:  1,
		PaletteSize:       512, //DEFAULT_MAX_PALETTE_SIZE
		Lookback:          1,
		Divisor:           30,           //CONTEXT_TREE_COUNT_DIV
		MinSize:           50,           //CONTEXT_TREE_MIN_SUBTREE_SIZE
		SplitThreashold:   5461 * 8 * 8, //CONTEXT_TREE_SPLIT_THRESHOLD
		AlphaZeroLossless: true,
		ChanceCutoff:      2,
		ChanceAlpha:       19,
		CrcCheck:          false,
		ChannelCompact:    1,
		YCoCg:             1,
		FrameShape:        1,
		Lossy:             0,
	}
}

func (e *FlifEncoder) Destroy() {
	if e.enc == nil {
		return
	}

	e.once.Do(func() {
		flifDestroyEncoder(e.enc)
	})
}

func (e *FlifEncoder) setOptions() {
	interlaced := 0
	crc := 0
	if e.Interlaced {
		interlaced = 1
	}
	if e.CrcCheck {
		crc = 1
	}

	flifEncoderSetInterlaced(e.enc, interlaced)
	flifEncoderSetLearnRepeat(e.enc, e.LearnRepeat)
	flifEncoderSetAutoColorBuckets(e.enc, e.AutoColorBuckets)
	flifEncoderSetPaletteSize(e.enc, e.PaletteSize)
	flifEncoderSetLookback(e.enc, e.Lookback)
	flifEncoderSetDivisor(e.enc, e.Divisor)
	flifEncoderSetMinSize(e.enc, e.MinSize)
	flifEncoderSetSplitThreshold(e.enc, e.SplitThreashold)
	if e.AlphaZeroLossless {
		flifEncoderSetAlphaZeroLossless(e.enc)
	}
	flifEncoderSetChanceCutoff(e.enc, e.ChanceCutoff)
	flifEncoderSetChanceAlpha(e.enc, e.ChanceAlpha)
	flifEncoderSetCrcCheck(e.enc, crc)
	flifEncoderSetChannelCompact(e.enc, e.ChannelCompact)
	flifEncoderSetYcocg(e.enc, e.YCoCg)
	flifEncoderSetFrameShape(e.enc, e.FrameShape)
	flifEncoderSetLossy(e.enc, e.Lossy)
}

func (e *FlifEncoder) AddImage(image *FlifImage) {
	for _, img := range image.images {
		flifEncoderAddImage(e.enc, img)
	}
}

func (e *FlifEncoder) EncodeFile(name string) int {
	e.setOptions()
	return flifEncoderEncodeFile(e.enc, name)
}

func (e *FlifEncoder) EncodeMemory(data *[]byte) int {
	e.setOptions()
	return flifEncoderEncodeMemory(e.enc, data)
}
