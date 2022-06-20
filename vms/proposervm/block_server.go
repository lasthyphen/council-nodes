// Copyright (C) 2019-2021, Dijets, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package proposervm

import (
	"github.com/lasthyphen/dijetsgo/ids"
	"github.com/lasthyphen/dijetsgo/snow/consensus/snowman"
	"github.com/lasthyphen/dijetsgo/vms/proposervm/indexer"
)

var _ indexer.BlockServer = &VM{}

// Note: this is a contention heavy call that should be avoided
// for frequent/repeated indexer ops
func (vm *VM) GetFullPostForkBlock(blkID ids.ID) (snowman.Block, error) {
	vm.ctx.Lock.Lock()
	defer vm.ctx.Lock.Unlock()

	return vm.getPostForkBlock(blkID)
}

func (vm *VM) Commit() error {
	vm.ctx.Lock.Lock()
	defer vm.ctx.Lock.Unlock()

	return vm.db.Commit()
}
