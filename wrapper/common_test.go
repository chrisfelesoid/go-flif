// Copyright 2016 chrisfelesoid. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package wrapper

import (
	"reflect"
	"testing"

	"github.com/chrisfelesoid/go-flif/internal/wrapper"
)

func TestNewFlifImage(t *testing.T) {
	mock := &mockCommonWrapper{}
	common = mock
	defer func() {
		common = &wrapper.CcommonWrapper{}
	}()

	w := 0
	h := 0
	im := (*wrapper.CflifImage)(nil)
	mock.flifCreateImage = func(width, height int) *wrapper.CflifImage {
		w = width
		h = height
		return im
	}

	// error in FlifCreateImage
	f := NewFlifImage(10, 20)
	if f != nil {
		t.Errorf("want %v, got %v", nil, f)
	}
	if w != 10 {
		t.Errorf("want %v, got %v", 10, w)
	}
	if h != 20 {
		t.Errorf("want %v, got %v", 20, h)
	}

	// success
	w = 0
	h = 0
	im = &wrapper.CflifImage{}
	f = NewFlifImage(10, 20)
	if f == nil {
		t.Errorf("must not nil")
	}
	if f.images[0] != im {
		t.Errorf("want %v, got %v", im, f.images[0])
	}
	if w != 10 {
		t.Errorf("want %v, got %v", 10, w)
	}
	if h != 20 {
		t.Errorf("want %v, got %v", 20, h)
	}
}

func TestNewFlifImageHDR(t *testing.T) {
	mock := &mockCommonWrapper{}
	common = mock
	defer func() {
		common = &wrapper.CcommonWrapper{}
	}()

	w := 0
	h := 0
	im := (*wrapper.CflifImage)(nil)
	mock.flifCreateImageHDR = func(width, height int) *wrapper.CflifImage {
		w = width
		h = height
		return im
	}

	// error in FlifCreateImage
	f := NewFlifImageHDR(10, 20)
	if f != nil {
		t.Errorf("want %v, got %v", nil, f)
	}
	if w != 10 {
		t.Errorf("want %v, got %v", 10, w)
	}
	if h != 20 {
		t.Errorf("want %v, got %v", 20, h)
	}

	// success
	w = 0
	h = 0
	im = &wrapper.CflifImage{}
	f = NewFlifImageHDR(10, 20)
	if f == nil {
		t.Errorf("must not nil")
	}
	if f.images[0] != im {
		t.Errorf("want %v, got %v", im, f.images[0])
	}
	if w != 10 {
		t.Errorf("want %v, got %v", 10, w)
	}
	if h != 20 {
		t.Errorf("want %v, got %v", 20, h)
	}
}

func TestNewFlifImageFromRGBA(t *testing.T) {
	mock := &mockCommonWrapper{}
	common = mock
	defer func() {
		common = &wrapper.CcommonWrapper{}
	}()

	w := 0
	h := 0
	d := ([]byte)(nil)
	st := 0
	im := (*wrapper.CflifImage)(nil)
	mock.flifImportImageRGBA = func(width, height int, data []byte, stride int) *wrapper.CflifImage {
		w = width
		h = height
		d = data
		st = stride
		return im
	}

	// error in FlifCreateImage
	data := []byte{0x11, 0x22}
	f := NewFlifImageFromRGBA(10, 20, data)
	if f != nil {
		t.Errorf("want %v, got %v", nil, f)
	}
	if w != 10 {
		t.Errorf("want %v, got %v", 10, w)
	}
	if h != 20 {
		t.Errorf("want %v, got %v", 20, h)
	}
	if !reflect.DeepEqual(d, data) {
		t.Errorf("want %v, got %v", data, d)
	}
	if st != 40 {
		t.Errorf("want %v, got %v", 40, st)
	}

	// success
	w = 0
	h = 0
	d = nil
	st = 0
	im = &wrapper.CflifImage{}
	mock.flifImportImageRGBA = func(width, height int, data []byte, stride int) *wrapper.CflifImage {
		w = width
		h = height
		d = data
		st = stride
		return im
	}
	data = []byte{0x11, 0x22}
	f = NewFlifImageFromRGBA(10, 20, data)
	if f == nil {
		t.Errorf("must not nil")
	}
	if f.images[0] != im {
		t.Errorf("want %v, got %v", im, f.images[0])
	}
	if w != 10 {
		t.Errorf("want %v, got %v", 10, w)
	}
	if h != 20 {
		t.Errorf("want %v, got %v", 20, h)
	}
	if !reflect.DeepEqual(d, data) {
		t.Errorf("want %v, got %v", data, d)
	}
	if st != 40 {
		t.Errorf("want %v, got %v", 40, st)
	}
}

