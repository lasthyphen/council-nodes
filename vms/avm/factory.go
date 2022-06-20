// Copyright (C) 2019-2021, Dijets, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package avm

import (
	"github.com/lasthyphen/dijetsgo/snow"
)

type Factory struct {
	TxFee            uint64
	CreateAssetTxFee uint64
}

func (f *Factory) New(*snow.Context) (interface{}, error) {
	return &VM{Factory: *f}, nil
}
