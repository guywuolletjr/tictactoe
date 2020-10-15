package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSetGame{}

type MsgSetGame struct {
  ID      string      		`json:"id" yaml:"id"`
  Creator sdk.AccAddress 	`json:"creator" yaml:"creator"`
  Player1 sdk.AccAddress 	`json:"Player1" yaml:"Player1"`
  Player2 sdk.AccAddress 	`json:"Player2" yaml:"Player2"`
}

func NewMsgSetGame(creator sdk.AccAddress, id string, Player1 sdk.AccAddress, Player2 sdk.AccAddress) MsgSetGame {
  return MsgSetGame{
    ID: id,
		Creator: creator,
    Player1: Player1,
    Player2: Player2,
	}
}

func (msg MsgSetGame) Route() string {
  return RouterKey
}

func (msg MsgSetGame) Type() string {
  return "SetGame"
}

func (msg MsgSetGame) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgSetGame) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg MsgSetGame) ValidateBasic() error {
  if msg.Creator.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }
  return nil
}
