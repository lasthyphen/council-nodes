// Copyright (C) 2019-2021, Dijets, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package codec

import "github.com/lasthyphen/dijetsgo/utils/wrappers"

// Codec marshals and unmarshals
type Codec interface {
	MarshalInto(interface{}, *wrappers.Packer) error
	Unmarshal([]byte, interface{}) error
}
