package types

import (
	"crypto/sha256"
	"encoding/hex"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

type MsgCreateQuestion struct {
	Id          string         `json:"id"`
	Creator     sdk.AccAddress `json:"creator"`
	Description string         `json:"description"`
	Reward      sdk.Coins      `json:"reward"`
}

func NewMsgCreateQuestion(description string, reward sdk.Coins, creator sdk.AccAddress) MsgCreateQuestion {

	var descriptionHash = sha256.Sum256([]byte(description))
	var descriptionHashString = hex.EncodeToString(descriptionHash[:])

	return MsgCreateQuestion{
		Id: descriptionHashString,
		Description: description,
		Reward:      reward,
		Creator:     creator,
	}
}

const CreateScavengeConst = "CreateQuestion"

// Route should return the name of the module
func (msg MsgCreateQuestion) Route() string { return RouterKey }

// Type should return the action
func (msg MsgCreateQuestion) Type() string { return CreateScavengeConst }

func (msg MsgCreateQuestion) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgCreateQuestion) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg MsgCreateQuestion) ValidateBasic() error {
	if msg.Creator.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Creator.String())
	}
	if len(msg.Description) == 0 || msg.Reward.IsZero() {
		return sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "Description and/or Reward cannot be empty")
	}
	return nil
}
