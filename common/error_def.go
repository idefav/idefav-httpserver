package common

import "errors"

var (
	NotFoundError = errors.New("not found")
	RuntimeError  = errors.New("runtime error")
)