func TestNewFlifImageFromRGB(t *testing.T) {
	mock := &mockCommonWrapper{}
	common = mock
	defer func() {
		common = &wrapper.CcommonWrapper{}
	}()

	w := 0
	h := 0
	d := ([]byte)(nil)
	st := 0
	im := (*wrapper.CflifImage)(nil)
	mock.flifImportImageRGB = func(width, height int, data []byte, stride int) *wrapper.CflifImage {
		w = width
		h = height
		d = data
		st = stride
		return im
	}

	// error in FlifCreateImage
	data := []byte{0x11, 0x22}
	f := NewFlifImageFromRGB(10, 20, data)
	if f != nil {
		t.Errorf("want %v, got %v", nil, f)
	}
	if w != 10 {
		t.Errorf("want %v, got %v", 10, w)
	}
	if h != 20 {
		t.Errorf("want %v, got %v", 20, h)
	}
	if !reflect.DeepEqual(d, data) {
		t.Errorf("want %v, got %v", data, d)
	}
	if st != 30 {
		t.Errorf("want %v, got %v", 30, st)
	}

	// success
	w = 0
	h = 0
	d = nil
	st = 0
	im = &wrapper.CflifImage{}
	data = []byte{0x11, 0x22}
	f = NewFlifImageFromRGB(10, 20, data)
	if f == nil {
		t.Errorf("must not nil")
	}
	if f.images[0] != im {
		t.Errorf("want %v, got %v", im, f.images[0])
	}
	if w != 10 {
		t.Errorf("want %v, got %v", 10, w)
	}
	if h != 20 {
		t.Errorf("want %v, got %v", 20, h)
	}
	if !reflect.DeepEqual(d, data) {
		t.Errorf("want %v, got %v", data, d)
	}
	if st != 30 {
		t.Errorf("want %v, got %v", 30, st)
	}
}

func TestNewFlifImageFromGray(t *testing.T) {
	mock := &mockCommonWrapper{}
	common = mock
	defer func() {
		common = &wrapper.CcommonWrapper{}
	}()

	w := 0
	h := 0
	d := ([]byte)(nil)
	st := 0
	im := (*wrapper.CflifImage)(nil)
	mock.flifImportImageGRAY = func(width, height int, data []byte, stride int) *wrapper.CflifImage {
		w = width
		h = height
		d = data
		st = stride
		return im
	}

	// error in FlifCreateImage
	data := []byte{0x11, 0x22}
	f := NewFlifImageFromGray(10, 20, data)
	if f != nil {
		t.Errorf("want nil, got %v", f)
	}
	if w != 10 {
		t.Errorf("want %v, got %v", 10, w)
	}
	if h != 20 {
		t.Errorf("want %v, got %v", 20, h)
	}
	if !reflect.DeepEqual(d, data) {
		t.Errorf("want %v, got %v", data, d)
	}
	if st != 10 {
		t.Errorf("want %v, got %v", 10, st)
	}

	// success
	w = 0
	h = 0
	d = nil
	st = 0
	im = &wrapper.CflifImage{}
	data = []byte{0x11, 0x22}
	f = NewFlifImageFromGray(10, 20, data)
	if f == nil {
		t.Errorf("must not nil")
	}
	if f.images[0] != im {
		t.Errorf("want %v, got %v", im, f.images[0])
	}
	if w != 10 {
		t.Errorf("want %v, got %v", 10, w)
	}
	if h != 20 {
		t.Errorf("want %v, got %v", 20, h)
	}
	if !reflect.DeepEqual(d, data) {
		t.Errorf("want %v, got %v", data, d)
	}
	if st != 10 {
		t.Errorf("want %v, got %v", 10, st)
	}
}

