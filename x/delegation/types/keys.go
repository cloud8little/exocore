package types

import (
	"github.com/ExocoreNetwork/exocore/utils"
	"github.com/ExocoreNetwork/exocore/x/operator/types"
	"golang.org/x/xerrors"
	"strings"

	assetstypes "github.com/ExocoreNetwork/exocore/x/assets/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
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

	prefixPendingUndelegations

	// used to store the undelegation hold count
	prefixUndelegationOnHold

	prefixAssociatedOperatorByStaker

	prefixForUndelegationID
)

var (
	// KeyPrefixRestakerDelegationInfo restakerID = clientChainAddr+'_'+ExoCoreChainIndex
	// KeyPrefixRestakerDelegationInfo
	// key-value:
	// restakerID +'/'+assetID+'/'+operatorAddr -> DelegationAmounts
	KeyPrefixRestakerDelegationInfo = []byte{prefixRestakerDelegationInfo}

	// KeyPrefixStakersByOperator key->value: operatorAddr+'/'+assetID -> stakerList
	KeyPrefixStakersByOperator = []byte{prefixStakersByOperator}

	// KeyPrefixUndelegationID key-value:
	// prefixForUndelegationID -> uint64
	// We use an incrementing number to identify undelegations because we support different
	// assets across multiple client chains and the Exocore chain.
	KeyPrefixUndelegationID = []byte{prefixForUndelegationID}

	// KeyPrefixUndelegationInfo singleRecordKey = operatorAddr+BlockHeight+UndelegationID+txHash
	// it can be constructed by GetUndelegationRecordKey
	// singleRecordKey -> UndelegationRecord
	KeyPrefixUndelegationInfo = []byte{prefixUndelegationInfo}
	// KeyPrefixStakerUndelegationInfo restakerID+'/'+assetID+'/'+UndelegationID -> singleRecordKey
	KeyPrefixStakerUndelegationInfo = []byte{prefixStakerUndelegationInfo}
	// KeyPrefixPendingUndelegations
	// key=epochIdentifierLength+completedEpochIdentifier+completedEpochNumber+UndelegationID
	// it can be constructed by GetPendingUndelegationRecordKey
	// key -> singleRecordKey
	KeyPrefixPendingUndelegations = []byte{prefixPendingUndelegations}

	// KeyPrefixAssociatedOperatorByStaker stakerID -> operator address
	KeyPrefixAssociatedOperatorByStaker = []byte{prefixAssociatedOperatorByStaker}
)

func IteratorPrefixForStakerAsset(stakerID, assetID string) []byte {
	tmp := []byte(strings.Join([]string{stakerID, assetID}, "/"))
	tmp = append(tmp, '/')
	return tmp
}

func ParseStakerAssetIDAndOperator(key []byte) (keys *SingleDelegationInfoReq, err error) {
	stringList, err := assetstypes.ParseJoinedStoreKey(key, 3)
	if err != nil {
		return nil, err
	}
	return &SingleDelegationInfoReq{StakerId: stringList[0], AssetId: stringList[1], OperatorAddr: stringList[2]}, nil
}

// GetUndelegationRecordKey returns the key for the undelegation record. The caller must ensure that the parameters
// are valid; this function performs no validation whatsoever.
func GetUndelegationRecordKey(blockHeight, undelegationID uint64, txHash string, operatorAddr string) []byte {
	// we can use `Must` here because we stored this record ourselves.
	operatorAccAddress := sdk.MustAccAddressFromBech32(operatorAddr)
	return utils.AppendMany(
		// operator address,20bytes
		operatorAccAddress,
		// Append the height,8bytes
		sdk.Uint64ToBigEndian(blockHeight),
		// Append the undelegationID,8bytes
		sdk.Uint64ToBigEndian(undelegationID),
		// Append txHash,32bytes
		common.HexToHash(txHash).Bytes(),
	)
}

