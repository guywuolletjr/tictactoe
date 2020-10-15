package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/google/uuid"
)

var _ sdk.Msg = &MsgCreateGame{}

type MsgCreateGame struct {
  ID      string
  Creator sdk.AccAddress `json:"creator" yaml:"creator"`
  Id string `json:"Id" yaml:"Id"`
  Player1 string `json:"Player1" yaml:"Player1"`
  Player2 string `json:"Player2" yaml:"Player2"`
  Board string `json:"Board" yaml:"Board"`
  Winner string `json:"Winner" yaml:"Winner"`
}

func NewMsgCreateGame(creator sdk.AccAddress, Id string, Player1 string, Player2 string, Board string, Winner string) MsgCreateGame {
  return MsgCreateGame{
    ID: uuid.New().String(),
		Creator: creator,
    Id: Id,
    Player1: Player1,
    Player2: Player2,
    Board: Board,
    Winner: Winner,
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