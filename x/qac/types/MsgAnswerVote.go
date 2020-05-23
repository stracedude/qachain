package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

type MsgAnswerVote struct {
	Answer string         `json:"answer"`
	Voter  sdk.AccAddress `json:"voter"`
}

func NewMsgAnswerVote(answer string, voter sdk.AccAddress) MsgAnswerVote {
	return MsgAnswerVote{
		Answer: answer,
		Voter:  voter,
	}
}

// Route should return the name of the module
func (msg MsgAnswerVote) Route() string { return RouterKey }

// Type should return the action
func (msg MsgAnswerVote) Type() string { return "AnswerVote" }

func (msg MsgAnswerVote) ValidateBasic() error {
	if msg.Voter.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Voter.String())
	}
	if len(msg.Answer) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "Description and/or Reward cannot be empty")
	}
	return nil
}

func (msg MsgAnswerVote) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners defines whose signature is required
func (msg MsgAnswerVote) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Voter}
}
