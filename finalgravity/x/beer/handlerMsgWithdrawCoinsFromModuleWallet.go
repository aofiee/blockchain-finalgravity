package beer

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/crypto"

	// sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/aofiee/finalgravity/x/beer/types"
)

func handleMsgCreateWithdrawCoinsFromModuleWallet(ctx sdk.Context, k Keeper, msg types.MsgCreateWithdrawCoinsFromModuleWallet) (*sdk.Result, error) {
	fmt.Printf("handleMsgWithdrawCoinsFromModuleWallet\n")
	var wallet = types.WithdrawCoinsFromModuleWallet{
		WithdrawID: msg.WithdrawID,
		Sender:     msg.Sender,
		Reciever:   msg.Reciever,
		Amount:     msg.Amount,
	}
	k.CreateWitdraw(ctx, wallet)
	moduleAcct := sdk.AccAddress(crypto.AddressHash([]byte(types.ModuleName)))
	err := k.CoinKeeper.SendCoins(ctx, moduleAcct, msg.Reciever, msg.Amount)
	fmt.Printf("err %v\n", err)
	if err != nil {
		return nil, err
	}
	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
