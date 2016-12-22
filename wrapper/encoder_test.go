// Copyright 2016 chrisfelesoid. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package wrapper

import (
	"testing"

	"reflect"

	"github.com/chrisfelesoid/go-flif/internal/wrapper"
)

func TestNewFlifEncoder(t *testing.T) {
	mock := &mockEncoderWrapper{}
	encoder = mock
	defer func() {
		encoder = &wrapper.CencoderWrapper{}
	}()

	enc := (*wrapper.CflifEncoder)(nil)
	mock.flifCreateEncoder = func() *wrapper.CflifEncoder {
		return enc
	}

	// error in FlifCreateEncoder
	fe := NewFlifEncoder()
	if fe != nil {
		t.Errorf("want %v, got %v", nil, fe)
	}

	// success
	enc = &wrapper.CflifEncoder{}
	fe = NewFlifEncoder()
	if fe == nil {
		t.Errorf("must not nil")
	}
	if fe.enc != enc {
		t.Errorf("want %v, got %v", enc, fe.enc)
	}
	if fe.Interlaced != true {
		t.Errorf("want %v, got %v", true, fe.Interlaced)
	}
	if fe.LearnRepeat != 2 {
		t.Errorf("want %v, got %v", 2, fe.LearnRepeat)
	}
	if fe.AutoColorBuckets != 1 {
		t.Errorf("want %v, got %v", 1, fe.AutoColorBuckets)
	}
	if fe.PaletteSize != 512 {
		t.Errorf("want %v, got %v", 512, fe.PaletteSize)
	}
	if fe.Lookback != 1 {
		t.Errorf("want %v, got %v", 1, fe.Lookback)
	}
	if fe.Divisor != 30 {
		t.Errorf("want %v, got %v", 30, fe.Divisor)
	}
	if fe.MinSize != 50 {
		t.Errorf("want %v, got %v", 50, fe.MinSize)
	}
	if fe.SplitThreashold != 5461*8*8 {
		t.Errorf("want %v, got %v", 5461*8*8, fe.SplitThreashold)
	}
	if fe.AlphaZeroLossless != true {
		t.Errorf("want %v, got %v", true, fe.AlphaZeroLossless)
	}
	if fe.ChanceCutoff != 2 {
		t.Errorf("want %v, got %v", 2, fe.ChanceCutoff)
	}
	if fe.ChanceAlpha != 19 {
		t.Errorf("want %v, got %v", 19, fe.ChanceAlpha)
	}
	if fe.CrcCheck != false {
		t.Errorf("want %v, got %v", false, fe.CrcCheck)
	}
	if fe.ChannelCompact != 1 {
		t.Errorf("want %v, got %v", 1, fe.ChannelCompact)
	}
	if fe.YCoCg != 1 {
		t.Errorf("want %v, got %v", 1, fe.YCoCg)
	}
	if fe.FrameShape != 1 {
		t.Errorf("want %v, got %v", 1, fe.FrameShape)
	}
	if fe.Lossy != 0 {
		t.Errorf("want %v, got %v", 0, fe.Lossy)
	}
}

func TestFlifEncoder_Destroy(t *testing.T) {
	mock := &mockEncoderWrapper{}
	encoder = mock
	defer func() {
		encoder = &wrapper.CencoderWrapper{}
	}()

	enc := (*wrapper.CflifEncoder)(nil)
	call := 0
	mock.flifDestroyEncoder = func(encoder *wrapper.CflifEncoder) {
		enc = encoder
		call++
	}

	// no encoder
	fe := &FlifEncoder{}
	fe.Destroy()
	if call != 0 {
		t.Errorf("want %v, got %v", 0, call)
	}

	// success
	enc = nil
	call = 0
	cenc := &wrapper.CflifEncoder{}
	fe = &FlifEncoder{enc: cenc}
	fe.Destroy()
	if enc != cenc {
		t.Errorf("want %v, got %v", cenc, enc)
	}
	if call != 1 {
		t.Errorf("want %v, got %v", 1, call)
	}
	if fe.enc != nil {
		t.Errorf("want %v, got %v", nil, fe.enc)
	}
	// retry
	fe.enc = cenc
	fe.Destroy()
	if call != 1 {
		t.Errorf("want %v, got %v", 1, call)
	}
}

