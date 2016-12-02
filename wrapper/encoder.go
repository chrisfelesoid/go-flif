// Copyright 2016 chrisfelesoid. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package wrapper

import (
	"fmt"
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

func NewFlifEncoder() *FlifEncoder {
	p := wrapper.CflifCreateEncoder()
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
		wrapper.CflifDestroyEncoder(e.enc)
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

	wrapper.CflifEncoderSetInterlaced(e.enc, interlaced)
	wrapper.CflifEncoderSetLearnRepeat(e.enc, e.LearnRepeat)
	wrapper.CflifEncoderSetAutoColorBuckets(e.enc, e.AutoColorBuckets)
	wrapper.CflifEncoderSetPaletteSize(e.enc, e.PaletteSize)
	wrapper.CflifEncoderSetLookback(e.enc, e.Lookback)
	wrapper.CflifEncoderSetDivisor(e.enc, e.Divisor)
	wrapper.CflifEncoderSetMinSize(e.enc, e.MinSize)
	wrapper.CflifEncoderSetSplitThreshold(e.enc, e.SplitThreashold)
	if e.AlphaZeroLossless {
		wrapper.CflifEncoderSetAlphaZeroLossless(e.enc)
	}
	wrapper.CflifEncoderSetChanceCutoff(e.enc, e.ChanceCutoff)
	wrapper.CflifEncoderSetChanceAlpha(e.enc, e.ChanceAlpha)
	wrapper.CflifEncoderSetCrcCheck(e.enc, crc)
	wrapper.CflifEncoderSetChannelCompact(e.enc, e.ChannelCompact)
	wrapper.CflifEncoderSetYcocg(e.enc, e.YCoCg)
	wrapper.CflifEncoderSetFrameShape(e.enc, e.FrameShape)
	wrapper.CflifEncoderSetLossy(e.enc, e.Lossy)
}

func (e *FlifEncoder) AddImage(image *FlifImage) {
	for _, img := range image.images {
		wrapper.CflifEncoderAddImage(e.enc, img)
	}
}

func (e *FlifEncoder) EncodeFile(name string) error {
	e.setOptions()
	st := wrapper.CflifEncoderEncodeFile(e.enc, name)
	if st == 0 {
		return fmt.Errorf("encode error: %d", st)
	}
	return nil
}

func (e *FlifEncoder) Encode() ([]byte, error) {
	var data []byte
	e.setOptions()
	st := wrapper.CflifEncoderEncodeMemory(e.enc, &data)
	if st == 0 {
		return nil, fmt.Errorf("encode error: %d", st)
	}
	return data, nil
}
