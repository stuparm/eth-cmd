package gen

import (
	"github.com/spf13/cobra"
	"github.com/stuparm/eth-cmd/cmd/cli/flags"
)

var CmdGen = &cobra.Command{
	Use:   "gen",
	Short: "Generate ",
	RunE:  func(cmd *cobra.Command, args []string) error { return execute(cmd, args) },
}

func init() {
	flags.RegisterFlag(CmdGen, flags.Output)
}

func execute(cmd *cobra.Command, args []string) error {
	
	return nil
}
