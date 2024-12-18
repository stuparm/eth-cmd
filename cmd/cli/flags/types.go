package flags

type FlagType string

const (
	StringFlagType   FlagType = "string"
	IntFlagType      FlagType = "int"
	HexFlagType      FlagType = "hex"
	DurationFlagType FlagType = "duration"
	BoolFlagType     FlagType = "bool"
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
	BlockNumber = CmdFlag{
		Name:      "block-number",
		Shorthand: "b",
		Usage:     "--block-number <block-number>",
		Type:      HexFlagType,
	}
	Throttle = CmdFlag{
		Name:      "throttle",
		Shorthand: "t",
		Usage:     "--throttle <throttle>",
		Type:      DurationFlagType,
	}
	Limit = CmdFlag{
		Name:      "limit",
		Shorthand: "l",
		Usage:     "--limit <limit>",
		Type:      IntFlagType,
	}
	Output = CmdFlag{
		Name:      "output",
		Shorthand: "o",
		Usage:     "--output <output>",
		Type:      StringFlagType,
	}
	Summary = CmdFlag{
		Name:      "summary",
		Shorthand: "",
		Usage:     "--summary",
		Type:      BoolFlagType,
	}
	FromBlockNumber = CmdFlag{
		Name:      "from-block-number",
		Shorthand: "",
		Usage:     "--from-block-number <from-block>",
		Type:      HexFlagType,
	}
	ToBlockNumber = CmdFlag{
		Name:      "to-block-number",
		Shorthand: "",
		Usage:     "--to-block-number <to-block>",
		Type:      HexFlagType,
	}
)
