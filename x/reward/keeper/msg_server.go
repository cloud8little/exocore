package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/imua-xyz/imuachain/utils"
	"github.com/imua-xyz/imuachain/x/reward/types"
)

type msgServer struct {
	Keeper
}

func (k Keeper) UpdateParams(ctx context.Context, params *types.MsgUpdateParams) (*types.MsgUpdateParamsResponse, error) {
	c := sdk.UnwrapSDKContext(ctx)
	if utils.IsMainnet(c.ChainID()) && k.authority != params.Authority {
		return nil, govtypes.ErrInvalidSigner.Wrapf(
			"invalid authority; expected %s, got %s",
			k.authority, params.Authority,
		)
	}

	k.Logger(c).Info(
		"UpdateParams request",
		"authority", k.authority,
		"params.AUthority", params.Authority,
	)

	err := k.SetParams(c, &params.Params)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

var _ types.MsgServer = msgServer{}
