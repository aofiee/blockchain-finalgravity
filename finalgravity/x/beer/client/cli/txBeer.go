package cli

import (
	"bufio"
	"fmt"
	"time"

	"github.com/spf13/cobra"
	// "github.com/tendermint/tendermint/crypto"

	"github.com/aofiee/finalgravity/x/beer/types"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
)

//GetCmdCreateBeer function
func GetCmdCreateBeer(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "create-beer [BrewerID] [BeerName] [TypeOfBeer] [StyleBeer] [TagLine] [Appearance] [Taste] [Aroma] [Strength] [Story] [ImageLabel]",
		Short: "Creates a new Beer",
		Args:  cobra.ExactArgs(11),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsBrewerID := string(args[0])
			argsBeerName := string(args[1])
			argsTypeOfBeer := string(args[2])
			argsStyleBeer := string(args[3])
			argsTagLine := string(args[4])
			argsAppearance := string(args[5])
			argsTaste := string(args[6])
			argsAroma := string(args[7])
			argsStrength := string(args[8])
			argsStory := string(args[9])
			argsImageLabel := string(args[10])

			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			createAt := fmt.Sprintf("%s", time.Now())
			msg := types.NewMsgCreateBeer(cliCtx.GetFromAddress(), argsBrewerID, argsBeerName, argsTypeOfBeer, argsStyleBeer, argsTagLine, argsAppearance, argsTaste, argsAroma, argsStrength, argsStory, argsImageLabel, createAt)
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}