func TestFlifEncoder_AddImage(t *testing.T) {
	mock := &mockEncoderWrapper{}
	encoder = mock
	defer func() {
		encoder = &wrapper.CencoderWrapper{}
	}()

	enc := (*wrapper.CflifEncoder)(nil)
	imgs := ([]*wrapper.CflifImage)(nil)
	call := 0
	mock.flifEncoderAddImage = func(encoder *wrapper.CflifEncoder, image *wrapper.CflifImage) {
		enc = encoder
		imgs = append(imgs, image)
		call++
	}

	// error nil image
	fe := &FlifEncoder{}
	fe.AddImage(nil)
	if call != 0 {
		t.Errorf("want %v, got %v", 0, call)
	}

	// error no image
	call = 0
	fe = &FlifEncoder{}
	fi := &FlifImage{}
	fe.AddImage(fi)
	if call != 0 {
		t.Errorf("want %v, got %v", 0, call)
	}

	// success
	enc = nil
	imgs = nil
	call = 0
	cenc := &wrapper.CflifEncoder{}
	fis := []*wrapper.CflifImage{
		&wrapper.CflifImage{},
		&wrapper.CflifImage{},
		&wrapper.CflifImage{},
	}
	fe = &FlifEncoder{enc: cenc}
	fi = &FlifImage{images: fis}
	fe.AddImage(fi)
	if call != len(fis) {
		t.Errorf("want %v, got %v", len(fis), call)
	}
	if !reflect.DeepEqual(imgs, fis) {
		t.Errorf("want %v, got %v", fis, imgs)
	}
	if enc != cenc {
		t.Errorf("want %v, got %v", cenc, enc)
	}
}

func TestFlifEncoder_EncodeFile(t *testing.T) {
	bkSetOptions := setEncoderOptions
	mock := &mockEncoderWrapper{}
	encoder = mock
	defer func() {
		encoder = &wrapper.CencoderWrapper{}
		setEncoderOptions = bkSetOptions
	}()
	ocall := 0
	setEncoderOptions = func(e *FlifEncoder) {
		ocall++
	}

	enc := (*wrapper.CflifEncoder)(nil)
	fn := ""
	stat := 0
	mock.flifEncoderEncodeFile = func(encoder *wrapper.CflifEncoder, filename string) int {
		enc = encoder
		fn = filename
		return stat
	}

	// error in FlifEncoderEncodeFile
	stat = 0
	name := "hoge"
	fe := &FlifEncoder{}
	err := fe.EncodeFile(name)
	if err != ErrUnknown {
		t.Errorf("another error occurred: %v", err)
	}
	if ocall != 1 {
		t.Errorf("want %v, got %v", 1, ocall)
	}

	// success
	ocall = 0
	stat = 1
	name = "hoge"
	cenc := &wrapper.CflifEncoder{}
	fe = &FlifEncoder{enc: cenc}
	err = fe.EncodeFile(name)
	if err != nil {
		t.Errorf("error occurred: %v", err)
	}
	if ocall != 1 {
		t.Errorf("want %v, got %v", 1, ocall)
	}
	if enc != cenc {
		t.Errorf("want %v, got %v", cenc, enc)
	}
	if fn != name {
		t.Errorf("want %v, got %v", name, fn)
	}
}

