package seedchain_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "seedchain/testutil/keeper"
	"seedchain/testutil/nullify"
	"seedchain/x/seedchain"
	"seedchain/x/seedchain/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.SeedchainKeeper(t)
	seedchain.InitGenesis(ctx, *k, genesisState)
	got := seedchain.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
