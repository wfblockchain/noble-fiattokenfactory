package keeper_test

import (
	"testing"

	testkeeper "github.com/wfblockchain/noble-fiattokenfactory/testutil/keeper"
	"github.com/wfblockchain/noble-fiattokenfactory/x/fiattokenfactory/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.FiatTokenfactoryKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
