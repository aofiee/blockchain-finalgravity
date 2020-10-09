package keeper

import (
	"github.com/aofiee/finalgravity/x/beer/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	// abci "github.com/tendermint/tendermint/abci/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/tendermint/tendermint/crypto"
)

//CreateWitdraw function
func (k Keeper) CreateWitdraw(ctx sdk.Context, wallet types.WithdrawCoinsFromModuleWallet) {
	store := ctx.KVStore(k.storeKey)
	key := []byte(types.BeerWalletPrefix + wallet.WithdrawID)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(wallet)
	store.Set(key, value)
}

//GetModuleBalance function
func GetModuleBalance(ctx sdk.Context, k Keeper) ([]byte, error) {
	moduleAcct := sdk.AccAddress(crypto.AddressHash([]byte(types.ModuleName)))
	totalCoin := k.CoinKeeper.GetCoins(ctx, moduleAcct)
	var wallet = types.BeerWallet{
		Creator: moduleAcct,
		Amount:  totalCoin,
	}
	res, err := codec.MarshalJSONIndent(k.cdc, wallet)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}
	return res, nil
}
