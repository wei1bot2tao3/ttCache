package OSCache

import "time"

type Cache interface {
	Get(key string) (interface{}, bool)
	Set(key string, value interface{})
}

type item struct {
	val      any
	deadline time.Time
}
