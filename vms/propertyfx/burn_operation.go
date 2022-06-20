// Copyright (C) 2019-2021, Dijets, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package propertyfx

import (
	"github.com/lasthyphen/dijetsgo/snow"
	"github.com/lasthyphen/dijetsgo/vms/components/verify"
	"github.com/lasthyphen/dijetsgo/vms/secp256k1fx"
)

type BurnOperation struct {
	secp256k1fx.Input `serialize:"true"`
}

func (op *BurnOperation) InitCtx(ctx *snow.Context) {}

func (op *BurnOperation) Outs() []verify.State { return nil }
