package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stracedude/qac/x/qac/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

// NewQuerier creates a new querier for qac clients.
func NewQuerier(k Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, error) {
		switch path[0] {
		case types.QueryListQuestion:
			return listQuestion(ctx, k)
		case types.QueryGetQuestion:
			return getQuestion(ctx, path[1:], k)
		case types.QueryGetOwnQuestion:
			return getOwnQuestion(ctx, path[1:], k)

			// TODO: Put the modules query routes
		default:
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "unknown qac query endpoint")
		}
	}
}

func RemovePrefixFromHash(key []byte, prefix []byte) (hash []byte) {
	hash = key[len(prefix):]
	return hash
}

func listQuestion(ctx sdk.Context, k Keeper) ([]byte, error) {
	var questionList types.QuestionsList

	iterator := k.GetQuestionIterator(ctx)

	for ; iterator.Valid(); iterator.Next() {
		questionHash := RemovePrefixFromHash(iterator.Key(), []byte(types.QuestionPrefix))
		question, err := k.GetQuestion(ctx, string(questionHash))
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrPanic, err.Error())
		}
		questionList = append(questionList, question)
	}

	res, err := codec.MarshalJSONIndent(k.cdc, questionList)
	if err != nil {
		return res, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

func getQuestion(ctx sdk.Context, path []string, k Keeper) (res []byte, sdkError error) {
	questionHash := path[0]
	question, err := k.GetQuestion(ctx, questionHash)
	if err != nil {
		return nil, err
	}

	var view = types.QuestionsView{
		Id: question.Id,
		Reward: question.Reward,
		Creator: question.Creator,
		Description: question.Description,
	}

	answerIterator := k.GetAnswerIterator(ctx)

	for ; answerIterator.Valid(); answerIterator.Next() {
		answerHash := RemovePrefixFromHash(answerIterator.Key(), []byte(types.AnswerPrefix))
		answer, err := k.GetAnswer(ctx, string(answerHash))
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrPanic, err.Error())
		}

		if answer.QuestionId == questionHash {
			view.AnswerList = append(view.AnswerList, answer)
		}
	}

	res, err = codec.MarshalJSONIndent(k.cdc, view)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

func getOwnQuestion(ctx sdk.Context, path []string, k Keeper) (res []byte, sdkError error) {
	hash := path[0]
	var questionList types.QuestionsList

	iterator := k.GetQuestionIterator(ctx)

	for ; iterator.Valid(); iterator.Next() {
		questionHash := RemovePrefixFromHash(iterator.Key(), []byte(types.QuestionPrefix))
		question, err := k.GetQuestion(ctx, string(questionHash))
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrPanic, err.Error())
		}

		if question.Creator.String() == hash {
			questionList = append(questionList, question)
		}
	}

	res, err := codec.MarshalJSONIndent(k.cdc, questionList)
	if err != nil {
		return res, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}