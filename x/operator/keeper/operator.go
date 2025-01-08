package keeper

import (
	"fmt"
	"github.com/ExocoreNetwork/exocore/x/epochs/types"
	"golang.org/x/xerrors"
	"math"
	"strings"

	errorsmod "cosmossdk.io/errors"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	assetstype "github.com/ExocoreNetwork/exocore/x/assets/types"
	operatortypes "github.com/ExocoreNetwork/exocore/x/operator/types"
)

// SetOperatorInfo is used to store the operator's information on the chain.
// There is no current way implemented to delete an operator's registration or edit it.
// TODO: implement operator edit function, which should allow editing:
// approve address?
// commission, subject to limits and once within 24 hours.
// client chain earnings addresses (maybe append only?)
func (k *Keeper) SetOperatorInfo(
	ctx sdk.Context, addr string, info *operatortypes.OperatorInfo,
) (err error) {
	// #nosec G703 // already validated in `ValidateBasic`
	opAccAddr, err := sdk.AccAddressFromBech32(addr)
	if err != nil {
		return errorsmod.Wrap(err, "SetOperatorInfo: error occurred when parse acc address from Bech32")
	}
	// if already registered, this request should go to EditOperator.
	// TODO: EditOperator needs to be implemented.
	if k.IsOperator(ctx, opAccAddr) {
		return errorsmod.Wrap(
			operatortypes.ErrOperatorAlreadyExists,
			fmt.Sprintf("SetOperatorInfo: operator already exists, address: %s", opAccAddr),
		)
	}
	// TODO: add minimum commission rate module parameter and check that commission exceeds it.
	info.Commission.UpdateTime = ctx.BlockTime()

	if info.ClientChainEarningsAddr != nil {
		for _, data := range info.ClientChainEarningsAddr.EarningInfoList {
			if data.ClientChainEarningAddr == "" {
				return errorsmod.Wrap(
					operatortypes.ErrParameterInvalid,
					"SetOperatorInfo: client chain earning address is empty",
				)
			}
			if !k.assetsKeeper.ClientChainExists(ctx, data.LzClientChainID) {
				return errorsmod.Wrap(
					operatortypes.ErrParameterInvalid,
					"SetOperatorInfo: client chain not found",
				)
			}
		}
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), operatortypes.KeyPrefixOperatorInfo)
	bz := k.cdc.MustMarshal(info)
	store.Set(opAccAddr, bz)
	return nil
}

func (k *Keeper) OperatorInfo(ctx sdk.Context, addr string) (info *operatortypes.OperatorInfo, err error) {
	opAccAddr, err := sdk.AccAddressFromBech32(addr)
	if err != nil {
		return nil, errorsmod.Wrap(err, "GetOperatorInfo: error occurred when parse acc address from Bech32")
	}
	store := prefix.NewStore(ctx.KVStore(k.storeKey), operatortypes.KeyPrefixOperatorInfo)
	// key := common.HexToAddress(incentive.Contract)
	value := store.Get(opAccAddr)
	if value == nil {
		return nil, errorsmod.Wrap(operatortypes.ErrNoKeyInTheStore, fmt.Sprintf("GetOperatorInfo: key is %s", opAccAddr))
	}
	ret := operatortypes.OperatorInfo{}
	k.cdc.MustUnmarshal(value, &ret)
	return &ret, nil
}

// AllOperators return the list of all operators' detailed information
func (k *Keeper) AllOperators(ctx sdk.Context) []operatortypes.OperatorDetail {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), operatortypes.KeyPrefixOperatorInfo)
	iterator := sdk.KVStorePrefixIterator(store, nil)
	defer iterator.Close()

	ret := make([]operatortypes.OperatorDetail, 0)
	for ; iterator.Valid(); iterator.Next() {
		var operatorInfo operatortypes.OperatorInfo
		operatorAddr := sdk.AccAddress(iterator.Key())
		k.cdc.MustUnmarshal(iterator.Value(), &operatorInfo)
		ret = append(ret, operatortypes.OperatorDetail{
			OperatorAddress: operatorAddr.String(),
			OperatorInfo:    operatorInfo,
		})
	}
	return ret
}

func (k Keeper) IsOperator(ctx sdk.Context, addr sdk.AccAddress) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), operatortypes.KeyPrefixOperatorInfo)
	return store.Has(addr)
}

