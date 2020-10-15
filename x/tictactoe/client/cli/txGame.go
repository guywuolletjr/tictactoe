package cli

import (
	"bufio"
    
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/guywuolletjr/tictactoe/x/tictactoe/types"
)

func GetCmdCreateGame(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "create-Game [Id] [Player1] [Player2] [Board] [Winner]",
		Short: "Creates a new Game",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsId := string(args[0] )
			argsPlayer1 := string(args[1] )
			argsPlayer2 := string(args[2] )
			argsBoard := string(args[3] )
			argsWinner := string(args[4] )
			
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			msg := types.NewMsgCreateGame(cliCtx.GetFromAddress(), string(argsId), string(argsPlayer1), string(argsPlayer2), string(argsBoard), string(argsWinner))
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}


func GetCmdSetGame(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "set-Game [id]  [Id] [Player1] [Player2] [Board] [Winner]",
		Short: "Set a new Game",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id := args[0]
			argsId := string(args[1])
			argsPlayer1 := string(args[2])
			argsPlayer2 := string(args[3])
			argsBoard := string(args[4])
			argsWinner := string(args[5])
			
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			msg := types.NewMsgSetGame(cliCtx.GetFromAddress(), id, string(argsId), string(argsPlayer1), string(argsPlayer2), string(argsBoard), string(argsWinner))
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

func GetCmdDeleteGame(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "delete-Game [id]",
		Short: "Delete a new Game by ID",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {

			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

			msg := types.NewMsgDeleteGame(args[0], cliCtx.GetFromAddress())
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}
