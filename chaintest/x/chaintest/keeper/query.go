package keeper

import (
	"chaintest/x/chaintest/types"
)

var _ types.QueryServer = Keeper{}
