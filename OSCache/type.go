package OSCache

import (
	"context"
	"time"
)

type Cache interface {
	Get(ctx context.Context, key string) (interface{}, bool)
	SetGos(ctx context.Context, key string, value interface{}, expiration time.Duration) error
	SetOnceGo(ctx context.Context, key string, value interface{}, expiration time.Duration) error
	SetTimeOut(ctx context.Context, key string, value interface{}, expiration time.Duration) error
	Delete(ctx context.Context) (interface{}, error)
}

type item struct {
	val      any
	deadline time.Time
}
