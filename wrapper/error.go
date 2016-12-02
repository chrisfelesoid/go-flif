// Copyright 2016 chrisfelesoid. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package wrapper

import "errors"

var (
	ErrUnknown = errors.New("unknown error")
)

func checkResult(stat int) error {
	if stat == 0 {
		return ErrUnknown
	}
	return nil
}
