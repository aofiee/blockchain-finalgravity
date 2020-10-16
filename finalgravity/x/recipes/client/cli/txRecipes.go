package cli

import (
	"bufio"
	"fmt"

	"github.com/spf13/cobra"
	// "github.com/tendermint/tendermint/crypto"

	"github.com/aofiee/finalgravity/x/recipes/types"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
)

//GetCmdCreateRecipes function
func GetCmdCreateRecipes(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "create-recipes [BrewerID] [BeerID] [BatchNo] [JSON Properties] [JSON Mashing] [JSON Boil] [JSON Ferment] [JSON Ages] [Factory] [ImageLabel]",
		Short: "Creates a new recipes",
		Args:  cobra.MinimumNArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {

			BrewerID := string(args[0])
			BeerID := string(args[1])
			BatchNo := string(args[2])
			Factory := string(args[8])
			ImageLabel := string(args[9])

			var properties types.MsgProperties
			cdc.UnmarshalJSON([]byte(args[3]), &properties)
			var mashing types.MsgMashing
			cdc.UnmarshalJSON([]byte(args[4]), &mashing)
			var boil types.MsgBoil
			cdc.UnmarshalJSON([]byte(args[5]), &boil)
			var ferment types.MsgFerment
			cdc.UnmarshalJSON([]byte(args[6]), &ferment)
			var ages types.MsgAges
			cdc.UnmarshalJSON([]byte(args[7]), &ages)
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			//check properties
			if properties.OriginalGravity == "" {
				errMsg := fmt.Sprintf("properties is %v", properties)
				return sdkerrors.Wrap(types.ErrPropertiesOriginalGravityEmpty, errMsg)
			}
			if properties.ExpectedFinalGravity == "" {
				errMsg := fmt.Sprintf("properties is %v", properties)
				return sdkerrors.Wrap(types.ErrPropertiesExpectedFinalGravityEmpty, errMsg)
			}
			if properties.FinalGravity == "" {
				errMsg := fmt.Sprintf("properties is %v", properties)
				return sdkerrors.Wrap(types.ErrPropertiesFinalGravityEmpty, errMsg)
			}
			if properties.TotalLiquor == "" {
				errMsg := fmt.Sprintf("properties is %v", properties)
				return sdkerrors.Wrap(types.ErrPropertiesTotalLiquorEmpty, errMsg)
			}
			if properties.Makes == "" {
				errMsg := fmt.Sprintf("properties is %v", properties)
				return sdkerrors.Wrap(types.ErrPropertiesMakesEmpty, errMsg)
			}
			if properties.ReadyToDrink == "" {
				errMsg := fmt.Sprintf("properties is %v", properties)
				return sdkerrors.Wrap(types.ErrPropertiesReadyToDrinkEmpty, errMsg)
			}
			if properties.EstimatedABV == "" {
				errMsg := fmt.Sprintf("properties is %v", properties)
				return sdkerrors.Wrap(types.ErrPropertiesEstimatedABVEmpty, errMsg)
			}
			if properties.BitternessRating == "" {
				errMsg := fmt.Sprintf("properties is %v", properties)
				return sdkerrors.Wrap(types.ErrPropertiesBitternessRatingEmpty, errMsg)
			}
			if properties.ColourRating == "" {
				errMsg := fmt.Sprintf("properties is %v", properties)
				return sdkerrors.Wrap(types.ErrPropertiesColourRatingEmpty, errMsg)
			}

			if mashing.Liquor == "" {
				errMsg := fmt.Sprintf("mashing is %v", mashing)
				return sdkerrors.Wrap(types.ErrMashingLiquorEmpty, errMsg)
			}
			if mashing.MashTime == "" {
				errMsg := fmt.Sprintf("mashing is %v", mashing)
				return sdkerrors.Wrap(types.ErrMashingMashTimeEmpty, errMsg)
			}
			if mashing.Temperature == "" {
				errMsg := fmt.Sprintf("mashing is %v", mashing)
				return sdkerrors.Wrap(types.ErrMashingTemperatureEmpty, errMsg)
			}
			if len(mashing.GrainBillProperties) == 0 {
				errMsg := fmt.Sprintf("mashing is %v", mashing)
				return sdkerrors.Wrap(types.ErrMashingGrainBillPropertiesEmpty, errMsg)
			}
			for i := range mashing.GrainBillProperties {
				if mashing.GrainBillProperties[i].Name == "" {
					errMsg := fmt.Sprintf("mashing is %v", mashing)
					return sdkerrors.Wrap(types.ErrMashingGrainBillPropertiesNameEmpty, errMsg)
				}
				if mashing.GrainBillProperties[i].Quantity == "" {
					errMsg := fmt.Sprintf("mashing is %v", mashing)
					return sdkerrors.Wrap(types.ErrMashingGrainBillPropertiesQuantityEmpty, errMsg)
				}
			}

			if boil.Liquor == "" {
				errMsg := fmt.Sprintf("boil is %v", boil)
				return sdkerrors.Wrap(types.ErrBoilLiquorEmpty, errMsg)
			}
			if boil.BoilTime == "" {
				errMsg := fmt.Sprintf("boil is %v", boil)
				return sdkerrors.Wrap(types.ErrBoilBoilTimeEmpty, errMsg)
			}
			if len(boil.HopsProperties) == 0 {
				errMsg := fmt.Sprintf("boil is %v", boil)
				return sdkerrors.Wrap(types.ErrBoilHopsPropertiesEmpty, errMsg)
			}
			for i := range boil.HopsProperties {
				if boil.HopsProperties[i].Name == "" {
					errMsg := fmt.Sprintf("boil is %v", boil)
					return sdkerrors.Wrap(types.ErrBoilHopsPropertiesNameEmpty, errMsg)
				}
				if boil.HopsProperties[i].Quantity == "" {
					errMsg := fmt.Sprintf("boil is %v", boil)
					return sdkerrors.Wrap(types.ErrBoilHopsPropertiesQuantityEmpty, errMsg)
				}
				if boil.HopsProperties[i].IBU == "" {
					errMsg := fmt.Sprintf("boil is %v", boil)
					return sdkerrors.Wrap(types.ErrBoilHopsPropertiesIBUEmpty, errMsg)
				}
				if boil.HopsProperties[i].WhenToAdd == "" {
					errMsg := fmt.Sprintf("boil is %v", boil)
					return sdkerrors.Wrap(types.ErrBoilHopsPropertiesWhenToAddEmpty, errMsg)
				}
			}
			for i := range boil.OtherProperties {
				if boil.OtherProperties[i].Name == "" {
					errMsg := fmt.Sprintf("boil is %v", boil)
					return sdkerrors.Wrap(types.ErrBoilOtherPropertiesNameEmpty, errMsg)
				}
				if boil.OtherProperties[i].Quantity == "" {
					errMsg := fmt.Sprintf("boil is %v", boil)
					return sdkerrors.Wrap(types.ErrBoilOtherPropertiesQuantityEmpty, errMsg)
				}
				if boil.OtherProperties[i].WhenToAdd == "" {
					errMsg := fmt.Sprintf("boil is %v", boil)
					return sdkerrors.Wrap(types.ErrBoilOtherPropertiesWhenToAddEmpty, errMsg)
				}
			}

			if len(ferment.YeastProperties) == 0 {
				errMsg := fmt.Sprintf("ferment is %v", ferment)
				return sdkerrors.Wrap(types.ErrFermentYeastPropertiesEmpty, errMsg)
			}
			for i := range ferment.YeastProperties {
				if ferment.YeastProperties[i].Name == "" {
					errMsg := fmt.Sprintf("ferment is %v", ferment)
					return sdkerrors.Wrap(types.ErrFermentYeastPropertiesNameEmpty, errMsg)
				}
				if ferment.YeastProperties[i].Quantity == "" {
					errMsg := fmt.Sprintf("ferment is %v", ferment)
					return sdkerrors.Wrap(types.ErrFermentYeastPropertiesQuantityEmpty, errMsg)
				}
			}
			if ferment.Fermentation.Temperature == "" {
				errMsg := fmt.Sprintf("ferment is %v", ferment)
				return sdkerrors.Wrap(types.ErrFermentFermentationTemperatureEmpty, errMsg)
			}
			if ferment.Fermentation.Conditioning == "" {
				errMsg := fmt.Sprintf("ferment is %v", ferment)
				return sdkerrors.Wrap(types.ErrFermentFermentationConditioningEmpty, errMsg)
			}
			if len(ferment.HopsProperties) == 0 {
				errMsg := fmt.Sprintf("ferment is %v", ferment)
				return sdkerrors.Wrap(types.ErrFermentHopsPropertiesEmpty, errMsg)
			}
			for i := range ferment.HopsProperties {
				if ferment.HopsProperties[i].Name == "" {
					errMsg := fmt.Sprintf("ferment is %v", ferment)
					return sdkerrors.Wrap(types.ErrFermentHopsPropertiesNameEmpty, errMsg)
				}
				if ferment.HopsProperties[i].Quantity == "" {
					errMsg := fmt.Sprintf("ferment is %v", ferment)
					return sdkerrors.Wrap(types.ErrFermentHopsPropertiesQuantityEmpty, errMsg)
				}
				if ferment.HopsProperties[i].IBU == "" {
					errMsg := fmt.Sprintf("ferment is %v", ferment)
					return sdkerrors.Wrap(types.ErrFermentHopsPropertiesIBUEmpty, errMsg)
				}
				if ferment.HopsProperties[i].WhenToAdd == "" {
					errMsg := fmt.Sprintf("ferment is %v", ferment)
					return sdkerrors.Wrap(types.ErrFermentHopsPropertiesWhenToAddEmpty, errMsg)
				}
			}
			for i := range ferment.OtherProperties {
				if ferment.OtherProperties[i].Name == "" {
					errMsg := fmt.Sprintf("ferment is %v", ferment)
					return sdkerrors.Wrap(types.ErrFermentOtherPropertiesNameEmpty, errMsg)
				}
				if ferment.OtherProperties[i].Quantity == "" {
					errMsg := fmt.Sprintf("ferment is %v", ferment)
					return sdkerrors.Wrap(types.ErrFermentOtherPropertiesQuantityEmpty, errMsg)
				}
				if ferment.OtherProperties[i].WhenToAdd == "" {
					errMsg := fmt.Sprintf("ferment is %v", ferment)
					return sdkerrors.Wrap(types.ErrFermentOtherPropertiesWhenToAddEmpty, errMsg)
				}
			}
			if ages.BestBefore == "" {
				errMsg := fmt.Sprintf("ages is %v", ages)
				return sdkerrors.Wrap(types.ErrAgesBestBeforeEmpty, errMsg)
			}
			if ages.ExpiredDate == "" {
				errMsg := fmt.Sprintf("ages is %v", ages)
				return sdkerrors.Wrap(types.ErrAgesBestExpiredDateEmpty, errMsg)
			}
			if ages.BottledDate == "" {
				errMsg := fmt.Sprintf("ages is %v", ages)
				return sdkerrors.Wrap(types.ErrAgesBottledDateEmpty, errMsg)
			}
			if ages.BestAfter == "" {
				errMsg := fmt.Sprintf("ages is %v", ages)
				return sdkerrors.Wrap(types.ErrAgesBestAfterEmpty, errMsg)
			}
			/*
				cliCtx.PrintOutput(properties)
				cliCtx.PrintOutput(mashing)
				cliCtx.PrintOutput(boil)
				cliCtx.PrintOutput(ferment)
				cliCtx.PrintOutput(ages)
			*/

			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))
			msg := types.NewMsgCreateRecipes(cliCtx.GetFromAddress(), BrewerID, BeerID, BatchNo, properties, mashing, boil, ferment, ages, Factory, ImageLabel)
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
			// fmt.Printf("Address %v\n", toAddress)
			msg := types.NewMsgCreateWithdrawCoinsFromModuleWallet(toAddress, coins)
			// fmt.Printf("msg %v\n", msg)
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}