type UndelegationKeyFields struct {
	BlockHeight  uint64
	TxNonce      uint64
	TxHash       string
	OperatorAddr string
}

func ParseUndelegationRecordKey(key []byte) (field *UndelegationKeyFields, err error) {
	expectLength := types.AccAddressLength + 2*types.ByteLengthForUint64 + common.HashLength
	if len(key) != expectLength {
		return nil, xerrors.Errorf(
			"invalid undelegation record key, expectedLength:%d,actualLength:%d",
			expectLength, len(key))
	}
	// operator accAddress: 20bytes
	startIndex := 0
	operatorAccAddr := sdk.AccAddress(key[startIndex : startIndex+types.AccAddressLength])
	// the height type is uint64: 8bytes
	startIndex += types.AccAddressLength
	height := sdk.BigEndianToUint64(key[startIndex : startIndex+types.ByteLengthForUint64])
	// the nonce type is uint64: 8bytes
	startIndex += types.ByteLengthForUint64
	txNonce := sdk.BigEndianToUint64(key[startIndex : startIndex+types.ByteLengthForUint64])
	// txHash: 32bytes
	startIndex += types.ByteLengthForUint64
	txHash := common.BytesToHash(key[startIndex : startIndex+common.HashLength])
	return &UndelegationKeyFields{
		OperatorAddr: operatorAccAddr.String(),
		BlockHeight:  height,
		TxNonce:      txNonce,
		TxHash:       txHash.String(),
	}, nil
}

func GetStakerUndelegationRecordKey(stakerID, assetID string, lzNonce uint64) []byte {
	return []byte(strings.Join([]string{stakerID, assetID, hexutil.EncodeUint64(lzNonce)}, "/"))
}

type PendingUndelegationKeyFields struct {
	EpochIdentifier string
	EpochNumber     uint64
	TxNonce         uint64
}

func GetPendingUndelegationRecordKey(epochIdentifier string, epochNumber int64, nonce uint64) []byte {
	return utils.AppendMany(
		// length of identifier,8bytes
		sdk.Uint64ToBigEndian(uint64(len(epochIdentifier))),
		// epoch identifier, length = len(epochIdentifier)
		[]byte(epochIdentifier),
		// Append the epoch number,8bytes
		sdk.Uint64ToBigEndian(uint64(epochNumber)),
		// Append the nonce,8bytes
		sdk.Uint64ToBigEndian(nonce),
	)
}

func ParsePendingUndelegationKey(key []byte) (field *PendingUndelegationKeyFields, err error) {
	if len(key) <= 3*types.ByteLengthForUint64 {
		return nil, xerrors.New("ParsePendingUndelegationKey,key length is too short to contain epoch info and nonce")
	}
	identifierLen := sdk.BigEndianToUint64(key[0:types.ByteLengthForUint64])
	if uint64(len(key)) != uint64(3*types.ByteLengthForUint64)+identifierLen {
		return nil, xerrors.Errorf("ParsePendingUndelegationKey,key length is invalid,expect:%d,actual:%d", uint64(3*types.ByteLengthForUint64)+identifierLen, len(key))
	}
	epochIdentifier := string(key[types.ByteLengthForUint64 : types.ByteLengthForUint64+identifierLen])
	epochNumber := sdk.BigEndianToUint64(key[types.ByteLengthForUint64+identifierLen : types.ByteLengthForUint64*2+identifierLen])
	txNonce := sdk.BigEndianToUint64(key[types.ByteLengthForUint64*2+identifierLen:])
	return &PendingUndelegationKeyFields{
		EpochIdentifier: epochIdentifier,
		EpochNumber:     epochNumber,
		TxNonce:         txNonce,
	}, nil
}

// GetUndelegationOnHoldKey returns the key for the undelegation hold count
func GetUndelegationOnHoldKey(recordKey []byte) []byte {
	return append([]byte{prefixUndelegationOnHold}, recordKey...)
}
