package financiamento

import (
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/gabrielscota/defihub/testutil/sample"
	financiamentosimulation "github.com/gabrielscota/defihub/x/financiamento/simulation"
	"github.com/gabrielscota/defihub/x/financiamento/types"
)

// avoid unused import issue
var (
	_ = financiamentosimulation.FindAccount
	_ = rand.Rand{}
	_ = sample.AccAddress
	_ = sdk.AccAddress{}
	_ = simulation.MsgEntryKind
)

const (
	opWeightMsgSolicitarEmprestimo = "op_weight_msg_solicitar_emprestimo"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSolicitarEmprestimo int = 100

	opWeightMsgAprovarEmprestimo = "op_weight_msg_aprovar_emprestimo"
	// TODO: Determine the simulation weight value
	defaultWeightMsgAprovarEmprestimo int = 100

	opWeightMsgPagarEmprestimo = "op_weight_msg_pagar_emprestimo"
	// TODO: Determine the simulation weight value
	defaultWeightMsgPagarEmprestimo int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	financiamentoGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&financiamentoGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ simtypes.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgSolicitarEmprestimo int
	simState.AppParams.GetOrGenerate(opWeightMsgSolicitarEmprestimo, &weightMsgSolicitarEmprestimo, nil,
		func(_ *rand.Rand) {
			weightMsgSolicitarEmprestimo = defaultWeightMsgSolicitarEmprestimo
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSolicitarEmprestimo,
		financiamentosimulation.SimulateMsgSolicitarEmprestimo(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgAprovarEmprestimo int
	simState.AppParams.GetOrGenerate(opWeightMsgAprovarEmprestimo, &weightMsgAprovarEmprestimo, nil,
		func(_ *rand.Rand) {
			weightMsgAprovarEmprestimo = defaultWeightMsgAprovarEmprestimo
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgAprovarEmprestimo,
		financiamentosimulation.SimulateMsgAprovarEmprestimo(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgPagarEmprestimo int
	simState.AppParams.GetOrGenerate(opWeightMsgPagarEmprestimo, &weightMsgPagarEmprestimo, nil,
		func(_ *rand.Rand) {
			weightMsgPagarEmprestimo = defaultWeightMsgPagarEmprestimo
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgPagarEmprestimo,
		financiamentosimulation.SimulateMsgPagarEmprestimo(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgSolicitarEmprestimo,
			defaultWeightMsgSolicitarEmprestimo,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				financiamentosimulation.SimulateMsgSolicitarEmprestimo(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgAprovarEmprestimo,
			defaultWeightMsgAprovarEmprestimo,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				financiamentosimulation.SimulateMsgAprovarEmprestimo(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgPagarEmprestimo,
			defaultWeightMsgPagarEmprestimo,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				financiamentosimulation.SimulateMsgPagarEmprestimo(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
