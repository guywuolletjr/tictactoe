package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSetGame{}

type MsgPlayGame struct {
	ID      	string      		`json:"id" yaml:"id"`
	Player 		sdk.AccAddress 	`json:"Player" yaml:"Player"`
	Move			uint 						`json:"Move" yaml:"Move"`
}

func NewMsgPlayGame(id string, Player sdk.AccAddress, Move uint) MsgPlayGame {
  return MsgSetGame{
    ID: id,
    Player: Player,
    Move: Move,
	}
}

func (msg MsgPlayGame) Route() string {
  return RouterKey
}

func (msg MsgPlayGame) Type() string {
  return "PlayGame"
}

func (msg MsgPlayGame) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Player)}
}

func (msg MsgPlayGame) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg MsgPlayGame) ValidateBasic() error {
  if msg.Player.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }
  return nil
}
