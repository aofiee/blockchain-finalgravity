package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/google/uuid"
)

var _ sdk.Msg = &MsgCreateBeer{}

//MsgCreateBeer struct
type MsgCreateBeer struct {
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

//NewMsgCreateBeer function
func NewMsgCreateBeer(creator sdk.AccAddress, beername string, typeofbeer string, stylebeer string, tagline string, appearance string, taste string, aroma string, strength string, story string, imagelabel string, createdat string) MsgCreateBeer {
	return MsgCreateBeer{
		Creator:    creator,
		BeerID:     uuid.New().String(),
		BeerName:   beername,
		TypeOfBeer: typeofbeer,
		StyleBeer:  stylebeer,
		TagLine:    tagline,
		Appearance: appearance,
		Taste:      taste,
		Aroma:      aroma,
		Strength:   strength,
		Story:      story,
		ImageLabel: imagelabel,
		CreatedAt:  createdat,
	}
}

//CreateBeerConst const
const CreateBeerConst = "CreateBeer"

//Route function
func (msg MsgCreateBeer) Route() string {
	return RouterKey
}

//Type function
func (msg MsgCreateBeer) Type() string {
	return CreateBeerConst
}

//GetSigners function
func (msg MsgCreateBeer) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

//GetSignBytes function
func (msg MsgCreateBeer) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

//ValidateBasic function
func (msg MsgCreateBeer) ValidateBasic() error {
	if msg.Creator.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
	}
	return nil
}
