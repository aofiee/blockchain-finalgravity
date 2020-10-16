package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

//Recipes struct
type Recipes struct {
	Creator    sdk.AccAddress `json:"Creator" yaml:"Creator"`
	BrewerID   string         `json:"BrewerID" yaml:"BrewerID"`
	RecipesID  string         `json:"RecipesID" yaml:"RecipesID"`
	BeerID     string         `json:"BeerID" yaml:"BeerID"`
	BatchNo    string         `json:"BatchNo" yaml:"BatchNo"`
	Properties Properties     `json:"Properties" yaml:"Properties"`
	Mashing    Mashing        `json:"Mashing" yaml:"Mashing"`
	Boil       Boil           `json:"Boil" yaml:"Boil"`
	Ferment    Ferment        `json:"Ferment" yaml:"Ferment"`
	Ages       Ages           `json:"Ages" yaml:"Ages"`
	Factory    string         `json:"Factory" yaml:"Factory"`
	ImageLabel string         `json:"ImageLabel" yaml:"ImageLabel"`
}

//Ages struct
type Ages struct {
	BestBefore  string `json:"BestBefore" yaml:"BestBefore"`
	ExpiredDate string `json:"ExpiredDate" yaml:"ExpiredDate"`
	BottledDate string `json:"BottledDate" yaml:"BottledDate"`
	BestAfter   string `json:"BestAfter" yaml:"BestAfter"`
}

//Ferment struct
type Ferment struct {
	YeastProperties []YeastProperties `json:"YeastProperties" yaml:"YeastProperties"`
	Fermentation    Fermentation      `json:"Fermentation" yaml:"Fermentation"`
	HopsProperties  []HopsProperties  `json:"HopsProperties" yaml:"HopsProperties"`
	OtherProperties []OtherProperties `json:"OtherProperties" yaml:"OtherProperties"`
}

//Fermentation struct
type Fermentation struct {
	Temperature  string `json:"Temperature" yaml:"Temperature"`
	Conditioning string `json:"Conditioning" yaml:"Conditioning"`
}

//YeastProperties struct
type YeastProperties struct {
	Name     string `json:"Name" yaml:"Name"`
	Quantity string `json:"Quantity" yaml:"Quantity"`
}

//Boil struct
type Boil struct {
	Liquor          string            `json:"Liquor" yaml:"Liquor"`
	BoilTime        string            `json:"BoilTime" yaml:"BoilTime"`
	HopsProperties  []HopsProperties  `json:"HopsProperties" yaml:"HopsProperties"`
	OtherProperties []OtherProperties `json:"OtherProperties" yaml:"OtherProperties"`
}

//OtherProperties struct
type OtherProperties struct {
	Name      string `json:"Name" yaml:"Name"`
	Quantity  string `json:"Quantity" yaml:"Quantity"`
	WhenToAdd string `json:"WhenToAdd" yaml:"WhenToAdd"`
}

//HopsProperties struct
type HopsProperties struct {
	Name      string `json:"Name" yaml:"Name"`
	Quantity  string `json:"Quantity" yaml:"Quantity"`
	IBU       string `json:"IBU" yaml:"IBU"`
	WhenToAdd string `json:"WhenToAdd" yaml:"WhenToAdd"`
}

//Mashing struct
type Mashing struct {
	Liquor              string                `json:"Liquor" yaml:"Liquor"`
	MashTime            string                `json:"MashTime" yaml:"MashTime"`
	Temperature         string                `json:"Temperature" yaml:"Temperature"`
	GrainBillProperties []GrainBillProperties `json:"GrainBillProperties" yaml:"GrainBillProperties"`
}

//GrainBillProperties struct
type GrainBillProperties struct {
	Name     string `json:"Name" yaml:"Name"`
	Quantity string `json:"Quantity" yaml:"Quantity"`
}

//Properties struct
type Properties struct {
	OriginalGravity      string `json:"OriginalGravity" yaml:"OriginalGravity"`
	ExpectedFinalGravity string `json:"ExpectedFinalGravity" yaml:"ExpectedFinalGravity"`
	FinalGravity         string `json:"FinalGravity" yaml:"FinalGravity"`
	TotalLiquor          string `json:"TotalLiquor" yaml:"TotalLiquor"`
	Makes                string `json:"Makes" yaml:"Makes"`
	ReadyToDrink         string `json:"ReadyToDrink" yaml:"ReadyToDrink"`
	EstimatedABV         string `json:"EstimatedABV" yaml:"EstimatedABV"`
	BitternessRating     string `json:"BitternessRating" yaml:"BitternessRating"`
	ColourRating         string `json:"ColourRating" yaml:"ColourRating"`
}

//RecipesWallet struct
type RecipesWallet struct {
	Creator sdk.AccAddress `json:"Creator" yaml:"Creator"`
	Amount  sdk.Coins      `json:"Amount" yaml:"Amount"`
}

//WithdrawCoinsFromModuleWallet struct
type WithdrawCoinsFromModuleWallet struct {
	WithdrawID string         `json:"WithdrawID" yaml:"WithdrawID"`
	Sender     sdk.AccAddress `json:"Sender" yaml:"Sender"`
	Reciever   sdk.AccAddress `json:"Reciever" yaml:"Reciever"`
	Amount     sdk.Coins      `json:"Amount" yaml:"Amount"`
}
