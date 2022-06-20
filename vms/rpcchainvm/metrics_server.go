// Copyright (C) 2019-2021, Dijets, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package rpcchainvm

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/lasthyphen/dijetsgo/api/proto/vmproto"
)

func (vm *VMServer) Gather(context.Context, *emptypb.Empty) (*vmproto.GatherResponse, error) {
	mfs, err := vm.ctx.Metrics.Gather()
	return &vmproto.GatherResponse{MetricFamilies: mfs}, err
}
