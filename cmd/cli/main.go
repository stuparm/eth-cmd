package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/stuparm/eth-cmd/cmd/cli/gen"
	"github.com/stuparm/eth-cmd/cmd/cli/states"
	"os"
)

var rootCmd = &cobra.Command{
	Use:               "eth-cmd",
	Short:             "eth-cmd is a cli tool",
	CompletionOptions: cobra.CompletionOptions{DisableDefaultCmd: true},
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

func init() {
	rootCmd.AddCommand(states.CmdStates)
	rootCmd.AddCommand(gen.CmdGen)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(os.Stderr, err)
		os.Exit(1)
	}
}
