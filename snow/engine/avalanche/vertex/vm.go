// Copyright (C) 2019-2021, Dijets, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package vertex

import (
	"github.com/lasthyphen/dijetsgo/ids"
	"github.com/lasthyphen/dijetsgo/snow/consensus/snowstorm"
	"github.com/lasthyphen/dijetsgo/snow/engine/common"
)

// DAGVM defines the minimum functionality that an avalanche VM must
// implement
type DAGVM interface {
	common.VM

	// Return any transactions that have not been sent to consensus yet
	PendingTxs() []snowstorm.Tx

	// Convert a stream of bytes to a transaction or return an error
	ParseTx(tx []byte) (snowstorm.Tx, error)

	// Retrieve a transaction that was submitted previously
	GetTx(ids.ID) (snowstorm.Tx, error)
}
