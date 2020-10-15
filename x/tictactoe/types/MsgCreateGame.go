package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/google/uuid"
)

var _ sdk.Msg = &MsgCreateGame{}

type MsgCreateGame struct {
	ID      string      		`json:"id" yaml:"id"`
  Creator sdk.AccAddress 	`json:"creator" yaml:"creator"`
  Player1 sdk.AccAddress 	`json:"Player1" yaml:"Player1"`
  Player2 sdk.AccAddress 	`json:"Player2" yaml:"Player2"`
}

func NewMsgCreateGame(creator sdk.AccAddress, id string, Player1 sdk.AccAddress, Player2 sdk.AccAddress) MsgCreateGame {
  return MsgCreateGame{
    ID: id,
		Creator: creator,
    Player1: Player1,
    Player2: Player2,
	}
}

func (msg MsgCreateGame) Route() string {
  return RouterKey
}

func (msg MsgCreateGame) Type() string {
  return "CreateGame"
}

func (msg MsgCreateGame) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgCreateGame) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg MsgCreateGame) ValidateBasic() error {
  if msg.Creator.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }
  return nil
}
