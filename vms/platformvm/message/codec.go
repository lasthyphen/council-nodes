// Copyright (C) 2019-2021, Dijets, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package message

import (
	"github.com/lasthyphen/dijetsgo/codec"
	"github.com/lasthyphen/dijetsgo/codec/linearcodec"
	"github.com/lasthyphen/dijetsgo/codec/reflectcodec"
	"github.com/lasthyphen/dijetsgo/utils/units"
	"github.com/lasthyphen/dijetsgo/utils/wrappers"
)

const (
	codecVersion   uint16 = 0
	maxMessageSize        = 512 * units.KiB
	maxSliceLen           = maxMessageSize
)

// Codec does serialization and deserialization
var c codec.Manager

func init() {
	c = codec.NewManager(maxMessageSize)
	lc := linearcodec.New(reflectcodec.DefaultTagName, maxSliceLen)

	errs := wrappers.Errs{}
	errs.Add(
		lc.RegisterType(&Tx{}),
		c.RegisterCodec(codecVersion, lc),
	)
	if errs.Errored() {
		panic(errs.Err)
	}
}
