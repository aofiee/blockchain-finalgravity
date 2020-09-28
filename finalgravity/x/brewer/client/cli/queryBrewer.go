package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/aofiee/finalgravity/x/brewer/types"
	"github.com/spf13/cobra"
)

func GetCmdListBrewer(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "list-brewer",
		Short: "list all brewer",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/"+types.QueryListBrewer, queryRoute), nil)
			if err != nil {
				fmt.Printf("could not list Post\n%s\n", err.Error())
				return nil
			}
			var out []types.Brewer
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}

func GetCmdRetriveBrewerByID(queryRoute string, cdc *codec.Codec)  *cobra.Command {
	return &cobra.Command{
		Use:   "get-brewer [BrewerID]",
		Short: "Get brewer by BrewerID",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			BrewerID := args[0]
			path := fmt.Sprintf("custom/%s/"+types.QueryGetBrewerByID+"/%s/", queryRoute,BrewerID)
			res, _, err := cliCtx.QueryWithData(path, nil)
			if err != nil {
				fmt.Printf("could not list BrewerID\n%s\n", err.Error())
				return nil
			}
			//fmt.Printf("GetCmdRetriveModuleAddress %v",res)
			var out types.Brewer
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
		
	}
}