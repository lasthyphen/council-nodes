// Copyright (C) 2019-2021, Dijets, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package avalanche

// Factory returns new instances of Consensus
type Factory interface {
	New() Consensus
}