func (k *Keeper) HandleOptedInfo(ctx sdk.Context, operatorAddr, avsAddr string, handleFunc func(info *operatortypes.OptedInfo)) error {
	opAccAddr, err := sdk.AccAddressFromBech32(operatorAddr)
	if err != nil {
		return errorsmod.Wrap(err, "HandleOptedInfo: error occurred when parse acc address from Bech32")
	}
	store := prefix.NewStore(ctx.KVStore(k.storeKey), operatortypes.KeyPrefixOperatorOptedAVSInfo)
	infoKey := assetstype.GetJoinedStoreKey(operatorAddr, strings.ToLower(avsAddr))
	// get info from the store
	value := store.Get(infoKey)
	if value == nil {
		return errorsmod.Wrap(operatortypes.ErrNoKeyInTheStore, fmt.Sprintf("HandleOptedInfo: key is %s", opAccAddr))
	}
	info := &operatortypes.OptedInfo{}
	k.cdc.MustUnmarshal(value, info)
	// call the handleFunc
	handleFunc(info)
	// restore the info after handling
	bz := k.cdc.MustMarshal(info)
	store.Set(infoKey, bz)
	return nil
}

func (k *Keeper) SetOptedInfo(ctx sdk.Context, operatorAddr, avsAddr string, info *operatortypes.OptedInfo) error {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), operatortypes.KeyPrefixOperatorOptedAVSInfo)

	// check operator address validation
	_, err := sdk.AccAddressFromBech32(operatorAddr)
	if err != nil {
		return assetstype.ErrInvalidOperatorAddr
	}
	infoKey := assetstype.GetJoinedStoreKey(operatorAddr, strings.ToLower(avsAddr))

	bz := k.cdc.MustMarshal(info)
	store.Set(infoKey, bz)
	return nil
}

func (k *Keeper) GetOptedInfo(ctx sdk.Context, operatorAddr, avsAddr string) (info *operatortypes.OptedInfo, err error) {
	opAccAddr, err := sdk.AccAddressFromBech32(operatorAddr)
	if err != nil {
		return nil, errorsmod.Wrap(err, "GetOptedInfo: error occurred when parse acc address from Bech32")
	}
	store := prefix.NewStore(ctx.KVStore(k.storeKey), operatortypes.KeyPrefixOperatorOptedAVSInfo)
	infoKey := assetstype.GetJoinedStoreKey(operatorAddr, strings.ToLower(avsAddr))
	value := store.Get(infoKey)
	if value == nil {
		return nil, errorsmod.Wrap(operatortypes.ErrNoKeyInTheStore, fmt.Sprintf("GetOptedInfo: operator is %s, avs address is %s", opAccAddr, avsAddr))
	}

	ret := operatortypes.OptedInfo{}
	k.cdc.MustUnmarshal(value, &ret)
	return &ret, nil
}

func (k *Keeper) IsOptedIn(ctx sdk.Context, operatorAddr, avsAddr string) bool {
	optedInfo, err := k.GetOptedInfo(ctx, operatorAddr, avsAddr)
	if err != nil {
		return false
	}
	return optedInfo.OptedOutHeight == operatortypes.DefaultOptedOutHeight
}

func (k *Keeper) IsActive(ctx sdk.Context, operatorAddr sdk.AccAddress, avsAddr string) bool {
	optedInfo, err := k.GetOptedInfo(ctx, operatorAddr.String(), avsAddr)
	if err != nil {
		// not opted in
		return false
	}
	if optedInfo.OptedOutHeight != operatortypes.DefaultOptedOutHeight {
		// opted out
		return false
	}
	if optedInfo.Jailed {
		// frozen - either temporarily or permanently
		return false
	}
	return true
}

func (k *Keeper) IterateOptInfo(ctx sdk.Context, isUpdate bool, iteratePrefix []byte, opFunc func(key []byte, optedInfo *operatortypes.OptedInfo) error) error {
	// get all opted-in info
	store := prefix.NewStore(ctx.KVStore(k.storeKey), operatortypes.KeyPrefixOperatorOptedAVSInfo)
	iterator := sdk.KVStorePrefixIterator(store, iteratePrefix)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var optedInfo operatortypes.OptedInfo
		k.cdc.MustUnmarshal(iterator.Value(), &optedInfo)
		err := opFunc(iterator.Key(), &optedInfo)
		if err != nil {
			return err
		}
		if isUpdate {
			bz := k.cdc.MustMarshal(&optedInfo)
			store.Set(iterator.Key(), bz)
		}
	}
	return nil
}

func (k *Keeper) GetOptedInAVSForOperator(ctx sdk.Context, operatorAddr string) ([]string, error) {
	avsList := make([]string, 0)
	opFunc := func(key []byte, optedInfo *operatortypes.OptedInfo) error {
		if optedInfo.OptedOutHeight == operatortypes.DefaultOptedOutHeight {
			keys, err := assetstype.ParseJoinedStoreKey(key, 2)
			if err != nil {
				return err
			}
			avsList = append(avsList, keys[1])
		}
		return nil
	}
	err := k.IterateOptInfo(ctx, false, []byte(operatorAddr), opFunc)
	if err != nil {
		return nil, err
	}
	return avsList, nil
}

