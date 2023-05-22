package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "seedchain/testutil/keeper"
	"seedchain/x/seedchain/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.SeedchainKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
