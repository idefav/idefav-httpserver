package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"idefav-httpserver/cfg"
	"io"
	"log"
	"net/http"
)

const (
	SUCCESS = 0
	FAIL    = 1
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type DispatchHandler struct {
	RequestMapping map[string]*Request
}

func NewDespatchHandler(handlers ...HandlerMapping) *DispatchHandler {
	var requestMapping = map[string]*Request{}
	for _, h := range handlers {
		requestMapping[h.Path()] = NewRequest(h)
	}
	return &DispatchHandler{
		RequestMapping: requestMapping,
	}
}

func (d *DispatchHandler) MatchHandler1(path, method string) (string, *Request, error) {
	if d.RequestMapping == nil {
		return "", nil, fmt.Errorf("request mapping is nil, %w", NotFoundError)
	}
	request, ok := d.RequestMapping[path]
	if !ok {
		return "", nil, fmt.Errorf("path not match, %w", NotFoundError)
	}
	// 判断method
	if method != request.Method {
		return "", nil, fmt.Errorf("method not match, %w", NotFoundError)
	}
	return path, request, nil
}

func (d *DispatchHandler) MatchHandler2(req *http.Request) (string, *Request, error) {
	path := req.RequestURI
	method := req.Method
	return d.MatchHandler1(path, method)
}

func (d DispatchHandler) ErrorHandler(err error) (int, *Request) {
	var code = FAIL
	var message = ""
	message = fmt.Sprintf("error: %v", err)
	if errors.Is(err, NotFoundError) {
		code = http.StatusNotFound
	}
	if errors.Is(err, RuntimeError) {
		code = http.StatusInternalServerError
	}

	_, request, err := d.MatchHandler1(cfg.ERROR_HANDLER, http.MethodGet)
	if err != nil {
		return code, &Request{
			Path:    "/error",
			Method:  http.MethodGet,
			Handler: &ErrorHandler{Code: code, Message: message},
		}
	}
	return code, request
}

type Request struct {
	Path    string
	Method  string
	Handler HandlerMapping
}

func NewRequest(h HandlerMapping) *Request {
	return &Request{
		Path:    h.Name(),
		Method:  h.Method(),
		Handler: h,
	}
}

type HandlerMapping interface {
	Name() string
	Path() string
	Method() string
	Handler(writer http.ResponseWriter, request *http.Request) (int, *Response)
}

func (d *DispatchHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	_, handler, err := d.MatchHandler2(request)
	status := http.StatusOK
	var response = &Response{}
	if handler == nil {
		err = NotFoundError
	}

	if err != nil {
		code, req := d.ErrorHandler(err)
		status = code
		_, response = req.Handler.Handler(writer, request)
	} else {
		status, response = handler.Handler.Handler(writer, request)
	}
	d.accessLog(request, status)
	d.responseOfJson(writer, status, response)
}

func (d *DispatchHandler) accessLog(request *http.Request, status int) {
	log.Printf("%s %s %s - %d",
		request.RemoteAddr,
		request.RequestURI,
		request.Method,
		status,
	)
}

func (d *DispatchHandler) responseOfJson(writer http.ResponseWriter, status int, response *Response) {
	resp, _ := json.Marshal(response)
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(status)
	_, _ = io.WriteString(writer, string(resp))
}
