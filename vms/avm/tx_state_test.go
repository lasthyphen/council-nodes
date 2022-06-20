// Copyright (C) 2019-2021, Dijets, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package avm

import (
	"testing"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/stretchr/testify/assert"

	"github.com/lasthyphen/dijetsgo/database"
	"github.com/lasthyphen/dijetsgo/database/memdb"
	"github.com/lasthyphen/dijetsgo/ids"
	"github.com/lasthyphen/dijetsgo/utils/crypto"
	"github.com/lasthyphen/dijetsgo/utils/units"
	"github.com/lasthyphen/dijetsgo/vms/components/djtx"
	"github.com/lasthyphen/dijetsgo/vms/secp256k1fx"
)

func TestTxState(t *testing.T) {
	assert := assert.New(t)

	db := memdb.New()
	codec, err := staticCodec()
	assert.NoError(err)

	s := NewTxState(db, codec).(*txState)

	_, err = s.GetTx(ids.Empty)
	assert.Equal(database.ErrNotFound, err)

	tx := &Tx{UnsignedTx: &BaseTx{BaseTx: djtx.BaseTx{
		NetworkID:    networkID,
		BlockchainID: chainID,
		Ins: []*djtx.TransferableInput{{
			UTXOID: djtx.UTXOID{
				TxID:        ids.Empty,
				OutputIndex: 0,
			},
			Asset: djtx.Asset{ID: assetID},
			In: &secp256k1fx.TransferInput{
				Amt: 20 * units.KiloDjtx,
				Input: secp256k1fx.Input{
					SigIndices: []uint32{
						0,
					},
				},
			},
		}},
	}}}
	err = tx.SignSECP256K1Fx(codec, [][]*crypto.PrivateKeySECP256K1R{{keys[0]}})
	assert.NoError(err)

	err = s.PutTx(ids.Empty, tx)
	assert.NoError(err)

	loadedTx, err := s.GetTx(ids.Empty)
	assert.NoError(err)
	assert.Equal(tx.ID(), loadedTx.ID())

	s.txCache.Flush()

	loadedTx, err = s.GetTx(ids.Empty)
	assert.NoError(err)
	assert.Equal(tx.ID(), loadedTx.ID())

	err = s.DeleteTx(ids.Empty)
	assert.NoError(err)

	_, err = s.GetTx(ids.Empty)
	assert.Equal(database.ErrNotFound, err)

	s.txCache.Flush()

	_, err = s.GetTx(ids.Empty)
	assert.Equal(database.ErrNotFound, err)
}

func TestMeteredTxState(t *testing.T) {
	assert := assert.New(t)

	db := memdb.New()
	codec, err := staticCodec()
	assert.NoError(err)

	_, err = NewMeteredTxState(db, codec, prometheus.NewRegistry())
	assert.NoError(err)
}
