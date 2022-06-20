// Copyright (C) 2019-2021, Dijets, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package block

import (
	"github.com/lasthyphen/dijetsgo/codec"
	"github.com/lasthyphen/dijetsgo/codec/linearcodec"
	"github.com/lasthyphen/dijetsgo/utils/wrappers"
)

const version = 0

var c codec.Manager

func init() {
	lc := linearcodec.NewDefault()
	c = codec.NewDefaultManager()

	errs := wrappers.Errs{}
	errs.Add(
		lc.RegisterType(&statelessBlock{}),
		lc.RegisterType(&option{}),

		c.RegisterCodec(version, lc),
	)
	if errs.Errored() {
		panic(errs.Err)
	}
}
