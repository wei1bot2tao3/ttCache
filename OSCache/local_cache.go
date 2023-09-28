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

// NewBuildInMapCache 返回一个实例
func NewBuildInMapCache(size int) *BuildInMapCache {
	res := &BuildInMapCache{
		data: make(map[string]*item, size),
	}
	return res
}

// Set 设置过期时间 每设置一个key就开一个goroutine来监控
func (b *BuildInMapCache) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
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

func (b *BuildInMapCache) Get(ctx context.Context, key string) (interface{}, error) {
	b.mutex.RLock()
	defer b.mutex.RUnlock()

	res, ok := b.data[key]
	if !ok {
		return nil, errs.NewErrNotfound(key)
	}

	return res, nil
}
func (b *BuildInMapCache) delete(key string) (interface{}, error) {
	itm, ok := b.data[key]
	if !ok {
		return nil, errs.NewErrNotfound(key)
	}
	delete(b.data, key)

	return itm, nil
}
func (b *BuildInMapCache) Delete(ctx context.Context, key string) (interface{}, error) {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	_, ok := b.data[key]
	if !ok {
		return nil, errs.NewErrNotfound(key)
	}

	itm, err := b.delete(key)

	return itm, err
}

func (b *BuildInMapCache) Close() error {

	select {
	case b.close <- struct{}{}:

	default:
		return errs.ErrCacheCloes
	}
	return nil

}

// CacheOneGo  一个goroutine来轮训过期时间
type CacheOneGo struct {
	Cache Cache
}

// NewBuildInMapCacheOneGo 开启一个go去轮训 所有时间过期的key 然后close后会过期
func NewBuildInMapCacheOneGo(res *BuildInMapCache, interval time.Duration) *CacheOneGo {

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

	return &CacheOneGo{
		Cache: res,
	}
}

func (c *CacheOneGo) Get(ctx context.Context, key string) (interface{}, error) {
	return c.Cache.Get(ctx, key)
}

func (c *CacheOneGo) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	return c.Cache.Set(ctx, key, value, expiration)
}

func (c *CacheOneGo) Delete(ctx context.Context, key string) (interface{}, error) {
	return c.Cache.Delete(ctx, key)
}

// CacheGos 一个key一个goroutine
type CacheGos struct {
	Cache Cache
}

func NewBuildInMapCacheGos(res *BuildInMapCache) *CacheGos {
	return &CacheGos{
		Cache: res,
	}
}

func (c *CacheGos) Get(ctx context.Context, key string) (interface{}, error) {

	return c.Cache.Get(ctx, key)
}

func (c *CacheGos) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	err := c.Cache.Set(ctx, key, value, expiration)
	if err != nil {

		return err
	}
	if expiration > 0 {
		go func(ctx context.Context, key string, expiration time.Duration) {
			time.AfterFunc(expiration, func() {
				c.Delete(ctx, key)
			})
		}(ctx, key, expiration)
	}
	return err
}

func (c *CacheGos) Delete(ctx context.Context, key string) (interface{}, error) {
	val, err := c.Cache.Delete(ctx, key)
	return val, err
}

// CacheNoGo 懒惰删除
type CacheNoGo struct {
	Cache Cache
}

func NewBuildInMapCacheNoGo(res *BuildInMapCache) *CacheNoGo {
	return &CacheNoGo{
		Cache: res,
	}
}

func (c *CacheNoGo) Get(ctx context.Context, key string) (interface{}, error) {
	Map, err := c.Get(ctx, key)
	if err != nil {
		return nil, err
	}
	node := Map.(*item)
	if !node.deadline.IsZero() && node.deadline.Before(time.Now()) {
		_, err = c.Cache.Delete(ctx, key)
		if err != nil {
			return nil, err
		}
		return nil, errs.NewErrKeyExists(key)
	}
	return Map, err

}

func (c *CacheNoGo) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	return c.Cache.Set(ctx, key, value, expiration)
}

func (c *CacheNoGo) Delete(ctx context.Context, key string) (interface{}, error) {
	return c.Cache.Delete(ctx, key)
}
