package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateW3Func{}

func NewMsgCreateW3Func(creator string, cond string, code string) *MsgCreateW3Func {
	return &MsgCreateW3Func{
		Creator: creator,
		Cond:    cond,
		Code:    code,
	}
}

func (msg *MsgCreateW3Func) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
