package keeper

import (
	"context"

	"github.com/ExocoreNetwork/exocore/x/assets/types"

	sdkmath "cosmossdk.io/math"
	operatortypes "github.com/ExocoreNetwork/exocore/x/operator/types"
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Keeper struct {
	storeKey      storetypes.StoreKey
	cdc           codec.BinaryCodec
	historicalCtx types.CreateQueryContext
	// other keepers
	assetsKeeper     operatortypes.AssetsKeeper
	delegationKeeper operatortypes.DelegationKeeper
	oracleKeeper     operatortypes.OracleKeeper
	avsKeeper        operatortypes.AVSKeeper

	hooks       operatortypes.OperatorHooks // set separately via call to SetHooks
	slashKeeper operatortypes.SlashKeeper   // for jailing and unjailing check TODO(mm)
}

func NewKeeper(
	storeKey storetypes.StoreKey,
	cdc codec.BinaryCodec,
	historicalCtx types.CreateQueryContext,
	assetsKeeper operatortypes.AssetsKeeper,
	delegationKeeper operatortypes.DelegationKeeper,
	oracleKeeper operatortypes.OracleKeeper,
	avsKeeper operatortypes.AVSKeeper,
	slashKeeper operatortypes.SlashKeeper,
) Keeper {
	return Keeper{
		storeKey:         storeKey,
		cdc:              cdc,
		historicalCtx:    historicalCtx,
		assetsKeeper:     assetsKeeper,
		delegationKeeper: delegationKeeper,
		oracleKeeper:     oracleKeeper,
		avsKeeper:        avsKeeper,
		slashKeeper:      slashKeeper,
	}
}

func (k *Keeper) OracleInterface() operatortypes.OracleKeeper {
	return k.oracleKeeper
}

func (k Keeper) GetUnbondingExpirationBlockNumber(_ sdk.Context, _ sdk.AccAddress, startHeight uint64) uint64 {
	return startHeight + operatortypes.UnbondingExpiration
}

// OperatorKeeper interface will be implemented by deposit keeper
type OperatorKeeper interface {
	// RegisterOperator handle the registerOperator txs from msg service
	RegisterOperator(ctx context.Context, req *operatortypes.RegisterOperatorReq) (*operatortypes.RegisterOperatorResponse, error)

	IsOperator(ctx sdk.Context, addr sdk.AccAddress) bool

	GetUnbondingExpirationBlockNumber(ctx sdk.Context, OperatorAddress sdk.AccAddress, startHeight uint64) uint64

	OptIn(ctx sdk.Context, operatorAddress sdk.AccAddress, AVSAddr string) error

	OptOut(ctx sdk.Context, OperatorAddress sdk.AccAddress, AVSAddr string) error

	NoInstantaneousSlash(ctx sdk.Context, operatorAddress sdk.AccAddress, AVSAddr, slashContract, slashID string, occurredSateHeight int64, slashProportion sdkmath.LegacyDec) error
}
