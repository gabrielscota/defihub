package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/gabrielscota/defihub/testutil/keeper"
	"github.com/gabrielscota/defihub/x/financiamento/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.FinanciamentoKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