func TestFlifImage_Destroy(t *testing.T) {
	mock := &mockCommonWrapper{}
	common = mock
	defer func() {
		common = &wrapper.CcommonWrapper{}
	}()

	ncall := 0
	im := (*wrapper.CflifImage)(nil)
	mock.flifDestroyImage = func(image *wrapper.CflifImage) {
		im = image
		ncall++
	}
	// do nothing
	f := &FlifImage{}
	f.Destroy()
	if im != nil {
		t.Errorf("want %v, get %v", nil, im)
	}

	// destroy
	ncall = 0
	im = nil
	cimg := &wrapper.CflifImage{}
	f = &FlifImage{
		images: []*wrapper.CflifImage{cimg},
	}
	f.Destroy()
	if im != cimg {
		t.Errorf("want %v, got %v", cimg, im)
	}
	if f.images != nil {
		t.Errorf("want %v, got %v", nil, f.images)
	}
	if ncall != 1 {
		t.Errorf("want %v, got %v", 1, ncall)
	}
	// retry
	f.Destroy()
	if ncall != 1 {
		t.Errorf("want %v, got %v", 1, ncall)
	}
}

func TestFlifImage_GetImageCount(t *testing.T) {
	f := &FlifImage{}
	if f.GetImageCount() != 0 {
		t.Errorf("want %v, got %v", 0, f.GetImageCount())
	}

	f = &FlifImage{
		images: make([]*wrapper.CflifImage, 4),
	}
	if f.GetImageCount() != 4 {
		t.Errorf("want %v, got %v", 4, f.GetImageCount())
	}
}

func TestFlifImage_GetWidth(t *testing.T) {
	mock := &mockCommonWrapper{}
	common = mock
	defer func() {
		common = &wrapper.CcommonWrapper{}
	}()

	w := 123
	im := (*wrapper.CflifImage)(nil)
	mock.flifImageGetWidth = func(image *wrapper.CflifImage) int {
		im = image
		return w
	}
	// no images
	f := &FlifImage{}
	if f.GetWidth() != 0 {
		t.Errorf("want %v, got %v", 0, f.GetWidth())
	}
	if im != nil {
		t.Errorf("want %v, got %v", nil, im)
	}

	// success
	im = nil
	cimg := &wrapper.CflifImage{}
	f = &FlifImage{
		images: []*wrapper.CflifImage{cimg},
	}
	if f.GetWidth() != w {
		t.Errorf("want %v, got %v", w, f.GetWidth())
	}
	if im != cimg {
		t.Errorf("want %v, got %v", cimg, im)
	}
}

func TestFlifImage_GetHeight(t *testing.T) {
	mock := &mockCommonWrapper{}
	common = mock
	defer func() {
		common = &wrapper.CcommonWrapper{}
	}()

	h := 123
	im := (*wrapper.CflifImage)(nil)
	mock.flifImageGetHeight = func(image *wrapper.CflifImage) int {
		im = image
		return h
	}
	// no images
	f := &FlifImage{}
	if f.GetHeight() != 0 {
		t.Errorf("want %v, got %v", 0, f.GetHeight())
	}
	if im != nil {
		t.Errorf("want %v, got %v", nil, im)
	}

	// success
	im = nil
	cimg := &wrapper.CflifImage{}
	f = &FlifImage{
		images: []*wrapper.CflifImage{cimg},
	}
	if f.GetHeight() != h {
		t.Errorf("want %v, got %v", h, f.GetHeight())
	}
	if im != cimg {
		t.Errorf("want %v, got %v", cimg, im)
	}
}

