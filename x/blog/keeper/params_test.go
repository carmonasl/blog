package keeper_test

import (
	"testing"

	testkeeper "github.com/carmonasl/blog/testutil/keeper"
	"github.com/carmonasl/blog/x/blog/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.BlogKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
