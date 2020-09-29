package brewer

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	// sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/aofiee/finalgravity/x/brewer/types"
	"github.com/aofiee/finalgravity/x/brewer/keeper"
	"github.com/tendermint/tendermint/crypto"
)

func handleMsgCreateBrewer(ctx sdk.Context, k keeper.Keeper, msg types.MsgCreateBrewer) (*sdk.Result, error) {
	var brewer = types.Brewer{
		Creator:      msg.Creator,
		BrewerID:     msg.BrewerID,
		TypeOfBrewer: msg.TypeOfBrewer,
		Address:      msg.Address,
		Telephone:    msg.Telephone,
		Email:        msg.Email,
		Story:        msg.Story,
		LogoURL:      msg.LogoURL,
		CoverURL:     msg.CoverURL,
		Founded:      msg.Founded,
		Founder:      msg.Founder,
		SiteURL:      msg.SiteURL,
	}
	k.CreateBrewer(ctx, brewer)
	moduleAcct := sdk.AccAddress(crypto.AddressHash([]byte(types.ModuleName)))
	fmt.Printf("moduleAcct %v\n",moduleAcct)
	totalCoin := k.CoinKeeper.GetCoins(ctx, moduleAcct)
	fmt.Printf("totalCoin %v\n",totalCoin)
	payment, _ := sdk.ParseCoins("200rune")
	if err := k.CoinKeeper.SendCoins(ctx, msg.Creator, moduleAcct, payment); err != nil {
		return nil, err
	}
	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
