package router_ex

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type RouterExContext struct {
	request    *http.Request
	writer     http.ResponseWriter
	Path       string
	Method     string
	Params     map[string]string
	StatusCode int
}

func NewRouterExContext(request *http.Request, writer http.ResponseWriter) *RouterExContext {
	return &RouterExContext{request: request, writer: writer, Path: request.RequestURI, Method: request.Method}
}

func (c RouterExContext) PostForm(key string) string {
	return c.request.PostForm.Get(key)
}

func (c RouterExContext) RequestBody(body interface{}) (interface{}, error) {
	bytes, err := ioutil.ReadAll(c.request.Body)
	err = json.Unmarshal(bytes, body)
	return body, err
}

func (c RouterExContext) Query(key string) string {
	return c.request.URL.Query().Get(key)
}

func (c RouterExContext) Status(code int) {
	c.StatusCode = code
}

func (c RouterExContext) Header(key string) string {
	return c.request.Header.Get(key)
}

func (c RouterExContext) SetHeader(key string, value string) {
	c.request.Header.Set(key, value)
}

func (c RouterExContext) Success(code string, msg string, data interface{}) {

}

func (c RouterExContext) Fail(code string, msg string) {
	panic("implement me")
}
