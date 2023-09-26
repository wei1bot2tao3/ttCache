package OSCache

import (
	"context"
	"time"
)

type Cache interface {
	Get(ctx context.Context, key string) (interface{}, bool)
	Set(ctx context.Context, key string, value interface{}) error
	SetTimeOut(ctx context.Context, key string, value interface{}, time time.Duration) error
	Delete(ctx context.Context)
}

type item struct {
	val      any
	deadline time.Time
}
