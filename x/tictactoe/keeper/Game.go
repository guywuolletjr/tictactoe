package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/guywuolletjr/tictactoe/x/tictactoe/types"
    "github.com/cosmos/cosmos-sdk/codec"
)

// CreateGame creates a Game
func (k Keeper) CreateGame(ctx sdk.Context, Game types.Game) {
	store := ctx.KVStore(k.storeKey)
	key := []byte(types.GamePrefix + Game.ID)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(Game)
	store.Set(key, value)
}

// GetGame returns the Game information
func (k Keeper) GetGame(ctx sdk.Context, key string) (types.Game, error) {
	store := ctx.KVStore(k.storeKey)
	var Game types.Game
	byteKey := []byte(types.GamePrefix + key)
	err := k.cdc.UnmarshalBinaryLengthPrefixed(store.Get(byteKey), &Game)
	if err != nil {
		return Game, err
	}
	return Game, nil
}

// SetGame sets a Game
func (k Keeper) SetGame(ctx sdk.Context, Game types.Game) {
	GameKey := Game.ID
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(Game)
	key := []byte(types.GamePrefix + GameKey)
	store.Set(key, bz)
}

// DeleteGame deletes a Game
func (k Keeper) DeleteGame(ctx sdk.Context, key string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete([]byte(types.GamePrefix + key))
}

//
// Functions used by querier
//

func listGame(ctx sdk.Context, k Keeper) ([]byte, error) {
	var GameList []types.Game
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, []byte(types.GamePrefix))
	for ; iterator.Valid(); iterator.Next() {
		var Game types.Game
		k.cdc.MustUnmarshalBinaryLengthPrefixed(store.Get(iterator.Key()), &Game)
		GameList = append(GameList, Game)
	}
	res := codec.MustMarshalJSONIndent(k.cdc, GameList)
	return res, nil
}

func getGame(ctx sdk.Context, path []string, k Keeper) (res []byte, sdkError error) {
	key := path[0]
	Game, err := k.GetGame(ctx, key)
	if err != nil {
		return nil, err
	}

	res, err = codec.MarshalJSONIndent(k.cdc, Game)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

// Get creator of the item
func (k Keeper) GetGameOwner(ctx sdk.Context, key string) sdk.AccAddress {
	Game, err := k.GetGame(ctx, key)
	if err != nil {
		return nil
	}
	return Game.Creator
}


// Check if the key exists in the store
func (k Keeper) GameExists(ctx sdk.Context, key string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has([]byte(types.GamePrefix + key))
}
