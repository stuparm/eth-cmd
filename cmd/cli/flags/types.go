package flags

type FlagType string

const (
	StringFlagType FlagType = "string"
	IntFlagType    FlagType = "int"
)

type CmdFlag struct {
	Name      string
	Default   string
	Shorthand string
	Usage     string
	Type      FlagType
}

var (
	RPCUrl = CmdFlag{
		Name:      "rpc-url",
		Shorthand: "r",
		Usage:     "--rpc-url <rpc-url>",
		Type:      StringFlagType,
	}

	Timeout = CmdFlag{
		Name:      "timeout",
		Shorthand: "t",
		Usage:     "--timeout <timeout>",
		Type:      IntFlagType,
	}
)
