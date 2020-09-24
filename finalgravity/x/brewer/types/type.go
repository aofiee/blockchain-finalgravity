package types

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/google/uuid"
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

//NewBrewer func
func NewBrewer() Brewer {
	return Brewer{
		BrewerID:     uuid.New().String(),
		TypeOfBrewer: "Home Brew",
		Address:      "44/261 Passorn Onut Prawet Prawet Bangkok 10250 Thailand.",
		Telephone:    "+6692 590 5444",
		Email:        "aofiee666@gmail.com",
		Story:        "Punk IPA is the beer that kick-started it. This light, golden classic has been subverted with new world hops to create an explosion of flavour. Bursts of caramel and tropical fruit with an all-out riot of grapefruit, pineapple and lychee, precede a spiky bitter finish.This is the beer that started it all - and it’s not done yet...",
		LogoURL:      "https://www.brewdog.com/static/version1600328468/frontend/Born/arcticFox/en_US/images/logo.svg",
		CoverURL:     "https://www.brewdog.com/static/version1600328468/frontend/Born/arcticFox/en_US/images/logo.svg",
		Founded:      "2020-09-18",
		Founder:      "Khomkrid Lerdprasert",
		SiteURL:      "https://www.aofiee.dev",
	}
}

func (b Brewer) String() string {
	return strings.TrimSpace(fmt.Sprintf(`BrewerID: %s`, b.BrewerID))
}
