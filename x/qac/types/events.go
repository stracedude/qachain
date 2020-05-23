package types

// qac module event types
const (
	EventTypeCreateQuestion = "CreateQuestion"
	EventTypeCreateAnswer   = "CreateAnswer"
	EventTypeAnswerVote     = "AnswerVote"
	EventTypeAnswerReward   = "AnswerReward"

	AttributeDescription         = "description"
	AttributeQuestion            = "question"
	AttributeQuestionHash        = "questionHash"
	AttributeAnswerHash       = "answerHash"
	AttributeReward              = "reward"
	AttributeResponder           = "responder"
	AttributeAnswerResponderHash = "answerHash"

	AttributeValueCategory = ModuleName
)
