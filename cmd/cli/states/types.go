package states

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type Block struct {
	Transactions []common.Hash `json:"transactions"`
}

type TracerConfig struct {
	DiffMode bool `json:"diffMode"`
}

type TracerParams struct {
	Tracer       string        `json:"tracer"`
	TracerConfig *TracerConfig `json:"tracerConfig"`
}

// TODO: Add options to configure the tracer params
func NewStateTracerParams() *TracerParams {
	return &TracerParams{
		Tracer: "prestateTracer",
		TracerConfig: &TracerConfig{
			DiffMode: true,
		},
	}
}

type AddressAccount map[common.Address]*types.Account

type States struct {
	Pre  []*AddressAccount `json:"pre"`
	Post []*AddressAccount `json:"post"`
}
