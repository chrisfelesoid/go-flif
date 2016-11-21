// Copyright 2016 chrisfelesoid. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package flif

import (
	"errors"
	"image"
	"image/color"
	"io"
	"io/ioutil"

	"github.com/chrisfelesoid/go-flif/wrapper"
)

const flifHeader = "FLIF"

const (
	chgray = 1
	chrgb  = 3
	chrgba = 4
)

func copyRow(fi *wrapper.FlifImage, img image.Image) error {
	h := fi.GetHeight()
	w := fi.GetWidth()
	for y := 0; y < h; y++ {
		buf := readRow(fi, y)
		for x := 0; x < w; x++ {
			switch chkimg := img.(type) {
			case *image.RGBA:
				r := buf[4*x+0]
				g := buf[4*x+1]
				b := buf[4*x+2]
				a := buf[4*x+3]
				chkimg.SetRGBA(x, y, color.RGBA{r, g, b, a})
			case *image.RGBA64:
				r := uint16(buf[8*x+0])<<8 | uint16(buf[8*x+1])
				g := uint16(buf[8*x+2])<<8 | uint16(buf[8*x+3])
				b := uint16(buf[8*x+4])<<8 | uint16(buf[8*x+5])
				a := uint16(buf[8*x+6])<<8 | uint16(buf[8*x+7])
				chkimg.SetRGBA64(x, y, color.RGBA64{r, g, b, a})
			default:
				return errors.New("fail copy pixels")
			}
		}
	}
	return nil
}

func readRow(fi *wrapper.FlifImage, row int) []byte {
	if fi.GetDepth() <= 8 {
		return fi.ReadRowRGBA8(row, 0)
	}
	return fi.ReadRowRGBA16(row, 0)
}

func Decode(r io.Reader) (image.Image, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	d := wrapper.NewFlifDecoder()
	d.DecodeMemory(b)
	fi := d.GetImage()

	var img image.Image
	if fi.GetDepth() <= 8 {
		img = image.NewRGBA(image.Rect(0, 0, fi.GetWidth(), fi.GetHeight()))
	} else {
		img = image.NewRGBA64(image.Rect(0, 0, fi.GetWidth(), fi.GetHeight()))
	}

	err = copyRow(fi, img)
	if err != nil {
		return nil, err
	}

	return img, nil
}

func DecodeConfig(r io.Reader) (image.Config, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return image.Config{}, err
	}
	info := wrapper.GetInfo(b)
	if info == nil {
		return image.Config{}, errors.New("cannnot get flif information")
	}

	var cm color.Model

	switch info.Channel {
	// internal color model. read only RGBA values.
	// case chgray:
	// 	if info.Depth <= 8 {
	// 		cm = color.GrayModel
	// 	} else {
	// 		cm = color.Gray16Model
	// 	}
	case chgray, chrgb, chrgba:
		if info.Depth <= 8 {
			cm = color.RGBAModel
		} else {
			cm = color.RGBA64Model
		}
	}
	return image.Config{
		ColorModel: cm,
		Width:      info.Width,
		Height:     info.Height,
	}, nil
}

func init() {
	image.RegisterFormat("flif", flifHeader, Decode, DecodeConfig)
}
