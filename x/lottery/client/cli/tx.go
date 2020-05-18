package cli

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/cosmos/cosmos-sdk/x/auth"
	authclient "github.com/cosmos/cosmos-sdk/x/auth/client"
	"github.com/spf13/cobra"

	"github.com/freemanjackal/lottery-cosmoschain/x/lottery/types"
	"strconv"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd(storeKey string, cdc *codec.Codec) *cobra.Command {
	lotteryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "lottery transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}
	lotteryCmd.AddCommand(flags.PostCommands(
		GetCmdBet(cdc),
		CreateLotteryCmdRequest(cdc),
		GetCmdCloseLottery(cdc),
		GetCmdSetChannel(cdc),
	)...)

	return lotteryCmd
}

// GetCmdRequest implements the request command handler.
func GetCmdBet(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "bet [number] [price]",
		Short: "Make a new bet on a number",
		Args:  cobra.ExactArgs(2),
		Long: strings.TrimSpace(
			fmt.Sprintf(`Make a new lottery bet.
Example:
$ %s tx bet on number 1000000dfsbsdfdf/transfer/uatom
`,
				version.ClientName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(authclient.GetTxEncoder(cdc))

			price, err := sdk.ParseCoins(args[1])
			//number := 26
			if err != nil {
				return err
			}

			number, err := strconv.ParseUint(args[0], 10, 64)
			//number := 26
			if err != nil {
				return err
			}
			msg := types.NewMsgPlayLottery(
				cliCtx.GetFromAddress(),
				price,
				number,
			)

			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			return authclient.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}

	return cmd
}

func CreateLotteryCmdRequest(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create lottery",
		Args:  cobra.ExactArgs(0),
		Long: strings.TrimSpace(
			fmt.Sprintf(`Make a new lottery bet.
Example:
$ %s tx bet on number 1000000dfsbsdfdf/transfer/uatom
`,
				version.ClientName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(authclient.GetTxEncoder(cdc))

			//amount, err := sdk.ParseCoins(args[0])
			//number := 26
			//initial amount is zero
			var amount = sdk.Coins{sdk.NewInt64Coin("stake", 0)}
			
			msg := types.NewMsgCreateLottery(
				cliCtx.GetFromAddress(),
				types.Open,
				amount,
				)

			err := msg.ValidateBasic()
			if err != nil {
				return err
			}

			return authclient.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}

	return cmd
}

func GetCmdCloseLottery(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "close",
		Short: "close lottery and open a new one",
		Args:  cobra.ExactArgs(0),
		Long: strings.TrimSpace(
			fmt.Sprintf(`close lottery and open new one, tranfer funds to these new lottery.
Example:
$ %s tx close 
`,
				version.ClientName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(authclient.GetTxEncoder(cdc))

			
			msg := types.NewMsgCloseLottery(
				cliCtx.GetFromAddress(),
			)

			err := msg.ValidateBasic()
			if err != nil {
				return err
			}

			return authclient.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}

	return cmd
}

// GetCmdSetChannel implements the set channel command handler.
func GetCmdSetChannel(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-channel [chain-id] [port] [channel-id]",
		Short: "Register a verified channel",
		Args:  cobra.ExactArgs(3),
		Long: strings.TrimSpace(
			fmt.Sprintf(`Register a verified channel.
Example:
$ %s tx lottery set-cahnnel bandchain lottery dbdfgsdfsd
`,
				version.ClientName,
			),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(authclient.GetTxEncoder(cdc))

			msg := types.NewMsgSetSourceChannel(
				args[0],
				args[1],
				args[2],
				cliCtx.GetFromAddress(),
			)

			err := msg.ValidateBasic()
			if err != nil {
				return err
			}

			return authclient.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}

	return cmd
}