func TestFlifImage_GetChannel(t *testing.T) {
	mock := &mockCommonWrapper{}
	common = mock
	defer func() {
		common = &wrapper.CcommonWrapper{}
	}()

	ch := 2
	im := (*wrapper.CflifImage)(nil)
	mock.flifImageGetNbChannels = func(image *wrapper.CflifImage) int {
		im = image
		return ch
	}
	// no images
	f := &FlifImage{}
	if f.GetChannel() != 0 {
		t.Errorf("want %v, got %v", 0, f.GetChannel())
	}
	if im != nil {
		t.Errorf("want %v, got %v", nil, im)
	}

	// success
	im = nil
	cimg := &wrapper.CflifImage{}
	f = &FlifImage{
		images: []*wrapper.CflifImage{cimg},
	}
	if f.GetChannel() != ch {
		t.Errorf("want %v, got %v", ch, f.GetChannel())
	}
	if im != cimg {
		t.Errorf("want %v, got %v", cimg, im)
	}
}

func TestFlifImage_GetDepth(t *testing.T) {
	mock := &mockCommonWrapper{}
	common = mock
	defer func() {
		common = &wrapper.CcommonWrapper{}
	}()

	d := 8
	im := (*wrapper.CflifImage)(nil)
	mock.flifImageGetDepth = func(image *wrapper.CflifImage) int {
		im = image
		return d
	}
	// no images
	f := &FlifImage{}
	if f.GetDepth() != 0 {
		t.Errorf("want %v, got %v", 0, f.GetDepth())
	}
	if im != nil {
		t.Errorf("want %v, got %v", nil, im)
	}

	// success
	im = nil
	cimg := &wrapper.CflifImage{}
	f = &FlifImage{
		images: []*wrapper.CflifImage{cimg},
	}
	if f.GetDepth() != d {
		t.Errorf("want %v, got %v", d, f.GetDepth())
	}
	if im != cimg {
		t.Errorf("want %v, got %v", cimg, im)
	}
}

func TestFlifImage_GetFrameDelay(t *testing.T) {
	mock := &mockCommonWrapper{}
	common = mock
	defer func() {
		common = &wrapper.CcommonWrapper{}
	}()

	d := 1
	im := (*wrapper.CflifImage)(nil)
	mock.flifImageGetFrameDelay = func(image *wrapper.CflifImage) int {
		im = image
		return d
	}
	// no images
	f := &FlifImage{}
	if f.GetFrameDelay() != 0 {
		t.Errorf("want %v, got %v", 0, f.GetFrameDelay())
	}
	if im != nil {
		t.Errorf("want %v, got %v", nil, im)
	}

	// success
	im = nil
	cimg := &wrapper.CflifImage{}
	f = &FlifImage{
		images: []*wrapper.CflifImage{cimg},
	}
	if f.GetFrameDelay() != d {
		t.Errorf("want %v, got %v", d, f.GetFrameDelay())
	}
	if im != cimg {
		t.Errorf("want %v, got %v", cimg, im)
	}
}

func TestFlifImage_SetFrameDelay(t *testing.T) {
	mock := &mockCommonWrapper{}
	common = mock
	defer func() {
		common = &wrapper.CcommonWrapper{}
	}()

	d := 0
	im := (*wrapper.CflifImage)(nil)
	mock.flifImageSetFrameDelay = func(image *wrapper.CflifImage, delay int) {
		im = image
		d = delay
	}
	// no images
	f := &FlifImage{}
	f.SetFrameDelay(10)
	if d != 0 {
		t.Errorf("want %v, got %v", 0, d)
	}
	if im != nil {
		t.Errorf("want %v, got %v", nil, im)
	}

	// success
	d = 0
	im = nil
	cimg := &wrapper.CflifImage{}
	f = &FlifImage{
		images: []*wrapper.CflifImage{cimg},
	}
	f.SetFrameDelay(10)
	if d != 10 {
		t.Errorf("want %v, got %v", 10, d)
	}
	if im != cimg {
		t.Errorf("want %v, got %v", cimg, im)
	}
}

