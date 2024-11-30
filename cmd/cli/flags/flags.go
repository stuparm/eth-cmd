package flags

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/spf13/cobra"
	"strconv"
)

func RegisterFlag(cmd *cobra.Command, flag CmdFlag) {
	switch flag.Type {
	case StringFlagType:
		val := ""
		cmd.Flags().StringVarP(&val, flag.Name, flag.Shorthand, "", flag.Usage)
		return
	case IntFlagType:
		val := 0
		cmd.Flags().IntVarP(&val, flag.Name, flag.Shorthand, 0, flag.Usage)
		return
	case HexFlagType:
		val := ""
		cmd.Flags().StringVarP(&val, flag.Name, flag.Shorthand, "", flag.Usage)
		return
	}

	panic(fmt.Sprintf("Unsupported flag type: %s", flag.Type))
}

func ReadFlag[T any](cmd *cobra.Command, flag CmdFlag) T {
	switch flag.Type {
	case StringFlagType:
		value := cmd.Flag(flag.Name).Value.String()
		return any(value).(T)
	case IntFlagType:
		valueStr := cmd.Flag(flag.Name).Value.String()
		value, err := strconv.Atoi(valueStr)
		if err != nil {
			panic(err)
		}
		return any(value).(T)
	case HexFlagType:
		valueStr := cmd.Flag(flag.Name).Value.String()
		value := hexutil.Big(*hexutil.MustDecodeBig(valueStr))
		return any(value).(T)
	}

	panic(fmt.Sprintf("Unsupported flag type: %s", flag.Type))
}
