package router

import (
	"errors"
	"idefav-httpserver/context"
	"idefav-httpserver/models"
	"net/http"
)

const (
	DEFAULT_ROUTER = "Default"
)

type Interface interface {
	GetName() string
	Add(handler models.HandlerMapping)
	NewContext(request *http.Request, writer http.ResponseWriter) context.Interface
	Match() (models.HandlerMapping, error)
}

type RouterMap map[string]Interface

var Routers = RouterMap{}

func AddRouter(r Interface) {
	Routers[r.GetName()] = r
}

func GetRouter(name string) (Interface, error) {
	if name == "" {
		name = DEFAULT_ROUTER
	}
	component, ok := Routers[name]
	if !ok {
		defaultRouter, ok := Routers[DEFAULT_ROUTER]
		if !ok {
			return nil, errors.New("router not found")
		}
		return defaultRouter, nil
	}
	return component, nil
}
