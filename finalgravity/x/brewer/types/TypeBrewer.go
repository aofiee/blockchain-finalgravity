package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

//Brewer struct
type Brewer struct {
	Creator      sdk.AccAddress `json:"creator" yaml:"creator"`
	BrewerID     string         `json:"BrewerID" yaml:"BrewerID"`
	TypeOfBrewer string         `json:"TypeOfBrewer" yaml:"TypeOfBrewer"`
	Address      string         `json:"Address" yaml:"Address"`
	Telephone    string         `json:"Telephone" yaml:"Telephone"`
	Email        string         `json:"Email" yaml:"Email"`
	Story        string         `json:"Story" yaml:"Story"`
	LogoURL      string         `json:"LogoURL" yaml:"LogoURL"`
	CoverURL     string         `json:"CoverURL" yaml:"CoverURL"`
	Founded      string         `json:"Founded" yaml:"Founded"`
	Founder      string         `json:"Founder" yaml:"Founder"`
	SiteURL      string         `json:"SiteURL" yaml:"SiteURL"`
}

//BrewerWallet struct
type BrewerWallet struct {
	Creator sdk.AccAddress `json:"creator" yaml:"creator"`
	Amount  sdk.Coins      `json:"Amount" yaml:"Amount"`
}

//WithdrawCoinsFromModuleWallet struct
type WithdrawCoinsFromModuleWallet struct {
	WithdrawID string         `json:"WithdrawID" yaml:"WithdrawID"`
	Sender     sdk.AccAddress `json:"Sender" yaml:"Sender"`
	Reciever   sdk.AccAddress `json:"Reciever" yaml:"Reciever"`
	Amount     sdk.Coins      `json:"Amount" yaml:"Amount"`
}
