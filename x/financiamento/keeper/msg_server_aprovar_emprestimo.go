package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gabrielscota/defihub/x/financiamento/types"
)

func (k msgServer) AprovarEmprestimo(goCtx context.Context, msg *types.MsgAprovarEmprestimo) (*types.MsgAprovarEmprestimoResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Obter o empréstimo
	emprestimo, found := k.GetEmprestimo(ctx, msg.EmprestimoId)
	if !found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrNotFound, "empréstimo não encontrado: %d", msg.EmprestimoId)
	}

	if emprestimo.Credor != "" {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "empréstimo já aprovado")
	}

	// Definir o credor
	emprestimo.Credor = msg.Creator

	// Atualizar o empréstimo
	k.SetEmprestimo(ctx, emprestimo)

	// Transferir fundos do credor para o devedor
	credorAddr, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}

	devedorAddr, err := sdk.AccAddressFromBech32(emprestimo.Devedor)
	if err != nil {
		return nil, err
	}

	amount := sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewIntFromUint64(emprestimo.Valor))

	err = k.bankKeeper.SendCoins(ctx, credorAddr, devedorAddr, sdk.NewCoins(amount))
	if err != nil {
		return nil, err
	}

	return &types.MsgAprovarEmprestimoResponse{}, nil
}