func TestFlifEncoder_Encode(t *testing.T) {
	bkSetOptions := setEncoderOptions
	mock := &mockEncoderWrapper{}
	encoder = mock
	defer func() {
		encoder = &wrapper.CencoderWrapper{}
		setEncoderOptions = bkSetOptions
	}()
	ocall := 0
	setEncoderOptions = func(e *FlifEncoder) {
		ocall++
	}

	enc := (*wrapper.CflifEncoder)(nil)
	rd := ([]byte)(nil)
	stat := 0
	mock.flifEncoderEncodeMemory = func(encoder *wrapper.CflifEncoder, data *[]byte) int {
		enc = encoder
		*data = rd
		return stat
	}

	// error in FlifEncoderEncodeMemory
	stat = 0
	fe := &FlifEncoder{}
	d, err := fe.Encode()
	if err != ErrUnknown {
		t.Errorf("another error occurred: %v", err)
	}
	if d != nil {
		t.Errorf("want %v, got %v", nil, d)
	}
	if ocall != 1 {
		t.Errorf("want %v, got %v", 1, ocall)
	}

	// success
	ocall = 0
	enc = nil
	rd = []byte{0x11, 0x22}
	stat = 1
	cenc := &wrapper.CflifEncoder{}
	fe = &FlifEncoder{enc: cenc}
	d, err = fe.Encode()
	if err != nil {
		t.Errorf("error occurred: %v", err)
	}
	if enc != cenc {
		t.Errorf("want %v, got %v", cenc, enc)
	}
	if !reflect.DeepEqual(d, rd) {
		t.Errorf("want %v, got %v", rd, d)
	}
}

