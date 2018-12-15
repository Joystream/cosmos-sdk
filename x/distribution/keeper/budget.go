package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/distribution/types"
)

// AllocateCommunityFunds adds or creates to allocation for a beneficiary in a budget distinfo
func (k Keeper) AllocateBudget(ctx sdk.Context, amount sdk.Coins, beneficiary sdk.AccAddress) {
	distInfo := k.GetOrCreateBudgetDistInfo(ctx, beneficiary)
	distInfo.AddAllocation(amount)
	k.SetBudgetDistInfo(ctx, distInfo)
}

// SpendFromCommunityPool - Withdraws full budget from community pool, none if not enough funds
func (k Keeper) SpendFromCommunityPool(ctx sdk.Context, beneficiary sdk.AccAddress) bool {
	distInfo := k.GetOrCreateBudgetDistInfo(ctx, beneficiary)

	// Does this beneficiary have any allocations?
	if !distInfo.HasAllocation() {
		return false
	}

	feePool := k.GetFeePool(ctx)

	// Ensure we can withdraw full amount from community pool
	success, communityPool, withdrawl := distInfo.WithdrawFromPool(feePool.CommunityPool)
	if !success {
		return false
	}

	// Make sure we did not over withdraw - this should be an invariant check?
	if communityPool.HasNegative() {
		return false
	}

	feePool.CommunityPool = communityPool
	k.SetFeePool(ctx, feePool)
	k.bankKeeper.AddCoins(ctx, beneficiary, withdrawl)
	k.RemoveBudgetDistInfo(ctx, distInfo)
	return true
}

// GetOrCreateBudgetDistInfo gets the budget distribution info for a beneficiary if exists or creates new
func (k Keeper) GetOrCreateBudgetDistInfo(ctx sdk.Context, beneficiary sdk.AccAddress) (distInfo types.BudgetDistInfo) {
	store := ctx.KVStore(k.storeKey)
	b := store.Get(GetBudgetDistInfoKey(beneficiary))
	if b == nil {
		// construct a new dist info
		return types.NewBudgetDistInfo(beneficiary)
	}

	k.cdc.MustUnmarshalBinaryLengthPrefixed(b, &distInfo)
	return
}

// SetBudgetDistInfo sets the budget distribution info
func (k Keeper) SetBudgetDistInfo(ctx sdk.Context, distInfo types.BudgetDistInfo) {
	store := ctx.KVStore(k.storeKey)
	b := k.cdc.MustMarshalBinaryLengthPrefixed(distInfo)
	distInfoKey := GetBudgetDistInfoKey(distInfo.GetBeneficiary())
	store.Set(distInfoKey, b)
}

func (k Keeper) RemoveBudgetDistInfo(ctx sdk.Context, distInfo types.BudgetDistInfo) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(GetBudgetDistInfoKey(distInfo.GetBeneficiary()))
}
