package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/google/uuid"
)

// TODO: Describe your actions, these will implment the interface of `sdk.Msg`

// verify interface at compile time
var _ sdk.Msg = &MsgCreateBrewer{}

// MsgCreateBrewer - struct for unjailing jailed validator
type MsgCreateBrewer struct {
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

// NewMsgCreateBrewer creates a new MsgCreateBrewer instance
func NewMsgCreateBrewer(creator sdk.AccAddress, typebrewer string, addr string, tel string, email string, story string, logoURL string, coverURL string, founded string, founder string, siteURL string) MsgCreateBrewer {
	return MsgCreateBrewer{
		Creator:      creator,
		BrewerID:     uuid.New().String(),
		TypeOfBrewer: typebrewer,
		Address:      addr,
		Telephone:    tel,
		Email:        email,
		Story:        story,
		LogoURL:      logoURL,
		CoverURL:     coverURL,
		Founded:      founded,
		Founder:      founder,
		SiteURL:      siteURL,
	}
}

//CreateBrewerConst const
const CreateBrewerConst = "CreateBrewer"

// Route should return the name of the module
func (msg MsgCreateBrewer) Route() string { return RouterKey }

// Type should return the action
func (msg MsgCreateBrewer) Type() string { return CreateBrewerConst }

// GetSigners should return the name of the module
func (msg MsgCreateBrewer) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

// GetSignBytes gets the bytes for the message signer to sign on
func (msg MsgCreateBrewer) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic validity check for the AnteHandler
func (msg MsgCreateBrewer) ValidateBasic() error {
	if msg.Creator.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing validator address")
	}
	return nil
}
