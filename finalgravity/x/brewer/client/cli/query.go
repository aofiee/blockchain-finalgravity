package cli

import (
	"fmt"
	// "strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	// "github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	// sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/aofiee/finalgravity/x/brewer/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string, cdc *codec.Codec) *cobra.Command {
	// Group brewer queries under a subcommand
	brewerQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	brewerQueryCmd.AddCommand(
		flags.GetCommands(
	// TODO: Add query Cmds
			GetCmdListBrewer(queryRoute, cdc),
			GetCmdRetriveBrewerByID(queryRoute, cdc),
		)...,
	)

	return brewerQueryCmd
}

// TODO: Add Query Commands
