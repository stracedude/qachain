package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stracedude/qac/x/qac/types"
)

// Keeper of the qac store
type Keeper struct {
	CoinKeeper types.BankKeeper

	storeKey sdk.StoreKey
	cdc      *codec.Codec
}

// NewKeeper creates a qac keeper
func NewKeeper(cdc *codec.Codec, key sdk.StoreKey, coinKeeper types.BankKeeper) Keeper {
	keeper := Keeper{
		storeKey:   key,
		cdc:        cdc,
		CoinKeeper: coinKeeper,
	}
	return keeper
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) GetQuestion(ctx sdk.Context, questionHash string) (types.Question, error) {
	store := ctx.KVStore(k.storeKey)
	var question types.Question
	byteKey := []byte(types.QuestionPrefix + questionHash)
	err := k.cdc.UnmarshalBinaryLengthPrefixed(store.Get(byteKey), &question)
	if err != nil {
		return question, err
	}
	return question, nil
}

func (k Keeper) SetQuestion(ctx sdk.Context, question types.Question) {
	questionHash := question.Id
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(question)
	key := []byte(types.QuestionPrefix + questionHash)
	store.Set(key, bz)
}

func (k Keeper) GetAnswer(ctx sdk.Context, questionHash string) (types.Answer, error) {
	store := ctx.KVStore(k.storeKey)
	var answer types.Answer
	byteKey := []byte(types.AnswerPrefix + questionHash)
	err := k.cdc.UnmarshalBinaryLengthPrefixed(store.Get(byteKey), &answer)
	if err != nil {
		return answer, err
	}
	return answer, nil
}

func (k Keeper) SetAnswer(ctx sdk.Context, answer types.Answer) {
	answerHash := answer.Id
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(answer)
	key := []byte(types.AnswerPrefix + answerHash)
	store.Set(key, bz)
}

func (k Keeper) AddAnswerVote(ctx sdk.Context, answerHash string, vote sdk.AccAddress) (int, error) {
	answer, err := k.GetAnswer(ctx, answerHash)
	if err != nil {
		return -1, err
	}

	answer.Voted = append(answer.Voted, vote)
	k.SetAnswer(ctx, answer)
	return len(answer.Voted), nil
}

func (k Keeper) GetQuestionIterator(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return sdk.KVStorePrefixIterator(store, []byte(types.QuestionPrefix))
}

func (k Keeper) GetAnswerIterator(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return sdk.KVStorePrefixIterator(store, []byte(types.AnswerPrefix))
}