func TestFlifImage_SetMetadata(t *testing.T) {
	mock := &mockCommonWrapper{}
	common = mock
	defer func() {
		common = &wrapper.CcommonWrapper{}
	}()

	im := (*wrapper.CflifImage)(nil)
	n := ""
	d := ([]byte)(nil)
	mock.flifImageSetMetadata = func(image *wrapper.CflifImage, name string, data []byte) {
		im = image
		n = name
		d = data
	}

	// no images
	f := &FlifImage{}
	data := []byte{0x11, 0x22}

	f.SetMetadata("hoge", data)
	if d != nil {
		t.Errorf("want %v, got %v", nil, d)
	}

	// success
	im = nil
	n = ""
	d = nil
	cimg := &wrapper.CflifImage{}
	f = &FlifImage{
		images: []*wrapper.CflifImage{cimg},
	}
	data = []byte{0x11, 0x22}

	f.SetMetadata("hoge", data)
	if im != cimg {
		t.Errorf("want %v, got %v", cimg, im)
	}
	if n != "hoge" {
		t.Errorf("want %v, got %v", "hoge", n)
	}
	if !reflect.DeepEqual(d, data) {
		t.Errorf("want %v, got %v", data, d)
	}
}

func TestFlifImage_GetMetadata(t *testing.T) {
	mock := &mockCommonWrapper{}
	common = mock
	defer func() {
		common = &wrapper.CcommonWrapper{}
	}()

	im := (*wrapper.CflifImage)(nil)
	n := ""
	d := []byte{0x11, 0x22}
	mock.flifImageGetMetadata = func(image *wrapper.CflifImage, name string, data *[]byte) {
		im = image
		n = name
		*data = d
	}
	// no images
	f := &FlifImage{}
	data := f.GetMetadata("hoge")
	if data != nil {
		t.Errorf("want %v, got %v", nil, data)
	}
	if im != nil {
		t.Errorf("want %v, got %v", nil, im)
	}

	// success
	im = nil
	n = ""
	cimg := &wrapper.CflifImage{}
	f = &FlifImage{
		images: []*wrapper.CflifImage{cimg},
	}
	data = f.GetMetadata("hoge")
	if !reflect.DeepEqual(d, data) {
		t.Errorf("want %v, got %v", d, data)
	}
	if im != cimg {
		t.Errorf("want %v, got %v", cimg, im)
	}
}

func TestFlifImage_WriteRowRGBA8(t *testing.T) {
	mock := &mockCommonWrapper{}
	common = mock
	defer func() {
		common = &wrapper.CcommonWrapper{}
	}()

	im := (*wrapper.CflifImage)(nil)
	r := 0
	d := ([]byte)(nil)
	w := 0
	h := 0
	mock.flifImageWriteRowRGBA8 = func(image *wrapper.CflifImage, row int, data []byte) {
		im = image
		r = row
		d = data
	}
	mock.flifImageGetHeight = func(image *wrapper.CflifImage) int {
		return h
	}
	mock.flifImageGetWidth = func(image *wrapper.CflifImage) int {
		return w
	}
	// no images (GetHeight is 0)
	f := &FlifImage{}
	data := []byte{0x11, 0x22, 0x33, 0x44}
	err := f.WriteRowRGBA8(0, data, 0)
	if err != ErrOutOfRange {
		t.Errorf("another error occurred:%v", err)
	}

	// error GetHeight is 0
	cimg := &wrapper.CflifImage{}
	f = &FlifImage{
		images: []*wrapper.CflifImage{cimg},
	}
	err = f.WriteRowRGBA8(0, data, 0)
	if err != ErrOutOfRange {
		t.Errorf("another error occurred:%v", err)
	}

	// error out row
	cimg = &wrapper.CflifImage{}
	f = &FlifImage{
		images: []*wrapper.CflifImage{cimg},
	}
	h = 1
	err = f.WriteRowRGBA8(2, data, 0)
	if err != ErrOutOfRange {
		t.Errorf("another error occurred:%v", err)
	}

	// error out row
	cimg = &wrapper.CflifImage{}
	f = &FlifImage{
		images: []*wrapper.CflifImage{cimg},
	}
	h = 1
	err = f.WriteRowRGBA8(-1, data, 0)
	if err != ErrOutOfRange {
		t.Errorf("another error occurred:%v", err)
	}

	// error GetWidth is 0
	cimg = &wrapper.CflifImage{}
	f = &FlifImage{
		images: []*wrapper.CflifImage{cimg},
	}
	h = 1
	err = f.WriteRowRGBA8(0, data, 0)
	if err != ErrOutOfRange {
		t.Errorf("another error occurred:%v", err)
	}

	// error data length
	cimg = &wrapper.CflifImage{}
	f = &FlifImage{
		images: []*wrapper.CflifImage{cimg},
	}
	h = 1
	w = 2
	err = f.WriteRowRGBA8(0, data, 0)
	if err != ErrOutOfRange {
		t.Errorf("another error occurred:%v", err)
	}

	// error image index
	cimg = &wrapper.CflifImage{}
	f = &FlifImage{
		images: []*wrapper.CflifImage{cimg},
	}
	h = 1
	w = 1
	err = f.WriteRowRGBA8(0, data, 1)
	if err != ErrOutOfRange {
		t.Errorf("another error occurred:%v", err)
	}

	// error image index
	cimg = &wrapper.CflifImage{}
	f = &FlifImage{
		images: []*wrapper.CflifImage{cimg},
	}
	h = 1
	w = 1
	err = f.WriteRowRGBA8(0, data, -1)
	if err != ErrOutOfRange {
		t.Errorf("another error occurred:%v", err)
	}

	// success
	cimg = &wrapper.CflifImage{}
	f = &FlifImage{
		images: []*wrapper.CflifImage{cimg},
	}
	h = 2
	w = 1
	err = f.WriteRowRGBA8(1, data, 0)
	if err != nil {
		t.Errorf("error occurred:%v", err)
	}
	if r != 1 {
		t.Errorf("want %v, got %v", 1, r)
	}
	if !reflect.DeepEqual(d, data) {
		t.Errorf("want %v, got %v", data, d)
	}
}

