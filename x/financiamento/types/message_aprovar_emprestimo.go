package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgAprovarEmprestimo{}

func NewMsgAprovarEmprestimo(creator string, emprestimoId uint64) *MsgAprovarEmprestimo {
	return &MsgAprovarEmprestimo{
		Creator:      creator,
		EmprestimoId: emprestimoId,
	}
}

func (msg *MsgAprovarEmprestimo) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
