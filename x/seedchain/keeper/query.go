package keeper

import (
	"seedchain/x/seedchain/types"
)

var _ types.QueryServer = Keeper{}
