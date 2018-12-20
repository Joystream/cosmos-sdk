package distribution

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

// query endpoints supported by the governance Querier
const (
	QueryCommunityPool = "feepool"
)

func NewQuerier(keeper Keeper, cdc *codec.Codec) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, sdk.Error) {
		switch path[0] {
		case QueryCommunityPool:
			return queryFeePool(ctx, path[1:], req, keeper, cdc)
		default:
			return nil, sdk.ErrUnknownRequest("unknown distribution query endpoint")
		}
	}
}

func queryFeePool(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper, cdc *codec.Codec) ([]byte, sdk.Error) {

	bz, err := codec.MarshalJSONIndent(cdc, keeper.GetFeePool(ctx))
	if err != nil {
		return nil, sdk.ErrInternal(sdk.AppendMsgToErr("could not marshal result to JSON", err.Error()))
	}
	return bz, nil

}
