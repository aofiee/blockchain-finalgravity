package beer

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/crypto"

	"github.com/aofiee/finalgravity/x/beer/keeper"
	"github.com/aofiee/finalgravity/x/beer/types"
)

func handleMsgCreateBeer(ctx sdk.Context, k keeper.Keeper, msg types.MsgCreateBeer) (*sdk.Result, error) {
	var beer = types.Beer{
		Creator:    msg.Creator,
		BrewerID:   msg.BrewerID,
		BeerID:     msg.BeerID,
		BeerName:   msg.BeerName,
		TypeOfBeer: msg.TypeOfBeer,
		StyleBeer:  msg.StyleBeer,
		TagLine:    msg.TagLine,
		Appearance: msg.Appearance,
		Taste:      msg.Taste,
		Aroma:      msg.Aroma,
		Strength:   msg.Strength,
		Story:      msg.Story,
		ImageLabel: msg.ImageLabel,
		CreatedAt:  msg.CreatedAt,
	}
	k.CreateBeer(ctx, beer)
	moduleAcct := sdk.AccAddress(crypto.AddressHash([]byte(types.ModuleName)))
	fmt.Printf("moduleAcct %v\n", moduleAcct)
	totalCoin := k.CoinKeeper.GetCoins(ctx, moduleAcct)
	fmt.Printf("totalCoin %v\n", totalCoin)
	payment, _ := sdk.ParseCoins("100rune")
	if err := k.CoinKeeper.SendCoins(ctx, msg.Creator, moduleAcct, payment); err != nil {
		return nil, err
	}
	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
