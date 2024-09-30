package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/devdammit/kepler/testutil/keeper"
	"github.com/devdammit/kepler/x/kepler/keeper"
	kepler "github.com/devdammit/kepler/x/kepler/module"
	"github.com/devdammit/kepler/x/kepler/types"
	"github.com/stretchr/testify/require"
)

func setupWithSimpleW3Func(t *testing.T) (types.MsgServer, keeper.Keeper, sdk.Context) {
	k, ctx := keepertest.KeplerKeeper(t)

	kepler.InitGenesis(ctx, k, *types.DefaultGenesis())

	server := keeper.NewMsgServerImpl(k)

	w3Func := types.W3Func{
		Code:       "SendWebhook('http://localhost:8080', 'Hello, World!')",
		Cond:       "automation.Counter == 0 ? true : global.Time - automation.ExecutedAt > duration('1m')",
		ExecutedAt: "0001-01-01T00:00:00Z",
		Counter:    1,
	}
	k.AppendW3Func(ctx, w3Func)
	require.Equal(t, uint64(1), k.AppendW3Func(ctx, w3Func))

	return server, k, ctx
}

func TestExecuteW3Func(t *testing.T) {
	_, k, ctx := setupWithSimpleW3Func(t)

	err := k.ExecuteW3Func(ctx.WithBlockTime(ctx.BlockTime().AddDate(0, 0, 1)), 1)
	require.NoError(t, err)

	automation, err := k.GetW3Func(ctx, 1)
	require.NotNil(t, automation)
	require.Equal(t, uint64(2), automation.Counter)
}
