package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// RouterKey is they name of the lottery module
const RouterKey = ModuleName

// MsgSetSoruceChannel is a message for setting source channel to other chain
type MsgSetSourceChannel struct {
	ChainName     string         `json:"chain_name"`
	SourcePort    string         `json:"source_port"`
	SourceChannel string         `json:"source_channel"`
	Signer        sdk.AccAddress `json:"signer"`
}

func NewMsgSetSourceChannel(
	chainName, sourcePort, sourceChannel string,
	signer sdk.AccAddress,
) MsgSetSourceChannel {
	return MsgSetSourceChannel{
		ChainName:     chainName,
		SourcePort:    sourcePort,
		SourceChannel: sourceChannel,
		Signer:        signer,
	}
}

// Route implements the sdk.Msg interface for MsgSetSourceChannel.
func (msg MsgSetSourceChannel) Route() string { return RouterKey }

// Type implements the sdk.Msg interface for MsgSetSourceChannel.
func (msg MsgSetSourceChannel) Type() string { return "set_source_channel" }

// ValidateBasic implements the sdk.Msg interface for MsgSetSourceChannel.
func (msg MsgSetSourceChannel) ValidateBasic() error {
	// TODO: Add validate basic
	return nil
}

// GetSigners implements the sdk.Msg interface for MsgSetSourceChannel.
func (msg MsgSetSourceChannel) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Signer}
}

// GetSignBytes implements the sdk.Msg interface for MsgSetSourceChannel.
func (msg MsgSetSourceChannel) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// MsgCreateLottery is a message to create lottery game with open status
//	AccumulatedAmount sdk.Coins 		`json:"amount"`
type MsgCreateLottery struct {
	Creator  sdk.AccAddress `json:"creator"`
	Status   LotteryStatus `json:"status"`
	Amount   sdk.Coins    `json:"amount"`
}

// 
func NewMsgCreateLottery(
	creator sdk.AccAddress,
	status LotteryStatus, 
	amount sdk.Coins,
) MsgCreateLottery {
	return MsgCreateLottery{
		Creator:  creator,
		Status: status,
		Amount: amount,

	}
}
// Route implements the sdk.Msg interface for MsgCreateLottery.
func (msg MsgCreateLottery) Route() string { return RouterKey }

// Type implements the sdk.Msg interface for MsgCreateLottery.
func (msg MsgCreateLottery) Type() string { return "play_lottery" }

// ValidateBasic implements the sdk.Msg interface for MsgBuyGold.
func (msg MsgCreateLottery) ValidateBasic() error {
	if msg.Creator.Empty() {
		return sdkerrors.Wrapf(ErrInvalidBasicMsg, "MsgCreateLottery: Sender address must not be empty.")
	}

	//TODO: validate max_number is not empty or null
	//if msg.MaxNumber.Empty() {
	//	return sdkerrors.Wrapf(ErrInvalidBasicMsg, "MsgCreateLottery: MaxNumber must not be empty.")
	//}
	
	return nil
}

// GetSigners implements the sdk.Msg interface for MsgBuyGold.
func (msg MsgCreateLottery) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Creator}
}

// GetSignBytes implements the sdk.Msg interface for MsgBuyGold.
func (msg MsgCreateLottery) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// MsgPlayLottery is a message for play a number in the open lottery
type MsgPlayLottery struct {
	Player  sdk.AccAddress `json:"player"`
	Amount sdk.Coins       `json:"amount"`
	Number uint64          `json:"number"`
}

// NewMsgBuyGold creates a new MsgBuyGold instance.
func NewMsgPlayLottery(
	player sdk.AccAddress,
	amount sdk.Coins,
	number uint64,
) MsgPlayLottery {
	return MsgPlayLottery{
		Player:  player,
		Amount: amount,
		Number: number,

	}
}

// Route implements the sdk.Msg interface for MsgPlayLottery.
func (msg MsgPlayLottery) Route() string { return RouterKey }

// Type implements the sdk.Msg interface for MsgPlayLottery.
func (msg MsgPlayLottery) Type() string { return "play_lottery" }

// ValidateBasic implements the sdk.Msg interface for MsgBuyGold.
func (msg MsgPlayLottery) ValidateBasic() error {
	if msg.Player.Empty() {
		return sdkerrors.Wrapf(ErrInvalidBasicMsg, "MsgPlayLottery: Sender address must not be empty.")
	}
	if msg.Amount.Empty() {
		return sdkerrors.Wrapf(ErrInvalidBasicMsg, "MsgPlayLottery: Amount must not be empty.")
	}
	//if msg.Number == NullInt64 {
	//	return sdkerrors.Wrapf(ErrInvalidBasicMsg, "MsgPlayLottery: Number must not be empty.")
	//}
	return nil
}

// GetSigners implements the sdk.Msg interface for MsgBuyGold.
func (msg MsgPlayLottery) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Player}
}

// GetSignBytes implements the sdk.Msg interface for MsgBuyGold.
func (msg MsgPlayLottery) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// MsgCreateLottery is a message to create lottery game with open status
type MsgCLoseLottery struct {
	Closer  sdk.AccAddress `json:"creator"`
}

// 
func NewMsgCloseLottery(
	closer sdk.AccAddress,
) MsgCLoseLottery {
	return MsgCLoseLottery{
		Closer:  closer,

	}
}
// Route implements the sdk.Msg interface for MsgCLoseLottery.
func (msg MsgCLoseLottery) Route() string { return RouterKey }

// Type implements the sdk.Msg interface for MsgCLoseLottery.
func (msg MsgCLoseLottery) Type() string { return "close_lottery" }

// ValidateBasic implements the sdk.Msg interface for MsgCLoseLottery.
func (msg MsgCLoseLottery) ValidateBasic() error {
	if msg.Closer.Empty() {
		return sdkerrors.Wrapf(ErrInvalidBasicMsg, "MsgCLoseLottery: closer address must not be empty.")
	}
	//TODO: closer be the same user as creator
	
	return nil
}

// GetSigners implements the sdk.Msg interface for MsgBuyGold.
func (msg MsgCLoseLottery) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Closer}
}

// GetSignBytes implements the sdk.Msg interface for MsgBuyGold.
func (msg MsgCLoseLottery) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}