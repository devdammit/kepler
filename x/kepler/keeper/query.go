package keeper

import (
	"github.com/devdammit/kepler/x/kepler/types"
)

var _ types.QueryServer = Keeper{}
