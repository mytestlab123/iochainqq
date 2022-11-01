package iochainqq_test

import (
	"testing"

	keepertest "github.com/mytestlab123/iochainqq/testutil/keeper"
	"github.com/mytestlab123/iochainqq/testutil/nullify"
	"github.com/mytestlab123/iochainqq/x/iochainqq"
	"github.com/mytestlab123/iochainqq/x/iochainqq/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.IochainqqKeeper(t)
	iochainqq.InitGenesis(ctx, *k, genesisState)
	got := iochainqq.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
