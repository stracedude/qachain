package cli

import (
	"bufio"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/stracedude/qac/x/qac/types"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd(cdc *codec.Codec) *cobra.Command {
	qacTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	qacTxCmd.AddCommand(flags.PostCommands(
		GetCmdCreateQuestion(cdc),
		GetCmdCreateAnswer(cdc),
		GetCmdAnswerVote(cdc),

	)...)

	return qacTxCmd
}

func GetCmdCreateQuestion(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "createQuestion [reward] [description]",
		Short: "Creates a new question with a reward",
		Args:  cobra.ExactArgs(2), // Does your request require arguments
		RunE: func(cmd *cobra.Command, args []string) error {

			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

			reward, err := sdk.ParseCoins(args[0])
			if err != nil {
				return err
			}

			var description = args[1]

			msg := types.NewMsgCreateQuestion(description, reward, cliCtx.GetFromAddress())
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

func GetCmdCreateAnswer(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "createAnswer [questHash] [answer]",
		Short: "Creates a new answer for question by hash ",
		Args:  cobra.ExactArgs(2), // Does your request require arguments
		RunE: func(cmd *cobra.Command, args []string) error {

			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

			var questHash = args[0]
			var answer = args[1]

			msg := types.NewMsgCreateAnswer(questHash, answer, cliCtx.GetFromAddress())
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

func GetCmdAnswerVote(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "AnswerVote [answerHash]",
		Short: "Vote answer by hash",
		Args:  cobra.ExactArgs(1), // Does your request require arguments
		RunE: func(cmd *cobra.Command, args []string) error {

			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

			var answerHash = args[0]

			msg := types.NewMsgAnswerVote(answerHash, cliCtx.GetFromAddress())
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}
}

