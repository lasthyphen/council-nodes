// Copyright (C) 2019-2021, Dijets, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package snowman

import (
	"github.com/lasthyphen/dijetsgo/snow"
	"github.com/lasthyphen/dijetsgo/snow/consensus/snowball"
	"github.com/lasthyphen/dijetsgo/snow/consensus/snowman"
	"github.com/lasthyphen/dijetsgo/snow/engine/common"
	"github.com/lasthyphen/dijetsgo/snow/engine/snowman/block"
	"github.com/lasthyphen/dijetsgo/snow/validators"
)

// Config wraps all the parameters needed for a snowman engine
type Config struct {
	common.AllGetsServer

	Ctx        *snow.ConsensusContext
	VM         block.ChainVM
	Sender     common.Sender
	Validators validators.Set
	Params     snowball.Parameters
	Consensus  snowman.Consensus
}
