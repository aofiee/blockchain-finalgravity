package keeper

import (
	"fmt"

	"github.com/aofiee/finalgravity/x/recipes/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	// abci "github.com/tendermint/tendermint/abci/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
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

// GetRecipes function
func (k Keeper) GetRecipes(ctx sdk.Context, recipesID string) types.Recipes {
	store := ctx.KVStore(k.storeKey)
	fmt.Printf("types.RecipesPrefix + recipesID %v", types.RecipesPrefix+recipesID)
	bz := store.Get([]byte(types.RecipesPrefix + recipesID))
	var recipes types.Recipes
	fmt.Printf("recipes %s", bz)
	k.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &recipes)
	return recipes
}

//GetRecipesByID function
func GetRecipesByID(ctx sdk.Context, path []string, k Keeper) ([]byte, error) {
	recipes := k.GetRecipes(ctx, path[0])
	fmt.Printf("GetRecipesByID path[1] %v\n", path[0])
	res, err := codec.MarshalJSONIndent(k.cdc, recipes)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}
