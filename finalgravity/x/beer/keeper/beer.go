package keeper

import (
	"github.com/aofiee/finalgravity/x/beer/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

//CreateBeer function
func (k Keeper) CreateBeer(ctx sdk.Context, beer types.Beer) {
	store := ctx.KVStore(k.storeKey)
	key := []byte(types.BeerPrefix + beer.BeerID)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(beer)
	store.Set(key, value)
}

func listBeer(ctx sdk.Context, k Keeper) ([]byte, error) {
	var beerList []types.Beer
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, []byte(types.BeerPrefix))
	for ; iterator.Valid(); iterator.Next() {
		var beer types.Beer
		k.cdc.MustUnmarshalBinaryLengthPrefixed(store.Get(iterator.Key()), &beer)
		beerList = append(beerList, beer)
	}
	res := codec.MustMarshalJSONIndent(k.cdc, beerList)
	return res, nil
}
