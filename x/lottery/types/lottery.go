package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type LotteryStatus uint8

const (
	
	Open LotteryStatus = iota
	Closed
)

var StandardPrice = sdk.Coins{sdk.NewInt64Coin("stake", 1)}
var MaxNUmber = uint64(1000)

type Lottery struct {
	AccumulatedAmount sdk.Coins      `json:"accumulatedAmount"`
	Price   sdk.Coins       `json:"price"`
	Status LotteryStatus    `json:"status"`
	WinningNumber uint64     `json:"winningNumber"`
	MaxNUmber uint64 		`json:"maxNumber"`
	Id uint64  				`json:"id"`
}

func NewLottery(Open LotteryStatus, id uint64, amount sdk.Coins) Lottery {
	return Lottery{
		Status: Open,
		Price: StandardPrice,
		MaxNUmber: MaxNUmber,
		AccumulatedAmount: amount,
		Id: id,
	}
}


type Bet struct {
	Player  sdk.AccAddress `json:"player"`
	LotteryNumber uint64     `json:"lotteryNumber"`
	Lottery Lottery 		`json:"lotteryId"`
}

func NewBet(player sdk.AccAddress, number uint64, lottery Lottery) Bet {
	return Bet{
		Player: player,
		LotteryNumber: number,
		Lottery: lottery,
	}
}