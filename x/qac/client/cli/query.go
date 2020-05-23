package cli

import (
	"fmt"
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/stracedude/qac/x/qac/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string, cdc *codec.Codec) *cobra.Command {
	// Group qac queries under a subcommand
	qacQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	qacQueryCmd.AddCommand(
		flags.GetCommands(
			GetCmdListQuestion(queryRoute, cdc),
			GetCmdGetQuestion(queryRoute, cdc),
			GetCmdGetOwnQuestion(queryRoute, cdc),
			)...,
	)

	return qacQueryCmd
}

func GetCmdListQuestion(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "list",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/"+types.QueryListQuestion, queryRoute), nil)
			if err != nil {
				fmt.Printf("could not get list question\n%s\n", err.Error())
				return nil
			}

			var out types.QuestionsList
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}

func GetCmdGetQuestion(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "get [questionHash]",
		Short: "get q by hash",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			var hash = args[0]
			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s/%s", queryRoute, types.QueryGetQuestion, hash), nil)
			if err != nil {
				fmt.Printf("could not get question\n%s\n", err.Error())
				return nil
			}

			var out types.QuestionsView
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}

func GetCmdGetOwnQuestion(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "getOwnQuestion",
		Short: "get q by acc",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s/%s", queryRoute, types.QueryGetOwnQuestion, cliCtx.GetFromAddress()), nil)
			if err != nil {
				fmt.Printf("could not get question\n%s\n", err.Error())
				return nil
			}

			var out types.QuestionsList
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}

