package demo

import (
	"idefav-httpserver/cfg"
	"idefav-httpserver/handler"
	"net/http"
	"os"
)

func init() {
	handler.Get("/headerz", func(writer http.ResponseWriter, request *http.Request) (interface{}, error) {
		for headerName, headerValues := range request.Header {
			for _, v := range headerValues {
				writer.Header().Add(headerName, v)
			}
		}
		version := os.Getenv(cfg.VERSION)
		if version != "" {
			writer.Header().Add(cfg.VERSION, version)
		}
		return "Ok", nil
	})

	handler.Get("/demo", func(writer http.ResponseWriter, request *http.Request) (interface{}, error) {
		return "Demo", nil
	})

	handler.Get("/panic", func(writer http.ResponseWriter, request *http.Request) (interface{}, error) {
		panic("demo panic")
	})
}
