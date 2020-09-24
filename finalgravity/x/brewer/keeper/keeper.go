package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/aofiee/finalgravity/x/brewer/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Keeper of the brewer store
type Keeper struct {
	storeKey   sdk.StoreKey
	cdc        *codec.Codec
	paramspace types.ParamSubspace
}

// NewKeeper creates a brewer keeper
func NewKeeper(cdc *codec.Codec, key sdk.StoreKey, paramspace types.ParamSubspace) Keeper {
	keeper := Keeper{
		storeKey:   key,
		cdc:        cdc,
		paramspace: paramspace.WithKeyTable(types.ParamKeyTable()),
	}
	return keeper
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// GetBrewer returns the pubkey from the adddress-pubkey relation
func (k Keeper) GetBrewer(ctx sdk.Context, key string) (types.Brewer, error) {
	store := ctx.KVStore(k.storeKey)
	var brewer types.Brewer
	byteKey := []byte(key)
	err := k.cdc.UnmarshalBinaryLengthPrefixed(store.Get(byteKey), &brewer)
	if err != nil {
		return brewer, err
	}
	return brewer, nil
}

// SetBrewer the entire Brewer metadata struct for a name
func (k Keeper) SetBrewer(ctx sdk.Context, key string, brewer types.Brewer) {
	if brewer.Creator.Empty() {
		return
	}
	fmt.Printf("brewer is %v", brewer)
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(brewer)
	store.Set([]byte(key), bz)
}

// DeleteBrewer the entire Brewer metadata struct for a name
func (k Keeper) DeleteBrewer(ctx sdk.Context, name string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete([]byte(name))
}

// GetBrewerIterator an iterator over all names in which the keys are the names and the values are the brewer
func (k Keeper) GetBrewerIterator(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return sdk.KVStorePrefixIterator(store, []byte{})
}
