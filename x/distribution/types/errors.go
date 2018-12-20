// nolint
package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type CodeType = sdk.CodeType

const (
	DefaultCodespace       sdk.CodespaceType = "DISTR"
	CodeInvalidInput       CodeType          = 103
	CodeNoDistributionInfo CodeType          = 104
)

func ErrNilDelegatorAddr(codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(codespace, CodeInvalidInput, "delegator address is nil")
}
func ErrNilWithdrawAddr(codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(codespace, CodeInvalidInput, "withdraw address is nil")
}
func ErrNilValidatorAddr(codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(codespace, CodeInvalidInput, "validator address is nil")
}
func ErrNoDelegationDistInfo(codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(codespace, CodeNoDistributionInfo, "no delegation distribution info")
}
func ErrNoValidatorDistInfo(codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(codespace, CodeNoDistributionInfo, "no validator distribution info")
}
func ErrWithdrawerHasNoAllocation(codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(codespace, CodeNoDistributionInfo, "withdrawer has no budget allocated")
}
func ErrPoolDoesNotHaveEnoughFunds(codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(codespace, CodeNoDistributionInfo, "pool does not have enough funds")
}
