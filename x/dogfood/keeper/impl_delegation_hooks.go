package keeper

import (
	delegationtypes "github.com/ExocoreNetwork/exocore/x/delegation/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// DelegationHooksWrapper is the wrapper structure that implements the delegation hooks for the
// dogfood keeper.
type DelegationHooksWrapper struct {
	keeper *Keeper
}

// Interface guard
var _ delegationtypes.DelegationHooks = DelegationHooksWrapper{}

// DelegationHooks returns the delegation hooks wrapper. It follows the "accept interfaces,
// return concretes" pattern.
func (k *Keeper) DelegationHooks() DelegationHooksWrapper {
	return DelegationHooksWrapper{k}
}

// AfterDelegation is called after a delegation is made.
func (wrapper DelegationHooksWrapper) AfterDelegation(
	sdk.Context, sdk.AccAddress,
) {
	// we do nothing here, since the vote power for all operators is calculated
	// in the end separately. even if we knew the amount of the delegation, the
	// average exchange rate for the epoch is unknown.
}

// AfterUndelegationStarted is called after an undelegation is started.
func (wrapper DelegationHooksWrapper) AfterUndelegationStarted(
	ctx sdk.Context, operator sdk.AccAddress, recordKey []byte,
) error {
	var unbondingCompletionEpoch int64
	if wrapper.keeper.operatorKeeper.IsOperatorOptingOutFromChainID(
		ctx, operator, ctx.ChainID(),
	) {
		// if the operator is opting out, we need to use the finish epoch of the opt out.
		unbondingCompletionEpoch = wrapper.keeper.GetOperatorOptOutFinishEpoch(ctx, operator)
		// even if the operator opts back in, the undelegated vote power does not reappear
		// in the picture. slashable events between undelegation and opt in cannot occur
		// because the operator is not in the validator set.
	} else {
		// otherwise, we use the default unbonding completion epoch.
		unbondingCompletionEpoch = wrapper.keeper.GetUnbondingCompletionEpoch(ctx)
		// if the operator opts out after this, the undelegation will mature before the opt out.
		// so this is not a concern.
	}
	wrapper.keeper.AppendUndelegationToMature(ctx, unbondingCompletionEpoch, recordKey)
	return wrapper.keeper.delegationKeeper.IncrementUndelegationHoldCount(ctx, recordKey)
}