func TestFlifImage_ReadRowRGBA8(t *testing.T) {
	mock := &mockCommonWrapper{}
	common = mock
	defer func() {
		common = &wrapper.CcommonWrapper{}
	}()

	im := (*wrapper.CflifImage)(nil)
	r := 0
	d := ([]byte)(nil)
	h := 0
	mock.flifImageReadRowRGBA8 = func(image *wrapper.CflifImage, row int) []byte {
		im = image
		r = row
		return d
	}
	mock.flifImageGetHeight = func(image *wrapper.CflifImage) int {
		return h
	}

	// no images (GetHeight is 0)
	f := &FlifImage{}
	data, err := f.ReadRowRGBA8(0, 0)
	if err != ErrOutOfRange {
		t.Errorf("another error occurred:%v", err)
	}
	if data != nil {
		t.Errorf("want %v, got %v", nil, data)
	}

	// error GetHeight is 0
	cimg := &wrapper.CflifImage{}
	f = &FlifImage{
		images: []*wrapper.CflifImage{cimg},
	}
	data, err = f.ReadRowRGBA8(0, 0)
	if err != ErrOutOfRange {
		t.Errorf("another error occurred:%v", err)
	}

	// error out row
	cimg = &wrapper.CflifImage{}
	f = &FlifImage{
		images: []*wrapper.CflifImage{cimg},
	}
	h = 1
	data, err = f.ReadRowRGBA8(1, 0)
	if err != ErrOutOfRange {
		t.Errorf("another error occurred:%v", err)
	}

	// error out row
	cimg = &wrapper.CflifImage{}
	f = &FlifImage{
		images: []*wrapper.CflifImage{cimg},
	}
	h = 1
	data, err = f.ReadRowRGBA8(-1, 0)
	if err != ErrOutOfRange {
		t.Errorf("another error occurred:%v", err)
	}

	// error image index
	cimg = &wrapper.CflifImage{}
	f = &FlifImage{
		images: []*wrapper.CflifImage{cimg},
	}
	h = 1
	data, err = f.ReadRowRGBA8(0, 1)
	if err != ErrOutOfRange {
		t.Errorf("another error occurred:%v", err)
	}

	// error image index
	cimg = &wrapper.CflifImage{}
	f = &FlifImage{
		images: []*wrapper.CflifImage{cimg},
	}
	h = 1
	data, err = f.ReadRowRGBA8(0, -1)
	if err != ErrOutOfRange {
		t.Errorf("another error occurred:%v", err)
	}

	// error get data
	cimg = &wrapper.CflifImage{}
	f = &FlifImage{
		images: []*wrapper.CflifImage{cimg},
	}
	h = 1
	data, err = f.ReadRowRGBA8(0, 0)
	if err != ErrUnknown {
		t.Errorf("another error occurred:%v", err)
	}
	if data != nil {
		t.Errorf("want %v, got %v", nil, data)
	}

	// success
	d = []byte{0x11, 0x22, 0x33, 0x44}
	cimg = &wrapper.CflifImage{}
	f = &FlifImage{
		images: []*wrapper.CflifImage{cimg},
	}
	h = 1
	data, err = f.ReadRowRGBA8(0, 0)
	if err != nil {
		t.Errorf("error occurred:%v", err)
	}
	if !reflect.DeepEqual(d, data) {
		t.Errorf("want %v, got %v", d, data)
	}
}

