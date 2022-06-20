// Copyright (C) 2019-2021, Dijets, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package common

import (
	"github.com/lasthyphen/dijetsgo/ids"
)

// Fx wraps an instance of a feature extension
type Fx struct {
	ID ids.ID
	Fx interface{}
}
