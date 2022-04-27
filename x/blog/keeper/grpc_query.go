package keeper

import (
	"github.com/carmonasl/blog/x/blog/types"
)

var _ types.QueryServer = Keeper{}
