package keeper

import (
	//"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	//transfer "github.com/cosmos/cosmos-sdk/x/ibc/20-transfer"

	"github.com/freemanjackal/lottery-cosmoschain/x/lottery/types"
)

func (k Keeper) AddBet(ctx sdk.Context, player sdk.AccAddress, amount sdk.Coins,  number uint64) (uint64, error) {
	betID := k.GetNextBetCount(ctx)

	lotteryID := k.GetLotteryCount(ctx)

	lottery,err := k.GetLottery(ctx, lotteryID)
	if err != nil {
		return 0, err
	}
	// TODO: Config chain name
	//collateralChain := "band-cosmoshub"

	// TODO: Support only 1 coin
	if len(amount) != 1 {
		return 0, sdkerrors.Wrapf(types.ErrOnlyOneDenomAllowed, "%d denoms included", len(amount))
	}

	
	// escrow source tokens. It fails if balance insufficient.
	//if amount[1] < types.StandardPrice{
	//	return 0, sdkerrors.Wrapf(types.ErrorInvalidAmount, "%d invalid amount", len(amount[1]))
	
	k.SetBet(ctx, betID, types.NewBet(player, number, lottery), lotteryID)

	return betID, nil
}


func (k Keeper) CreateLottery(ctx sdk.Context, status types.LotteryStatus, amount sdk.Coins) (uint64, error) {
	lotteryID := k.GetNextLotteryCount(ctx)
	
	
	k.SetLottery(ctx, lotteryID, types.NewLottery(status,lotteryID, amount))

	return lotteryID, nil
}


// SetBet saves the given Bet to the open lottery to the store without performing any validation.
func (k Keeper) SetBet(ctx sdk.Context, id uint64, bet types.Bet, lotteryID uint64) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.BetStoreKey(id, lotteryID), k.cdc.MustMarshalBinaryBare(bet))
}


func (k Keeper) SetLottery(ctx sdk.Context, id uint64, lottery types.Lottery) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.LotteryStoreKey(id), k.cdc.MustMarshalBinaryBare(lottery))
}

// GetLottery gets the given Lottery from the store
func (k Keeper) GetLottery(ctx sdk.Context, id uint64) (types.Lottery, error) {
	store := ctx.KVStore(k.storeKey)
	if !store.Has(types.LotteryStoreKey(id)) {
		return types.Lottery{}, sdkerrors.Wrapf(sdkerrors.ErrKeyNotFound, "lottery %d not found", id)
	}
	bz := store.Get(types.LotteryStoreKey(id))
	var lottery types.Lottery
	k.cdc.MustUnmarshalBinaryBare(bz, &lottery)
	return lottery, nil
}

// Get an iterator over all bets
func (k Keeper) GetBetsIterator(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return sdk.KVStorePrefixIterator(store, []byte{})
}


func (k Keeper) GetWinners(ctx sdk.Context, winningNumber uint64)([]sdk.AccAddress, error) {
	result := make([]sdk.AccAddress, 1)
	//bets := make([]types.Bet, 1)

	//result = append(s, "d")
	loteryId:= k.GetLotteryCount(ctx)

	//store := ctx.KVStore(k.storeKey)
	//if !store.Has(types.BetStoreKey(loterryId,loterryId)) {
	//	return nil, sdkerrors.Wrapf(sdkerrors.ErrKeyNotFound, "winners %d not found, check errors", loterryId)
	//}

	//TODo: ccheck an efficient way to get the bets of open lottery
	//bz := store.Get(types.BetStoreKey(loterryId,loterryId))


	iterator := k.GetBetsIterator(ctx)

	for ; iterator.Valid(); iterator.Next() {
		var bet types.Bet
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &bet)
		//bet := types.Bet(iterator.Value())

		if bet.LotteryNumber == winningNumber && loteryId == bet.Lottery.Id{
            result = append(result, bet.Player)
        }

		//bets = append(bets, types.Bet(iterator.Value()))
	}
	//k.GetBetsIterator(ctx)
	//var bets []types.Bet
	//k.cdc.MustUnmarshalBinaryBare(bz, &bets)

	 //for _, bet := range bets {
      //  if bet.LotteryNumber == winningNumber {
       //     result = append(result, bet.Player)
        //}
    //}

	return result, nil

}

//close lottery
func (k Keeper) CloseLottery(ctx sdk.Context)() {
	lotteryId := k.GetLotteryCount(ctx)
	lottery,err := k.GetLottery(ctx, lotteryId)
	if err != nil {
		panic( "error closing lottery")
	}
	lottery.Status = types.Closed

	k.SetLottery(ctx, lotteryId, lottery)
}
