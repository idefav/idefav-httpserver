package headerz

import (
	"idefav-httpserver/cfg"
	"idefav-httpserver/handler"
	"io"
	"net/http"
	"os"
)

type HeaderHandler struct {
}

func (h *HeaderHandler) Name() string {
	return "Headerz"
}

func (h *HeaderHandler) Path() string {
	return "/headerz"
}

func (h *HeaderHandler) Method() string {
	return http.MethodGet
}

func (h *HeaderHandler) Handler(writer http.ResponseWriter, request *http.Request) (int, *handler.Response) {
	for headerName, headerValues := range request.Header {
		for _, v := range headerValues {
			writer.Header().Add(headerName, v)
		}
	}
	version := os.Getenv(cfg.VERSION)
	if version != "" {
		writer.Header().Add(cfg.VERSION, version)
	}
	return http.StatusOK, &handler.Response{
		Code:    handler.SUCCESS,
		Message: cfg.OK,
	}
}

func (h *HeaderHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	for headerName, headerValues := range request.Header {
		for _, v := range headerValues {
			writer.Header().Add(headerName, v)
		}
	}
	version := os.Getenv(cfg.VERSION)
	if version != "" {
		writer.Header().Add(cfg.VERSION, version)
	}

	_, _ = io.WriteString(writer, cfg.OK)
}
