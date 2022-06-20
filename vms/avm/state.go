// Copyright (C) 2019-2021, Dijets, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package avm

import (
	"github.com/prometheus/client_golang/prometheus"

	"github.com/lasthyphen/dijetsgo/cache"
	"github.com/lasthyphen/dijetsgo/codec"
	"github.com/lasthyphen/dijetsgo/database"
	"github.com/lasthyphen/dijetsgo/database/prefixdb"
	"github.com/lasthyphen/dijetsgo/vms/components/djtx"
)

const (
	txDeduplicatorSize = 8192
)

var (
	utxoStatePrefix            = []byte("utxo")
	statusStatePrefix          = []byte("status")
	singletonStatePrefix       = []byte("singleton")
	txStatePrefix              = []byte("tx")
	_                    State = &state{}
)

// State persistently maintains a set of UTXOs, transaction, statuses, and
// singletons.
type State interface {
	djtx.UTXOState
	djtx.StatusState
	djtx.SingletonState
	TxState

	DeduplicateTx(tx *UniqueTx) *UniqueTx
}

type state struct {
	djtx.UTXOState
	djtx.StatusState
	djtx.SingletonState
	TxState

	uniqueTxs cache.Deduplicator
}

func NewState(db database.Database, genesisCodec, codec codec.Manager) State {
	utxoDB := prefixdb.New(utxoStatePrefix, db)
	statusDB := prefixdb.New(statusStatePrefix, db)
	singletonDB := prefixdb.New(singletonStatePrefix, db)
	txDB := prefixdb.New(txStatePrefix, db)

	return &state{
		UTXOState:      djtx.NewUTXOState(utxoDB, codec),
		StatusState:    djtx.NewStatusState(statusDB),
		SingletonState: djtx.NewSingletonState(singletonDB),
		TxState:        NewTxState(txDB, genesisCodec),

		uniqueTxs: &cache.EvictableLRU{
			Size: txDeduplicatorSize,
		},
	}
}

func NewMeteredState(db database.Database, genesisCodec, codec codec.Manager, metrics prometheus.Registerer) (State, error) {
	utxoDB := prefixdb.New(utxoStatePrefix, db)
	statusDB := prefixdb.New(statusStatePrefix, db)
	singletonDB := prefixdb.New(singletonStatePrefix, db)
	txDB := prefixdb.New(txStatePrefix, db)

	utxoState, err := djtx.NewMeteredUTXOState(utxoDB, codec, metrics)
	if err != nil {
		return nil, err
	}

	statusState, err := djtx.NewMeteredStatusState(statusDB, metrics)
	if err != nil {
		return nil, err
	}

	txState, err := NewMeteredTxState(txDB, genesisCodec, metrics)
	return &state{
		UTXOState:      utxoState,
		StatusState:    statusState,
		SingletonState: djtx.NewSingletonState(singletonDB),
		TxState:        txState,

		uniqueTxs: &cache.EvictableLRU{
			Size: txDeduplicatorSize,
		},
	}, err
}

// UniqueTx de-duplicates the transaction.
func (s *state) DeduplicateTx(tx *UniqueTx) *UniqueTx {
	return s.uniqueTxs.Deduplicate(tx).(*UniqueTx)
}
