package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/stuparm/evm-storage/cmd/cli/states"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "evm-storage",
	Short: "evm-storage is a cli tool to analyze storage of evm contracts",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Executing evm-storage cmd")
	},
}

func init() {
	rootCmd.AddCommand(states.CmdStates)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
