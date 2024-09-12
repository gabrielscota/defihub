package defihub_test

import (
	"testing"

	keepertest "github.com/gabrielscota/defihub/testutil/keeper"
	"github.com/gabrielscota/defihub/testutil/nullify"
	defihub "github.com/gabrielscota/defihub/x/defihub/module"
	"github.com/gabrielscota/defihub/x/defihub/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.DefihubKeeper(t)
	defihub.InitGenesis(ctx, k, genesisState)
	got := defihub.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
