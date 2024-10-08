package financiamento_test

import (
	"testing"

	keepertest "github.com/gabrielscota/defihub/testutil/keeper"
	"github.com/gabrielscota/defihub/testutil/nullify"
	financiamento "github.com/gabrielscota/defihub/x/financiamento/module"
	"github.com/gabrielscota/defihub/x/financiamento/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.FinanciamentoKeeper(t)
	financiamento.InitGenesis(ctx, k, genesisState)
	got := financiamento.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
