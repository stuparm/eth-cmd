package states

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/stuparm/eth-cmd/cmd/cli/flags"
	"github.com/stuparm/eth-cmd/cmd/cli/rpc"
	"math/big"
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
	flags.RegisterFlag(CmdStates, flags.FromBlockNumber)
	flags.RegisterFlag(CmdStates, flags.ToBlockNumber)
	flags.RegisterFlag(CmdStates, flags.Throttle)
	flags.RegisterFlag(CmdStates, flags.Limit)
	flags.RegisterFlag(CmdStates, flags.Output)
	flags.RegisterFlag(CmdStates, flags.Summary)
}

func execute(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()

	rpcURL := flags.ReadFlag[string](cmd, flags.RPCUrl)
	blockNumber := flags.ReadFlag[hexutil.Big](cmd, flags.BlockNumber)
	fromBlockNumber := flags.ReadFlag[hexutil.Big](cmd, flags.FromBlockNumber)
	toBlockNumber := flags.ReadFlag[hexutil.Big](cmd, flags.ToBlockNumber)
	throttle := flags.ReadFlag[time.Duration](cmd, flags.Throttle)
	limit := flags.ReadFlag[int](cmd, flags.Limit)
	output := flags.ReadFlag[string](cmd, flags.Output)
	withSummary := flags.ReadFlag[bool](cmd, flags.Summary)

	err := validateBlockParams(blockNumber, fromBlockNumber, toBlockNumber)
	if err != nil {
		return errors.Wrap(err, "invalid block parameters")
	}

	if fromBlockNumber.ToInt().Int64() == 0 {
		fromBlockNumber = blockNumber
	}

	if toBlockNumber.ToInt().Int64() == 0 {
		toBlockNumber = blockNumber
	}

	client := rpc.NewClient(ctx, rpcURL)

	blocks := make([]Block, 0)
	for {
		if fromBlockNumber.ToInt().Int64() > toBlockNumber.ToInt().Int64() {
			break
		}

		block, err := rpc.NewCaller[Block](client).
			Call(ctx, "eth_getBlockByNumber", fromBlockNumber.String(), false)
		if err != nil {
			return errors.Wrap(err, "failed to call rpc method (getBlockByNumber)")
		}
		blocks = append(blocks, block)

		fromBlockNumber = hexutil.Big(*big.NewInt(0).Add(fromBlockNumber.ToInt(), big.NewInt(1)))
	}

	txs := make([]common.Hash, 0)
	for _, b := range blocks {
		txs = append(txs, b.Transactions...)
	}

	preStates := NewBlockStates()
	postStates := NewBlockStates()
	i := 1
	for _, tx := range txs {
		txState, err := rpc.NewCaller[States](client).
			Call(ctx, "debug_traceTransaction", tx.String(), NewStateTracerParams())
		if err != nil {
			return errors.Wrap(err, "failed to call rpc method (debug_traceTransaction)")
		}

		for addr, acc := range txState.Post {
			postStates.AddAccountState(addr, acc)
		}

		for addr, acc := range txState.Pre {
			preStates.AddAccountState(addr, acc)
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
	if err := reporter.WriteBlockStates(preStates, postStates); err != nil {
		return errors.Wrap(err, "failed to write block states")
	}

	if withSummary {
		summary := NewSummarizer().Summarize(preStates, postStates)
		if err := reporter.WriteSummary(summary); err != nil {
			return errors.Wrap(err, "failed to write summary")
		}
	}

	return nil
}

func validateBlockParams(blockNumber, fromBlockNumber, toBlockNumber hexutil.Big) error {
	if blockNumber.ToInt().Int64() == 0 && fromBlockNumber.ToInt().Int64() == 0 && toBlockNumber.ToInt().Int64() == 0 {
		return errors.New("at least some block parameter must be provided")
	}

	if blockNumber.ToInt().Int64() != 0 && (fromBlockNumber.ToInt().Int64() != 0 || toBlockNumber.ToInt().Int64() != 0) {
		return errors.New("block-number cannot be used with from-block-number or to-block-number")
	}

	if fromBlockNumber.ToInt().Int64() != 0 && toBlockNumber.ToInt().Int64() == 0 {
		return errors.New("from-block-number must be used with to-block-number")
	}

	if fromBlockNumber.ToInt().Int64() == 0 && toBlockNumber.ToInt().Int64() != 0 {
		return errors.New("to-block-number must be used with from-block-number")
	}

	return nil
}
