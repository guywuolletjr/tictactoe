package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Game struct {
		Creator sdk.AccAddress 	`json:"creator" yaml:"creator"`
		ID      uint         		`json:"id" yaml:"id"`
    Player1 sdk.AccAddress 	`json:"Player1" yaml:"Player1"`
    Player2 sdk.AccAddress 	`json:"Player2" yaml:"Player2"`
    Board 	[9]uint 				`json:"Board" yaml:"Board"`
    Winner 	uint 						`json:"Winner" yaml:"Winner"`
}
