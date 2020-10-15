package rest

import (
	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client/context"
)

// RegisterRoutes registers tictactoe-related REST handlers to a router
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router) {
  // this line is used by starport scaffolding # 1
		r.HandleFunc("/tictactoe/Game", createGameHandler(cliCtx)).Methods("POST")
		r.HandleFunc("/tictactoe/Game", listGameHandler(cliCtx, "tictactoe")).Methods("GET")
		r.HandleFunc("/tictactoe/Game/{key}", getGameHandler(cliCtx, "tictactoe")).Methods("GET")
		r.HandleFunc("/tictactoe/Game", setGameHandler(cliCtx)).Methods("PUT")
		r.HandleFunc("/tictactoe/Game", deleteGameHandler(cliCtx)).Methods("DELETE")

		
}
