package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

// RegisterCodec registers concrete types on codec
func RegisterCodec(cdc *codec.Codec) {
  // this line is used by starport scaffolding # 1
		cdc.RegisterConcrete(MsgCreateGame{}, "tictactoe/CreateGame", nil)
		cdc.RegisterConcrete(MsgSetGame{}, "tictactoe/SetGame", nil)
		cdc.RegisterConcrete(MsgDeleteGame{}, "tictactoe/DeleteGame", nil)
}

// ModuleCdc defines the module codec
var ModuleCdc *codec.Codec

func init() {
	ModuleCdc = codec.New()
	RegisterCodec(ModuleCdc)
	codec.RegisterCrypto(ModuleCdc)
	ModuleCdc.Seal()
}
