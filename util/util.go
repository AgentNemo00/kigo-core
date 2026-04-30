package util

import (
	"context"
	"fmt"
)

func WaitingWithContext[T any](ctx context.Context, ch chan T) (*T, error) {
	for {
		select {
			case <- ctx.Done():
				return nil, ctx.Err()
			case value, ok := <-ch:
				if !ok {
					return nil, fmt.Errorf("channel closed")
				}
				return &value, nil
		}
	}
}
