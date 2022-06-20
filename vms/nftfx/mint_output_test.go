// Copyright (C) 2019-2021, Dijets, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package nftfx

import (
	"testing"

	"github.com/lasthyphen/dijetsgo/vms/components/verify"
)

func TestMintOutputState(t *testing.T) {
	intf := interface{}(&MintOutput{})
	if _, ok := intf.(verify.State); !ok {
		t.Fatalf("should be marked as state")
	}
}
