package types

import (
	"crypto/sha256"
	"encoding/hex"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

type MsgCreateAnswer struct {
	Id           string         `json:"id"`
	QuestionHash string         `json:"question"`
	Responder    sdk.AccAddress `json:"responder"`
	Answer       string         `json:"answer"`
}

func NewMsgCreateAnswer(qhash string, answer string, responder sdk.AccAddress) MsgCreateAnswer {
	var answerHash = sha256.Sum256([]byte(answer))
	var answerHashString = hex.EncodeToString(answerHash[:])

	return MsgCreateAnswer{
		Id:           answerHashString,
		QuestionHash: qhash,
		Answer:       answer,
		Responder:    responder,
	}
}

const CreateAnswerConst = "CreateAnswer"

// Route should return the name of the module
func (msg MsgCreateAnswer) Route() string { return RouterKey }

// Type should return the action
func (msg MsgCreateAnswer) Type() string { return CreateAnswerConst }

func (msg MsgCreateAnswer) ValidateBasic() error {
	if msg.Responder.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Responder.String())
	}
	if len(msg.Answer) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "Answer cannot be empty")
	}
	return nil
}

func (msg MsgCreateAnswer) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg MsgCreateAnswer) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Responder)}
}
