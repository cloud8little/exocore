package withdraw

import (
	errorsmod "cosmossdk.io/errors"
	exocmn "github.com/ExocoreNetwork/exocore/precompiles/common"
	"github.com/ExocoreNetwork/exocore/x/assets/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
)

const (
	// MethodWithdraw defines the ABI method name for the withdrawal transaction.
	MethodWithdraw = "withdrawPrinciple"
)

// Withdraw assets to the staker, that will change the state in withdraw module.
func (p Precompile) Withdraw(
	ctx sdk.Context,
	_ common.Address,
	contract *vm.Contract,
	_ vm.StateDB,
	method *abi.Method,
	args []interface{},
) ([]byte, error) {
	// check the invalidation of caller contract
	err := p.assetsKeeper.CheckExocoreLzAppAddr(ctx, contract.CallerAddress)
	if err != nil {
		return nil, errorsmod.Wrap(err, exocmn.ErrContractCaller)
	}

	withdrawParam, err := p.GetWithdrawParamsFromInputs(ctx, args)
	if err != nil {
		return nil, err
	}

	err = p.withdrawKeeper.Withdraw(ctx, withdrawParam)
	if err != nil {
		return nil, err
	}
	// get the latest asset state of staker to return.
	stakerID, assetID := types.GetStakeIDAndAssetID(withdrawParam.ClientChainLzID, withdrawParam.WithdrawAddress, withdrawParam.AssetsAddress)
	info, err := p.assetsKeeper.GetStakerSpecifiedAssetInfo(ctx, stakerID, assetID)
	if err != nil {
		return nil, err
	}
	return method.Outputs.Pack(true, info.TotalDepositAmount.BigInt())
}
