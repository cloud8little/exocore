package network

import (
	sdkmath "cosmossdk.io/math"
	assetstypes "github.com/ExocoreNetwork/exocore/x/assets/types"
	delegationtypes "github.com/ExocoreNetwork/exocore/x/delegation/types"
	dogfoodtypes "github.com/ExocoreNetwork/exocore/x/dogfood/types"
	operatortypes "github.com/ExocoreNetwork/exocore/x/operator/types"
	oracletypes "github.com/ExocoreNetwork/exocore/x/oracle/types"
)

const (
	// TestEVMChainID represents the LayerZero chain ID for the test EVM chain
	TestEVMChainID = 101
	// EVMAddressLength is the standard length of EVM addresses in bytes
	EVMAddressLength = 20
)

var (
	// DefaultGenStateAssets only includes two assets, one for ETH and the other for NST ETH
	// For the contract address of asset-ETH we filled with the address of USDT, that's ok for test
	// we bond both tokens to the price of ETH in oracle module
	DefaultGenStateAssets = assetstypes.GenesisState{
		Params: assetstypes.Params{
			ExocoreLzAppAddress:    "0x3e108c058e8066da635321dc3018294ca82ddedf",
			ExocoreLzAppEventTopic: assetstypes.DefaultParams().ExocoreLzAppEventTopic,
		},
		ClientChains: []assetstypes.ClientChainInfo{
			{
				Name:             "Example EVM chain",
				MetaInfo:         "Example EVM chain metaInfo",
				LayerZeroChainID: TestEVMChainID,
				AddressLength:    EVMAddressLength,
			},
		},
		Tokens: []assetstypes.StakingAssetInfo{
			NewTestToken("ETH", "Ethereum native token", "0xdac17f958d2ee523a2206206994597c13d831ec7", TestEVMChainID, 5000),
			NewTestToken("NST ETH", "native restaking ETH", "0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee", TestEVMChainID, 5000),
		},
	}

	DefaultGenStateOperator = operatortypes.GenesisState{}
	// DefaultGenStateOperator = *operatortypes.DefaultGenesis()

	DefaultGenStateDelegation = delegationtypes.GenesisState{}
	// DefaultGenStateDelegation = *delegationtypes.DefaultGenesis()

	DefaultGenStateDogfood = *dogfoodtypes.DefaultGenesis()

	DefaultGenStateOracle = *oracletypes.DefaultGenesis()
)

func init() {
	// bond assetsIDs of ETH, NSTETH to ETH price
	DefaultGenStateOracle.Params.Tokens[1].AssetID = "0xdac17f958d2ee523a2206206994597c13d831ec7_0x65,0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee_0x65"
	// set ETH tokenfeeder's 'StartBaseBlock' to 10
	DefaultGenStateOracle.Params.TokenFeeders[1].StartBaseBlock = 10
}

func NewTestToken(name, metaInfo, address string, chainID uint64, amount int64) assetstypes.StakingAssetInfo {
	return assetstypes.StakingAssetInfo{
		AssetBasicInfo: assetstypes.AssetInfo{
			Name:             name,
			MetaInfo:         metaInfo,
			Address:          address,
			LayerZeroChainID: chainID,
		},
		StakingTotalAmount: sdkmath.NewInt(amount),
	}
}
