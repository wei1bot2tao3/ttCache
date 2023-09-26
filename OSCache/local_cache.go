package OSCache

import (
	"context"
	"ttCache/internal/errs"

	"sync"
	"time"
)

// BuildInMapCache 一个缓存Map 类型的缓存
type BuildInMapCache struct {
	data      map[string]*item
	mutex     sync.RWMutex
	close     chan struct{}
	onEvicted func(key string, val any)
}

func (b *BuildInMapCache) Set(ctx context.Context, key string, value interface{}) error {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	newItem := &item{
		val: value,
	}
	if _, ok := b.data[key]; ok {
		return errs.ErrKeyExists
	}
	b.data[key] = newItem
	return nil
}
func (b *BuildInMapCache) SetTimeOut(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	//TODO implement me
	panic("implement me")
}
func (b *BuildInMapCache) Get(ctx context.Context, key string) (interface{}, error) {
	b.mutex.RLock()
	defer b.mutex.RUnlock()
	res, ok := b.data[key]
	if !ok {
		return nil, errs.ErrKeyNotFound
	}
	return res, nil
}

func (b *BuildInMapCache) Delete(ctx context.Context) (interface{}, error) {
	//TODO implement me
	panic("implement me")
}

func NewBuildInMapCache() *BuildInMapCache {
	res := &BuildInMapCache{
		data: make(map[string]*item),
	}
	return res
}
