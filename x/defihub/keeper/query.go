package keeper

import (
	"github.com/gabrielscota/defihub/x/defihub/types"
)

var _ types.QueryServer = Keeper{}
