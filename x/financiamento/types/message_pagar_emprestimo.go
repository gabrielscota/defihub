package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgPagarEmprestimo{}

func NewMsgPagarEmprestimo(creator string, emprestimoId uint64, valor uint64) *MsgPagarEmprestimo {
	return &MsgPagarEmprestimo{
		Creator:      creator,
		EmprestimoId: emprestimoId,
		Valor:        valor,
	}
}

func (msg *MsgPagarEmprestimo) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
