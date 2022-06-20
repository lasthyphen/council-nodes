// Copyright (C) 2019-2021, Dijets, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package genesis

import (
	"time"

	"github.com/lasthyphen/dijetsgo/utils/units"
	"github.com/lasthyphen/dijetsgo/vms/platformvm/reward"
)

var (
	mainnetGenesisConfigJSON = `{
		"networkID": 1,
		"allocations": [
			{
				"ethAddr": "0x1B00f59fff05F6591c13e32740377eAE72661061",
				"djtxAddr": "X-dijets1uxkc262hvgsvstguvawwaanmsh8zvtcsd64l5e",
				"initialAmount": 2400000000000000,
				"unlockSchedule": [
					{
						"amount": 360000000000,
						"locktime": 1646352000
					}
				]
			},
			{
				"ethAddr": "0xd86b355443158939c2f1b2A00961F8453b33E74E",
				"djtxAddr": "X-dijets1v8wat5z4cxh7wh873d7n6d9m6mpnynr8sgl059",
				"initialAmount": 24000000000000000,
				"unlockSchedule": [
					{
						"amount": 360000000000,
						"locktime": 1646352000
					}
				]
			},
			{
				"ethAddr": "0x80231567cD6E270c8360B80b97034Ec26dad83b8",
				"djtxAddr": "X-dijets16yd4ams4xdp9c6ht9zfnp90225ukwmqnj964sw",
				"initialAmount": 2400000000000000,
				"unlockSchedule": [
					{
						"amount": 360000000000,
						"locktime": 1646352000
					}
				]
			},
			{
				"ethAddr": "0x2cf25c0323f52bb8d3f46ebc2dfcd34ba057b781",
				"djtxAddr": "X-dijets1mc6ar37ggvvh80ezkwnnyrkfummc32k0z9v9la",
				"initialAmount": 240000000000,
				"unlockSchedule": [
					{
						"amount": 360000000000,
						"locktime": 1646352000
					}
				]
			}
		],
		"startTime": 1599696000,
		"initialStakeDuration": 31536000,
		"initialStakeDurationOffset": 54000,
		"initialStakedFunds": [
			"X-dijets1uxkc262hvgsvstguvawwaanmsh8zvtcsd64l5e"
		],
		"initialStakers": [
			{
				"nodeID": "NodeID-7Xhw2mDxuDS44j42TCB6U5579esbSt3Lg",
				"rewardAddress": "X-dijets1uxkc262hvgsvstguvawwaanmsh8zvtcsd64l5e",
				"delegationFee": 1000000
			},
			{
				"nodeID": "NodeID-MFrZFVCXPv5iCn6M9K6XduxGTYp891xXZ",
				"rewardAddress": "X-dijets1uxkc262hvgsvstguvawwaanmsh8zvtcsd64l5e",
				"delegationFee": 500000
			},
			{
				"nodeID": "NodeID-NFBbbJ4qCmNaCzeW7sxErhvWqvEQMnYcN",
				"rewardAddress": "X-dijets1uxkc262hvgsvstguvawwaanmsh8zvtcsd64l5e",
				"delegationFee": 250000
			},
			{
				"nodeID": "NodeID-GWPcbFJZFfZreETSoWjPimr846mXEKCtu",
				"rewardAddress": "X-dijets1uxkc262hvgsvstguvawwaanmsh8zvtcsd64l5e",
				"delegationFee": 125000
			},
			{
				"nodeID": "NodeID-P7oB2McjBGgW2NXXWVYjV8JEDFoW9xDE5",
				"rewardAddress": "X-dijets1uxkc262hvgsvstguvawwaanmsh8zvtcsd64l5e",
				"delegationFee": 62500
			}
		],
		"cChainGenesis": "{\"config\":{\"chainId\":98200,\"homesteadBlock\":0,\"daoForkBlock\":0,\"daoForkSupport\":true,\"eip150Block\":0,\"eip150Hash\":\"0x2086799aeebeae135c246c65021c82b4e15a2c451340993aacfd2751886514f0\",\"eip155Block\":0,\"eip158Block\":0,\"byzantiumBlock\":0,\"constantinopleBlock\":0,\"petersburgBlock\":0,\"istanbulBlock\":0,\"muirGlacierBlock\":0},\"nonce\":\"0x0\",\"timestamp\":\"0x0\",\"extraData\":\"0x00\",\"gasLimit\":\"0x5f5e100\",\"difficulty\":\"0x0\",\"mixHash\":\"0x0000000000000000000000000000000000000000000000000000000000000000\",\"coinbase\":\"0x0000000000000000000000000000000000000000\",\"alloc\":{\"0100000000000000000000000000000000000000\":{\"code\":\"0x7300000000000000000000000000000000000000003014608060405260043610603d5760003560e01c80631e010439146042578063b6510bb314606e575b600080fd5b605c60048036036020811015605657600080fd5b503560b1565b60408051918252519081900360200190f35b818015607957600080fd5b5060af60048036036080811015608e57600080fd5b506001600160a01b03813516906020810135906040810135906060013560b6565b005b30cd90565b836001600160a01b031681836108fc8690811502906040516000604051808303818888878c8acf9550505050505015801560f4573d6000803e3d6000fd5b505050505056fea26469706673582212201eebce970fe3f5cb96bf8ac6ba5f5c133fc2908ae3dcd51082cfee8f583429d064736f6c634300060a0033\",\"balance\":\"0x0\"}},\"number\":\"0x0\",\"gasUsed\":\"0x0\",\"parentHash\":\"0x0000000000000000000000000000000000000000000000000000000000000000\"}",
		"message": "Pen Xen Ole Fren"
	}`

	// MainnetParams are the params used for mainnet
	MainnetParams = Params{
		TxFeeConfig: TxFeeConfig{
			TxFee:                 units.MilliDjtx,
			CreateAssetTxFee:      10 * units.MilliDjtx,
			CreateSubnetTxFee:     1 * units.Djtx,
			CreateBlockchainTxFee: 1 * units.Djtx,
		},
		StakingConfig: StakingConfig{
			UptimeRequirement: .8, // 80%
			MinValidatorStake: 2 * units.KiloDjtx,
			MaxValidatorStake: 3 * units.MegaDjtx,
			MinDelegatorStake: 25 * units.Djtx,
			MinDelegationFee:  20000, // 2%
			MinStakeDuration:  2 * 7 * 24 * time.Hour,
			MaxStakeDuration:  365 * 24 * time.Hour,
			RewardConfig: reward.Config{
				MaxConsumptionRate: .12 * reward.PercentDenominator,
				MinConsumptionRate: .10 * reward.PercentDenominator,
				MintingPeriod:      365 * 24 * time.Hour,
				SupplyCap:          100 * units.MegaDjtx,
			},
		},
	}
)
