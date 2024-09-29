package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/devdammit/kepler/x/kepler/types"
)

func (k msgServer) CreateW3Func(goCtx context.Context, msg *types.MsgCreateW3Func) (*types.MsgCreateW3FuncResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var w3Func = types.W3Func{
		Creator: msg.Creator,
		Cond:    msg.Cond,
		Code:    msg.Code,
	}

	id := k.AppendW3Func(ctx, w3Func)

	return &types.MsgCreateW3FuncResponse{
		Id: id,
	}, nil
}
