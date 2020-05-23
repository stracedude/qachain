package types

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"strings"
)

var MinQuestionReward = sdk.Coins{sdk.NewInt64Coin("qactoken", 1)}

type Question struct {
	Id          string         `json:"id"`
	Creator     sdk.AccAddress `json:"creator"`
	Description string         `json:"description"`
	Answer      string         `json:"answer"`
	Reward      sdk.Coins      `json:"reward"`
}

func NewQuestion() Question {
	return Question{
		Reward: MinQuestionReward,
	}
}

func (q Question) String() string {
	return strings.TrimSpace(fmt.Sprintf(`Id: %s Creator: %s
Description: %s
Reward: %s Answer: %s`, q.Id, q.Creator, q.Description, q.Reward, q.Answer))
}

type Answer struct {
	Id         string           `json:"id"`
	QuestionId string           `json:"qid"`
	Responder  sdk.AccAddress   `json:"responder"`
	Answer     string           `json:"answer"`
	Voted      []sdk.AccAddress `json:"voted"`
}

func (a Answer) String() string {
	return strings.TrimSpace(fmt.Sprintf(`Id: %s
QuestionId: %s
Responder: %s Answer: %s Rating: %c`, a.Id, a.QuestionId, a.Responder, a.Answer, len(a.Voted)))
}
