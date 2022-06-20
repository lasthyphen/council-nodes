// Copyright (C) 2019-2021, Dijets, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package primary

import (
	"context"
	"fmt"
	"time"

	"github.com/lasthyphen/dijetsgo/genesis"
	"github.com/lasthyphen/dijetsgo/ids"
	"github.com/lasthyphen/dijetsgo/utils/constants"
	"github.com/lasthyphen/dijetsgo/utils/units"
	"github.com/lasthyphen/dijetsgo/vms/components/djtx"
	"github.com/lasthyphen/dijetsgo/vms/secp256k1fx"
)

func ExampleWallet() {
	ctx := context.Background()
	kc := secp256k1fx.NewKeychain(genesis.EWOQKey)

	// NewWallet fetches the available UTXOs owned by [kc] on the network that
	// [LocalAPIURI] is hosting.
	walletSyncStartTime := time.Now()
	wallet, err := NewWallet(ctx, LocalAPIURI, kc)
	if err != nil {
		fmt.Printf("failed to initialize wallet with: %s\n", err)
		return
	}
	fmt.Printf("synced wallet in %s\n", time.Since(walletSyncStartTime))

	// Get the P-chain and the X-chain wallets
	pWallet := wallet.P()
	xWallet := wallet.X()

	// Pull out useful constants to use when issuing transactions.
	xChainID := xWallet.BlockchainID()
	djtxAssetID := xWallet.DJTXAssetID()
	owner := &secp256k1fx.OutputOwners{
		Threshold: 1,
		Addrs: []ids.ShortID{
			genesis.EWOQKey.PublicKey().Address(),
		},
	}

	// Send 100 schmeckles to the P-chain.
	exportStartTime := time.Now()
	exportTxID, err := xWallet.IssueExportTx(
		constants.PlatformChainID,
		[]*djtx.TransferableOutput{
			{
				Asset: djtx.Asset{
					ID: djtxAssetID,
				},
				Out: &secp256k1fx.TransferOutput{
					Amt:          100 * units.Schmeckle,
					OutputOwners: *owner,
				},
			},
		},
	)
	if err != nil {
		fmt.Printf("failed to issue X->P export transaction with: %s\n", err)
		return
	}
	fmt.Printf("issued X->P export %s in %s\n", exportTxID, time.Since(exportStartTime))

	// Import the 100 schmeckles from the X-chain into the P-chain.
	importStartTime := time.Now()
	importTxID, err := pWallet.IssueImportTx(xChainID, owner)
	if err != nil {
		fmt.Printf("failed to issue X->P import transaction with: %s\n", err)
		return
	}
	fmt.Printf("issued X->P import %s in %s\n", importTxID, time.Since(importStartTime))
}
