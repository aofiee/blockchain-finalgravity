package brewer

import (
	"fmt"

	"github.com/aofiee/finalgravity/x/brewer/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	// "github.com/aofiee/finalgravity/x/brewer/keeper"
)

// NewHandler creates an sdk.Handler for all the brewer type messages
func NewHandler(k Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())
		switch msg := msg.(type) {
		// TODO: Define your msg cases
		//
		//Example:
		// case Msg<Action>:
		// 	return handleMsg<Action>(ctx, k, msg)
		case types.MsgCreateBrewer:
			// fmt.Printf("MsgCreateBrewer")
			return handleMsgCreateBrewer(ctx, k, msg)
		case types.MsgCreateWithdrawCoinsFromModuleWallet:
			// fmt.Printf("brewer types.MsgCreateWithdrawCoinsFromModuleWallet %v", msg)
			return handleMsgCreateWithdrawCoinsFromModuleWallet(ctx, k, msg)
		default:
			fmt.Printf("default")
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", ModuleName, msg)
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}
