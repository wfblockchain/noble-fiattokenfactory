package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	errors "github.com/pkg/errors"
)

const TypeMsgBurn = "burn"

var _ sdk.Msg = &MsgBurn{}

func NewMsgBurn(from string, amount sdk.Coin) *MsgBurn {
	return &MsgBurn{
		From:   from,
		Amount: amount,
	}
}

func (msg *MsgBurn) Route() string {
	return RouterKey
}

func (msg *MsgBurn) Type() string {
	return TypeMsgBurn
}

func (msg *MsgBurn) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{from}
}

func (msg *MsgBurn) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgBurn) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.From)
	if err != nil {
		return errors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid from address (%s)", err)
	}

	if msg.Amount.IsNil() {
		return errors.Wrap(sdkerrors.ErrInvalidCoins, "burn amount cannot be nil")
	}

	if msg.Amount.IsNegative() {
		return errors.Wrap(sdkerrors.ErrInvalidCoins, "burn amount cannot be negative")
	}

	if msg.Amount.IsZero() {
		return errors.Wrap(sdkerrors.ErrInvalidCoins, "burn amount cannot be zero")
	}

	return nil
}
