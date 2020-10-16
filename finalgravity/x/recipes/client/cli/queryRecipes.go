package cli

import (
	"fmt"

	"github.com/aofiee/finalgravity/x/recipes/types"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/spf13/cobra"
)

//GetCmdListRecipes function
func GetCmdListRecipes(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "list-recipes",
		Short: "list all recipes",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/"+types.QueryListRecipes, queryRoute), nil)
			if err != nil {
				fmt.Printf("could not list Post\n%s\n", err.Error())
				return nil
			}
			var out []types.Recipes
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}

//GetCmdRetriveRecipesWallet function
func GetCmdRetriveRecipesWallet(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "get-wallet",
		Short: "Get recipes wallet",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			path := fmt.Sprintf("custom/%s/"+types.QueryGetRecipesWallet, queryRoute)
			res, _, err := cliCtx.QueryWithData(path, nil)
			fmt.Printf("path %v\n", path)
			if err != nil {
				fmt.Printf("could not list RecipesID\n%s\n", err.Error())
				return nil
			}
			var out types.RecipesWallet
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}
