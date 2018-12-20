package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// BudgetDistInfo distribution information for a budget
type BudgetDistInfo struct {
	BeneficiaryAddr     sdk.AccAddress `json:"beneficiary_addr"`
	WithdrawlAllocation sdk.Coins      `json:"allocation"` // funds allocated remaining to be withdrawn
}

func NewBudgetDistInfo(beneficiary sdk.AccAddress) BudgetDistInfo {
	return BudgetDistInfo{beneficiary, sdk.Coins{}}
}

func (distInfo BudgetDistInfo) GetBeneficiary() sdk.AccAddress {
	return distInfo.BeneficiaryAddr
}

// Add adds coins to the allocation for the budget beneficiary
func (distInfo *BudgetDistInfo) AddAllocation(amount sdk.Coins) {
	distInfo.WithdrawlAllocation = distInfo.WithdrawlAllocation.Plus(amount)
}

func (distInfo BudgetDistInfo) HasAllocation() bool {
	return !distInfo.WithdrawlAllocation.IsZero()
}

func (distInfo *BudgetDistInfo) WithdrawFromPool(pool DecCoins) (bool, DecCoins, sdk.Coins) {
	newPool := pool.Minus(NewDecCoins(distInfo.WithdrawlAllocation))

	if newPool.HasNegative() {
		return false, pool, sdk.Coins{}
	}

	withdrawl := distInfo.WithdrawlAllocation
	distInfo.WithdrawlAllocation = sdk.Coins{}

	return true, newPool, withdrawl
}
