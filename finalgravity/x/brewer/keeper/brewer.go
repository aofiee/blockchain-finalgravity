package keeper

import (
	"fmt"

	"github.com/aofiee/finalgravity/x/brewer/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	// abci "github.com/tendermint/tendermint/abci/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

//CreateBrewer function
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

//BytesToString function
func BytesToString(data []byte) string {
	return string(data[:])
}

// GetBrewer function
func (k Keeper) GetBrewer(ctx sdk.Context, brewerID string) types.Brewer {
	store := ctx.KVStore(k.storeKey)
	// fmt.Printf("brewerID from brewer %v\n",brewerID)
	bz := store.Get([]byte(types.BrewerPrefix + brewerID))
	// o := BytesToString(bz)
	// fmt.Printf("o %v\n\n",o)
	// fmt.Printf("bz %v\n",bz)
	var brewer types.Brewer
	// k.cdc.MustUnmarshalBinaryBare(bz, &brewer) ในกรณีไม่ใช้ prefix
	k.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &brewer)
	return brewer
}

//GetBrewerByID function
func GetBrewerByID(ctx sdk.Context, path []string, k Keeper) ([]byte, error) {
	brewer := k.GetBrewer(ctx, path[0])
	fmt.Printf("GetBrewerByID path[1] %v\n", path[0])
	res, err := codec.MarshalJSONIndent(k.cdc, brewer)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}
