package financiamento

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "github.com/gabrielscota/defihub/api/defihub/financiamento"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: modulev1.Query_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "Params",
					Use:       "params",
					Short:     "Shows the parameters of the module",
				},
				// this line is used by ignite scaffolding # autocli/query
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service:              modulev1.Msg_ServiceDesc.ServiceName,
			EnhanceCustomCommand: true, // only required if you want to use the custom command
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "UpdateParams",
					Skip:      true, // skipped because authority gated
				},
				{
					RpcMethod:      "SolicitarEmprestimo",
					Use:            "solicitar-emprestimo [valor] [prazo] [taxa]",
					Short:          "Send a SolicitarEmprestimo tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "valor"}, {ProtoField: "prazo"}, {ProtoField: "taxa"}},
				},
				{
					RpcMethod:      "AprovarEmprestimo",
					Use:            "aprovar-emprestimo [emprestimo-id]",
					Short:          "Send a AprovarEmprestimo tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "emprestimoId"}},
				},
				{
					RpcMethod:      "PagarEmprestimo",
					Use:            "pagar-emprestimo [emprestimo-id] [valor]",
					Short:          "Send a PagarEmprestimo tx",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "emprestimoId"}, {ProtoField: "valor"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
