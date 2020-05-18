package lottery

import (
	"github.com/freemanjackal/lottery-cosmoschain/x/lottery/keeper"
	"github.com/freemanjackal/lottery-cosmoschain/x/lottery/types"
)

const (
	ModuleName = types.ModuleName
	RouterKey  = types.RouterKey
	StoreKey   = types.StoreKey
)

var (
	NewKeeper     = keeper.NewKeeper
	RegisterCodec = types.RegisterCodec
	NewQuerier    = keeper.NewQuerier
)

type (
	Keeper              = keeper.Keeper
	MsgPlayLottery          = types.MsgPlayLottery
	MsgCreateLottery    = types.MsgCreateLottery
	MsgCLoseLottery     = types.MsgCLoseLottery
	MsgSetSourceChannel = types.MsgSetSourceChannel
)
