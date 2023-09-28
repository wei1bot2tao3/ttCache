package errs

import "errors"

var (
	ErrKeyNotFound = errors.New("cache 不存在")
	ErrKeyExists   = errors.New("key 已经存在，重复")
	ErrCacheCloes  = errors.New("重复关闭")
)

func NewErrKeyExists(key string) error {
	return errors.New("key:" + key + "已经存在，重复")
}

func NewErrNotfound(key string) error {
	return errors.New("key:" + key + "不存在")
}

func NewExpiredKeyError(key string) error {
	return errors.New("key:" + key + "过期删除")
}
