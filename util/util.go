package util

import "context"

func WaitingWithContext(ctx context.Context, fn func() error) error {
	for {
		select {
			case <- ctx.Done():
				return ctx.Err()
			default:
				return fn()
		}
	}
}