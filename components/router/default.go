package router

import (
	"fmt"
	"idefav-httpserver/common"
	"idefav-httpserver/context"
	"idefav-httpserver/models"
	"net/http"
)

type RequestMapping map[string]models.HandlerMapping

type DefaultRouter struct {
	Name           string
	Context        *context.DefaultContext
	RequestMapping RequestMapping
}

func (d DefaultRouter) NewContext(request *http.Request, writer http.ResponseWriter) context.Interface {
	defaultContext := context.NewDefaultContext(request, writer)
	d.Context = defaultContext
	return defaultContext
}

func (d DefaultRouter) Add(mapping models.HandlerMapping) {
	if d.RequestMapping == nil {
		d.RequestMapping = RequestMapping{}
	}
	d.RequestMapping[mapping.GetPath()] = mapping
}

func (d DefaultRouter) GetName() string {
	return d.Name
}

func (d DefaultRouter) Match() (models.HandlerMapping, error) {

	path := d.Context.Path
	method := d.Context.Method

	if d.RequestMapping == nil {
		return nil, fmt.Errorf("request mapping is nil, %w", common.NotFoundError)
	}
	handlerMapping, ok := d.RequestMapping[path]
	if !ok {
		return nil, fmt.Errorf("path not match, %w", common.NotFoundError)
	}
	// 判断method
	if method != handlerMapping.GetMethod() {
		return nil, fmt.Errorf("method not match, %w", common.NotFoundError)
	}
	return handlerMapping, nil
}

func init() {
	AddRouter(&DefaultRouter{
		Name:           DEFAULT_ROUTER,
		RequestMapping: RequestMapping{},
	})
}
