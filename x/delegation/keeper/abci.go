package keeper

import (
	assetstypes "github.com/ExocoreNetwork/exocore/x/assets/types"
	"github.com/ExocoreNetwork/exocore/x/delegation/types"
	abci "github.com/cometbft/cometbft/abci/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// EndBlock : Completes expired pending undelegation events based on epoch information.
// This function is triggered at the end of every block. It queries the completable
// pending undelegations and completes them.
// We use EndBlock instead of epoch hooks to trigger completion because we want
// expired pending undelegations that are still held to be completed per block,
// rather than per epoch.
// However, the implementation can be switched to an epoch hook if it is deemed
// better to complete expired pending undelegations that are still held by epoch.
func (k *Keeper) EndBlock(
	originalCtx sdk.Context, _ abci.RequestEndBlock,
) []abci.ValidatorUpdate {
	logger := k.Logger(originalCtx)
	records, err := k.GetCompletablePendingUndelegations(originalCtx)
	if err != nil {
		// When encountering an error while retrieving pending undelegation, skip the undelegation at the given height without causing the node to stop running.
		logger.Error("Error in GetCompletablePendingUndelegations during the delegation's EndBlock execution", "error", err)
		return []abci.ValidatorUpdate{}
	}
	if len(records) == 0 {
		return []abci.ValidatorUpdate{}
	}
	for i := range records {
		record := records[i] // avoid implicit memory aliasing
		cc, writeCache := originalCtx.CacheContext()
		// we can use `Must` here because we stored this record ourselves.
		operatorAccAddress := sdk.MustAccAddressFromBech32(record.OperatorAddr)
		// TODO check if the operator has been slashed or frozen

		recordAmountNeg := record.Amount.Neg()
		// update delegation state
		deltaAmount := &types.DeltaDelegationAmounts{
			WaitUndelegationAmount: recordAmountNeg,
		}
		_, err = k.UpdateDelegationState(cc, record.StakerId, record.AssetId, record.OperatorAddr, deltaAmount)
		if err != nil {
			logger.Error("Error in UpdateDelegationState during the delegation's EndBlock execution", "error", err)
			continue
		}

		// update the staker state
		if record.AssetId == assetstypes.ExocoreAssetID {
			stakerAddrHex, _, err := assetstypes.ParseID(record.StakerId)
			if err != nil {
				logger.Error(
					"failed to parse staker ID",
					"error", err,
				)
				continue
			}
			stakerAddrBytes, err := hexutil.Decode(stakerAddrHex)
			if err != nil {
				logger.Error(
					"failed to decode staker address",
					"error", err,
				)
				continue
			}
			stakerAddr := sdk.AccAddress(stakerAddrBytes)
			if err := k.bankKeeper.UndelegateCoinsFromModuleToAccount(
				cc, types.DelegatedPoolName, stakerAddr,
				sdk.NewCoins(
					sdk.NewCoin(assetstypes.ExocoreAssetDenom, record.ActualCompletedAmount),
				),
			); err != nil {
				logger.Error(
					"failed to undelegate coins from module to account",
					"error", err,
				)
				continue
			}
		} else {
			err = k.assetsKeeper.UpdateStakerAssetState(cc, record.StakerId, record.AssetId, assetstypes.DeltaStakerSingleAsset{
				WithdrawableAmount:        record.ActualCompletedAmount,
				PendingUndelegationAmount: recordAmountNeg,
			})
			if err != nil {
				logger.Error("Error in UpdateStakerAssetState during the delegation's EndBlock execution", "error", err)
				continue
			}
		}

		// update the operator state
		err = k.assetsKeeper.UpdateOperatorAssetState(cc, operatorAccAddress, record.AssetId, assetstypes.DeltaOperatorSingleAsset{
			PendingUndelegationAmount: recordAmountNeg,
		})
		if err != nil {
			logger.Error("Error in UpdateOperatorAssetState during the delegation's EndBlock execution", "error", err)
			continue
		}

		// delete the Undelegation records that have been completed
		err = k.DeleteUndelegationRecord(cc, record)
		if err != nil {
			logger.Error("Error in DeleteUndelegationRecord during the delegation's EndBlock execution", "error", err)
			continue
		}
		// when calling `writeCache`, events are automatically emitted on the parent context
		writeCache()
	}
	return []abci.ValidatorUpdate{}
}
