package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/aofiee/finalgravity/x/brewer/types"
)

func (k Keeper) CreateBrewer(ctx sdk.Context, brewer types.Brewer) {
	store := ctx.KVStore(k.storeKey)
	key := []byte(types.BrewerPrefix + brewer.BrewerID)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(brewer)
	store.Set(key, value)
}

func listBrewer(ctx sdk.Context, k Keeper) ([]byte, error) {
	var brewerList []types.Brewer
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, []byte(types.BrewerPrefix))
	for ; iterator.Valid(); iterator.Next() {
		var brewer types.Brewer
		k.cdc.MustUnmarshalBinaryLengthPrefixed(store.Get(iterator.Key()), &brewer)
		brewerList = append(brewerList, brewer)
	}
	res := codec.MustMarshalJSONIndent(k.cdc, brewerList)
	return res, nil
}