// Copyright (C) 2019-2021, Dijets, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package indexer

import (
	"github.com/lasthyphen/dijetsgo/database/versiondb"
	"github.com/lasthyphen/dijetsgo/ids"
	"github.com/lasthyphen/dijetsgo/snow/consensus/snowman"
)

// BlockServer represents all requests heightIndexer can issue
// against ProposerVM. All methods must be thread-safe.
type BlockServer interface {
	versiondb.Commitable

	// Note: this is a contention heavy call that should be avoided
	// for frequent/repeated indexer ops
	GetFullPostForkBlock(blkID ids.ID) (snowman.Block, error)
}
