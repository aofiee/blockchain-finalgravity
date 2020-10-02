package keeper

import (
	"fmt"

	"github.com/aofiee/finalgravity/x/brewer/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	// abci "github.com/tendermint/tendermint/abci/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/tendermint/tendermint/crypto"
)

//CreateWitdraw function
func (k Keeper) CreateWitdraw(ctx sdk.Context, wallet types.WithdrawCoinsFromModuleWallet) {
	fmt.Printf("CreateWitdraw %v\n", wallet)
	store := ctx.KVStore(k.storeKey)
	key := []byte(types.BrewerWalletPrefix + wallet.WithdrawID)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(wallet)
	store.Set(key, value)
}

//GetModuleBalance function
func GetModuleBalance(ctx sdk.Context, k Keeper) ([]byte, error) {
	moduleAcct := sdk.AccAddress(crypto.AddressHash([]byte(types.ModuleName)))
	fmt.Printf("moduleAcct %v\n", moduleAcct)
	totalCoin := k.CoinKeeper.GetCoins(ctx, moduleAcct)
	fmt.Printf("totalCoin %v\n", totalCoin)
	var wallet = types.BrewerWallet{
		Creator: moduleAcct,
		Amount:  totalCoin,
	}
	res, err := codec.MarshalJSONIndent(k.cdc, wallet)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}
	return res, nil
}
