package keeper

import (
	"github.com/collinpowell/checkers/x/checkers/types"
)

var _ types.QueryServer = Keeper{}
