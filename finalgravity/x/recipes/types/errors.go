package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// TODO: Fill out some custom errors for the module
// You can see how they are constructed below:
var (
	ErrInvalid                             = sdkerrors.Register(ModuleName, 1, "custom error message")
	ErrPropertiesColourRatingEmpty         = sdkerrors.Register(ModuleName, 2, "Properties ColourRating is Empty!")
	ErrPropertiesOriginalGravityEmpty      = sdkerrors.Register(ModuleName, 3, "Properties OriginalGravity is Empty!")
	ErrPropertiesExpectedFinalGravityEmpty = sdkerrors.Register(ModuleName, 4, "Properties ExpectedFinalGravity is Empty!")
	ErrPropertiesFinalGravityEmpty         = sdkerrors.Register(ModuleName, 5, "Properties FinalGravity is Empty!")
	ErrPropertiesTotalLiquorEmpty          = sdkerrors.Register(ModuleName, 6, "Properties TotalLiquor is Empty!")
	ErrPropertiesMakesEmpty                = sdkerrors.Register(ModuleName, 7, "Properties Makes is Empty!")
	ErrPropertiesReadyToDrinkEmpty         = sdkerrors.Register(ModuleName, 8, "Properties ReadyToDrink is Empty!")
	ErrPropertiesEstimatedABVEmpty         = sdkerrors.Register(ModuleName, 9, "Properties EstimatedABV is Empty!")
	ErrPropertiesBitternessRatingEmpty     = sdkerrors.Register(ModuleName, 10, "Properties BitternessRating is Empty!")

	ErrMashingLiquorEmpty                      = sdkerrors.Register(ModuleName, 11, "Mashing Liquor is Empty")
	ErrMashingMashTimeEmpty                    = sdkerrors.Register(ModuleName, 12, "Mashing MashTime is Empty")
	ErrMashingTemperatureEmpty                 = sdkerrors.Register(ModuleName, 13, "Mashing Temperature is Empty")
	ErrMashingGrainBillPropertiesEmpty         = sdkerrors.Register(ModuleName, 14, "Mashing GrainBillProperties is Empty")
	ErrMashingGrainBillPropertiesNameEmpty     = sdkerrors.Register(ModuleName, 15, "Mashing GrainBillProperties Name is Empty")
	ErrMashingGrainBillPropertiesQuantityEmpty = sdkerrors.Register(ModuleName, 16, "Mashing GrainBillProperties Quantity is Empty")

	ErrBoilLiquorEmpty                   = sdkerrors.Register(ModuleName, 17, "Boil Liquor is Empty")
	ErrBoilBoilTimeEmpty                 = sdkerrors.Register(ModuleName, 18, "Boil BoilTime is Empty")
	ErrBoilHopsPropertiesEmpty           = sdkerrors.Register(ModuleName, 19, "Boil HopsProperties is Empty")
	ErrBoilHopsPropertiesNameEmpty       = sdkerrors.Register(ModuleName, 20, "Boil HopsProperties Name is Empty")
	ErrBoilHopsPropertiesQuantityEmpty   = sdkerrors.Register(ModuleName, 21, "Boil HopsProperties Quantity is Empty")
	ErrBoilHopsPropertiesIBUEmpty        = sdkerrors.Register(ModuleName, 22, "Boil HopsProperties IBU is Empty")
	ErrBoilHopsPropertiesWhenToAddEmpty  = sdkerrors.Register(ModuleName, 23, "Boil HopsProperties WhenToAdd is Empty")
	ErrBoilOtherPropertiesEmpty          = sdkerrors.Register(ModuleName, 24, "Boil OtherProperties is Empty")
	ErrBoilOtherPropertiesNameEmpty      = sdkerrors.Register(ModuleName, 25, "Boil OtherProperties Name is Empty")
	ErrBoilOtherPropertiesQuantityEmpty  = sdkerrors.Register(ModuleName, 26, "Boil OtherProperties Quantity is Empty")
	ErrBoilOtherPropertiesWhenToAddEmpty = sdkerrors.Register(ModuleName, 27, "Boil OtherProperties WhenToAdd is Empty")

	ErrFermentYeastPropertiesEmpty          = sdkerrors.Register(ModuleName, 28, "Ferment YeastProperties is Empty")
	ErrFermentYeastPropertiesNameEmpty      = sdkerrors.Register(ModuleName, 29, "Ferment YeastProperties Name is Empty")
	ErrFermentYeastPropertiesQuantityEmpty  = sdkerrors.Register(ModuleName, 30, "Ferment YeastProperties Quantity is Empty")
	ErrFermentFermentationEmpty             = sdkerrors.Register(ModuleName, 31, "Ferment Fermentation is Empty")
	ErrFermentFermentationTemperatureEmpty  = sdkerrors.Register(ModuleName, 32, "Ferment Fermentation Temperature is Empty")
	ErrFermentFermentationConditioningEmpty = sdkerrors.Register(ModuleName, 33, "Ferment Fermentation Conditioning is Empty")
	ErrFermentHopsPropertiesEmpty           = sdkerrors.Register(ModuleName, 34, "Ferment HopsProperties is Empty")
	ErrFermentHopsPropertiesNameEmpty       = sdkerrors.Register(ModuleName, 35, "Ferment HopsProperties Name is Empty")
	ErrFermentHopsPropertiesQuantityEmpty   = sdkerrors.Register(ModuleName, 36, "Ferment HopsProperties Quantity is Empty")
	ErrFermentHopsPropertiesIBUEmpty        = sdkerrors.Register(ModuleName, 37, "Ferment HopsProperties IBU is Empty")
	ErrFermentHopsPropertiesWhenToAddEmpty  = sdkerrors.Register(ModuleName, 38, "Ferment HopsProperties WhenToAdd is Empty")
	ErrFermentOtherPropertiesEmpty          = sdkerrors.Register(ModuleName, 39, "Ferment OtherProperties is Empty")
	ErrFermentOtherPropertiesNameEmpty      = sdkerrors.Register(ModuleName, 40, "Ferment OtherProperties Name is Empty")
	ErrFermentOtherPropertiesQuantityEmpty  = sdkerrors.Register(ModuleName, 41, "Ferment OtherProperties Quantity is Empty")
	ErrFermentOtherPropertiesWhenToAddEmpty = sdkerrors.Register(ModuleName, 42, "Ferment OtherProperties WhenToAdd is Empty")

	ErrAgesBestBeforeEmpty      = sdkerrors.Register(ModuleName, 43, "Ages BestBefore is Empty")
	ErrAgesBestExpiredDateEmpty = sdkerrors.Register(ModuleName, 44, "Ages ExpiredDate is Empty")
	ErrAgesBottledDateEmpty     = sdkerrors.Register(ModuleName, 45, "Ages BottledDate is Empty")
	ErrAgesBestAfterEmpty       = sdkerrors.Register(ModuleName, 46, "Ages BestAfter is Empty")
)
