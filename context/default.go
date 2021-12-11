package context

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type DefaultContext struct {
	request    *http.Request
	writer     http.ResponseWriter
	Path       string
	Method     string
	StatusCode int
}

func NewDefaultContext(request *http.Request, writer http.ResponseWriter) *DefaultContext {
	return &DefaultContext{request: request, writer: writer, Path: request.RequestURI, Method: request.Method}
}

func (c DefaultContext) PostForm(key string) string {
	return c.request.PostForm.Get(key)
}

func (c DefaultContext) RequestBody(body interface{}) (interface{}, error) {
	bytes, err := ioutil.ReadAll(c.request.Body)
	err = json.Unmarshal(bytes, body)
	return body, err
}

func (c DefaultContext) Query(key string) string {
	return c.request.URL.Query().Get(key)
}

func (c DefaultContext) Status(code int) {
	c.StatusCode = code
}

func (c DefaultContext) Header(key string) string {
	return c.request.Header.Get(key)
}

func (c DefaultContext) SetHeader(key string, value string) {
	c.request.Header.Set(key, value)
}

func (c DefaultContext) Success(code string, msg string, data interface{}) {

}

func (c DefaultContext) Fail(code string, msg string) {
	panic("implement me")
}
