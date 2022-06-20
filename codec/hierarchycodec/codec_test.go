// Copyright (C) 2019-2021, Dijets, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package hierarchycodec

import (
	"testing"

	"github.com/lasthyphen/dijetsgo/codec"
)

func TestVectors(t *testing.T) {
	for _, test := range codec.Tests {
		c := NewDefault()
		test(c, t)
	}
}
