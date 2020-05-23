package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// TODO: Fill out some custom errors for the module
// You can see how they are constructed below:
var (
	ErrDuplicateError = sdkerrors.Register(ModuleName, 1, "Duplicate error")
)
