package keeper

import (
	"fmt"
	"math/big"
	"slices"
	"strconv"
	"strings"

	errorsmod "cosmossdk.io/errors"
	sdkmath "cosmossdk.io/math"

	"github.com/ExocoreNetwork/exocore/x/avs/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/evmos/evmos/v16/x/evm/statedb"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetAVSSupportedAssets returns a map of assets supported by the AVS. The avsAddr supplied must be hex.
func (k *Keeper) GetAVSSupportedAssets(ctx sdk.Context, avsAddr string) (map[string]interface{}, error) {
	if !common.IsHexAddress(avsAddr) {
		return nil, errorsmod.Wrap(types.ErrInvalidAddr, fmt.Sprintf("GetAVSSupportedAssets: key is %s", avsAddr))
	}
	avsInfo, err := k.GetAVSInfo(ctx, avsAddr)
	if err != nil {
		return nil, errorsmod.Wrap(err, fmt.Sprintf("GetAVSSupportedAssets: key is %s", avsAddr))
	}
	assetIDList := avsInfo.Info.AssetIDs
	ret := make(map[string]interface{})

	for _, assetID := range assetIDList {
		asset, err := k.assetsKeeper.GetStakingAssetInfo(ctx, assetID)
		if err != nil {
			return nil, errorsmod.Wrap(err, fmt.Sprintf("[GetAVSSupportedAssets] GetStakingAssetInfo: key is %s", assetID))
		}
		ret[assetID] = asset
	}

	return ret, nil
}

// GetAVSSlashContract returns the address of the contract that will be used to slash the AVS.
// The avsAddr supplied must be hex.
func (k *Keeper) GetAVSSlashContract(ctx sdk.Context, avsAddr string) (string, error) {
	avsInfo, err := k.GetAVSInfo(ctx, avsAddr)
	if err != nil {
		return "", errorsmod.Wrap(err, fmt.Sprintf("GetAVSSlashContract: key is %s", avsAddr))
	}
	if avsInfo.Info.SlashAddr == (common.Address{}).String() {
		return "", nil
	}
	return avsInfo.Info.SlashAddr, nil
}

// GetAVSMinimumSelfDelegation returns the minimum self-delegation required for the AVS, on a per-operator basis.
// The avsAddr supplied must be hex.
func (k *Keeper) GetAVSMinimumSelfDelegation(ctx sdk.Context, avsAddr string) (sdkmath.LegacyDec, error) {
	avsInfo, err := k.GetAVSInfo(ctx, avsAddr)
	if err != nil {
		return sdkmath.LegacyZeroDec(), errorsmod.Wrap(err, fmt.Sprintf("GetAVSMinimumSelfDelegation: key is %s", avsAddr))
	}
	// #nosec G115
	return sdkmath.LegacyNewDec(int64(avsInfo.Info.MinSelfDelegation)), nil
}

// GetAVSUnbondingDuration returns the unbonding number of epochs for an AVS. The name is a misnomer,
// since it is not the duration but the number of epochs.
func (k *Keeper) GetAVSUnbondingDuration(ctx sdk.Context, avsAddr string) (uint64, error) {
	avsInfo, err := k.GetAVSInfo(ctx, avsAddr)
	if err != nil {
		return 0, errorsmod.Wrap(err, fmt.Sprintf("GetAVSUnbondingDuration: key is %s", avsAddr))
	}
	return avsInfo.Info.AvsUnbondingPeriod, nil
}

// GetEpochEndAVSs returns a list of hex AVS addresses for AVSs which are scheduled to start at the end of the
// current epoch, or the beginning of the next one. The address format returned is hex.
func (k *Keeper) GetEpochEndAVSs(ctx sdk.Context, epochIdentifier string, endingEpochNumber int64) []string {
	var avsList []string
	k.IterateAVSInfo(ctx, func(_ int64, avsInfo types.AVSInfo) (stop bool) {
		// consider the dogfood AVS as an example. it is scheduled to start at epoch 0.
		// the currentEpoch is 1, so we will return it.
		// consider another AVS which will start at epoch 5. the current epoch is 4.
		// it should be returned here, since the operator module should start tracking this.
		// #nosec G115
		if epochIdentifier == avsInfo.EpochIdentifier && endingEpochNumber >= int64(avsInfo.StartingEpoch)-1 {
			avsList = append(avsList, strings.ToLower(avsInfo.AvsAddress))
		}
		return false
	})

	return avsList
}

// GetAVSInfoByTaskAddress returns the AVS  which containing this task address
// A task contract address can only be used by one avs
// TODO:this function is frequently used while its implementation iterates over existing avs to find the target avs by task contract address,  we should use a reverse mapping to avoid iteration
func (k *Keeper) GetAVSInfoByTaskAddress(ctx sdk.Context, taskAddr string) types.AVSInfo {
	var avs types.AVSInfo
	if taskAddr == "" || taskAddr == (common.Address{}).String() {
		return avs
	}
	taskAddr = strings.ToLower(taskAddr)
	k.IterateAVSInfo(ctx, func(_ int64, avsInfo types.AVSInfo) (stop bool) {
		if taskAddr == avsInfo.GetTaskAddr() {
			avs = avsInfo
			return true // stop because we found the AVS
		}
		return false
	})
	return avs
}

// GetTaskStatisticalEpochEndAVSs returns the task list where the current block marks the end of their statistical period.
func (k *Keeper) GetTaskStatisticalEpochEndAVSs(ctx sdk.Context, epochIdentifier string, epochNumber int64) []types.TaskResultInfo {
	var taskResList []types.TaskResultInfo
	k.IterateResultInfo(ctx, func(_ int64, info types.TaskResultInfo) (stop bool) {
		avsInfo := k.GetAVSInfoByTaskAddress(ctx, info.TaskContractAddress)
		if avsInfo.AvsAddress == "" {
			return false
		}
		taskInfo, err := k.GetTaskInfo(ctx, strconv.FormatUint(info.TaskId, 10), info.TaskContractAddress)
		if err != nil {
			return false
		}
		// Determine if the statistical period has passed, the range of the statistical period is the num marked (StartingEpoch) add TaskStatisticalPeriod
		// #nosec G115
		if epochIdentifier == avsInfo.EpochIdentifier && epochNumber ==
			// #nosec G115
			int64(taskInfo.StartingEpoch)+int64(taskInfo.TaskResponsePeriod)+int64(taskInfo.TaskStatisticalPeriod) {
			taskResList = append(taskResList, info)
		}
		return false
	})
	return taskResList
}

// RegisterAVSWithChainID registers an AVS by its chainID.
// It is responsible for generating an AVS address based on the chainID.
// The following bare minimum parameters must be supplied:
// AssetIDs, EpochsUntilUnbonded, EpochIdentifier, MinSelfDelegation and StartingEpoch.
// This will ensure compatibility with all of the related AVS functions, like
// GetEpochEndAVSs, GetAVSSupportedAssets, and GetAVSMinimumSelfDelegation.
func (k Keeper) RegisterAVSWithChainID(
	oCtx sdk.Context, params *types.AVSRegisterOrDeregisterParams,
) (exists bool, avsAddr common.Address, err error) {
	// guard against errors
	ctx, writeFunc := oCtx.CacheContext()
	// remove the version number and validate
	params.ChainID = types.ChainIDWithoutRevision(params.ChainID)
	if len(params.ChainID) == 0 {
		return false, common.Address{}, errorsmod.Wrap(types.ErrNotNull, "RegisterAVSWithChainID: chainID is null")
	}
	avsAddrStr := types.GenerateAVSAddr(params.ChainID)
	avsAddr = common.HexToAddress(avsAddrStr)
	// check that the AVS is registered
	if isAvs, _ := k.IsAVS(ctx, avsAddrStr); isAvs {
		// negligible probability that an independent AVS without this chainID exists
		return true, avsAddr, nil
	}
	defer func() {
		if err == nil && !exists {
			// store the reverse lookup from AVSAddress to ChainID
			// (the forward can be generated on the fly by hashing).
			k.SetAVSAddrToChainID(ctx, avsAddr, params.ChainID)
			// write the cache
			writeFunc()
		}
	}()
	// Mark the account as occupied by a contract, so that any transactions that originate
	// from it are rejected in `state_transition.go`. This protects us against the very
	// rare case of address collision.
	if err := k.evmKeeper.SetAccount(
		ctx, avsAddr,
		statedb.Account{
			Balance:  big.NewInt(0),
			CodeHash: types.ChainIDCodeHash[:],
			Nonce:    k.evmKeeper.GetNewContractNonce(ctx),
		},
	); err != nil {
		return false, common.Address{}, err
	}
	// SetAVSInfo expects HexAddress for the AvsAddress
	params.AvsAddress = avsAddr
	params.Action = types.RegisterAction

	if err := k.UpdateAVSInfo(ctx, params); err != nil {
		return false, common.Address{}, err
	}
	return false, avsAddr, nil
}

// SetAVSAddressToChainID stores a lookup from the generated AVS address to the chainID.
func (k Keeper) SetAVSAddrToChainID(ctx sdk.Context, avsAddr common.Address, chainID string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixAVSAddressToChainID)
	store.Set(avsAddr[:], []byte(chainID))
	// emit an event
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeChainAvsCreated,
			sdk.NewAttribute(types.AttributeKeyChainID, chainID),
			sdk.NewAttribute(types.AttributeKeyAvsAddress, avsAddr.String()),
		),
	)
}

