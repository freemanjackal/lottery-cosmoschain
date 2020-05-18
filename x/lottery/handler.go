package lottery

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
//	"strconv"
	"strings"

	"github.com/bandprotocol/bandchain/chain/x/oracle"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	channel "github.com/cosmos/cosmos-sdk/x/ibc/04-channel"
	channeltypes "github.com/cosmos/cosmos-sdk/x/ibc/04-channel/types"

	"github.com/freemanjackal/lottery-cosmoschain/x/lottery/types"
)

// NewHandler creates the msg handler of this module, as required by Cosmos-SDK standard.
func NewHandler(keeper Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())
		switch msg := msg.(type) {
		case MsgCreateLottery:
			return handleCreateLottery(ctx, msg, keeper)
		case MsgPlayLottery:
			return handleMakeBet(ctx, msg, keeper)
		case MsgCLoseLottery:
			return handleCloseLottery(ctx, msg, keeper)
		case MsgSetSourceChannel:
			return handleSetSourceChannel(ctx, msg, keeper)
		case channeltypes.MsgPacket:
			var responsePacket oracle.OracleResponsePacketData
			if err := types.ModuleCdc.UnmarshalJSON(msg.GetData(), &responsePacket); err == nil {
				return handleOracleRespondPacketData(ctx, responsePacket, keeper)
			}
			return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "cannot unmarshal oracle packet data")
		default:
			return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized %s message type: %T", ModuleName, msg)
		}
	}
}


//create lottery at genesis
func handleCreateLottery(ctx sdk.Context, msg MsgCreateLottery, keeper Keeper) (*sdk.Result, error) {
	_, err := keeper.CreateLottery(ctx, msg.Status, msg.Amount)
	if err != nil {
		return nil, err
	}

	
	return  &sdk.Result{Events: ctx.EventManager().Events().ToABCIEvents()}, nil
}

func handleMakeBet(ctx sdk.Context, msg MsgPlayLottery, keeper Keeper) (*sdk.Result, error) {
	_, err := keeper.AddBet(ctx, msg.Player, msg.Amount, msg.Number)
	if err != nil {
		return nil, err
	}
	
	return &sdk.Result{Events: ctx.EventManager().Events().ToABCIEvents()}, nil
}

//close open lottery, distribute funds to winners, create a new Lottery
func handleCloseLottery(ctx sdk.Context, msg MsgCLoseLottery, keeper Keeper) (*sdk.Result, error) {
	
	// TODO: Set all bandchain parameter here
	bandChainID := "ibc-bandchain"
	port := "lottery"
	//script id of random  number generator
	oracleScriptID := oracle.OracleScriptID(6)
	calldata := make([]byte, 8)
	binary.LittleEndian.PutUint64(calldata, 1000000)
	askCount := int64(1)
	minCount := int64(1)
	//max_range := 10000

	channelID, err := keeper.GetChannel(ctx, bandChainID, port)

	if err != nil {
		return nil, sdkerrors.Wrapf(
			sdkerrors.ErrUnknownRequest,
			"not found channel to lottery",
		)
	}
	sourceChannelEnd, found := keeper.ChannelKeeper.GetChannel(ctx, port, channelID)
	if !found {
		return nil, sdkerrors.Wrapf(
			sdkerrors.ErrUnknownRequest,
			"unknown channel %s port lottery",
			channelID,
		)
	}
	destinationPort := sourceChannelEnd.Counterparty.PortID
	destinationChannel := sourceChannelEnd.Counterparty.ChannelID
	sequence, found := keeper.ChannelKeeper.GetNextSequenceSend(
		ctx, port, channelID,
	)
	if !found {
		return nil, sdkerrors.Wrapf(
			sdkerrors.ErrUnknownRequest,
			"unknown sequence number for channel %s port oracle",
			channelID,
		)
	}
	packet := oracle.NewOracleRequestPacketData(
		fmt.Sprintf("Getting winning number"), oracleScriptID, hex.EncodeToString(calldata),
		askCount, minCount,
	)
	err = keeper.ChannelKeeper.SendPacket(ctx, channel.NewPacket(packet.GetBytes(),
		sequence, port, channelID, destinationPort, destinationChannel,
		1000000000, // Arbitrarily high timeout for now
	))
	if err != nil {
		return nil, err
	}

	return &sdk.Result{Events: ctx.EventManager().Events().ToABCIEvents()}, nil
}

func handleSetSourceChannel(ctx sdk.Context, msg MsgSetSourceChannel, keeper Keeper) (*sdk.Result, error) {
	keeper.SetChannel(ctx, msg.ChainName, msg.SourcePort, msg.SourceChannel)
	return &sdk.Result{Events: ctx.EventManager().Events().ToABCIEvents()}, nil
}

func handleOracleRespondPacketData(ctx sdk.Context, packet oracle.OracleResponsePacketData, keeper Keeper) (*sdk.Result, error) {
	clientID := strings.Split(packet.ClientID, ":")
	if len(clientID) != 2 {
		return nil, sdkerrors.Wrapf(types.ErrUnknownClientID, "unknown client id %s", packet.ClientID)
	}
	//lotteryID, err := strconv.ParseUint(clientID[1], 10, 64)
	
	rawResult, err := hex.DecodeString(packet.Result)
	if err != nil {
		return nil, err
	}
	result, err := types.DecodeResult(rawResult)
	if err != nil {
		return nil, err
	}

	lotteryID  := keeper.GetLotteryCount(ctx)
	
	lottery,err  := keeper.GetLottery(ctx, lotteryID)
	if err != nil {
		return nil, err
	}

	amount:= lottery.AccumulatedAmount
	// TODO: Calculate collateral percentage
	winningNUmber := uint64(result.Px)
	//winners := make([]sdk.AccAddress, 1)
	winners,err := keeper.GetWinners(ctx, winningNUmber)

	//if no winners close lottery and pass funds t new lottery
	//var newAmount sdk.Coins
	if len(winners) == 0 {

		keeper.CloseLottery(ctx)
		keeper.CreateLottery(ctx, types.Open, amount)
	} else {
		//distribute funds to winners

		keeper.CloseLottery(ctx)
		keeper.CreateLottery(ctx, types.Open, sdk.Coins{sdk.NewCoin("stake", sdk.NewInt(0))})
	}

	
	return &sdk.Result{Events: ctx.EventManager().Events().ToABCIEvents()}, nil
}
