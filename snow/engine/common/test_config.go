// Copyright (C) 2019-2021, Dijets, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package common

import (
	"github.com/lasthyphen/dijetsgo/ids"
	"github.com/lasthyphen/dijetsgo/snow"
	"github.com/lasthyphen/dijetsgo/snow/validators"
)

// DefaultConfigTest returns a test configuration
func DefaultConfigTest() Config {
	isBootstrapped := false
	subnet := &SubnetTest{
		IsBootstrappedF: func() bool { return isBootstrapped },
		BootstrappedF:   func(ids.ID) { isBootstrapped = true },
	}

	return Config{
		Ctx:                            snow.DefaultConsensusContextTest(),
		Validators:                     validators.NewSet(),
		Beacons:                        validators.NewSet(),
		Sender:                         &SenderTest{},
		Bootstrapable:                  &BootstrapableTest{},
		Subnet:                         subnet,
		Timer:                          &TimerTest{},
		AncestorsMaxContainersSent:     2000,
		AncestorsMaxContainersReceived: 2000,
		SharedCfg:                      &SharedConfig{},
	}
}
