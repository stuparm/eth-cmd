package rpc

import (
	"context"
	"github.com/pkg/errors"
)

type Caller[T any] interface {
	Call(ctx context.Context, method string, params ...any) (T, error)
}

type caller[T any] struct {
	client Client
}

func NewCaller[T any](client Client) Caller[T] {
	return &caller[T]{
		client: client,
	}
}

func (c *caller[T]) Call(ctx context.Context, method string, params ...any) (T, error) {
	var result T
	err := c.client.Call(ctx, &result, method, params...)
	if err != nil {
		return result, errors.Wrap(err, "failed to call rpc method")
	}

	return result, nil
}
