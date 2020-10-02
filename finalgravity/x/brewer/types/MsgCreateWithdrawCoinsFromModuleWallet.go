package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/google/uuid"
	"github.com/tendermint/tendermint/crypto"
)

var _ sdk.Msg = &MsgCreateWithdrawCoinsFromModuleWallet{}

//MsgCreateWithdrawCoinsFromModuleWallet struct
type MsgCreateWithdrawCoinsFromModuleWallet struct {
	WithdrawID string         `json:"WithdrawID" yaml:"WithdrawID"`
	Sender     sdk.AccAddress `json:"Sender" yaml:"Sender"`
	Reciever   sdk.AccAddress `json:"Reciever" yaml:"Reciever"`
	Amount     sdk.Coins      `json:"Amount" yaml:"Amount"`
}

//NewMsgCreateWithdrawCoinsFromModuleWallet function
func NewMsgCreateWithdrawCoinsFromModuleWallet(recieverAddress sdk.AccAddress, amount sdk.Coins) MsgCreateWithdrawCoinsFromModuleWallet {
	return MsgCreateWithdrawCoinsFromModuleWallet{
		WithdrawID: uuid.New().String(),
		Sender:     sdk.AccAddress(crypto.AddressHash([]byte(ModuleName))),
		Reciever:   recieverAddress,
		Amount:     amount,
	}
}

//CreateWithdraw const
const CreateWithdraw = "CreateWithdraw"

//Route function
func (msg MsgCreateWithdrawCoinsFromModuleWallet) Route() string {
	return RouterKey
}

//Type function
func (msg MsgCreateWithdrawCoinsFromModuleWallet) Type() string {
	return CreateWithdraw
}

//GetSigners function
func (msg MsgCreateWithdrawCoinsFromModuleWallet) GetSigners() []sdk.AccAddress {
	fmt.Printf("msg.Reciever %v\n", msg.Reciever)
	return []sdk.AccAddress{msg.Reciever}
}

//GetSignBytes function
func (msg MsgCreateWithdrawCoinsFromModuleWallet) GetSignBytes() []byte {
	fmt.Printf("msg %v\n", msg)
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

//ValidateBasic function
func (msg MsgCreateWithdrawCoinsFromModuleWallet) ValidateBasic() error {
	if msg.Sender.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "Sender address can't be empty")
	}
	return nil
}
