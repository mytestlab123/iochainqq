package keeper_test

import (
	"testing"

	testkeeper "github.com/mytestlab123/iochainqq/testutil/keeper"
	"github.com/mytestlab123/iochainqq/x/iochainqq/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.IochainqqKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
