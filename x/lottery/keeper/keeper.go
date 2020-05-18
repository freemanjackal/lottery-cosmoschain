package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/freemanjackal/lottery-cosmoschain/x/lottery/types"
)

type Keeper struct {
	storeKey      sdk.StoreKey
	cdc           *codec.Codec
	BankKeeper    types.BankKeeper
	ChannelKeeper types.ChannelKeeper
}

// NewKeeper creates a new band consumer Keeper instance.
func NewKeeper(cdc *codec.Codec, key sdk.StoreKey, bankKeeper types.BankKeeper,
	channelKeeper types.ChannelKeeper,
) Keeper {
	return Keeper{
		storeKey:      key,
		cdc:           cdc,
		BankKeeper:    bankKeeper,
		ChannelKeeper: channelKeeper,
	}
}

// GetOrderCount returns the current number of all orders ever exist.
func (k Keeper) GetBetCount(ctx sdk.Context) uint64 {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.BetsCountStoreKey)
	if bz == nil {
		return 0
	}
	return binary.BigEndian.Uint64(bz)
}
// GetOrderCount returns the current number of all orders ever exist.
func (k Keeper) GetLotteryCount(ctx sdk.Context) uint64 {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.LotteryCountStoreKey)
	if bz == nil {
		return 0
	}
	return binary.BigEndian.Uint64(bz)
}

// GetNextBetCount increments and returns the current number of bets.
// If the global bet  count is not set, it initializes it with value 0.
func (k Keeper) GetNextBetCount(ctx sdk.Context) uint64 {
	betCount := k.GetBetCount(ctx)
	store := ctx.KVStore(k.storeKey)
	bz := sdk.Uint64ToBigEndian(betCount + 1)
	store.Set(types.BetsCountStoreKey, bz)
	return betCount + 1
}


// GetNextLotteryCount increments and returns the current number of lotteries.
// If the global lottery count is not set, it initializes it with value 0.
func (k Keeper) GetNextLotteryCount(ctx sdk.Context) uint64 {
	betCount := k.GetBetCount(ctx)
	store := ctx.KVStore(k.storeKey)
	bz := sdk.Uint64ToBigEndian(betCount + 1)
	store.Set(types.LotteryCountStoreKey, bz)
	return betCount + 1
}