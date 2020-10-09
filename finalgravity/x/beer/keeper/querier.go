package keeper

import (
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/aofiee/finalgravity/x/beer/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewQuerier creates a new querier for beer clients.
func NewQuerier(k Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, error) {
		switch path[0] {
		case types.QueryListBeer:
			return listBeer(ctx, k)
		case types.QueryGetBeerWallet:
			return GetModuleBalance(ctx, k)
		default:
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "unknown beer query endpoint")
		}
	}
}
