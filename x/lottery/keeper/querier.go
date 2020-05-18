package keeper

import (
	"fmt"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	abci "github.com/tendermint/tendermint/abci/types"
)

const (
	QueryLottery = "Lottery"
)

// NewQuerier is the module level router for state queries.
func NewQuerier(keeper Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err error) {
		switch path[0] {
		case QueryLottery:
			return queryLottery(ctx, path[1:], req, keeper)
		default:
			return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unknown nameservice query endpoint")
		}
	}
}

// queryLottery is a query function to get order by order ID.
func queryLottery(
	ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper,
) ([]byte, error) {
	if len(path) == 0 {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "must specify the order id")
	}
	id, err := strconv.ParseInt(path[0], 10, 64)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, fmt.Sprintf("wrong format for requestid %s", err.Error()))
	}
	lottery, err := keeper.GetLottery(ctx, uint64(id))
	if err != nil {
		return nil, err
	}
	return keeper.cdc.MustMarshalJSON(lottery), nil
}
