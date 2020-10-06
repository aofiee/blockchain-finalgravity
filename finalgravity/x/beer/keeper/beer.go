package keeper

import (
	"github.com/aofiee/finalgravity/x/beer/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

//CreateBeer function
func (k Keeper) CreateBeer(ctx sdk.Context, beer types.Beer) {
	store := ctx.KVStore(k.storeKey)
	key := []byte(types.BeerPrefix + beer.BeerID)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(beer)
	store.Set(key, value)
}
