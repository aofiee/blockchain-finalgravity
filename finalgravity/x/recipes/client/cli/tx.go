package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/aofiee/finalgravity/x/recipes/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd(cdc *codec.Codec) *cobra.Command {
	recipesTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	recipesTxCmd.AddCommand(flags.PostCommands(
		GetCmdCreateRecipes(cdc),
		GetCmdCreateWithdrawCoinsFromModuleWallet(cdc),
	)...)

	return recipesTxCmd
}
