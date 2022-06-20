// Copyright (C) 2019-2021, Dijets, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package nftfx

import (
	"github.com/lasthyphen/dijetsgo/vms/secp256k1fx"
)

type Credential struct {
	secp256k1fx.Credential `serialize:"true"`
}
