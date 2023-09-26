package errs

import "errors"

var (
	errKeyNotFound = errors.New("cache 不存在")
)
