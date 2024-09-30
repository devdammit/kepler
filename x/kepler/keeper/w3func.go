package keeper

import (
	"encoding/binary"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/devdammit/kepler/x/kepler/types"
)

func (k Keeper) ExecuteActiveW3Func(ctx sdk.Context) error {
	store := prefix.NewStore(runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx)), types.KeyPrefix(types.W3FuncKey))
	iterator := store.Iterator(nil, nil)

	defer iterator.Close()

	var executed []types.W3Func
	for ; iterator.Valid(); iterator.Next() {
		var w3Func types.W3Func
		k.cdc.MustUnmarshal(iterator.Value(), &w3Func)

		if k.ExecuteW3Func(ctx, w3Func.Id) == nil {
			executed = append(executed, w3Func)
		}
	}

	k.Logger().Info("executed active w3Funcs", "count", len(executed))

	return nil
}

func (k Keeper) MarkW3FuncExecuted(ctx sdk.Context, id uint64, executedAt string) {
	store := prefix.NewStore(runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx)), types.KeyPrefix(types.W3FuncKey))
	var w3Func types.W3Func
	k.cdc.MustUnmarshal(store.Get(GetW3FuncIDBytes(id)), &w3Func)

	w3Func.ExecutedAt = executedAt
	w3Func.Counter++

	appendedValue := k.cdc.MustMarshal(&w3Func)

	store.Set(GetW3FuncIDBytes(w3Func.Id), appendedValue)
}

func (k Keeper) AppendW3Func(ctx sdk.Context, w3Func types.W3Func) uint64 {
	count := k.GetW3FuncCount(ctx)

	k.Logger().Info("append w3Func", "id", count)
	w3Func.Id = count

	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.W3FuncKey))

	appendedValue := k.cdc.MustMarshal(&w3Func)

	store.Set(GetW3FuncIDBytes(w3Func.Id), appendedValue)

	k.SetW3FuncCount(ctx, count+1)

	return count
}

func (k Keeper) GetW3Func(ctx sdk.Context, id uint64) (*types.W3Func, error) {
	store := prefix.NewStore(runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx)), types.KeyPrefix(types.W3FuncKey))
	var w3Func types.W3Func
	k.cdc.MustUnmarshal(store.Get(GetW3FuncIDBytes(id)), &w3Func)
	return &w3Func, nil
}

// GetW3FuncCount get the total number of w3Func
func (k Keeper) GetW3FuncCount(ctx sdk.Context) uint64 {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.W3FuncCountKey)
	bz := store.Get(byteKey)
	if bz == nil {
		return 0
	}

	return binary.BigEndian.Uint64(bz)
}

// SetW3FuncCount set the total number of w3Func
func (k Keeper) SetW3FuncCount(ctx sdk.Context, count uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.W3FuncCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// GetW3FuncIDBytes returns the byte representation of the ID
func GetW3FuncIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}
