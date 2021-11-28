package models

import "net/http"

type HandlerMapping interface {
	GetName() string
	GetPath() string
	GetMethod() string
	Handler(writer http.ResponseWriter, request *http.Request) (interface{}, error)
}
