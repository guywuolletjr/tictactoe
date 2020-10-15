package tictactoe

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/guywuolletjr/tictactoe/x/tictactoe/types"
	"github.com/guywuolletjr/tictactoe/x/tictactoe/keeper"
)

func handleMsgCreateGame(ctx sdk.Context, k keeper.Keeper, msg types.MsgCreateGame) (*sdk.Result, error) {
	var Game = types.Game{
		Creator: msg.Creator,
		ID:      msg.ID,
    	Id: msg.Id,
    	Player1: msg.Player1,
    	Player2: msg.Player2,
    	Board: msg.Board,
    	Winner: msg.Winner,
	}
	k.CreateGame(ctx, Game)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
