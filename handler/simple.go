package handler

import (
	"net/http"
)

type SimpleHandler struct {
	Name   string
	Path   string
	Method string
	Proc   func(writer http.ResponseWriter, request *http.Request) (interface{}, error)
}

func (h *SimpleHandler) GetName() string {
	return h.Name
}

func (h *SimpleHandler) GetPath() string {
	return h.Path
}

func (h *SimpleHandler) GetMethod() string {
	return h.Method
}

func (h *SimpleHandler) Handler(writer http.ResponseWriter, request *http.Request) (interface{}, error) {
	return h.Proc(writer, request)
}
