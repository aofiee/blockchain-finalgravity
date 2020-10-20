package keeper

import (
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/aofiee/finalgravity/x/recipes/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewQuerier creates a new querier for recipes clients.
func NewQuerier(k Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, error) {
		switch path[0] {
		case types.QueryListRecipes:
			return listRecipes(ctx, k)
		case types.QueryGetRecipesWallet:
			return GetModuleBalance(ctx, k)
		case types.QueryGetRecipesByID:
			return GetRecipesByID(ctx, path[1:], k)
		default:
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "unknown recipes query endpoint")
		}
	}
}
