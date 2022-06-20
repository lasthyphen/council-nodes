// Copyright (C) 2019-2021, Dijets, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package platformvm

import (
	"errors"

	"github.com/lasthyphen/dijetsgo/vms/components/djtx"
)

var (
	errInvalidLocktime      = errors.New("invalid locktime")
	errNestedStakeableLocks = errors.New("shouldn't nest stakeable locks")
)

type StakeableLockOut struct {
	Locktime             uint64 `serialize:"true" json:"locktime"`
	djtx.TransferableOut `serialize:"true" json:"output"`
}

func (s *StakeableLockOut) Addresses() [][]byte {
	if addressable, ok := s.TransferableOut.(djtx.Addressable); ok {
		return addressable.Addresses()
	}
	return nil
}

func (s *StakeableLockOut) Verify() error {
	if s.Locktime == 0 {
		return errInvalidLocktime
	}
	if _, nested := s.TransferableOut.(*StakeableLockOut); nested {
		return errNestedStakeableLocks
	}
	return s.TransferableOut.Verify()
}

type StakeableLockIn struct {
	Locktime            uint64 `serialize:"true" json:"locktime"`
	djtx.TransferableIn `serialize:"true" json:"input"`
}

func (s *StakeableLockIn) Verify() error {
	if s.Locktime == 0 {
		return errInvalidLocktime
	}
	if _, nested := s.TransferableIn.(*StakeableLockIn); nested {
		return errNestedStakeableLocks
	}
	return s.TransferableIn.Verify()
}
