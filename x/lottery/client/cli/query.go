package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/spf13/cobra"

	"github.com/freemanjackal/lottery-cosmoschain/x/lottery/types"
)

// GetQueryCmd returns
func GetQueryCmd(storeKey string, cdc *codec.Codec) *cobra.Command {
	lotteryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the goldcdp module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	lotteryCmd.AddCommand(flags.GetCommands(
		GetCmdReadOrder(storeKey, cdc),
	)...)

	return lotteryCmd
}

// GetCmdReadOrder queries order by orderID
func GetCmdReadOrder(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:  "order",
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			orderID := args[0]

			res, _, err := cliCtx.QueryWithData(
				fmt.Sprintf("custom/%s/order/%s", queryRoute, orderID),
				nil,
			)
			if err != nil {
				fmt.Printf("read request fail - %s \n", orderID)
				return nil
			}

			var lottery types.Lottery
			if err := cdc.UnmarshalJSON(res, &lottery); err != nil {
				return err
			}
			return cliCtx.PrintOutput(lottery)
		},
	}
}
