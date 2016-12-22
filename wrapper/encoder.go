// Copyright 2016 chrisfelesoid. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package wrapper

import (
	"sync"

	"github.com/chrisfelesoid/go-flif/internal/wrapper"
)

type FlifEncoder struct {
	enc               *wrapper.CflifEncoder
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

var encoder wrapper.CencoderFunctionWrapper

func NewFlifEncoder() *FlifEncoder {
	p := encoder.FlifCreateEncoder()
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
		encoder.FlifDestroyEncoder(e.enc)
		e.enc = nil
	})
}

func (e *FlifEncoder) AddImage(image *FlifImage) {
	if image == nil {
		return
	}
	for _, img := range image.images {
		encoder.FlifEncoderAddImage(e.enc, img)
	}
}

func (e *FlifEncoder) EncodeFile(name string) error {
	setEncoderOptions(e)
	return checkResult(encoder.FlifEncoderEncodeFile(e.enc, name))
}

func (e *FlifEncoder) Encode() ([]byte, error) {
	var data []byte
	setEncoderOptions(e)

	err := checkResult(encoder.FlifEncoderEncodeMemory(e.enc, &data))
	if err != nil {
		return nil, err
	}
	return data, nil
}

var setEncoderOptions = func(e *FlifEncoder) {
	interlaced := 0
	crc := 0
	if e.Interlaced {
		interlaced = 1
	}
	if e.CrcCheck {
		crc = 1
	}

	encoder.FlifEncoderSetInterlaced(e.enc, interlaced)
	encoder.FlifEncoderSetLearnRepeat(e.enc, e.LearnRepeat)
	encoder.FlifEncoderSetAutoColorBuckets(e.enc, e.AutoColorBuckets)
	encoder.FlifEncoderSetPaletteSize(e.enc, e.PaletteSize)
	encoder.FlifEncoderSetLookback(e.enc, e.Lookback)
	encoder.FlifEncoderSetDivisor(e.enc, e.Divisor)
	encoder.FlifEncoderSetMinSize(e.enc, e.MinSize)
	encoder.FlifEncoderSetSplitThreshold(e.enc, e.SplitThreashold)
	if e.AlphaZeroLossless {
		encoder.FlifEncoderSetAlphaZeroLossless(e.enc)
	}
	encoder.FlifEncoderSetChanceCutoff(e.enc, e.ChanceCutoff)
	encoder.FlifEncoderSetChanceAlpha(e.enc, e.ChanceAlpha)
	encoder.FlifEncoderSetCrcCheck(e.enc, crc)
	encoder.FlifEncoderSetChannelCompact(e.enc, e.ChannelCompact)
	encoder.FlifEncoderSetYcocg(e.enc, e.YCoCg)
	encoder.FlifEncoderSetFrameShape(e.enc, e.FrameShape)
	encoder.FlifEncoderSetLossy(e.enc, e.Lossy)
}

func init() {
	encoder = &wrapper.CencoderWrapper{}
}
