package models

import (
	"github.com/opentracing/opentracing-go"
	"net/http"
)

type Context struct {
	Writer  http.ResponseWriter
	Request *http.Request
	Span    opentracing.Span
}

func NewContext(req *http.Request, writer http.ResponseWriter, span opentracing.Span) Context {
	return Context{
		Writer:  writer,
		Request: req,
		Span:    span,
	}
}
