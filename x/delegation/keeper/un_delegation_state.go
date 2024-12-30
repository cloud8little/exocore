package keeper

import (
	"fmt"
	"github.com/ExocoreNetwork/exocore/utils"
	"math"
	"strings"

	errorsmod "cosmossdk.io/errors"

	"github.com/ExocoreNetwork/exocore/x/delegation/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// AllUndelegations function returns all the undelegation records in the module.
// It is used during `ExportGenesis` to export the undelegation records.
func (k Keeper) AllUndelegations(ctx sdk.Context) (undelegations []types.UndelegationRecord, err error) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixUndelegationInfo)
	iterator := sdk.KVStorePrefixIterator(store, []byte{})
	defer iterator.Close()

	ret := make([]types.UndelegationRecord, 0)
	for ; iterator.Valid(); iterator.Next() {
		var undelegation types.UndelegationRecord
		k.cdc.MustUnmarshal(iterator.Value(), &undelegation)
		ret = append(ret, undelegation)
	}
	return ret, nil
}

// SetUndelegationRecords stores the provided undelegation records.
// The records are stored with 3 different keys:
// (1) recordKey == blockNumber + undelegationID + txHash + operatorAddress => record
// (2) operatorAccAddr + stakerID + assetID + undelegationID => recordKey
// (3) epochIdentifierLength + completedEpochIdentifier + completedEpochNumber + UndelegationID => recordKey
// If a record exists with the same key, it will be overwritten; however, that is not a big
// concern since the lzNonce and txHash are unique for each record.
func (k *Keeper) SetUndelegationRecords(ctx sdk.Context, records []types.UndelegationRecord) error {
	singleRecordStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixUndelegationInfo)
	stakerUndelegationStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixStakerUndelegationInfo)
	pendingUndelegationStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixPendingUndelegations)
	for i := range records {
		record := records[i]
		epochInfo, exist := k.epochsKeeper.GetEpochInfo(ctx, record.CompletedEpochIdentifier)
		if !exist {
			return errorsmod.Wrapf(types.ErrEpochIdentifierNotExist, "identifier:%s", record.CompletedEpochIdentifier)
		}
		if record.CompletedEpochNumber < epochInfo.CurrentEpoch {
			return errorsmod.Wrapf(types.ErrInvalidCompletionEpoch, "epochIdentifier:%s,currentEpochNumber:%d,CompleteEpochNumber:%d", record.CompletedEpochIdentifier, epochInfo.CurrentEpoch, record.CompletedEpochNumber)
		}
		bz := k.cdc.MustMarshal(&record)
		// todo: check if the following state can only be set once?
		singleRecKey := types.GetUndelegationRecordKey(record.BlockNumber, record.UndelegationId, record.TxHash, record.OperatorAddr)
		singleRecordStore.Set(singleRecKey, bz)

		stakerKey := types.GetStakerUndelegationRecordKey(record.StakerId, record.AssetId, record.UndelegationId)
		stakerUndelegationStore.Set(stakerKey, singleRecKey)

		pendingUndelegationKey := types.GetPendingUndelegationRecordKey(record.CompletedEpochIdentifier, record.CompletedEpochNumber, record.UndelegationId)
		pendingUndelegationStore.Set(pendingUndelegationKey, singleRecKey)
	}
	return nil
}

// DeleteUndelegationRecord deletes the undelegation record from the module.
// The deletion is performed from all the 3 stores.
func (k *Keeper) DeleteUndelegationRecord(ctx sdk.Context, record *types.UndelegationRecord) error {
	singleRecordStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixUndelegationInfo)
	stakerUndelegationStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixStakerUndelegationInfo)
	pendingUndelegationStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixPendingUndelegations)

	singleRecKey := types.GetUndelegationRecordKey(record.BlockNumber, record.UndelegationId, record.TxHash, record.OperatorAddr)
	singleRecordStore.Delete(singleRecKey)

	stakerKey := types.GetStakerUndelegationRecordKey(record.StakerId, record.AssetId, record.UndelegationId)
	stakerUndelegationStore.Delete(stakerKey)

	pendingUndelegationKey := types.GetPendingUndelegationRecordKey(record.CompletedEpochIdentifier, record.CompletedEpochNumber, record.UndelegationId)
	pendingUndelegationStore.Delete(pendingUndelegationKey)
	return nil
}

