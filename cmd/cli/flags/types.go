package flags

type FlagType string

const (
	StringFlagType FlagType = "string"
	IntFlagType    FlagType = "int"
	HexFlagType    FlagType = "hex"
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
	FromBlockNumber = CmdFlag{
		Name:      "from-block-number",
		Shorthand: "",
		Usage:     "--from-block-number <from-block>",
		Type:      IntFlagType,
	}
	ToBlockNumber = CmdFlag{
		Name:      "to-block-number",
		Shorthand: "",
		Usage:     "--to-block-number <to-block>",
		Type:      IntFlagType,
	}
)
