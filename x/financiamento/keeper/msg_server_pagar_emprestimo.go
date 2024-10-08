package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gabrielscota/defihub/x/financiamento/types"
)

func (k msgServer) PagarEmprestimo(goCtx context.Context, msg *types.MsgPagarEmprestimo) (*types.MsgPagarEmprestimoResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Obter o empréstimo
	emprestimo, found := k.GetEmprestimo(ctx, msg.EmprestimoId)
	if !found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrNotFound, "empréstimo não encontrado: %d", msg.EmprestimoId)
	}

	if emprestimo.Devedor != msg.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "somente o devedor pode pagar o empréstimo")
	}

	// Calcular o valor total a ser pago
	taxaDec, err := sdk.NewDecFromStr(emprestimo.Taxa)
	if err != nil {
		return nil, err
	}

	valorOriginal := sdk.NewDecFromInt(sdk.NewIntFromUint64(emprestimo.Valor))
	juros := valorOriginal.Mul(taxaDec)
	valorTotal := valorOriginal.Add(juros)

	// Verificar se o valor pago é correto
	valorPago := sdk.NewIntFromUint64(msg.Valor)
	if !valorPago.Equal(valorTotal.TruncateInt()) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "valor pago incorreto, esperado: %s", valorTotal.TruncateInt())
	}

	// Transferir os fundos
	devedorAddr, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}

	credorAddr, err := sdk.AccAddressFromBech32(emprestimo.Credor)
	if err != nil {
		return nil, err
	}

	amount := sdk.NewCoin(sdk.DefaultBondDenom, valorPago)

	err = k.bankKeeper.SendCoins(ctx, devedorAddr, credorAddr, sdk.NewCoins(amount))
	if err != nil {
		return nil, err
	}

	// Atualizar o empréstimo (opcional)
	// emprestimo.Status = "Pago"
	// k.SetEmprestimo(ctx, emprestimo)

	return &types.MsgPagarEmprestimoResponse{}, nil
}
