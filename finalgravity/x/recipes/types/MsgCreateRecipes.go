package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/google/uuid"
)

var _ sdk.Msg = &MsgCreateRecipes{}

//MsgCreateRecipes struct
type MsgCreateRecipes struct {
	Creator    sdk.AccAddress `json:"Creator" yaml:"Creator"`
	BrewerID   string         `json:"BrewerID" yaml:"BrewerID"`
	RecipesID  string         `json:"RecipesID" yaml:"RecipesID"`
	BeerID     string         `json:"BeerID" yaml:"BeerID"`
	BatchNo    string         `json:"BatchNo" yaml:"BatchNo"`
	Properties MsgProperties  `json:"Properties" yaml:"Properties"`
	Mashing    MsgMashing     `json:"Mashing" yaml:"Mashing"`
	Boil       MsgBoil        `json:"Boil" yaml:"Boil"`
	Ferment    MsgFerment     `json:"Ferment" yaml:"Ferment"`
	Ages       MsgAges        `json:"Ages" yaml:"Ages"`
	Factory    string         `json:"Factory" yaml:"Factory"`
	ImageLabel string         `json:"ImageLabel" yaml:"ImageLabel"`
}

//MsgAges struct
type MsgAges struct {
	BestBefore  string `json:"BestBefore" yaml:"BestBefore"`
	ExpiredDate string `json:"ExpiredDate" yaml:"ExpiredDate"`
	BottledDate string `json:"BottledDate" yaml:"BottledDate"`
	BestAfter   string `json:"BestAfter" yaml:"BestAfter"`
}

//MsgFerment struct
type MsgFerment struct {
	YeastProperties []MsgYeastProperties `json:"YeastProperties" yaml:"YeastProperties"`
	Fermentation    MsgFermentation      `json:"Fermentation" yaml:"Fermentation"`
	HopsProperties  []MsgHopsProperties  `json:"HopsProperties" yaml:"HopsProperties"`
	OtherProperties []MsgOtherProperties `json:"OtherProperties" yaml:"OtherProperties"`
}

//MsgFermentation struct
type MsgFermentation struct {
	Temperature  string `json:"Temperature" yaml:"Temperature"`
	Conditioning string `json:"Conditioning" yaml:"Conditioning"`
}

//MsgYeastProperties struct
type MsgYeastProperties struct {
	Name     string `json:"Name" yaml:"Name"`
	Quantity string `json:"Quantity" yaml:"Quantity"`
}

//MsgBoil struct
type MsgBoil struct {
	Liquor          string               `json:"Liquor" yaml:"Liquor"`
	BoilTime        string               `json:"BoilTime" yaml:"BoilTime"`
	HopsProperties  []MsgHopsProperties  `json:"HopsProperties" yaml:"HopsProperties"`
	OtherProperties []MsgOtherProperties `json:"OtherProperties" yaml:"OtherProperties"`
}

//MsgOtherProperties struct
type MsgOtherProperties struct {
	Name      string `json:"Name" yaml:"Name"`
	Quantity  string `json:"Quantity" yaml:"Quantity"`
	WhenToAdd string `json:"WhenToAdd" yaml:"WhenToAdd"`
}

//MsgHopsProperties struct
type MsgHopsProperties struct {
	Name      string `json:"Name" yaml:"Name"`
	Quantity  string `json:"Quantity" yaml:"Quantity"`
	IBU       string `json:"IBU" yaml:"IBU"`
	WhenToAdd string `json:"WhenToAdd" yaml:"WhenToAdd"`
}

//MsgMashing struct
type MsgMashing struct {
	Liquor              string                   `json:"Liquor" yaml:"Liquor"`
	MashTime            string                   `json:"MashTime" yaml:"MashTime"`
	Temperature         string                   `json:"Temperature" yaml:"Temperature"`
	GrainBillProperties []MsgGrainBillProperties `json:"GrainBillProperties" yaml:"GrainBillProperties"`
}

//MsgGrainBillProperties struct
type MsgGrainBillProperties struct {
	Name     string `json:"Name" yaml:"Name"`
	Quantity string `json:"Quantity" yaml:"Quantity"`
}

//MsgProperties struct
type MsgProperties struct {
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

//NewMsgCreateRecipes function
func NewMsgCreateRecipes(creator sdk.AccAddress, brewerID string, beerID string, batchNo string, properties MsgProperties, mashing MsgMashing, boil MsgBoil, ferment MsgFerment, ages MsgAges, factory string, imageLabel string) MsgCreateRecipes {
	return MsgCreateRecipes{
		Creator:    creator,
		BrewerID:   brewerID,
		RecipesID:  uuid.New().String(),
		BeerID:     beerID,
		BatchNo:    batchNo,
		Properties: properties,
		Mashing:    mashing,
		Boil:       boil,
		Ferment:    ferment,
		Ages:       ages,
		Factory:    factory,
		ImageLabel: imageLabel,
	}
}

//CreateReciepesConst const
const CreateReciepesConst = "CreateRecipes"

//Route function
func (msg MsgCreateRecipes) Route() string {
	return RouterKey
}

//Type function
func (msg MsgCreateRecipes) Type() string {
	return CreateReciepesConst
}

//GetSigners function
func (msg MsgCreateRecipes) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

//GetSignBytes function
func (msg MsgCreateRecipes) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

//ValidateBasic function
func (msg MsgCreateRecipes) ValidateBasic() error {
	if msg.Creator.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
	}
	return nil
}
