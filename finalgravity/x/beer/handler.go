package beer

import (
	"fmt"

	"github.com/aofiee/finalgravity/x/beer/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewHandler creates an sdk.Handler for all the beer type messages
func NewHandler(k Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())
		switch msg := msg.(type) {
		case types.MsgCreateBeer:
			// fmt.Printf("MsgCreateBeer")
			return handleMsgCreateBeer(ctx, k, msg)
		case types.MsgCreateWithdrawCoinsFromModuleWallet:
			// fmt.Printf("beer types.MsgCreateWithdrawCoinsFromModuleWallet %v", msg)
			return handleMsgCreateWithdrawCoinsFromModuleWallet(ctx, k, msg)
		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", ModuleName, msg)
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}