func TestFlifImage_WriteRowRGBA16(t *testing.T) {
	mock := &mockCommonWrapper{}
	common = mock
	defer func() {
		common = &wrapper.CcommonWrapper{}
	}()

	im := (*wrapper.CflifImage)(nil)
	r := 0
	d := ([]byte)(nil)
	w := 0
	h := 0
	mock.flifImageWriteRowRGBA16 = func(image *wrapper.CflifImage, row int, data []byte) {
		im = image
		r = row
		d = data
	}
	mock.flifImageGetHeight = func(image *wrapper.CflifImage) int {
		return h
	}
	mock.flifImageGetWidth = func(image *wrapper.CflifImage) int {
		return w
	}
	// no images (GetHeight is 0)
	f := &FlifImage{}
	data := []byte{0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88}
	err := f.WriteRowRGBA16(0, data, 0)
	if err != ErrOutOfRange {
		t.Errorf("another error occurred:%v", err)
	}

	// error GetHeight is 0
	cimg := &wrapper.CflifImage{}
	f = &FlifImage{
		images: []*wrapper.CflifImage{cimg},
	}
	err = f.WriteRowRGBA16(0, data, 0)
	if err != ErrOutOfRange {
		t.Errorf("another error occurred:%v", err)
	}

	// error out row
	cimg = &wrapper.CflifImage{}
	f = &FlifImage{
		images: []*wrapper.CflifImage{cimg},
	}
	h = 1
	err = f.WriteRowRGBA16(2, data, 0)
	if err != ErrOutOfRange {
		t.Errorf("another error occurred:%v", err)
	}

	// error out row
	cimg = &wrapper.CflifImage{}
	f = &FlifImage{
		images: []*wrapper.CflifImage{cimg},
	}
	h = 1
	err = f.WriteRowRGBA16(-1, data, 0)
	if err != ErrOutOfRange {
		t.Errorf("another error occurred:%v", err)
	}

	// error GetWidth is 0
	cimg = &wrapper.CflifImage{}
	f = &FlifImage{
		images: []*wrapper.CflifImage{cimg},
	}
	h = 1
	err = f.WriteRowRGBA16(0, data, 0)
	if err != ErrOutOfRange {
		t.Errorf("another error occurred:%v", err)
	}

	// error data length
	cimg = &wrapper.CflifImage{}
	f = &FlifImage{
		images: []*wrapper.CflifImage{cimg},
	}
	h = 1
	w = 2
	err = f.WriteRowRGBA16(0, data, 0)
	if err != ErrOutOfRange {
		t.Errorf("another error occurred:%v", err)
	}

	// error image index
	cimg = &wrapper.CflifImage{}
	f = &FlifImage{
		images: []*wrapper.CflifImage{cimg},
	}
	h = 1
	w = 1
	err = f.WriteRowRGBA16(0, data, 1)
	if err != ErrOutOfRange {
		t.Errorf("another error occurred:%v", err)
	}

	// error image index
	cimg = &wrapper.CflifImage{}
	f = &FlifImage{
		images: []*wrapper.CflifImage{cimg},
	}
	h = 1
	w = 1
	err = f.WriteRowRGBA16(0, data, -1)
	if err != ErrOutOfRange {
		t.Errorf("another error occurred:%v", err)
	}

	// success
	cimg = &wrapper.CflifImage{}
	f = &FlifImage{
		images: []*wrapper.CflifImage{cimg},
	}
	h = 2
	w = 1
	err = f.WriteRowRGBA16(1, data, 0)
	if err != nil {
		t.Errorf("error occurred:%v", err)
	}
	if r != 1 {
		t.Errorf("want %v, got %v", 1, r)
	}
	if !reflect.DeepEqual(data, d) {
		t.Errorf("want %v, got %v", data, d)
	}
}

