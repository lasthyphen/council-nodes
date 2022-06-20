// Copyright (C) 2019-2021, Dijets, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package djtx

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/lasthyphen/dijetsgo/database"
	"github.com/lasthyphen/dijetsgo/database/memdb"
	"github.com/lasthyphen/dijetsgo/ids"
	"github.com/lasthyphen/dijetsgo/snow/choices"
)

func TestStatusState(t *testing.T) {
	assert := assert.New(t)
	id0 := ids.GenerateTestID()

	db := memdb.New()
	s := NewStatusState(db)

	_, err := s.GetStatus(id0)
	assert.Equal(database.ErrNotFound, err)

	_, err = s.GetStatus(id0)
	assert.Equal(database.ErrNotFound, err)

	err = s.PutStatus(id0, choices.Accepted)
	assert.NoError(err)

	status, err := s.GetStatus(id0)
	assert.NoError(err)
	assert.Equal(choices.Accepted, status)

	err = s.DeleteStatus(id0)
	assert.NoError(err)

	_, err = s.GetStatus(id0)
	assert.Equal(database.ErrNotFound, err)

	err = s.PutStatus(id0, choices.Accepted)
	assert.NoError(err)

	s = NewStatusState(db)

	status, err = s.GetStatus(id0)
	assert.NoError(err)
	assert.Equal(choices.Accepted, status)
}