// GetUndelegationRecords returns the undelegation records for the provided record keys.
func (k *Keeper) GetUndelegationRecords(ctx sdk.Context, singleRecordKeys []string) (record []*types.UndelegationRecord, err error) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixUndelegationInfo)
	ret := make([]*types.UndelegationRecord, 0)
	for _, singleRecordKey := range singleRecordKeys {
		keyBytes := []byte(singleRecordKey)
		value := store.Get(keyBytes)
		if value == nil {
			return nil, errorsmod.Wrap(types.ErrNoKeyInTheStore, fmt.Sprintf("undelegation record key doesn't exist: key is %s", singleRecordKey))
		}
		undelegationRecord := types.UndelegationRecord{}
		k.cdc.MustUnmarshal(value, &undelegationRecord)
		ret = append(ret, &undelegationRecord)
	}
	return ret, nil
}

// IterateUndelegationsByOperator iterates over the undelegation records belonging to the
// provided operator and filter. If the filter is non-nil, it will only iterate over the
// records for which the block height is greater than or equal to the filter.
func (k *Keeper) IterateUndelegationsByOperator(
	ctx sdk.Context, operator string, heightFilter *uint64, isUpdate bool,
	opFunc func(undelegation *types.UndelegationRecord) error,
) error {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixUndelegationInfo)
	iterator := sdk.KVStorePrefixIterator(store, []byte(operator))
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		if heightFilter != nil {
			keyFields, err := types.ParseUndelegationRecordKey(iterator.Key())
			if err != nil {
				return err
			}
			if keyFields.BlockHeight < *heightFilter {
				continue
			}
		}
		undelegation := types.UndelegationRecord{}
		k.cdc.MustUnmarshal(iterator.Value(), &undelegation)
		err := opFunc(&undelegation)
		if err != nil {
			return err
		}

		if isUpdate {
			bz := k.cdc.MustMarshal(&undelegation)
			store.Set(iterator.Key(), bz)
		}
	}
	return nil
}

// GetStakerUndelegationRecKeys returns the undelegation record keys corresponding to the provided
// staker and asset.
func (k *Keeper) GetStakerUndelegationRecKeys(ctx sdk.Context, stakerID, assetID string) (recordKeyList []string, err error) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixStakerUndelegationInfo)
	iterator := sdk.KVStorePrefixIterator(store, []byte(strings.Join([]string{stakerID, assetID}, "/")))
	defer iterator.Close()

	ret := make([]string, 0)
	for ; iterator.Valid(); iterator.Next() {
		ret = append(ret, string(iterator.Value()))
	}
	return ret, nil
}

// GetStakerUndelegationRecords returns the undelegation records for the provided staker and asset.
func (k *Keeper) GetStakerUndelegationRecords(ctx sdk.Context, stakerID, assetID string) (records []*types.UndelegationRecord, err error) {
	recordKeys, err := k.GetStakerUndelegationRecKeys(ctx, stakerID, assetID)
	if err != nil {
		return nil, err
	}

	return k.GetUndelegationRecords(ctx, recordKeys)
}

// IterateUndelegationsByStakerAndAsset iterates over the undelegation records belonging to the provided
// stakerID and assetID. If the isUpdate is true, the undelegation record will be updated after the
// operation is performed.
func (k *Keeper) IterateUndelegationsByStakerAndAsset(
	ctx sdk.Context, stakerID, assetID string, isUpdate bool,
	opFunc func(undelegationKey string, undelegation *types.UndelegationRecord) (bool, error),
) error {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixStakerUndelegationInfo)
	iterator := sdk.KVStorePrefixIterator(store, types.IteratorPrefixForStakerAsset(stakerID, assetID))
	defer iterator.Close()
	undelegationInfoStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixUndelegationInfo)
	for ; iterator.Valid(); iterator.Next() {
		infoValue := undelegationInfoStore.Get(iterator.Value())
		if infoValue == nil {
			return errorsmod.Wrap(types.ErrNoKeyInTheStore, fmt.Sprintf("undelegation record key doesn't exist: key is %s", string(iterator.Value())))
		}
		undelegation := types.UndelegationRecord{}
		k.cdc.MustUnmarshal(infoValue, &undelegation)
		isBreak, err := opFunc(string(iterator.Value()), &undelegation)
		if err != nil {
			return err
		}
		if isUpdate {
			bz := k.cdc.MustMarshal(&undelegation)
			undelegationInfoStore.Set(iterator.Value(), bz)
		}
		if isBreak {
			break
		}
	}
	return nil
}

