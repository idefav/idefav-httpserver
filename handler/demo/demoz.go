package demo

import (
	"idefav-httpserver/handler"
	"net/http"
)

type DemoHandler struct {
}

func (d *DemoHandler) Name() string {
	return "Demo"
}

func (d *DemoHandler) Path() string {
	return "/demo"
}

func (d *DemoHandler) Method() string {
	return http.MethodGet
}

func (d *DemoHandler) Handler(writer http.ResponseWriter, request *http.Request) (int, *handler.Response) {
	return http.StatusOK, &handler.Response{
		Code:    handler.SUCCESS,
		Message: "Demo",
	}
}
