package keeper

import (
	// "fmt"

	abci "github.com/tendermint/tendermint/abci/types"

	// "github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/aofiee/finalgravity/x/brewer/types"
)

// NewQuerier creates a new querier for brewer clients.
func NewQuerier(k Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, error) {
		switch path[0] {
			// TODO: Put the modules query routes
		case types.QueryListBrewer:
			return listBrewer(ctx,k)
		default:
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "unknown brewer query endpoint")
		}
	}
}