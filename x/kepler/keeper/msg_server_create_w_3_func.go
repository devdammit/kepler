package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/devdammit/kepler/x/kepler/types"
)

func (k msgServer) CreateW3Func(goCtx context.Context, msg *types.MsgCreateW3Func) (*types.MsgCreateW3FuncResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgCreateW3FuncResponse{}, nil
}
