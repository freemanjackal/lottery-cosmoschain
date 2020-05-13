package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// RouterKey is they name of the goldcdp module
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

// MsgBuyGold is a message for creating order to buy gold
type MsgPlayLottery struct {
	Player  sdk.AccAddress `json:"player"`
	Amount sdk.Coins      `json:"amount"`
	Number uint8          `json:"number"`
}

// NewMsgBuyGold creates a new MsgBuyGold instance.
func NewMsgPlayLottery(
	player sdk.AccAddress,
	amount sdk.Coins,
	number uint8,
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
	if msg.Number.Empty() {
		return sdkerrors.Wrapf(ErrInvalidBasicMsg, "MsgPlayLottery: Number must not be empty.")
	}
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
