package handler

import (
	"errors"
	"idefav-httpserver/cfg"
	"net/http"
)

var (
	NotFoundError = errors.New("not found")
	RuntimeError  = errors.New("runtime error")
)

type ErrorHandler struct {
	Code    int
	Message string
}

func (e *ErrorHandler) Path() string {
	return "/error"
}

func (e *ErrorHandler) Method() string {
	return http.MethodGet
}

func (e *ErrorHandler) Name() string {
	return cfg.ERROR_HANDLER
}

func (e *ErrorHandler) Handler(writer http.ResponseWriter, request *http.Request) (int, *Response) {
	return http.StatusInternalServerError, &Response{
		Code:    e.Code,
		Message: e.Message,
	}
}
