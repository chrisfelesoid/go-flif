// Copyright 2016 chrisfelesoid. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package wrapper

import (
	"reflect"
	"testing"

	"github.com/chrisfelesoid/go-flif/internal/wrapper"
)

func TestNewFlifDecoder(t *testing.T) {
	mock := &mockDecoderWrapper{}
	decoder = mock
	defer func() {
		decoder = &wrapper.CdecoderWrapper{}
	}()

	dec := (*wrapper.CflifDecoder)(nil)
	mock.flifCreateDecoder = func() *wrapper.CflifDecoder {
		return dec
	}

	// error in FlifCreateDecoder
	fd := NewFlifDecoder()
	if fd != nil {
		t.Errorf("want %v, got %v", nil, fd)
	}

	// success
	dec = &wrapper.CflifDecoder{}
	fd = NewFlifDecoder()
	if fd == nil {
		t.Errorf("must not nil")
	}
	if fd.dec != dec {
		t.Errorf("want %v, got %v", dec, fd.dec)
	}
	// check default values
	if fd.CrcCheck {
		t.Errorf("want %v, got %v", false, fd.CrcCheck)
	}
	if fd.Quality != 100 {
		t.Errorf("want %v, got %v", 100, fd.Quality)
	}
	if fd.Scale != 1 {
		t.Errorf("want %v, got %v", 1, fd.Scale)
	}
	if fd.Width != 0 {
		t.Errorf("want %v, got %v", 0, fd.Width)
	}
	if fd.Height != 0 {
		t.Errorf("want %v, got %v", 0, fd.Height)
	}
	if fd.Fit {
		t.Errorf("want %v, got %v", false, fd.Fit)
	}
}

func TestFlifDecoder_Destroy(t *testing.T) {
	mock := &mockDecoderWrapper{}
	decoder = mock
	defer func() {
		decoder = &wrapper.CdecoderWrapper{}
	}()

	ncall := 0
	dec := (*wrapper.CflifDecoder)(nil)
	mock.flifDestroyDecoder = func(decoder *wrapper.CflifDecoder) {
		dec = decoder
		ncall++
	}

	// error no decoder
	fd := &FlifDecoder{}
	fd.Destroy()
	if dec != nil {
		t.Errorf("want %v, got %v", nil, dec)
	}

	// success
	ncall = 0
	cdec := &wrapper.CflifDecoder{}
	fd = &FlifDecoder{dec: cdec}
	fd.Destroy()
	if dec != cdec {
		t.Errorf("want %v, got %v", cdec, dec)
	}
	if fd.dec != nil {
		t.Errorf("want %v, got %v", nil, fd.dec)
	}
	// retry
	fd.dec = cdec
	fd.Destroy()
	if ncall != 1 {
		t.Errorf("want %v, got %v", 1, ncall)
	}
}

func TestFlifDecoder_Abort(t *testing.T) {
	mock := &mockDecoderWrapper{}
	decoder = mock
	defer func() {
		decoder = &wrapper.CdecoderWrapper{}
	}()

	dec := (*wrapper.CflifDecoder)(nil)
	stat := 0
	mock.flifAbortDecoder = func(decoder *wrapper.CflifDecoder) int {
		dec = decoder
		return stat
	}

	// error in FlifAbortDecoder
	stat = 0
	fd := &FlifDecoder{}
	err := fd.Abort()
	if err != ErrUnknown {
		t.Errorf("another error occurred: %v", err)
	}

	// success
	dec = nil
	stat = 1
	cdec := &wrapper.CflifDecoder{}
	fd = &FlifDecoder{dec: cdec}
	err = fd.Abort()
	if err != nil {
		t.Errorf("error occurred: %v", err)
	}
	if dec != cdec {
		t.Errorf("want %v, got %v", cdec, dec)
	}
}

