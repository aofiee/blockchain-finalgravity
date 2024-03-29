package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/google/uuid"
)

var _ sdk.Msg = &MsgCreateBrewer{}

//MsgCreateBrewer struct
type MsgCreateBrewer struct {
	Creator      sdk.AccAddress `json:"Creator" yaml:"Creator"`
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

//NewMsgCreateBrewer function
func NewMsgCreateBrewer(creator sdk.AccAddress, typeofbrewer string, address string, telephone string, email string, story string, logourl string, coverurl string, founded string, founder string, siteurl string) MsgCreateBrewer {
	return MsgCreateBrewer{
		Creator:      creator,
		BrewerID:     uuid.New().String(),
		TypeOfBrewer: typeofbrewer,
		Address:      address,
		Telephone:    telephone,
		Email:        email,
		Story:        story,
		LogoURL:      logourl,
		CoverURL:     coverurl,
		Founded:      founded,
		Founder:      founder,
		SiteURL:      siteurl,
	}
}

//CreateBrewerConst const
const CreateBrewerConst = "CreateBrewer"

//Route function
func (msg MsgCreateBrewer) Route() string {
	return RouterKey
}

//Type function
func (msg MsgCreateBrewer) Type() string {
	return CreateBrewerConst
}

//GetSigners function
func (msg MsgCreateBrewer) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

//GetSignBytes function
func (msg MsgCreateBrewer) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

//ValidateBasic function
func (msg MsgCreateBrewer) ValidateBasic() error {
	if msg.Creator.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
	}
	return nil
}
