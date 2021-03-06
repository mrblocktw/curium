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

package cli

import (
	"fmt"
	"github.com/bluzelle/curium/x/crud/internal/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/spf13/cobra"
)

func GetQueryCmd(storeKey string, cdc *codec.Codec) *cobra.Command {
	crudQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the crud module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	crudQueryCmd.AddCommand(client.GetCommands(GetCmdBLZQRead(storeKey, cdc))...)
	crudQueryCmd.AddCommand(client.GetCommands(GetCmdBLZQHas(storeKey, cdc))...)
	crudQueryCmd.AddCommand(client.GetCommands(GetCmdBLZQKeys(storeKey, cdc))...)

	return crudQueryCmd
}

func GetCmdBLZQRead(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "read [UUID] [key]",
		Short: "read UUID key",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			UUID := args[0]
			key := args[1]
			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/read/%s/%s", queryRoute, UUID, key), nil)

			if err != nil {
				fmt.Printf("could not read key - %s : %s", UUID, key)
				return nil
			}
			var out types.QueryResultRead
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}

func GetCmdBLZQHas(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "has [UUID] [key]",
		Short: "has UUID key",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			UUID := args[0]
			key := args[1]
			res, _, _ := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/has/%s/%s", queryRoute, UUID, key), nil)

			var out types.QueryResultHas
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}

func GetCmdBLZQKeys(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "keys [UUID]",
		Short: "keys UUID",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			UUID := args[0]
			res, _, _ := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/keys/%s", queryRoute, UUID), nil)

			var out types.QueryResultKeys
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}