func (k *Keeper) GetImpactfulAVSForOperator(ctx sdk.Context, operatorAddr string) ([]string, error) {
	avsList := make([]string, 0)
	opFunc := func(key []byte, optedInfo *operatortypes.OptedInfo) error {
		keys, err := assetstype.ParseJoinedStoreKey(key, 2)
		avsAddr := keys[1]
		if err != nil {
			return err
		}
		// add AVS currently opting in to the operator's list.
		if optedInfo.OptedOutHeight == operatortypes.DefaultOptedOutHeight {
			avsList = append(avsList, avsAddr)
		} else {
			// Add AVS that have opted out but are still within the unbonding duration,
			// and therefore still affect the operator, to the list.
			// #nosec G115
			epochNumber, err := k.GetEpochNumberByOptOutHeight(ctx, avsAddr, int64(optedInfo.OptedOutHeight))
			if err != nil {
				return err
			}
			epochInfo, err := k.avsKeeper.GetAVSEpochInfo(ctx, avsAddr)
			if err != nil {
				return err
			}
			unbondingDuration, err := k.avsKeeper.GetAVSUnbondingDuration(ctx, avsAddr)
			if err != nil {
				return err
			}
			// #nosec G115
			if epochNumber >= epochInfo.CurrentEpoch-int64(unbondingDuration) {
				avsList = append(avsList, avsAddr)
			}
		}
		return nil
	}
	err := k.IterateOptInfo(ctx, false, []byte(operatorAddr), opFunc)
	if err != nil {
		return nil, err
	}
	return avsList, nil
}

func (k Keeper) GetUnbondingExpiration(ctx sdk.Context, operator sdk.AccAddress) (string, int64, error) {
	// get the impactful AVSs for the operator
	avsList, err := k.GetImpactfulAVSForOperator(ctx, operator.String())
	if err != nil {
		return "", 0, err
	}
	// calculate the maximum unbonding expiration
	// Using minutes and the next epoch number as the default unbonding expiration.
	// This requires the Exocore chain to have minute-level epoch functionality enabled.
	minuteEpochInfo, exist := k.epochsKeeper.GetEpochInfo(ctx, types.MinuteEpochID)
	if !exist {
		return "", 0, errorsmod.Wrapf(operatortypes.ErrEpochIdentifierNotExist, "identifier:%s", types.MinuteEpochID)
	}
	retEpochIdentifier := types.MinuteEpochID
	retEpochNumber := minuteEpochInfo.CurrentEpoch + 1
	maxDurationSeconds := uint64(minuteEpochInfo.Duration)
	for _, avs := range avsList {
		epochInfo, err := k.avsKeeper.GetAVSEpochInfo(ctx, avs)
		if err != nil {
			return "", 0, err
		}
		unbondingDuration, err := k.avsKeeper.GetAVSUnbondingDuration(ctx, avs)
		if err != nil {
			return "", 0, err
		}
		if unbondingDuration+uint64(epochInfo.CurrentEpoch) > uint64(math.MaxInt64) {
			return "", 0, xerrors.New("the sum of unbondingDuration and the current epoch number exceeds the int64 range.")
		}
		durationSeconds := unbondingDuration * uint64(epochInfo.Duration)
		if durationSeconds > maxDurationSeconds {
			retEpochIdentifier = epochInfo.Identifier
			retEpochNumber = epochInfo.CurrentEpoch + int64(unbondingDuration)
			maxDurationSeconds = durationSeconds
		}
	}
	return retEpochIdentifier, retEpochNumber, nil
}

func (k *Keeper) SetAllOptedInfo(ctx sdk.Context, optedStates []operatortypes.OptedState) error {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), operatortypes.KeyPrefixOperatorOptedAVSInfo)
	for i := range optedStates {
		state := optedStates[i]
		bz := k.cdc.MustMarshal(&state.OptInfo)
		store.Set([]byte(state.Key), bz)
	}
	return nil
}

func (k *Keeper) GetAllOptedInfo(ctx sdk.Context) ([]operatortypes.OptedState, error) {
	ret := make([]operatortypes.OptedState, 0)
	opFunc := func(key []byte, optedInfo *operatortypes.OptedInfo) error {
		ret = append(ret, operatortypes.OptedState{
			Key:     string(key),
			OptInfo: *optedInfo,
		})
		return nil
	}
	err := k.IterateOptInfo(ctx, false, []byte{}, opFunc)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func (k *Keeper) GetOptedInOperatorListByAVS(ctx sdk.Context, avsAddr string) ([]string, error) {
	operatorList := make([]string, 0)
	opFunc := func(key []byte, optedInfo *operatortypes.OptedInfo) error {
		if optedInfo.OptedOutHeight == operatortypes.DefaultOptedOutHeight {
			keys, err := assetstype.ParseJoinedStoreKey(key, 2)
			if err != nil {
				return err
			}
			if strings.ToLower(avsAddr) == keys[1] {
				operatorList = append(operatorList, keys[0])
			}
		}
		return nil
	}
	err := k.IterateOptInfo(ctx, false, []byte{}, opFunc)
	if err != nil {
		return nil, err
	}
	return operatorList, nil
}
