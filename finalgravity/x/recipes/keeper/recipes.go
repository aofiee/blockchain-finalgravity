package keeper

import (
	"github.com/aofiee/finalgravity/x/recipes/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	// abci "github.com/tendermint/tendermint/abci/types"
)

//CreateRecipes function
func (k Keeper) CreateRecipes(ctx sdk.Context, recipes types.Recipes) {
	store := ctx.KVStore(k.storeKey)
	key := []byte(types.RecipesPrefix + recipes.RecipesID)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(recipes)
	store.Set(key, value)
}

func listRecipes(ctx sdk.Context, k Keeper) ([]byte, error) {
	var recipesList []types.Recipes
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, []byte(types.RecipesPrefix))
	for ; iterator.Valid(); iterator.Next() {
		var recipes types.Recipes
		k.cdc.MustUnmarshalBinaryLengthPrefixed(store.Get(iterator.Key()), &recipes)
		recipesList = append(recipesList, recipes)
	}
	res := codec.MustMarshalJSONIndent(k.cdc, recipesList)
	return res, nil
}
