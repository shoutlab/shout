package keeper

import (
	"github.com/shoutlab/shout/x/shout/types"
)

var _ types.QueryServer = Keeper{}
