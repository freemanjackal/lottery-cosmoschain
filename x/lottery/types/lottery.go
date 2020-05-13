package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type LotteryStatus uint8

const (
	
	Open LotteryStatus
	Closed
)

var StandardPrice = sdk.Coins{sdk.NewInt64Coin("lotteryToken", 1)}
var MaxNUmber = uint8(1000)

type Lottery struct {
	AccumulatedAmount sdk.Coins      `json:"accumulatedAmount"`
	Price   sdk.Coins       `json:"price"`
	Status LotteryStatus    `json:"status"`
	WinningNumber uint8     `json:"winningNumber"`
	MaxNUmber uint8 		`json:"maxNumber"`
}

func NewLottery(amount sdk.Coins) Lottery {
	return Lottery{
		Status: Open,
		Price: StandardPrice,
		MaxNUmber: MaxNUmber,
		AccumulatedAmount: amount
	}
}


type Bet struct {
	Player  sdk.AccAddress `json:"player"`
	LotteryNumber uint8     `json:"lotteryNumber"`
	Lottery Lottery 		`json:"lotteryId"`
}

func NewBet(player sdk.AccAddress, number uint8, lottery Lottery) Bet {
	return Bet{
		Player: player,
		LotteryNumber: number,
		Lottery: lottery
	}
}