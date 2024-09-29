package keeper

import (
	"context"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/devdammit/kepler/x/kepler/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ListW3Func(goCtx context.Context, req *types.QueryListW3FuncRequest) (*types.QueryListW3FuncResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.W3FuncKey))

	var w3Funcs []types.W3Func

	_, err := query.Paginate(store, req.Pagination, func(key []byte, value []byte) error {
		var w3Func types.W3Func
		if err := k.cdc.Unmarshal(value, &w3Func); err != nil {
			return err
		}

		w3Funcs = append(w3Funcs, w3Func)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryListW3FuncResponse{}, nil
}
