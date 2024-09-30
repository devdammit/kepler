package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/devdammit/kepler/x/kepler/types"
)

func (k msgServer) W3FuncExecuted(goCtx context.Context, msg *types.MsgW3FuncExecuted) (*types.MsgW3FuncExecutedResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	k.Keeper.MarkW3FuncExecuted(ctx, msg.Id, msg.ExecutedAt)

	return &types.MsgW3FuncExecutedResponse{}, nil
}
