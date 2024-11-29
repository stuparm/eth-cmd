package states

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/stuparm/evm-storage/cmd/cli/flags"
)

var CmdStates = &cobra.Command{
	Use:   "states [args]",
	Short: "state related commands",
	RunE: func(cmd *cobra.Command, args []string) error {
		rpc := flags.ReadFlag(cmd, flags.RPCUrl)
		timeout := flags.ReadFlag(cmd, flags.Timeout)

		fmt.Println("rpc-url: ", rpc)
		fmt.Println("timeout: ", timeout)

		return nil
	},
}

func init() {
	flags.RegisterFlag(CmdStates, flags.RPCUrl)
	flags.RegisterFlag(CmdStates, flags.Timeout)
}
