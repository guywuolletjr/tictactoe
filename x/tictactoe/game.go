package tictactoe

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Game struct {
  id        uint              `json:"id"`
  player1   sdk.AccAddress    `json:"player1"`
  player2   sdk.AccAddress    `json:"player2"`
  board     [9]int            `json:"board"`
  winner    uint              `json:"winner"`
}
