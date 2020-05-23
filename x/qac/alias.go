package qac

import (
	"github.com/stracedude/qac/x/qac/keeper"
	"github.com/stracedude/qac/x/qac/types"
)

const (
	// TODO: define constants that you would like exposed from your module

	ModuleName        = types.ModuleName
	RouterKey         = types.RouterKey
	StoreKey          = types.StoreKey
	DefaultParamspace = types.DefaultParamspace
	//QueryParams       = types.QueryParams
	QuerierRoute      = types.QuerierRoute
)

var (
	// functions aliases
	NewKeeper           = keeper.NewKeeper
	NewQuerier          = keeper.NewQuerier
	RegisterCodec       = types.RegisterCodec
	NewGenesisState     = types.NewGenesisState
	DefaultGenesisState = types.DefaultGenesisState
	ValidateGenesis     = types.ValidateGenesis
	// TODO: Fill out function aliases

	// variable aliases
	ModuleCdc     = types.ModuleCdc
	// TODO: Fill out variable aliases
	NewMsgCreateQuestion = types.NewMsgCreateQuestion
	NewMsgCreateAnswer = types.NewMsgCreateAnswer
	NewMsgAnswerVote = types.NewMsgAnswerVote

)

type (
	Keeper       = keeper.Keeper
	GenesisState = types.GenesisState
	Params       = types.Params

	MsgCreateQuestion = types.MsgCreateQuestion
	MsgCreateAnswer = types.MsgCreateAnswer
	MsgAnswerVote = types.MsgAnswerVote

	)
