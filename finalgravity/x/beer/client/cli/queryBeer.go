package cli

import (
	"fmt"

	"github.com/aofiee/finalgravity/x/beer/types"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/spf13/cobra"
)

//GetCmdListBeer function
func GetCmdListBeer(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "list-beer",
		Short: "list all beer",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/"+types.QueryListBeer, queryRoute), nil)
			if err != nil {
				fmt.Printf("could not list Post\n%s\n", err.Error())
				return nil
			}
			var out []types.Beer
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}

//GetCmdRetriveBeerWallet function
func GetCmdRetriveBeerWallet(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "get-wallet",
		Short: "Get beer wallet",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			path := fmt.Sprintf("custom/%s/"+types.QueryGetBeerWallet, queryRoute)
			res, _, err := cliCtx.QueryWithData(path, nil)
			fmt.Printf("path %v\n", path)
			if err != nil {
				fmt.Printf("could not list BrewerID\n%s\n", err.Error())
				return nil
			}
			var out types.BeerWallet
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}
