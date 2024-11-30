package rpc

import (
	"context"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/pkg/errors"
)

var _ Client = (*client)(nil)

type Client interface {
	Call(ctx context.Context, result any, method string, params ...any) error
}

type client struct {
	url string
}

func NewClient(ctx context.Context, url string) Client {
	return &client{
		url: url,
	}
}

func (c *client) Call(ctx context.Context, result any, method string, params ...any) error {
	rpcCLient, err := rpc.DialContext(ctx, c.url)
	if err != nil {
		return errors.Wrap(err, "failed to dial rpc")
	}

	err = rpcCLient.CallContext(ctx, &result, method, params...)
	if err != nil {
		return errors.Wrap(err, "failed to call rpc method")
	}

	return nil
}