// GetCompletablePendingUndelegations returns the undelegation records scheduled to completed at the end
// of the block. The pending undelegations should be expired and aren't held
func (k *Keeper) GetCompletablePendingUndelegations(ctx sdk.Context) ([]*types.UndelegationRecord, error) {
	records := make([]*types.UndelegationRecord, 0)
	expiredUndelegationOpFunc := func(recordKey []byte, record *types.UndelegationRecord) error {
		if k.GetUndelegationHoldCount(ctx, recordKey) > 0 {
			// The expired pending undelegation is held, so do not complete it;
			// then, continue addressing the other expired pending undelegations.
			k.Logger(ctx).Info("GetCompletablePendingUndelegations: the expired pending undelegation is held",
				"staker", record.StakerID,
				"assset", record.AssetID,
				"operator", record.OperatorAddr,
				"submittedBlockNumber", record.BlockNumber,
				"undelegationId", record.UndelegationId)
			return nil
		} else {
			records = append(records, record)
		}
		return nil
	}
	err := k.IterateExpiredPendingUndelegations(ctx, expiredUndelegationOpFunc)
	if err != nil {
		return nil, err
	}
	return records, nil
}

func (k *Keeper) IterateExpiredPendingUndelegations(
	ctx sdk.Context,
	opFunc func(recordKey []byte, undelegationRecord *types.UndelegationRecord) error,
) error {
	pendingUndelegationStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixPendingUndelegations)
	undelegationStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixUndelegationInfo)
	// iterate all pending undelegations across multiple epochs.
	allEpochs := k.epochsKeeper.AllEpochInfos(ctx)
	for _, epochInfo := range allEpochs {
		prefix := utils.AppendMany(
			sdk.Uint64ToBigEndian(uint64(len(epochInfo.Identifier))),
			[]byte(epochInfo.Identifier))

		iterator := sdk.KVStorePrefixIterator(pendingUndelegationStore, prefix)
		for ; iterator.Valid(); iterator.Next() {
			pendingUndelegationKeys, err := types.ParsePendingUndelegationKey(iterator.Key())
			if err != nil {
				return err
			}
			if pendingUndelegationKeys.EpochNumber > uint64(epochInfo.CurrentEpoch) {
				// These pending undelegations aren't expired, break the iteration
				break
			}
			// call opFunc to execute some operations for the expired pending undelegations
			value := undelegationStore.Get(iterator.Value())
			if value == nil {
				return errorsmod.Wrap(types.ErrNoKeyInTheStore, fmt.Sprintf("undelegation record key doesn't exist: key is %x", iterator.Value()))
			}
			undelegation := types.UndelegationRecord{}
			k.cdc.MustUnmarshal(value, &undelegation)
			err = opFunc(iterator.Value(), &undelegation)
			if err != nil {
				return err
			}
		}
		iterator.Close()
	}
	return nil
}

// IncrementUndelegationHoldCount increments the hold count for the undelegation record key.
func (k Keeper) IncrementUndelegationHoldCount(ctx sdk.Context, recordKey []byte) error {
	prev := k.GetUndelegationHoldCount(ctx, recordKey)
	if prev == math.MaxUint64 {
		return types.ErrCannotIncHoldCount
	}
	now := prev + 1
	store := ctx.KVStore(k.storeKey)
	store.Set(types.GetUndelegationOnHoldKey(recordKey), sdk.Uint64ToBigEndian(now))
	return nil
}

// GetUndelegationHoldCount returns the hold count for the undelegation record key.
func (k *Keeper) GetUndelegationHoldCount(ctx sdk.Context, recordKey []byte) uint64 {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.GetUndelegationOnHoldKey(recordKey))
	return sdk.BigEndianToUint64(bz)
}

// DecrementUndelegationHoldCount decrements the hold count for the undelegation record key.
func (k Keeper) DecrementUndelegationHoldCount(ctx sdk.Context, recordKey []byte) error {
	prev := k.GetUndelegationHoldCount(ctx, recordKey)
	if prev == 0 {
		return types.ErrCannotDecHoldCount
	}
	now := prev - 1
	store := ctx.KVStore(k.storeKey)
	store.Set(types.GetUndelegationOnHoldKey(recordKey), sdk.Uint64ToBigEndian(now))
	return nil
}

// IncrementUndelegationID increments the global undelegation ID.
func (k Keeper) IncrementUndelegationID(ctx sdk.Context) error {
	prev := k.GetUndelegationID(ctx)
	if prev == math.MaxUint64 {
		return types.ErrCannotIncUndelegationID
	}
	now := prev + 1
	store := ctx.KVStore(k.storeKey)
	store.Set(types.KeyPrefixUndelegationID, sdk.Uint64ToBigEndian(now))
	return nil
}

// GetUndelegationID returns the global undelegation ID.
func (k *Keeper) GetUndelegationID(ctx sdk.Context) uint64 {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.KeyPrefixUndelegationID)
	if bz == nil {
		// use 0 as the initial undelegation ID
		return 0
	}
	return sdk.BigEndianToUint64(bz)
}
