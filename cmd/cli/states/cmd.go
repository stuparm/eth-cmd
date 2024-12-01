package states

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/stuparm/eth-cmd/cmd/cli/flags"
	"github.com/stuparm/eth-cmd/cmd/cli/rpc"
	"time"
)

var CmdStates = &cobra.Command{
	Use:   "states [args]",
	Short: "state related commands",
	RunE:  func(cmd *cobra.Command, args []string) error { return execute(cmd, args) },
}

func init() {
	flags.RegisterFlag(CmdStates, flags.RPCUrl)
	flags.RegisterFlag(CmdStates, flags.BlockNumber)
	flags.RegisterFlag(CmdStates, flags.Throttle)
	flags.RegisterFlag(CmdStates, flags.Limit)
	flags.RegisterFlag(CmdStates, flags.Output)
	flags.RegisterFlag(CmdStates, flags.Summary)
}

func execute(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()

	rpcURL := flags.ReadFlag[string](cmd, flags.RPCUrl)
	blockNumber := flags.ReadFlag[hexutil.Big](cmd, flags.BlockNumber)
	throttle := flags.ReadFlag[time.Duration](cmd, flags.Throttle)
	limit := flags.ReadFlag[int](cmd, flags.Limit)
	output := flags.ReadFlag[string](cmd, flags.Output)
	withSummary := flags.ReadFlag[bool](cmd, flags.Summary)

	client := rpc.NewClient(ctx, rpcURL)

	// fetching the block by number
	block, err := rpc.NewCaller[Block](client).
		Call(ctx, "eth_getBlockByNumber", blockNumber.String(), false)
	if err != nil {
		return errors.Wrap(err, "failed to call rpc method (getBlockByNumber)")
	}

	// fetching the states for each transaction in the block
	blockStates := NewBlockStates()
	txs := block.Transactions

	i := 1
	for _, tx := range txs {
		txState, err := rpc.NewCaller[States](client).
			Call(ctx, "debug_traceTransaction", tx.String(), NewStateTracerParams())
		if err != nil {
			return errors.Wrap(err, "failed to call rpc method (debug_traceTransaction)")
		}

		for addr, acc := range txState.Post {
			blockStates.AddAccountState(addr, acc)
		}

		fmt.Println(fmt.Sprintf("Executing tx %d of %d", i, len(txs)))

		if throttle > 0 {
			time.Sleep(throttle)
		}
		if limit > 0 && i >= limit {
			break
		}

		i++
	}

	reporter := NewReporter(output)
	if err := reporter.WriteBlockStates(blockStates); err != nil {
		return errors.Wrap(err, "failed to write block states")
	}

	if withSummary {
		summary := NewSummarizer().Summarize(blockStates)
		if err := reporter.WriteSummary(summary); err != nil {
			return errors.Wrap(err, "failed to write summary")
		}
	}

	return nil
}
