package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gabrielscota/defihub/x/financiamento/types"
)

func (k msgServer) SolicitarEmprestimo(goCtx context.Context, msg *types.MsgSolicitarEmprestimo) (*types.MsgSolicitarEmprestimoResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Converter 'taxa' de string para sdk.Dec
	_, err := sdk.NewDecFromStr(msg.Taxa)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "taxa inválida: %v", err)
	}

	// Criar novo empréstimo
	emprestimo := types.Emprestimo{
		Valor:   msg.Valor,
		Prazo:   msg.Prazo,
		Taxa:    msg.Taxa,
		Devedor: msg.Creator,
		Credor:  "",
	}

	// Adicionar empréstimo ao armazenamento
	id := k.AppendEmprestimo(ctx, emprestimo)

	return &types.MsgSolicitarEmprestimoResponse{
		Id: id,
	}, nil
}
