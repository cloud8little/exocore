package types

import (
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/ethereum/go-ethereum/common"
)

// constants
const (
	// module name
	ModuleName = "native_token"

	// StoreKey to be used when creating the KVStore
	StoreKey = ModuleName

	// RouterKey to be used for message routing
	RouterKey = ModuleName
)

// ModuleAddress is the native module address for EVM
var ModuleAddress common.Address

func init() {
	ModuleAddress = common.BytesToAddress(authtypes.NewModuleAddress(ModuleName).Bytes())
}

const (
	prefixReStakerExocoreAddr = iota + 1
)

// KeyPrefixReStakerExoCoreAddr restakerID = clientChainAddr+'_'+ExoCoreChainIndex
// KeyPrefixReStakerExoCoreAddr key-value: restakerID->exoCoreAddr
var KeyPrefixReStakerExoCoreAddr = []byte{prefixReStakerExocoreAddr}
