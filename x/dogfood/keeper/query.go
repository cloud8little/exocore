package keeper

import (
	"github.com/ExocoreNetwork/exocore/x/dogfood/types"
)

var _ types.QueryServer = Keeper{}
