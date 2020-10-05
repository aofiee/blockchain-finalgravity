package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

//TypeOfBeer string
type TypeOfBeer string

//const available value for enum
const (
	Ale    TypeOfBeer = "Ale"
	Lager  TypeOfBeer = "Lager"
	Lambic TypeOfBeer = "Lambic"
)

//Beer struct
type Beer struct {
	Creator    sdk.AccAddress `json:"Creator" yaml:"Creator"`
	BeerID     string         `json:"BeerID" yaml:"BeerID"`
	BeerName   string         `json:"BeerName" yaml:"BeerName"`
	TypeOfBeer string         `json:"TypeOfBeer" yaml:"TypeOfBeer"`
	StyleBeer  string         `json:"StyleBeer" yaml:"StyleBeer"`
	TagLine    string         `json:"TagLine" yaml:"TagLine"`
	Appearance string         `json:"Appearance" yaml:"Appearance"`
	Taste      string         `json:"Taste" yaml:"Taste"`
	Aroma      string         `json:"Aroma" yaml:"Aroma"`
	Strength   string         `json:"Strength" yaml:"Strength"`
	Story      string         `json:"Story" yaml:"Story"`
	ImageLabel string         `json:"ImageLabel" yaml:"ImageLabel"`
	CreatedAt  string         `json:"CreatedAt" yaml:"CreatedAt"`
}

//BeerWallet struct
type BeerWallet struct {
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
