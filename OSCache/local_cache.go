package OSCache

import (
	"context"
	"sync"
	"time"
	"ttCache/internal/errs"
)

// BuildInMapCache 一个缓存Map 类型的缓存
type BuildInMapCache struct {
	data      map[string]*item
	mutex     sync.RWMutex
	close     chan struct{}
	onEvicted func(key string, val any)
}

// SetGos 设置过期时间 每设置一个key就开一个goroutine来监控
func (b *BuildInMapCache) SetGos(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	var dl time.Time
	if expiration > 0 {
		dl = time.Now().Add(expiration)
	}
	if _, ok := b.data[key]; ok {
		return errs.ErrKeyExists
	}
	b.data[key] = &item{
		val:      value,
		deadline: dl,
	}
	if expiration > 0 {
		time.AfterFunc(expiration, func() {
			b.mutex.Lock()
			defer b.mutex.Unlock()
			val, ok := b.data[key]
			// key 存在 有过期时间 过期时间超过了
			if ok && !val.deadline.IsZero() && val.deadline.Before(time.Now()) {
				delete(b.data, key)
			}

		})
	}

	return nil
}

func (b *BuildInMapCache) SetOneGo(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	b.mutex.Lock()
	defer b.mutex.Unlock()

	var dl time.Time
	if expiration > 0 {
		dl = time.Now().Add(expiration)
	}
	if _, ok := b.data[key]; ok {
		return errs.ErrKeyExists
	}
	b.data[key] = &item{
		val:      value,
		deadline: dl,
	}

	return nil
}

func (b *BuildInMapCache) SetTimeOut(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
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
func (b *BuildInMapCache) Get(ctx context.Context, key string) (interface{}, error) {
	b.mutex.RLock()
	defer b.mutex.RUnlock()
	res, ok := b.data[key]
	if !ok {
		return nil, errs.NewErrNotfound(key)
	}
	return res, nil
}

func (b *BuildInMapCache) Delete(ctx context.Context) (interface{}, error) {
	//TODO implement me
	panic("implement me")
}

func (b *BuildInMapCache) Close() error {

	select {
	case b.close <- struct{}{}:

	default:
		return errs.ErrCacheCloes
	}
	return nil

}

type ExpirationDecorator struct {
	cache Cache
}

func NewExpirationDecoratorOneGo(cache Cache) *ExpirationDecorator {
	return &ExpirationDecorator{
		cache: cache,
	}
}

func NewBuildInMapCacheGos(size int) *BuildInMapCache {
	return &BuildInMapCache{
		data: make(map[string]*item, size),
	}
}

// NewBuildInMapCacheOneGo 开启一个go去轮训 所有时间过期的key 然后close后会过期
func NewBuildInMapCacheOneGo(size int, interval time.Duration) *BuildInMapCache {
	res := &BuildInMapCache{
		data: make(map[string]*item, size),
	}
	go func() {
		ticker := time.NewTicker(interval)
		for {
			select {
			case t := <-ticker.C:
				res.mutex.Lock()
				i := 0
				for key, val := range res.data {
					if i > 1000 {
						break
					}
					if !val.deadline.IsZero() && val.deadline.Before(t) {
						delete(res.data, key)
					}
					i++
				}
				res.mutex.Unlock()
			case <-res.close:
				return

			}
		}
	}()
	return res
}
