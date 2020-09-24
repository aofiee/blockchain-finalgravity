package brewer

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	// sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/aofiee/finalgravity/x/brewer/types"
	"github.com/aofiee/finalgravity/x/brewer/keeper"
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
	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
