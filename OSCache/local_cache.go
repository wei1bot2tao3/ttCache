package OSCache

import (
	"sync"
)

// BuildInMapCache 一个缓存Map 类型的缓存
type BuildInMapCache struct {
	data      map[string]*item
	mutex     sync.RWMutex
	close     chan struct{}
	onEvicted func(key string, val any)
	//onEvicted  []func(key string,val any)
	//为啥不允许注册多个
}

func NewBuildInMapCache() *BuildInMapCache {
	res := &BuildInMapCache{
		data: make(map[string]*item),
	}
	return res
}

func (b *BuildInMapCache) Set(key string, value interface{}) {
	//TODO implement me
	panic("implement me")
}

func (b *BuildInMapCache) Get(key string) (interface{}, bool) {
	//TODO implement me
	panic("implement me")
}