func TestFlifDecoder_DecodeFile(t *testing.T) {
	bkSetOptions := setDecoderOptions
	mock := &mockDecoderWrapper{}
	decoder = mock
	defer func() {
		decoder = &wrapper.CdecoderWrapper{}
		setDecoderOptions = bkSetOptions
	}()

	scall := 0
	setDecoderOptions = func(d *FlifDecoder) {
		scall++
	}

	dec := (*wrapper.CflifDecoder)(nil)
	fn := ""
	stat := 0
	mock.flifDecoderDecodeFile = func(decoder *wrapper.CflifDecoder, filename string) int {
		dec = decoder
		fn = filename
		return stat
	}

	// error in FlifDecoderDecodeFile
	stat = 0
	fd := &FlifDecoder{}
	err := fd.DecodeFile("hoge")
	if err != ErrUnknown {
		t.Errorf("another error occurred: %v", err)
	}
	if dec != nil {
		t.Errorf("want %v, got %v", nil, dec)
	}
	if fn != "hoge" {
		t.Errorf("want %v, got %v", "hoge", fn)
	}
	if scall != 1 {
		t.Errorf("want %v, got %v", 1, scall)
	}

	// success
	scall = 0
	dec = nil
	stat = 1
	fn = ""
	cdec := &wrapper.CflifDecoder{}
	fd = &FlifDecoder{dec: cdec}
	err = fd.DecodeFile("hoge")
	if err != nil {
		t.Errorf("error occurred: %v", err)
	}
	if dec != cdec {
		t.Errorf("want %v, got %v", cdec, dec)
	}
	if fn != "hoge" {
		t.Errorf("want %v, got %v", "hoge", fn)
	}
	if scall != 1 {
		t.Errorf("want %v, got %v", 1, scall)
	}
}

func TestFlifDecoder_DecodeMemory(t *testing.T) {
	bkSetOptions := setDecoderOptions
	mock := &mockDecoderWrapper{}
	decoder = mock
	defer func() {
		decoder = &wrapper.CdecoderWrapper{}
		setDecoderOptions = bkSetOptions
	}()

	scall := 0
	setDecoderOptions = func(d *FlifDecoder) {
		scall++
	}

	dec := (*wrapper.CflifDecoder)(nil)
	d := ([]byte)(nil)
	stat := 0
	mock.flifDecoderDecodeMemory = func(decoder *wrapper.CflifDecoder, data []byte) int {
		dec = decoder
		d = data
		return stat
	}

	// error in FlifDecoderDecodeFile
	stat = 0
	d = nil
	data := []byte{0x11, 0x22}
	fd := &FlifDecoder{}
	err := fd.DecodeMemory(data)
	if err != ErrUnknown {
		t.Errorf("another error occurred: %v", err)
	}
	if dec != nil {
		t.Errorf("want %v, got %v", nil, dec)
	}
	if !reflect.DeepEqual(d, data) {
		t.Errorf("want %v, got %v", data, d)
	}
	if scall != 1 {
		t.Errorf("want %v, got %v", 1, scall)
	}

	// success
	scall = 0
	dec = nil
	stat = 1
	d = nil
	data = []byte{0x11, 0x22}
	cdec := &wrapper.CflifDecoder{}
	fd = &FlifDecoder{dec: cdec}
	err = fd.DecodeMemory(data)
	if err != nil {
		t.Errorf("error occurred: %v", err)
	}
	if dec != cdec {
		t.Errorf("want %v, got %v", cdec, dec)
	}
	if !reflect.DeepEqual(d, data) {
		t.Errorf("want %v, got %v", data, d)
	}
	if scall != 1 {
		t.Errorf("want %v, got %v", 1, scall)
	}
}

func TestFlifDecoder_GetImageCount(t *testing.T) {
	mock := &mockDecoderWrapper{}
	decoder = mock
	defer func() {
		decoder = &wrapper.CdecoderWrapper{}
	}()

	dec := (*wrapper.CflifDecoder)(nil)
	stat := 0
	mock.flifDecoderNumImages = func(decoder *wrapper.CflifDecoder) int {
		dec = decoder
		return stat
	}

	// success nil
	fd := &FlifDecoder{}
	n := fd.GetImageCount()
	if n != stat {
		t.Errorf("want %v, got %v", stat, n)
	}

	// success
	stat = 1
	cdec := &wrapper.CflifDecoder{}
	fd = &FlifDecoder{dec: cdec}
	n = fd.GetImageCount()
	if n != stat {
		t.Errorf("want %v, got %v", stat, n)
	}
	if dec != cdec {
		t.Errorf("want %v, got %v", cdec, dec)
	}
}

func TestFlifDecoder_GetLoopCount(t *testing.T) {
	mock := &mockDecoderWrapper{}
	decoder = mock
	defer func() {
		decoder = &wrapper.CdecoderWrapper{}
	}()

	dec := (*wrapper.CflifDecoder)(nil)
	stat := 0
	mock.flifDecoderGetNumLoops = func(decoder *wrapper.CflifDecoder) int {
		dec = decoder
		return stat
	}

	// success nil
	fd := &FlifDecoder{}
	n := fd.GetLoopCount()
	if n != stat {
		t.Errorf("want %v, got %v", stat, n)
	}

	// success
	stat = 1
	cdec := &wrapper.CflifDecoder{}
	fd = &FlifDecoder{dec: cdec}
	n = fd.GetLoopCount()
	if n != stat {
		t.Errorf("want %v, got %v", stat, n)
	}
	if dec != cdec {
		t.Errorf("want %v, got %v", cdec, dec)
	}
}

