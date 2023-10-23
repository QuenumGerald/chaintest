package chaintest_test

import (
	"testing"

	keepertest "chaintest/testutil/keeper"
	"chaintest/testutil/nullify"
	"chaintest/x/chaintest"
	"chaintest/x/chaintest/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.ChaintestKeeper(t)
	chaintest.InitGenesis(ctx, *k, genesisState)
	got := chaintest.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
