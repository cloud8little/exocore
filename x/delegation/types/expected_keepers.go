package types

import (
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	assetstype "github.com/imua-xyz/imuachain/x/assets/types"
	epochsTypes "github.com/imua-xyz/imuachain/x/epochs/types"
)

var _ SlashKeeper = VirtualSlashKeeper{}

type SlashKeeper interface {
	IsOperatorFrozen(ctx sdk.Context, opAddr sdk.AccAddress) bool
	OperatorAssetSlashedProportion(ctx sdk.Context, opAddr sdk.AccAddress, assetID string, startHeight, endHeight uint64) sdkmath.LegacyDec
}

// VirtualSlashKeeper todo: When the actual keeper functionality has not been implemented yet, temporarily use the virtual keeper.
type VirtualSlashKeeper struct{}

func (VirtualSlashKeeper) IsOperatorFrozen(_ sdk.Context, _ sdk.AccAddress) bool {
	return false
}

func (VirtualSlashKeeper) OperatorAssetSlashedProportion(_ sdk.Context, _ sdk.AccAddress, _ string, _, _ uint64) sdkmath.LegacyDec {
	return sdkmath.LegacyZeroDec()
}

// DelegationHooks are event hooks triggered by the delegation module
type DelegationHooks interface {
	// AfterDelegation we don't want the ability to cancel delegation or undelegation so no return type for
	// either
	// for delegation, we only care about the address of the operator to cache the event
	AfterDelegation(ctx sdk.Context, operator sdk.AccAddress)
	// AfterUndelegationStarted for undelegation, we use the address of the operator to figure out the list of impacted
	// chains for that operator. and we need the identifier to hold it until confirmed by subscriber
	AfterUndelegationStarted(ctx sdk.Context, addr sdk.AccAddress, recordKey []byte) error
}

type OperatorKeeper interface {
	IsOperator(ctx sdk.Context, addr sdk.AccAddress) bool
	GetUnbondingExpiration(ctx sdk.Context, operator sdk.AccAddress) (string, int64, error)
}

type AssetsKeeper interface {
	UpdateStakerAssetState(
		ctx sdk.Context, stakerID string, assetID string, changeAmount assetstype.DeltaStakerSingleAsset,
	) (info *assetstype.StakerAssetInfo, err error)
	UpdateOperatorAssetState(ctx sdk.Context, operatorAddr sdk.Address, assetID string, changeAmount assetstype.DeltaOperatorSingleAsset) (err error)
	GetStakerSpecifiedAssetInfo(ctx sdk.Context, stakerID string, assetID string) (info *assetstype.StakerAssetInfo, err error)
	GetOperatorSpecifiedAssetInfo(ctx sdk.Context, operatorAddr sdk.Address, assetID string) (info *assetstype.OperatorAssetInfo, err error)
	IsOperatorAssetExist(ctx sdk.Context, operatorAddr sdk.Address, assetID string) bool

	ClientChainExists(ctx sdk.Context, index uint64) bool
}

type BankKeeper interface {
	GetBalance(ctx sdk.Context, addr sdk.AccAddress, denom string) sdk.Coin
	DelegateCoinsFromAccountToModule(ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error
	UndelegateCoinsFromModuleToAccount(ctx sdk.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error
}

type AccountKeeper interface {
	GetSequence(ctx sdk.Context, addr sdk.AccAddress) (uint64, error)
}

// EpochsKeeper represents the expected keeper interface for the epochs module.
type EpochsKeeper interface {
	GetEpochInfo(ctx sdk.Context, epochIdentifier string) (epochsTypes.EpochInfo, bool)
	AllEpochInfos(ctx sdk.Context) []epochsTypes.EpochInfo
}