func TestFlifImage_ReadRowRGBA16(t *testing.T) {
	mock := &mockCommonWrapper{}
	common = mock
	defer func() {
		common = &wrapper.CcommonWrapper{}
	}()

	im := (*wrapper.CflifImage)(nil)
	r := 0
	d := ([]byte)(nil)
	h := 0
	mock.flifImageReadRowRGBA16 = func(image *wrapper.CflifImage, row int) []byte {
		im = image
		r = row
		return d
	}
	mock.flifImageGetHeight = func(image *wrapper.CflifImage) int {
		return h
	}

	// no images (GetHeight is 0)
	f := &FlifImage{}
	data, err := f.ReadRowRGBA16(0, 0)
	if err != ErrOutOfRange {
		t.Errorf("another error occurred:%v", err)
	}
	if data != nil {
		t.Errorf("want %v, got %v", nil, data)
	}

	// error GetHeight is 0
	cimg := &wrapper.CflifImage{}
	f = &FlifImage{
		images: []*wrapper.CflifImage{cimg},
	}
	data, err = f.ReadRowRGBA16(0, 0)
	if err != ErrOutOfRange {
		t.Errorf("another error occurred:%v", err)
	}

	// error out row
	cimg = &wrapper.CflifImage{}
	f = &FlifImage{
		images: []*wrapper.CflifImage{cimg},
	}
	h = 1
	data, err = f.ReadRowRGBA16(1, 0)
	if err != ErrOutOfRange {
		t.Errorf("another error occurred:%v", err)
	}

	// error out row
	cimg = &wrapper.CflifImage{}
	f = &FlifImage{
		images: []*wrapper.CflifImage{cimg},
	}
	h = 1
	data, err = f.ReadRowRGBA16(-1, 0)
	if err != ErrOutOfRange {
		t.Errorf("another error occurred:%v", err)
	}

	// error image index
	cimg = &wrapper.CflifImage{}
	f = &FlifImage{
		images: []*wrapper.CflifImage{cimg},
	}
	h = 1
	data, err = f.ReadRowRGBA16(0, 1)
	if err != ErrOutOfRange {
		t.Errorf("another error occurred:%v", err)
	}

	// error image index
	cimg = &wrapper.CflifImage{}
	f = &FlifImage{
		images: []*wrapper.CflifImage{cimg},
	}
	h = 1
	data, err = f.ReadRowRGBA16(0, -1)
	if err != ErrOutOfRange {
		t.Errorf("another error occurred:%v", err)
	}

	// error get data
	cimg = &wrapper.CflifImage{}
	f = &FlifImage{
		images: []*wrapper.CflifImage{cimg},
	}
	h = 1
	data, err = f.ReadRowRGBA16(0, 0)
	if err != ErrUnknown {
		t.Errorf("another error occurred:%v", err)
	}
	if data != nil {
		t.Errorf("want %v, got %v", nil, data)
	}

	// success
	d = []byte{0x11, 0x22, 0x33, 0x44}
	cimg = &wrapper.CflifImage{}
	f = &FlifImage{
		images: []*wrapper.CflifImage{cimg},
	}
	h = 1
	data, err = f.ReadRowRGBA16(0, 0)
	if err != nil {
		t.Errorf("error occurred:%v", err)
	}
	if !reflect.DeepEqual(d, data) {
		t.Errorf("want %v, got %v", d, data)
	}
}
