package models

type HandlerMapping interface {
	GetName() string
	GetPath() string
	GetMethod() string
	Handler(ctx *Context) (interface{}, error)
}
