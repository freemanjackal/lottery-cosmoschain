package goldcdp

import (
	"github.com/freemanjackal/lottery/x/lottery/keeper"
	"github.com/freemanjackal/lottery/x/lottery/types"
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
	MsgSetSourceChannel = types.MsgSetSourceChannel
)
