package keeper

import (
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	transfer "github.com/cosmos/cosmos-sdk/x/ibc/20-transfer"

	"github.com/freemanjackal/lottery/x/lottery/types"
)

func (k Keeper) AddBet(ctx sdk.Context, player sdk.AccAddress, amount sdk.Coins, uint8 number) (uint64, error) {
	betID := k.GetNextBetCount(ctx)

	lotteryID := k.GetLotteryCount(ctx)

	lottery = k.GetLottery(ctx, lotteryID)
	// TODO: Config chain name
	collateralChain := "band-cosmoshub"

	// TODO: Support only 1 coin
	if len(amount) != 1 {
		return 0, sdkerrors.Wrapf(types.ErrOnlyOneDenomAllowed, "%d denoms included", len(amount))
	}
	
	// escrow source tokens. It fails if balance insufficient.
	if amount[1] < types.StandardPrice{
		return 0, sdkerrors.Wrapf(types.ErrorInvalidAmount, "%d invalid amount", len(amount[1]))
	}
	escrowAddress := types.GetEscrowAddress()
	err = k.BankKeeper.SendCoins(ctx, player, escrowAddress, amount)
	if err != nil {
		return 0, err
	}
	k.SetBet(ctx, betID, types.NewBet(player, number, lottery))

	return betID, nil
}


func (k Keeper) CreateLottery(ctx sdk.Context, amount sdk.Coins) (uint64, error) {
	lotteryID := k.GetNextLotteryCount(ctx)
	
	
	k.SetLottery(ctx, lotteryID, types.NewLottery(amount))

	return lotteryID, nil
}


// SetBet saves the given Bet to the open lottery to the store without performing any validation.
func (k Keeper) SetBet(ctx sdk.Context, id uint64, bet types.Bet) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.BetStoreKey(id), k.cdc.MustMarshalBinaryBare(bet))
}


func (k Keeper) SetLottery(ctx sdk.Context, id uint64, lottery types.Lottery) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.LotteryStoreKey(id), k.cdc.MustMarshalBinaryBare(lottery))
}

// GetLottery gets the given Lottery from the store
func (k Keeper) GetLottery(ctx sdk.Context, id uint64) (types.Lottery, error) {
	store := ctx.KVStore(k.storeKey)
	if !store.Has(types.OrderStoreKey(id)) {
		return types.Lottery{}, sdkerrors.Wrapf(sdkerrors.ErrKeyNotFound, "lottery %d not found", id)
	}
	bz := store.Get(types.LotteryStoreKey(id))
	var lottery types.Lottery
	k.cdc.MustUnmarshalBinaryBare(bz, &order)
	return lottery, nil
}

func (k Keeper) CloseLottery(ctx sdk.Context)(id uint64) {
	
}
