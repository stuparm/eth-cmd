package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/stuparm/eth-cmd/cmd/cli/states"
	"os"
)

var rootCmd = &cobra.Command{
	Use:               "evm-storage",
	Short:             "evm-storage is a cli tool to analyze storage of evm contracts",
	CompletionOptions: cobra.CompletionOptions{DisableDefaultCmd: true},
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

func init() {
	rootCmd.AddCommand(states.CmdStates)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(os.Stderr, err)
		os.Exit(1)
	}
}
