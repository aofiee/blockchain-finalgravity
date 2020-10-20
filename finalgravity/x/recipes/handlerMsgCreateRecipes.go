package recipes

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	// sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/aofiee/finalgravity/x/recipes/keeper"
	"github.com/aofiee/finalgravity/x/recipes/types"
	"github.com/tendermint/tendermint/crypto"
)

func handleMsgCreateRecipes(ctx sdk.Context, k keeper.Keeper, msg types.MsgCreateRecipes) (*sdk.Result, error) {
	var properties = types.Properties{
		OriginalGravity:      msg.Properties.OriginalGravity,
		ExpectedFinalGravity: msg.Properties.ExpectedFinalGravity,
		FinalGravity:         msg.Properties.FinalGravity,
		TotalLiquor:          msg.Properties.TotalLiquor,
		Makes:                msg.Properties.Makes,
		ReadyToDrink:         msg.Properties.ReadyToDrink,
		EstimatedABV:         msg.Properties.EstimatedABV,
		BitternessRating:     msg.Properties.BitternessRating,
		ColourRating:         msg.Properties.ColourRating,
	}
	var GrainBillPropertiesList []types.GrainBillProperties
	for i := range msg.Mashing.GrainBillProperties {
		var grain = types.GrainBillProperties{
			Name:     msg.Mashing.GrainBillProperties[i].Name,
			Quantity: msg.Mashing.GrainBillProperties[i].Quantity,
		}
		GrainBillPropertiesList = append(GrainBillPropertiesList, grain)
	}
	var mashing = types.Mashing{
		Liquor:              msg.Mashing.Liquor,
		MashTime:            msg.Mashing.MashTime,
		Temperature:         msg.Mashing.Temperature,
		GrainBillProperties: GrainBillPropertiesList,
	}
	var boilHops []types.HopsProperties
	for i := range msg.Boil.HopsProperties {
		var hops = types.HopsProperties{
			Name:      msg.Boil.HopsProperties[i].Name,
			Quantity:  msg.Boil.HopsProperties[i].Quantity,
			IBU:       msg.Boil.HopsProperties[i].IBU,
			WhenToAdd: msg.Boil.HopsProperties[i].WhenToAdd,
		}
		boilHops = append(boilHops, hops)
	}
	var boilOthers []types.OtherProperties
	for i := range msg.Boil.OtherProperties {
		var others = types.OtherProperties{
			Name:      msg.Boil.OtherProperties[i].Name,
			Quantity:  msg.Boil.OtherProperties[i].Quantity,
			WhenToAdd: msg.Boil.OtherProperties[i].WhenToAdd,
		}
		boilOthers = append(boilOthers, others)
	}
	var boil = types.Boil{
		Liquor:          msg.Boil.Liquor,
		BoilTime:        msg.Boil.BoilTime,
		HopsProperties:  boilHops,
		OtherProperties: boilOthers,
	}
	var Fermentation = types.Fermentation{
		Temperature:  msg.Ferment.Fermentation.Temperature,
		Conditioning: msg.Ferment.Fermentation.Conditioning,
	}
	var yeastList []types.YeastProperties
	for i := range msg.Ferment.YeastProperties {
		var yeast = types.YeastProperties{
			Name:     msg.Ferment.YeastProperties[i].Name,
			Quantity: msg.Ferment.YeastProperties[i].Quantity,
		}
		yeastList = append(yeastList, yeast)
	}
	var fermentHops []types.HopsProperties
	for i := range msg.Ferment.HopsProperties {
		var hops = types.HopsProperties{
			Name:      msg.Ferment.HopsProperties[i].Name,
			Quantity:  msg.Ferment.HopsProperties[i].Quantity,
			IBU:       msg.Ferment.HopsProperties[i].IBU,
			WhenToAdd: msg.Ferment.HopsProperties[i].WhenToAdd,
		}
		fermentHops = append(fermentHops, hops)
	}
	var fermentOther []types.OtherProperties
	for i := range msg.Ferment.OtherProperties {
		var others = types.OtherProperties{
			Name:      msg.Ferment.OtherProperties[i].Name,
			Quantity:  msg.Ferment.OtherProperties[i].Quantity,
			WhenToAdd: msg.Ferment.OtherProperties[i].WhenToAdd,
		}
		fermentOther = append(fermentOther, others)
	}
	var ferment = types.Ferment{
		YeastProperties: yeastList,
		Fermentation:    Fermentation,
		HopsProperties:  fermentHops,
		OtherProperties: fermentOther,
	}
	var ages = types.Ages{
		BestBefore:  msg.Ages.BestBefore,
		ExpiredDate: msg.Ages.ExpiredDate,
		BottledDate: msg.Ages.BottledDate,
		BestAfter:   msg.Ages.BestAfter,
	}
	var recipes = types.Recipes{
		Creator:    msg.Creator,
		BrewerID:   msg.BrewerID,
		RecipesID:  msg.RecipesID,
		BeerID:     msg.BeerID,
		BatchNo:    msg.BatchNo,
		Properties: properties,
		Mashing:    mashing,
		Boil:       boil,
		Ferment:    ferment,
		Ages:       ages,
		Factory:    msg.Factory,
		ImageLabel: msg.ImageLabel,
	}
	k.CreateRecipes(ctx, recipes)
	moduleAcct := sdk.AccAddress(crypto.AddressHash([]byte(types.ModuleName)))
	// totalCoin := k.CoinKeeper.GetCoins(ctx, moduleAcct)
	payment, _ := sdk.ParseCoins("200rune")
	if err := k.CoinKeeper.SendCoins(ctx, msg.Creator, moduleAcct, payment); err != nil {
		return nil, err
	}
	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
