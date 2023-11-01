package keeper_test

import (
	"testing"

	testkeeper "chaintest/testutil/keeper"
	"chaintest/x/chaintest/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.ChaintestKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
