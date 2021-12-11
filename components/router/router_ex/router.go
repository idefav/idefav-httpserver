package router_ex

import (
	"errors"
	"idefav-httpserver/context"
	"idefav-httpserver/models"
	"net/http"
	"strings"
)

var RouterEx = newRouter()

type router struct {
	roots    map[string]*node
	Context  *RouterExContext
	handlers map[string]models.HandlerMapping
}

func (r *router) NewContext(request *http.Request, writer http.ResponseWriter) context.Interface {
	exContext := NewRouterExContext(request, writer)
	r.Context = exContext
	return exContext
}

func (r *router) GetName() string {
	return "AdvancedRouter"
}

func (r *router) Add(handler models.HandlerMapping) {
	RouterEx.addRoute(handler.GetMethod(), handler.GetPath(), handler)
}

func (r *router) Match() (models.HandlerMapping, error) {
	n, params := RouterEx.getRoute(r.Context.Method, r.Context.Path)
	if n != nil {
		r.Context.Params = params
		key := r.Context.Method + "-" + n.pattern
		return r.handlers[key], nil
	}
	return nil, errors.New("handler not match")
}

func newRouter() *router {
	return &router{
		roots:    make(map[string]*node),
		handlers: make(map[string]models.HandlerMapping),
	}
}

func parsePattern(pattern string) []string {
	vs := strings.Split(pattern, "/")

	parts := make([]string, 0)
	for _, item := range vs {
		if item != "" {
			parts = append(parts, item)
			if item[0] == '*' {
				break
			}
		}
	}
	return parts
}

func (r *router) addRoute(method string, pattern string, handler models.HandlerMapping) {
	parts := parsePattern(pattern)

	key := method + "-" + pattern
	_, ok := r.roots[method]
	if !ok {
		r.roots[method] = &node{}
	}
	r.roots[method].insert(pattern, parts, 0)
	r.handlers[key] = handler
}

func (r *router) getRoute(method string, path string) (*node, map[string]string) {
	searchParts := parsePattern(path)
	params := make(map[string]string)
	root, ok := r.roots[method]

	if !ok {
		return nil, nil
	}

	n := root.search(searchParts, 0)

	if n != nil {
		parts := parsePattern(n.pattern)
		for index, part := range parts {
			if part[0] == ':' {
				params[part[1:]] = searchParts[index]
			}
			if part[0] == '*' && len(part) > 1 {
				params[part[1:]] = strings.Join(searchParts[index:], "/")
				break
			}
		}
		return n, params
	}

	return nil, nil
}