func TestFlifDecoder_GetImage(t *testing.T) {
	mock := &mockDecoderWrapper{}
	decoder = mock
	defer func() {
		decoder = &wrapper.CdecoderWrapper{}
	}()

	icdec := (*wrapper.CflifDecoder)(nil)
	ic := 0
	gidec := (*wrapper.CflifDecoder)(nil)
	i := 0
	gfi := (*wrapper.CflifImage)(nil)
	gicall := 0
	mock.flifDecoderNumImages = func(decoder *wrapper.CflifDecoder) int {
		icdec = decoder
		return ic
	}
	mock.flifDecoderGetImage = func(decoder *wrapper.CflifDecoder, index int) *wrapper.CflifImage {
		gidec = decoder
		i = index
		gicall++
		return gfi
	}

	// no image
	fd := &FlifDecoder{}
	fi := fd.GetImage()
	if len(fi.images) != 0 {
		t.Errorf("want %v, got %v", 0, len(fi.images))
	}
	if gicall != 0 {
		t.Errorf("want %v, got %v", 0, gicall)
	}

	// success
	ic = 3
	gicall = 0
	gfi = &wrapper.CflifImage{}
	cdec := &wrapper.CflifDecoder{}
	fd = &FlifDecoder{dec: cdec}
	fi = fd.GetImage()
	if len(fi.images) != ic {
		t.Errorf("want %v, got %v", ic, len(fi.images))
	}
	if gicall != ic {
		t.Errorf("want %v, got %v", ic, gicall)
	}
	if icdec != cdec {
		t.Errorf("want %v, got %v", cdec, icdec)
	}
	if gidec != cdec {
		t.Errorf("want %v, got %v", cdec, gidec)
	}
}

func TestGetInfo(t *testing.T) {
	mock := &mockDecoderWrapper{}
	decoder = mock
	defer func() {
		decoder = &wrapper.CdecoderWrapper{}
	}()

	rin := (*wrapper.CflifInfo)(nil)
	d := ([]byte)(nil)
	mock.flifReadInfoFromMemory = func(data []byte) *wrapper.CflifInfo {
		d = data
		return rin
	}
	din := (*wrapper.CflifInfo)(nil)
	mock.flifDestroyInfo = func(info *wrapper.CflifInfo) {
		din = info
	}
	win := (*wrapper.CflifInfo)(nil)
	w := 0
	mock.flifInfoGetWidth = func(info *wrapper.CflifInfo) int {
		win = info
		return w
	}
	hin := (*wrapper.CflifInfo)(nil)
	h := 0
	mock.flifInfoGetHeight = func(info *wrapper.CflifInfo) int {
		hin = info
		return h
	}
	ncin := (*wrapper.CflifInfo)(nil)
	nc := 0
	mock.flifInfoGetNbChannels = func(info *wrapper.CflifInfo) int {
		ncin = info
		return nc
	}
	depin := (*wrapper.CflifInfo)(nil)
	dep := 0
	mock.flifInfoGetDepth = func(info *wrapper.CflifInfo) int {
		depin = info
		return dep
	}
	niin := (*wrapper.CflifInfo)(nil)
	ni := 0
	mock.flifInfoNumImages = func(info *wrapper.CflifInfo) int {
		niin = info
		return ni
	}

	// error in FlifReadInfoFromMemory
	data := []byte{0x11, 0x22}
	info := GetInfo(data)
	if info != nil {
		t.Errorf("want %v, got %v", nil, info)
	}
	if !reflect.DeepEqual(data, d) {
		t.Errorf("want %v, got %v", data, d)
	}

	//
	data = []byte{0x11, 0x22}
	rin = &wrapper.CflifInfo{}
	w = 10
	h = 20
	nc = 2
	dep = 8
	ni = 3
	info = GetInfo(data)
	if info == nil {
		t.Errorf("must not nil")
	}
	if !reflect.DeepEqual(d, data) {
		t.Errorf("want %v, got %v", data, d)
	}
	if din != rin {
		t.Errorf("want %v, got %v", rin, din)
	}
	if win != rin {
		t.Errorf("want %v, got %v", rin, win)
	}
	if info.Width != w {
		t.Errorf("want %v, got %v", w, info.Width)
	}
	if hin != rin {
		t.Errorf("want %v, got %v", rin, hin)
	}
	if info.Height != h {
		t.Errorf("want %v, got %v", h, info.Height)
	}
	if ncin != rin {
		t.Errorf("want %v, got %v", rin, ncin)
	}
	if info.Channel != nc {
		t.Errorf("want %v, got %v", nc, info.Channel)
	}
	if depin != rin {
		t.Errorf("want %v, got %v", rin, depin)
	}
	if info.Depth != dep {
		t.Errorf("want %v, got %v", dep, info.Depth)
	}
	if niin != rin {
		t.Errorf("want %v, got %v", rin, niin)
	}
	if info.ImageCount != ni {
		t.Errorf("want %v, got %v", ni, info.ImageCount)
	}
}

