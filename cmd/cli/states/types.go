package states

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
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

type Account struct {
	Code    hexutil.Bytes               `json:"code,omitempty"`
	Storage map[common.Hash]common.Hash `json:"storage,omitempty"`
	Balance *hexutil.Big                `json:"balance,omitempty"`
	Nonce   *big.Int                    `json:"nonce,omitempty"`
}

type States struct {
	Pre  map[common.Address]*Account `json:"pre"`
	Post map[common.Address]*Account `json:"post"`
}