// GetChainIDByAVSAddr returns the chainID for a given AVS address. It is a stateful
// function since it only returns the chainID if an AVS with the chainID was previously
// registered.
func (k Keeper) GetChainIDByAVSAddr(ctx sdk.Context, avsAddr string) (string, bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixAVSAddressToChainID)
	bz := store.Get(common.HexToAddress(avsAddr).Bytes())
	if bz == nil {
		return "", false
	}
	return string(bz), true
}

func (k *Keeper) GetAllChainIDInfos(ctx sdk.Context) ([]types.ChainIDInfo, error) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixAVSAddressToChainID)
	iterator := sdk.KVStorePrefixIterator(store, nil)
	defer iterator.Close()

	ret := make([]types.ChainIDInfo, 0)
	for ; iterator.Valid(); iterator.Next() {
		avsAddr := strings.ToLower(common.BytesToAddress(iterator.Key()).Hex())
		chainID := string(iterator.Value())

		ret = append(ret, types.ChainIDInfo{
			AvsAddress: avsAddr,
			ChainId:    chainID,
		})
	}
	return ret, nil
}

// IsWhitelisted checks if an operator is in the AVS whitelist.
// If the whitelist is empty, any operator is considered whitelisted.
// Returns true if the operator is whitelisted, false otherwise.
// Returns an error if the AVS address is invalid, the operator address is invalid,
// or if the operator is not in the whitelist.
func (k *Keeper) IsWhitelisted(ctx sdk.Context, avsAddr, operatorAddr string) (bool, error) {
	avsInfo, err := k.GetAVSInfo(ctx, avsAddr)
	if err != nil {
		return false, errorsmod.Wrap(err, fmt.Sprintf("IsWhitelisted: key is %s", avsAddr))
	}
	_, err = sdk.AccAddressFromBech32(operatorAddr)
	if err != nil {
		return false, errorsmod.Wrap(err, "IsWhitelisted: error occurred when parsing account acc address from Bech32")
	}
	// whitelist is disabled for avs of the dogfood type
	if len(avsInfo.Info.ChainId) != 0 && avsInfo.Info.AvsAddress == types.GenerateAVSAddr(avsInfo.Info.ChainId) {
		return true, nil
	}
	// Currently avs has no whitelist set and any operator can optin
	if len(avsInfo.Info.WhitelistAddress) == 0 {
		return true, nil
	}
	if !slices.Contains(avsInfo.Info.WhitelistAddress, operatorAddr) {
		return false, errorsmod.Wrap(err, fmt.Sprintf("operator %s not in whitelist", operatorAddr))
	}
	return true, nil
}
