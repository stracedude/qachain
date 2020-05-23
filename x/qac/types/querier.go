package types

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"strings"
)

// Query endpoints supported by the qac querier
const (
	QueryListQuestion = "list"
	QueryGetQuestion  = "get"
	QueryGetOwnQuestion  = "getOwnQuestions"
	QueryGetOwnAnswers  = "getOwnAnswers"
)

type QuestionsList []Question

// implement fmt.Stringer
func (list QuestionsList) String() string {
	var outString string
	for i := range list {
		outString = outString + "\n" + list[i].String()
	}
	return outString
}

type QuestionsView struct {
	Id          string         `json:"id"`
	Creator     sdk.AccAddress `json:"creator"`
	Description string         `json:"description"`
	Reward      sdk.Coins      `json:"reward"`
	AnswerList  []Answer       `json:"answers"`
}

func (v QuestionsView) String() string {
	return strings.TrimSpace(fmt.Sprintf(`Id: %s
QuestionId: %s \n
Responder: %s \n Reward: %s \n \n %c \n`, v.Id, v.Creator, v.Description, v.Reward, len(v.AnswerList)))
}