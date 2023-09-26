package errs

import "errors"

var (
	ErrKeyNotFound = errors.New("cache 不存在")
	ErrKeyExists   = errors.New("key 已经存在，重复")
)

func NewErrKeyExists(key string) error {
	return errors.New("key:" + key + "已经存在，重复")
}
