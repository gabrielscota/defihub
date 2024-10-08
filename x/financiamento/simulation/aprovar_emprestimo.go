package simulation

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/gabrielscota/defihub/x/financiamento/keeper"
	"github.com/gabrielscota/defihub/x/financiamento/types"
)

func SimulateMsgAprovarEmprestimo(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgAprovarEmprestimo{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the AprovarEmprestimo simulation

		return simtypes.NoOpMsg(types.ModuleName, sdk.MsgTypeURL(msg), "AprovarEmprestimo simulation not implemented"), nil, nil
	}
}
