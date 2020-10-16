package recipes

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	// sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/aofiee/finalgravity/x/recipes/keeper"
	"github.com/aofiee/finalgravity/x/recipes/types"
	"github.com/tendermint/tendermint/crypto"
)

func handleMsgCreateRecipes(ctx sdk.Context, k keeper.Keeper, msg types.MsgCreateRecipes) (*sdk.Result, error) {
	var recipes = types.Recipes{}
	k.CreateRecipes(ctx, recipes)
	moduleAcct := sdk.AccAddress(crypto.AddressHash([]byte(types.ModuleName)))
	// totalCoin := k.CoinKeeper.GetCoins(ctx, moduleAcct)
	payment, _ := sdk.ParseCoins("200rune")
	if err := k.CoinKeeper.SendCoins(ctx, msg.Creator, moduleAcct, payment); err != nil {
		return nil, err
	}
	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
