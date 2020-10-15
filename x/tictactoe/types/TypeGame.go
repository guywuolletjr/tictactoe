package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Game struct {
	Creator sdk.AccAddress `json:"creator" yaml:"creator"`
	ID      string         `json:"id" yaml:"id"`
    Id string `json:"Id" yaml:"Id"`
    Player1 string `json:"Player1" yaml:"Player1"`
    Player2 string `json:"Player2" yaml:"Player2"`
    Board string `json:"Board" yaml:"Board"`
    Winner string `json:"Winner" yaml:"Winner"`
}