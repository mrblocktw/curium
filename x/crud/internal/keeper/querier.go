// Copyright (C) 2020 Bluzelle
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License, version 3,
// as published by the Free Software Foundation.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package keeper

import (
	"github.com/bluzelle/curium/x/crud/internal/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

const (
	QueryRead = "read"
	QueryHas  = "has"
	QueryKeys = "keys"
)

func NewQuerier(keeper Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err sdk.Error) {
		switch path[0] {
		case QueryRead:
			return queryRead(ctx, path[1:], req, keeper)
		case QueryHas:
			return queryHas(ctx, path[1:], req, keeper)
		case QueryKeys:
			return queryKeys(ctx, path[1:], req, keeper)
		default:
			return nil, sdk.ErrUnknownRequest("unknown crud query endpoint")
		}
	}
}

func queryRead(ctx sdk.Context, path []string, _ abci.RequestQuery, keeper Keeper) ([]byte, sdk.Error) {
	value := keeper.GetBLZValue(ctx, path[0], path[1]).Value
	if len(value) == 0 {
		return []byte{}, sdk.ErrUnknownRequest("could not read key")
	}

	res, err := codec.MarshalJSONIndent(keeper.cdc, types.QueryResultRead{Value: value})
	if err != nil {
		panic("could not marshal result to JSON")
	}

	return res, nil
}

func queryHas(ctx sdk.Context, path []string, _ abci.RequestQuery, keeper Keeper) ([]byte, sdk.Error) {
	value := keeper.IsKeyPresent(ctx, path[0], path[1])

	res, err := codec.MarshalJSONIndent(keeper.cdc, types.QueryResultHas{Value: value})
	if err != nil {
		panic("could not marshal result to JSON")
	}

	return res, nil
}

func queryKeys(ctx sdk.Context, path []string, _ abci.RequestQuery, keeper Keeper) ([]byte, sdk.Error) {
	keys := keeper.GetKeys(ctx, path[0])

	res, err := codec.MarshalJSONIndent(keeper.cdc, types.QueryResultKeys{keys})
	if err != nil {
		panic("could not marshal result to JSON")
	}

	return res, nil
}
