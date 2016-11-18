// Copyright 2016 chrisfelesoid. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package flif

const flifHeader = "FLIF"

func Decode(r io.Reader) (image.Image, error) {
    return nil, nil
}

func DecodeConfig(r io.Reader) (image.Config, error) {
    return nil, nil
}

func init() {
    image.RegisterFormat("flif", flifHeader, Decode, DecodeConfig)
}