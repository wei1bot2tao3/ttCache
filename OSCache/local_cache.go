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
}

func NewBuildInMapCache() *BuildInMapCache {
	res := &BuildInMapCache{
		data: make(map[string]*item),
	}
	return res
}

// Set 建立一个
func (b *BuildInMapCache) Set(key string, value interface{}) {

}

func (b *BuildInMapCache) Get(key string) (interface{}, bool) {
	//TODO implement me
	panic("implement me")
}
