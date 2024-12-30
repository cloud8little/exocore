package keeper

import (
	"context"

	delegationtype "github.com/ExocoreNetwork/exocore/x/delegation/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ delegationtype.QueryServer = &Keeper{}

func (k *Keeper) QuerySingleDelegationInfo(ctx context.Context, req *delegationtype.SingleDelegationInfoReq) (*delegationtype.DelegationAmounts, error) {
	c := sdk.UnwrapSDKContext(ctx)
	return k.GetSingleDelegationInfo(c, req.StakerId, req.AssetId, req.OperatorAddr)
}

func (k *Keeper) QueryDelegationInfo(ctx context.Context, info *delegationtype.DelegationInfoReq) (*delegationtype.QueryDelegationInfoResponse, error) {
	c := sdk.UnwrapSDKContext(ctx)
	return k.GetDelegationInfo(c, info.StakerId, info.AssetId)
}

func (k *Keeper) QueryUndelegations(ctx context.Context, req *delegationtype.UndelegationsReq) (*delegationtype.UndelegationRecordList, error) {
	c := sdk.UnwrapSDKContext(ctx)
	undelegations, err := k.GetStakerUndelegationRecords(c, req.StakerId, req.AssetId)
	if err != nil {
		return nil, err
	}
	return &delegationtype.UndelegationRecordList{
		Undelegations: undelegations,
	}, nil
}

func (k *Keeper) QueryUndelegationsByHeight(ctx context.Context, req *delegationtype.UndelegationsByHeightReq) (*delegationtype.UndelegationRecordList, error) {
	c := sdk.UnwrapSDKContext(ctx)
	undelegations, err := k.GetCompletablePendingUndelegations(c, req.BlockHeight)
	if err != nil {
		return nil, err
	}
	return &delegationtype.UndelegationRecordList{
		Undelegations: undelegations,
	}, nil
}

func (k Keeper) QueryUndelegationHoldCount(ctx context.Context, req *delegationtype.UndelegationHoldCountReq) (*delegationtype.UndelegationHoldCountResponse, error) {
	c := sdk.UnwrapSDKContext(ctx)
	res := k.GetUndelegationHoldCount(c, []byte(req.RecordKey))
	return &delegationtype.UndelegationHoldCountResponse{HoldCount: res}, nil
}

func (k Keeper) QueryAssociatedOperatorByStaker(ctx context.Context, req *delegationtype.QueryAssociatedOperatorByStakerReq) (*delegationtype.QueryAssociatedOperatorByStakerResponse, error) {
	c := sdk.UnwrapSDKContext(ctx)
	operator, err := k.GetAssociatedOperator(c, req.StakerId)
	if err != nil {
		return nil, err
	}
	return &delegationtype.QueryAssociatedOperatorByStakerResponse{
		Operator: operator,
	}, nil
}

func (k Keeper) QueryAssociatedStakersByOperator(ctx context.Context, req *delegationtype.QueryAssociatedStakersByOperatorReq) (*delegationtype.QueryAssociatedStakersByOperatorResponse, error) {
	c := sdk.UnwrapSDKContext(ctx)
	stakers, err := k.GetAssociatedStakers(c, req.Operator)
	if err != nil {
		return nil, err
	}
	return &delegationtype.QueryAssociatedStakersByOperatorResponse{
		Stakers: stakers,
	}, nil
}

func (k Keeper) QueryDelegatedStakersByOperator(ctx context.Context, req *delegationtype.QueryDelegatedStakersByOperatorReq) (*delegationtype.QueryDelegatedStakersByOperatorResponse, error) {
	c := sdk.UnwrapSDKContext(ctx)
	stakers, err := k.GetStakersByOperator(c, req.Operator, req.AssetId)
	if err != nil {
		return nil, err
	}
	return &delegationtype.QueryDelegatedStakersByOperatorResponse{
		Count:   uint64(len(stakers.Stakers)),
		Stakers: stakers.Stakers,
	}, nil
}
