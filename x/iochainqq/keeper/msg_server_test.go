package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/mytestlab123/iochainqq/testutil/keeper"
	"github.com/mytestlab123/iochainqq/x/iochainqq/keeper"
	"github.com/mytestlab123/iochainqq/x/iochainqq/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.IochainqqKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
