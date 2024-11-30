package states

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/stuparm/evm-storage/cmd/cli/flags"
	"github.com/stuparm/evm-storage/cmd/cli/rpc"
	"time"
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
				if len(acc.Code) != 0 {
					blockStates.AddCode(addr)
				}
				if acc.Balance != nil {
					blockStates.AddBalance(addr)
				}
				if acc.Nonce != nil {
					blockStates.AddNonce(addr)
				}
				for key, _ := range acc.Storage {
					blockStates.AddStorage(addr, key)
				}
			}

			fmt.Println(fmt.Sprintf("Executing tx %d of %d", i, len(txs)))
			i++
			time.Sleep(1 * time.Second) // to avoid rate limiting, TODO: remove this

		}

		for address, i := range blockStates.balanceCounter {
			fmt.Println("Address:", address.Hex(), "Balance:", i)
		}
		for address, i := range blockStates.nonceCounter {
			fmt.Println("Address:", address.Hex(), "Nonce:", i)
		}
		for address, i := range blockStates.codesCounter {
			fmt.Println("Address:", address.Hex(), "Code:", i)
		}
		for address, storage := range blockStates.storageCounter {
			for key, i := range storage {
				fmt.Println("Address:", address.Hex(), "Storage:", key.Hex(), i)
			}
		}

		return nil
	},
}

func init() {
	flags.RegisterFlag(CmdStates, flags.RPCUrl)
	flags.RegisterFlag(CmdStates, flags.BlockNumber)

}
