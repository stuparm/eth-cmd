package flags

import (
	"fmt"
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
	}

	panic(fmt.Sprintf("Unsupported flag type: %s", flag.Type))
}

func ReadFlag(cmd *cobra.Command, flag CmdFlag) any {
	switch flag.Type {
	case StringFlagType:
		value := cmd.Flag(flag.Name).Value.String()
		return value
	case IntFlagType:
		valueStr := cmd.Flag(flag.Name).Value.String()
		value, err := strconv.Atoi(valueStr)
		if err != nil {
			panic(err)
		}
		return value
	}

	panic(fmt.Sprintf("Unsupported flag type: %s", flag.Type))
}
