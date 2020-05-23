package qac

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stracedude/qac/x/qac/types"
)

// NewHandler creates an sdk.Handler for all the qac type messages
func NewHandler(k Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())
		switch msg := msg.(type) {
		case MsgCreateQuestion:
			return handleMsgCreateQuestion(ctx, k, msg)
		case MsgCreateAnswer:
			return handleCreateAnswer(ctx, k, msg)
		case MsgAnswerVote:
			return handleAnswerVote(ctx, k, msg)
		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", ModuleName, msg)
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}

func handleMsgCreateQuestion(ctx sdk.Context, k Keeper, msg MsgCreateQuestion) (*sdk.Result, error) {

	var question = types.Question{
		Id:          msg.Id,
		Creator:     msg.Creator,
		Description: msg.Description,
		Reward:      msg.Reward,
	}

	_, err := k.GetQuestion(ctx, question.Id)
	if err == nil {
		return nil, sdkerrors.Wrap(types.ErrDuplicateError, "Question hash already exists")
	}

	k.SetQuestion(ctx, question)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeyAction, types.EventTypeCreateQuestion),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Creator.String()),
			sdk.NewAttribute(types.AttributeDescription, msg.Description),
			sdk.NewAttribute(types.AttributeQuestionHash, msg.Id),
			sdk.NewAttribute(types.AttributeReward, msg.Reward.String()),
		),
	)
	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

func handleCreateAnswer(ctx sdk.Context, k Keeper, msg MsgCreateAnswer) (*sdk.Result, error) {
	var answer = types.Answer{
		Id:         msg.Id,
		Responder:  msg.Responder,
		Answer:     msg.Answer,
		QuestionId: msg.QuestionHash,
	}

	_, err := k.GetAnswer(ctx, answer.Id)
	if err == nil {
		return nil, sdkerrors.Wrap(types.ErrDuplicateError, "Answer already exists")
	}

	k.SetAnswer(ctx, answer)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeyAction, types.EventTypeCreateAnswer),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Responder.String()),
			sdk.NewAttribute(types.AttributeAnswerHash, answer.Id),
			sdk.NewAttribute(types.AttributeQuestionHash, msg.QuestionHash),
		),
	)
	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}

func handleAnswerVote(ctx sdk.Context, k Keeper, msg MsgAnswerVote) (*sdk.Result, error) {
	count, err := k.AddAnswerVote(ctx, msg.Answer, msg.Voter)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "Answer not found")
	}

	if count == types.AnswerConfirmation {
		answer, _ := k.GetAnswer(ctx, msg.Answer)
		question , _ := k.GetQuestion(ctx, answer.QuestionId)
		sdkError := k.CoinKeeper.SendCoins(ctx,  question.Creator, answer.Responder, question.Reward)
		if sdkError != nil {
			return nil, sdkError
		}

		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				sdk.EventTypeMessage,
				sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
				sdk.NewAttribute(sdk.AttributeKeyAction, types.EventTypeAnswerReward),
				sdk.NewAttribute(sdk.AttributeKeySender, question.Creator.String()),
				sdk.NewAttribute(types.AttributeResponder, answer.Responder.String()),
				sdk.NewAttribute(types.AttributeAnswerResponderHash, answer.Id),
				sdk.NewAttribute(types.AttributeQuestionHash, answer.QuestionId),
			),
		)
		return &sdk.Result{Events: ctx.EventManager().Events()}, nil

	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeyAction, types.EventTypeAnswerReward),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Voter.String()),
		),
	)
	return &sdk.Result{Events: ctx.EventManager().Events()}, nil

}