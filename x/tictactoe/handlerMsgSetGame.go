package tictactoe

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/guywuolletjr/tictactoe/x/tictactoe/types"
	"github.com/guywuolletjr/tictactoe/x/tictactoe/keeper"
)

func handleMsgSetGame(ctx sdk.Context, k keeper.Keeper, msg types.MsgSetGame) (*sdk.Result, error) {
	var Game = types.Game{
		Creator: msg.Creator,
		ID:      msg.ID,
    	Id: msg.Id,
    	Player1: msg.Player1,
    	Player2: msg.Player2,
    	Board: msg.Board,
    	Winner: msg.Winner,
	}
	if !msg.Creator.Equals(k.GetGameOwner(ctx, msg.ID)) { // Checks if the the msg sender is the same as the current owner
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "Incorrect Owner") // If not, throw an error
	}

	k.SetGame(ctx, Game)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
