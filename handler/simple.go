package handler

import (
	"net/http"
)

type HandlerFunc func(writer http.ResponseWriter, request *http.Request) (interface{}, error)

type SimpleHandler struct {
	Name   string
	Path   string
	Method string
	Proc   HandlerFunc
}

func Get(path string, handler HandlerFunc) {
	AddHandler(http.MethodGet, path, handler)
}

func Post(path string, handler HandlerFunc) {
	AddHandler(http.MethodPost, path, handler)
}

func Put(path string, handler HandlerFunc) {
	AddHandler(http.MethodPut, path, handler)
}

func Delete(path string, handler HandlerFunc) {
	AddHandler(http.MethodDelete, path, handler)
}

func Options(path string, handler HandlerFunc) {
	AddHandler(http.MethodOptions, path, handler)
}

func Head(path string, handler HandlerFunc) {
	AddHandler(http.MethodHead, path, handler)
}

func Connect(path string, handler HandlerFunc) {
	AddHandler(http.MethodConnect, path, handler)
}

func Patch(path string, handler HandlerFunc) {
	AddHandler(http.MethodPatch, path, handler)
}

func Trace(path string, handler HandlerFunc) {
	AddHandler(http.MethodTrace, path, handler)
}

func AddHandler(method string, path string, handler HandlerFunc) {
	DefaultDispatchHandler.AddHandler(&SimpleHandler{
		Name:   method + ":" + path,
		Path:   path,
		Method: method,
		Proc:   handler,
	})
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
