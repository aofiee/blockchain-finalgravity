package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/aofiee/finalgravity/x/recipes/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string, cdc *codec.Codec) *cobra.Command {
	// Group recipes queries under a subcommand
	recipesQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	recipesQueryCmd.AddCommand(
		flags.GetCommands(
			GetCmdListRecipes(queryRoute, cdc),
			GetCmdRetriveRecipesWallet(queryRoute, cdc),
		)...,
	)

	return recipesQueryCmd
}

// TODO: Add Query Commands
