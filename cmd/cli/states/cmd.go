package states

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/stuparm/evm-storage/cmd/cli/flags"
	"github.com/stuparm/evm-storage/cmd/cli/rpc"
)

var CmdStates = &cobra.Command{
	Use:   "states [args]",
	Short: "state related commands",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := cmd.Context()

		rpcURL := flags.ReadFlag[string](cmd, flags.RPCUrl)
		blockNumber := flags.ReadFlag[hexutil.Big](cmd, flags.BlockNumber)

		client := rpc.NewClient(ctx, rpcURL)

		// fetching the block by number
		block, err := rpc.NewCaller[Block](client).
			Call(ctx, "eth_getBlockByNumber", blockNumber.String(), false)
		if err != nil {
			return errors.Wrap(err, "failed to call rpc method (getBlockByNumber)")
		}

		// fetching the states for each transaction in the block
		txs := block.Transactions
		for _, tx := range txs {
			states, err := rpc.NewCaller[States](client).
				Call(ctx, "debug_traceTransaction", tx.String(), NewStateTracerParams())
			if err != nil {
				return errors.Wrap(err, "failed to call rpc method (debug_traceTransaction)")
			}

			fmt.Println(states)
			break
		}

		return nil
	},
}

func init() {
	flags.RegisterFlag(CmdStates, flags.RPCUrl)
	flags.RegisterFlag(CmdStates, flags.BlockNumber)

}
