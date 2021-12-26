package demo

import (
	"idefav-httpserver/cfg"
	"idefav-httpserver/handler"
	"idefav-httpserver/models"
	"net/http"
	"os"
)

func init() {
	handler.DefaultDispatchHandler.AddHandler(&handler.SimpleHandler{
		Name:   "Headerz",
		Path:   "/headerz",
		Method: http.MethodGet,
		Proc: func(ctx *models.Context) (interface{}, error) {
			for headerName, headerValues := range ctx.Request.Header {
				for _, v := range headerValues {
					ctx.Writer.Header().Add(headerName, v)
				}
			}
			version := os.Getenv(cfg.VERSION)
			if version != "" {
				ctx.Writer.Header().Add(cfg.VERSION, version)
			}
			return "Ok", nil
		},
	})

	handler.DefaultDispatchHandler.AddHandler(&handler.SimpleHandler{
		Name:   "Demo",
		Path:   "/demo",
		Method: http.MethodGet,
		Proc: func(ctx *models.Context) (interface{}, error) {
			return "Demo", nil
		},
	})

	handler.DefaultDispatchHandler.AddHandler(&handler.SimpleHandler{
		Name:   "Panic",
		Path:   "/panic",
		Method: http.MethodGet,
		Proc: func(ctx *models.Context) (interface{}, error) {
			panic("demo panic")
		},
	})
}
