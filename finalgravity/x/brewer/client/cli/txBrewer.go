package cli

import (
	"bufio"
	"fmt"

	"github.com/spf13/cobra"
	// "github.com/tendermint/tendermint/crypto"

	"github.com/aofiee/finalgravity/x/brewer/types"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
)

//GetCmdCreateBrewer function
func GetCmdCreateBrewer(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "create-brewer [TypeOfBrewer] [Address] [Telephone] [Email] [Story] [LogoURL] [CoverURL] [Founded] [Founder] [SiteURL]",
		Short: "Creates a new brewer",
		Args:  cobra.ExactArgs(10),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsTypeOfBrewer := string(args[0])
			argsAddress := string(args[1])
			argsTelephone := string(args[2])
			argsEmail := string(args[3])
			argsStory := string(args[4])
			argsLogoURL := string(args[5])
			argsCoverURL := string(args[6])
			argsFounded := string(args[7])
			argsFounder := string(args[8])
			argsSiteURL := string(args[9])

			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

			msg := types.NewMsgCreateBrewer(cliCtx.GetFromAddress(), argsTypeOfBrewer, argsAddress, argsTelephone, argsEmail, argsStory, argsLogoURL, argsCoverURL, argsFounded, argsFounder, argsSiteURL)
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

//GetCmdCreateWithdrawCoinsFromModuleWallet function
func GetCmdCreateWithdrawCoinsFromModuleWallet(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "withdraw-to [Amount]",
		Short: "Withdraw Coins to Address",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			Amount := string(args[0])
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			toAddress := cliCtx.GetFromAddress()
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			coins, _ := sdk.ParseCoins(Amount)
			fmt.Printf("Address %v\n", toAddress)
			msg := types.NewMsgCreateWithdrawCoinsFromModuleWallet(toAddress, coins)
			fmt.Printf("msg %v\n", msg)
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}
