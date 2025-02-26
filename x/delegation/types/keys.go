package types

import (
	"strings"

	assetstypes "github.com/ExocoreNetwork/exocore/x/assets/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// constants
const (
	// ModuleName module name
	ModuleName = "delegation"

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
	prefixRestakerDelegationInfo = iota + 1
	prefixStakersByOperator
	prefixUndelegationInfo

	prefixStakerUndelegationInfo

	prefixWaitCompleteUndelegations

	// add for dogfood
	prefixUndelegationOnHold
)

var (
	// KeyPrefixRestakerDelegationInfo restakerID = clientChainAddr+'_'+ExoCoreChainIndex
	// KeyPrefixRestakerDelegationInfo
	// key-value:
	// restakerID +'/'+assetID+'/'+operatorAddr -> DelegationAmounts
	KeyPrefixRestakerDelegationInfo = []byte{prefixRestakerDelegationInfo}

	// KeyPrefixStakersByOperator key->value: operatorAddr+'/'+assetID -> stakerList
	KeyPrefixStakersByOperator = []byte{prefixStakersByOperator}

	// KeyPrefixUndelegationInfo singleRecordKey = operatorAddr+'/'+BlockHeight+'/'+LzNonce+'/'+txHash
	// singleRecordKey -> UndelegationRecord
	KeyPrefixUndelegationInfo = []byte{prefixUndelegationInfo}
	// KeyPrefixStakerUndelegationInfo restakerID+'/'+assetID+'/'+LzNonce -> singleRecordKey
	KeyPrefixStakerUndelegationInfo = []byte{prefixStakerUndelegationInfo}
	// KeyPrefixWaitCompleteUndelegations completeHeight +'/'+LzNonce -> singleRecordKey
	KeyPrefixWaitCompleteUndelegations = []byte{prefixWaitCompleteUndelegations}
)

func GetDelegationStateIteratorPrefix(stakerID, assetID string) []byte {
	tmp := []byte(strings.Join([]string{stakerID, assetID}, "/"))
	tmp = append(tmp, '/')
	return tmp
}

func ParseStakerAssetIDAndOperatorAddrFromKey(key []byte) (keys *SingleDelegationInfoReq, err error) {
	stringList, err := assetstypes.ParseJoinedStoreKey(key, 3)
	if err != nil {
		return nil, err
	}
	return &SingleDelegationInfoReq{StakerID: stringList[0], AssetID: stringList[1], OperatorAddr: stringList[2]}, nil
}

func GetUndelegationRecordKey(blockHeight, lzNonce uint64, txHash string, operatorAddr string) []byte {
	return []byte(strings.Join([]string{operatorAddr, hexutil.EncodeUint64(blockHeight), hexutil.EncodeUint64(lzNonce), txHash}, "/"))
}

type UndelegationKeyFields struct {
	BlockHeight  uint64
	LzNonce      uint64
	TxHash       string
	OperatorAddr string
}

func ParseUndelegationRecordKey(key []byte) (field *UndelegationKeyFields, err error) {
	stringList, err := assetstypes.ParseJoinedStoreKey(key, 4)
	if err != nil {
		return nil, err
	}
	height, err := hexutil.DecodeUint64(stringList[1])
	if err != nil {
		return nil, err
	}
	lzNonce, err := hexutil.DecodeUint64(stringList[2])
	if err != nil {
		return nil, err
	}
	return &UndelegationKeyFields{
		OperatorAddr: stringList[0],
		BlockHeight:  height,
		LzNonce:      lzNonce,
		TxHash:       stringList[3],
	}, nil
}

func GetStakerUndelegationRecordKey(stakerID, assetID string, lzNonce uint64) []byte {
	return []byte(strings.Join([]string{stakerID, assetID, hexutil.EncodeUint64(lzNonce)}, "/"))
}

func GetWaitCompleteRecordKey(height, lzNonce uint64) []byte {
	return []byte(strings.Join([]string{hexutil.EncodeUint64(height), hexutil.EncodeUint64(lzNonce)}, "/"))
}

// GetUndelegationOnHoldKey add for dogfood
func GetUndelegationOnHoldKey(recordKey []byte) []byte {
	return append([]byte{prefixUndelegationOnHold}, recordKey...)
}
