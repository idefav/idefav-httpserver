package handler

import (
	"idefav-httpserver/models"
)

type SimpleHandler struct {
	Name   string
	Path   string
	Method string
	Proc   func(ctx *models.Context) (interface{}, error)
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

func (h *SimpleHandler) Handler(ctx *models.Context) (interface{}, error) {
	return h.Proc(ctx)
}