func TestSetDecoderOptions(t *testing.T) {
	mock := &mockDecoderWrapper{}
	decoder = mock
	defer func() {
		decoder = &wrapper.CdecoderWrapper{}
	}()

	ccdec := (*wrapper.CflifDecoder)(nil)
	qdec := (*wrapper.CflifDecoder)(nil)
	sdec := (*wrapper.CflifDecoder)(nil)
	rdec := (*wrapper.CflifDecoder)(nil)
	fdec := (*wrapper.CflifDecoder)(nil)
	cc := -1
	q := -1
	s := -1
	w := -1
	h := -1

	reset := func() {
		ccdec = (*wrapper.CflifDecoder)(nil)
		qdec = (*wrapper.CflifDecoder)(nil)
		sdec = (*wrapper.CflifDecoder)(nil)
		rdec = (*wrapper.CflifDecoder)(nil)
		fdec = (*wrapper.CflifDecoder)(nil)
		cc = -1
		q = -1
		s = -1
		w = -1
		h = -1
	}
	mock.flifDecoderSetCrcCheck = func(decoder *wrapper.CflifDecoder, crcCheck int) {
		ccdec = decoder
		cc = crcCheck
	}
	mock.flifDecoderSetQuality = func(decoder *wrapper.CflifDecoder, quality int) {
		qdec = decoder
		q = quality
	}
	mock.flifDecoderSetScale = func(decoder *wrapper.CflifDecoder, scale int) {
		sdec = decoder
		s = scale
	}
	mock.flifDecoderSetResize = func(decoder *wrapper.CflifDecoder, width, height int) {
		rdec = decoder
		w = width
		h = height
	}
	mock.flifDecoderSetFit = func(decoder *wrapper.CflifDecoder, width, height int) {
		fdec = decoder
		w = width
		h = height
	}

	// check all, no crc, no fit
	cdec := &wrapper.CflifDecoder{}
	fd := &FlifDecoder{
		dec:     cdec,
		Quality: 50,
		Scale:   2,
		Width:   100,
		Height:  200,
	}
	setDecoderOptions(fd)
	if ccdec != cdec {
		t.Errorf("want %v, got %v", cdec, ccdec)
	}
	if cc != 0 {
		t.Errorf("want %v, got %v", 0, cc)
	}
	if qdec != cdec {
		t.Errorf("want %v, got %v", cdec, qdec)
	}
	if q != fd.Quality {
		t.Errorf("want %v, got %v", fd.Quality, q)
	}
	if sdec != cdec {
		t.Errorf("want %v, got %v", cdec, sdec)
	}
	if s != fd.Scale {
		t.Errorf("want %v, got %v", fd.Scale, s)
	}
	if rdec != cdec {
		t.Errorf("want %v, got %v", cdec, rdec)
	}
	if w != fd.Width {
		t.Errorf("want %v, got %v", fd.Width, w)
	}
	if h != fd.Height {
		t.Errorf("want %v, got %v", fd.Height, h)
	}
	// not exec
	if fdec != nil {
		t.Errorf("want %v, got %v", nil, fdec)
	}
	reset()

	// crc check
	cdec = &wrapper.CflifDecoder{}
	fd = &FlifDecoder{
		dec:      cdec,
		CrcCheck: true,
	}
	setDecoderOptions(fd)
	if cc != 1 {
		t.Errorf("want %v, got %v", 1, cc)
	}
	reset()

	// fit
	cdec = &wrapper.CflifDecoder{}
	fd = &FlifDecoder{
		dec:    cdec,
		Fit:    true,
		Width:  100,
		Height: 200,
	}
	setDecoderOptions(fd)
	if rdec != nil {
		t.Errorf("want %v, got %v", nil, rdec)
	}
	if fdec != cdec {
		t.Errorf("want %v, got %v", cdec, fdec)
	}
	if w != fd.Width {
		t.Errorf("want %v, got %v", fd.Width, w)
	}
	if h != fd.Height {
		t.Errorf("want %v, got %v", fd.Height, h)
	}
}