func TestSetEncoderOptions(t *testing.T) {
	mock := &mockEncoderWrapper{}
	encoder = mock
	defer func() {
		encoder = &wrapper.CencoderWrapper{}
	}()

	// TODO: too many...
	sinenc := (*wrapper.CflifEncoder)(nil)
	sin := -1
	sleenc := (*wrapper.CflifEncoder)(nil)
	sle := -1
	sauenc := (*wrapper.CflifEncoder)(nil)
	sau := -1
	spaenc := (*wrapper.CflifEncoder)(nil)
	spa := -1
	sloenc := (*wrapper.CflifEncoder)(nil)
	slo := -1
	sdienc := (*wrapper.CflifEncoder)(nil)
	sdi := -1
	smienc := (*wrapper.CflifEncoder)(nil)
	smi := -1
	sspenc := (*wrapper.CflifEncoder)(nil)
	ssp := -1
	salenc := (*wrapper.CflifEncoder)(nil)
	salcall := 0
	sccutenc := (*wrapper.CflifEncoder)(nil)
	sccut := -1
	scalenc := (*wrapper.CflifEncoder)(nil)
	scal := -1
	screnc := (*wrapper.CflifEncoder)(nil)
	scr := -1
	sccoenc := (*wrapper.CflifEncoder)(nil)
	scco := -1
	sycenc := (*wrapper.CflifEncoder)(nil)
	syc := -1
	sfrenc := (*wrapper.CflifEncoder)(nil)
	sfr := -1
	slosenc := (*wrapper.CflifEncoder)(nil)
	slos := -1

	reset := func() {
		sinenc = (*wrapper.CflifEncoder)(nil)
		sin = -1
		sleenc = (*wrapper.CflifEncoder)(nil)
		sle = -1
		sauenc = (*wrapper.CflifEncoder)(nil)
		sau = -1
		spaenc = (*wrapper.CflifEncoder)(nil)
		spa = -1
		sloenc = (*wrapper.CflifEncoder)(nil)
		slo = -1
		sdienc = (*wrapper.CflifEncoder)(nil)
		sdi = -1
		smienc = (*wrapper.CflifEncoder)(nil)
		smi = -1
		sspenc = (*wrapper.CflifEncoder)(nil)
		ssp = -1
		salenc = (*wrapper.CflifEncoder)(nil)
		salcall = 0
		sccutenc = (*wrapper.CflifEncoder)(nil)
		sccut = -1
		scalenc = (*wrapper.CflifEncoder)(nil)
		scal = -1
		screnc = (*wrapper.CflifEncoder)(nil)
		scr = -1
		sccoenc = (*wrapper.CflifEncoder)(nil)
		scco = -1
		sycenc = (*wrapper.CflifEncoder)(nil)
		syc = -1
		sfrenc = (*wrapper.CflifEncoder)(nil)
		sfr = -1
		slosenc = (*wrapper.CflifEncoder)(nil)
		slos = -1
	}

	mock.flifEncoderSetInterlaced = func(encoder *wrapper.CflifEncoder, interlaced int) {
		sinenc = encoder
		sin = interlaced
	}
	mock.flifEncoderSetLearnRepeat = func(encoder *wrapper.CflifEncoder, learnRepeat int) {
		sleenc = encoder
		sle = learnRepeat
	}
	mock.flifEncoderSetAutoColorBuckets = func(encoder *wrapper.CflifEncoder, autoColorBuckets int) {
		sauenc = encoder
		sau = autoColorBuckets
	}
	mock.flifEncoderSetPaletteSize = func(encoder *wrapper.CflifEncoder, paletteSize int) {
		spaenc = encoder
		spa = paletteSize
	}
	mock.flifEncoderSetLookback = func(encoder *wrapper.CflifEncoder, lookback int) {
		sloenc = encoder
		slo = lookback
	}
	mock.flifEncoderSetDivisor = func(encoder *wrapper.CflifEncoder, divisor int) {
		sdienc = encoder
		sdi = divisor
	}
	mock.flifEncoderSetMinSize = func(encoder *wrapper.CflifEncoder, minSize int) {
		smienc = encoder
		smi = minSize
	}
	mock.flifEncoderSetSplitThreshold = func(encoder *wrapper.CflifEncoder, splitThreshold int) {
		sspenc = encoder
		ssp = splitThreshold
	}
	mock.flifEncoderSetAlphaZeroLossless = func(encoder *wrapper.CflifEncoder) {
		salenc = encoder
		salcall++
	}
	mock.flifEncoderSetChanceCutoff = func(encoder *wrapper.CflifEncoder, chanceCutoff int) {
		sccutenc = encoder
		sccut = chanceCutoff
	}
	mock.flifEncoderSetChanceAlpha = func(encoder *wrapper.CflifEncoder, chanceAlpha int) {
		scalenc = encoder
		scal = chanceAlpha
	}
	mock.flifEncoderSetCrcCheck = func(encoder *wrapper.CflifEncoder, crcCheck int) {
		screnc = encoder
		scr = crcCheck
	}
	mock.flifEncoderSetChannelCompact = func(encoder *wrapper.CflifEncoder, channelCompact int) {
		sccoenc = encoder
		scco = channelCompact
	}
	mock.flifEncoderSetYcocg = func(encoder *wrapper.CflifEncoder, ycocg int) {
		sycenc = encoder
		syc = ycocg
	}
	mock.flifEncoderSetFrameShape = func(encoder *wrapper.CflifEncoder, frameShape int) {
		sfrenc = encoder
		sfr = frameShape
	}
	mock.flifEncoderSetLossy = func(encoder *wrapper.CflifEncoder, lossy int) {
		slosenc = encoder
		slos = lossy
	}

	// All
	cenc := &wrapper.CflifEncoder{}
	fe := &FlifEncoder{
		enc:              cenc,
		LearnRepeat:      20,
		AutoColorBuckets: 30,
		PaletteSize:      40,
		Lookback:         50,
		Divisor:          60,
		MinSize:          70,
		SplitThreashold:  80,
		ChanceCutoff:     90,
		ChanceAlpha:      100,
		ChannelCompact:   110,
		YCoCg:            120,
		FrameShape:       130,
		Lossy:            10,
	}
	setEncoderOptions(fe)
	// Interlaced
	if sinenc != cenc {
		t.Errorf("want %v, got %v", cenc, sinenc)
	}
	if sin != 0 {
		t.Errorf("want %v, got %v", 0, sin)
	}
	// LearnRepeat
	if sleenc != cenc {
		t.Errorf("want %v, got %v", cenc, sleenc)
	}
	if sle != fe.LearnRepeat {
		t.Errorf("want %v, got %v", fe.LearnRepeat, sle)
	}
	// AutoColorBuckets
	if sauenc != cenc {
		t.Errorf("want %v, got %v", cenc, sauenc)
	}
	if sau != fe.AutoColorBuckets {
		t.Errorf("want %v, got %v", fe.AutoColorBuckets, sau)
	}
	// PaletteSize
	if spaenc != cenc {
		t.Errorf("want %v, got %v", cenc, spaenc)
	}
	if spa != fe.PaletteSize {
		t.Errorf("want %v, got %v", fe.PaletteSize, spa)
	}
	// Lookback
	if sloenc != cenc {
		t.Errorf("want %v, got %v", cenc, sloenc)
	}
	if slo != fe.Lookback {
		t.Errorf("want %v, got %v", fe.Lookback, slo)
	}
	// Divisor
	if sdienc != cenc {
		t.Errorf("want %v, got %v", cenc, sdienc)
	}
	if sdi != fe.Divisor {
		t.Errorf("want %v, got %v", fe.Divisor, sdi)
	}
	// MinSize
	if smienc != cenc {
		t.Errorf("want %v, got %v", cenc, smienc)
	}
	if smi != fe.MinSize {
		t.Errorf("want %v, got %v", fe.MinSize, smi)
	}
	// SplitThreashold
	if sspenc != cenc {
		t.Errorf("want %v, got %v", cenc, sspenc)
	}
	if ssp != fe.SplitThreashold {
		t.Errorf("want %v, got %v", fe.SplitThreashold, ssp)
	}
	// AlphaZeroLossless
	if salenc != nil {
		t.Errorf("want %v, got %v", nil, salenc)
	}
	if salcall != 0 {
		t.Errorf("want %v, got %v", 0, salcall)
	}
	// ChanceCutoff
	if sccutenc != cenc {
		t.Errorf("want %v, got %v", cenc, sccutenc)
	}
	if sccut != fe.ChanceCutoff {
		t.Errorf("want %v, got %v", fe.ChanceCutoff, sccut)
	}
	// ChanceAlpha
	if scalenc != cenc {
		t.Errorf("want %v, got %v", cenc, scalenc)
	}
	if scal != fe.ChanceAlpha {
		t.Errorf("want %v, got %v", fe.ChanceAlpha, scal)
	}
	// CrcCheck
	if screnc != cenc {
		t.Errorf("want %v, got %v", cenc, screnc)
	}
	if scr != 0 {
		t.Errorf("want %v, got %v", 0, scr)
	}
	// ChannelCompact
	if sccoenc != cenc {
		t.Errorf("want %v, got %v", cenc, sccoenc)
	}
	if scco != fe.ChannelCompact {
		t.Errorf("want %v, got %v", fe.ChannelCompact, scco)
	}
	// YCoCg
	if sycenc != cenc {
		t.Errorf("want %v, got %v", cenc, sycenc)
	}
	if syc != fe.YCoCg {
		t.Errorf("want %v, got %v", fe.YCoCg, syc)
	}
	// FrameShape
	if sfrenc != cenc {
		t.Errorf("want %v, got %v", cenc, sfrenc)
	}
	if sfr != fe.FrameShape {
		t.Errorf("want %v, got %v", fe.FrameShape, sfr)
	}
	// Lossy
	if slosenc != cenc {
		t.Errorf("want %v, got %v", cenc, slosenc)
	}
	if slos != fe.Lossy {
		t.Errorf("want %v, got %v", fe.Lossy, slos)
	}
	reset()

	// interlaced
	cenc = &wrapper.CflifEncoder{}
	fe = &FlifEncoder{
		enc:        cenc,
		Interlaced: true,
	}
	setEncoderOptions(fe)
	if sinenc != cenc {
		t.Errorf("want %v, got %v", cenc, sinenc)
	}
	if sin != 1 {
		t.Errorf("want %v, got %v", 1, sin)
	}

	// crc check
	cenc = &wrapper.CflifEncoder{}
	fe = &FlifEncoder{
		enc:      cenc,
		CrcCheck: true,
	}
	setEncoderOptions(fe)
	if screnc != cenc {
		t.Errorf("want %v, got %v", cenc, screnc)
	}
	if scr != 1 {
		t.Errorf("want %v, got %v", 1, scr)
	}

	// AlphaZeroLossless
	cenc = &wrapper.CflifEncoder{}
	fe = &FlifEncoder{
		enc:               cenc,
		AlphaZeroLossless: true,
	}
	setEncoderOptions(fe)
	if salenc != cenc {
		t.Errorf("want %v, got %v", cenc, salenc)
	}
	if salcall != 1 {
		t.Errorf("want %v, got %v", 1, salcall)
	}
}
