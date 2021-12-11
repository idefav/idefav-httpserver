package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"idefav-httpserver/cfg"
	"idefav-httpserver/common"
	"idefav-httpserver/components/interceptor"
	"idefav-httpserver/components/router"
	"idefav-httpserver/models"
	"io"
	"log"
	"net/http"
)

const (
	SUCCESS = 0
	FAIL    = 1
)

type Handler interface {
	NewHandler() *models.HandlerMapping
}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

var DefaultDispatchHandler = NewDispatchHandler(nil)

type DispatchHandler struct {
	config *cfg.ServerConfig
}

func (d *DispatchHandler) GetRouter() router.Interface {
	getRouter, err := router.GetRouter(router.DEFAULT_ROUTER)
	if err != nil {
		log.Fatalf("no router found")
	}
	return getRouter
}

type Interface interface {
	AddHandler(handlers ...models.HandlerMapping)
	GetRouter() router.Interface
	Match(req *http.Request) (models.HandlerMapping, error)
}

func (d *DispatchHandler) AddHandler(handlers ...models.HandlerMapping) {
	getRouter := d.GetRouter()
	for _, h := range handlers {
		getRouter.Add(h)
	}
}

func NewDispatchHandler(config *cfg.ServerConfig, handlers ...models.HandlerMapping) *DispatchHandler {
	dispatchHandler := DispatchHandler{
		config: config,
	}
	getRouter := dispatchHandler.GetRouter()
	for _, h := range handlers {
		getRouter.Add(h)
	}
	return &dispatchHandler
}

func (d *DispatchHandler) Match(req *http.Request) (models.HandlerMapping, error) {
	getRouter := d.GetRouter()
	return getRouter.Match(req)
}

func (d DispatchHandler) ErrorHandler(err error) (int, *Response) {
	var code = FAIL
	var message = "unknown error"
	message = fmt.Sprintf("error: %v", err)
	if errors.Is(err, common.NotFoundError) {
		code = http.StatusNotFound
	} else if errors.Is(err, common.RuntimeError) {
		code = http.StatusInternalServerError
	}

	return code, &Response{
		Code:    code,
		Message: message,
	}
}

func (d *DispatchHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	defer func() {
		r := recover()
		if r != nil {
			var err error
			switch t := r.(type) {
			case string:
				err = errors.New(t)
			case error:
				err = t
			default:
				err = errors.New("unknown error")
			}
			d.responseOfJson(writer, http.StatusInternalServerError, &Response{
				Code:    FAIL,
				Message: err.Error(),
			})
		}
	}()
	h, err := d.Match(request)
	status := http.StatusOK
	if d.config.AccessLog {
		defer d.accessLog(request, status)
	}

	var response = &Response{}
	if h == nil {
		err = common.NotFoundError
	}

	if err != nil {
		code, resp := d.ErrorHandler(err)
		status = code
		response = resp

	} else {
		err := interceptor.Run(writer, request)
		if err != nil {
			code, resp := d.ErrorHandler(err)
			status = code
			response = resp
		} else {
			data, err2 := h.Handler(writer, request)
			if err2 != nil {
				code, resp := d.ErrorHandler(err)
				status = code
				response = resp
			} else {
				response.Data = data
				response.Code = SUCCESS
			}
		}

	}

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

func SetUpDispatchHandler(config *cfg.ServerConfig) *DispatchHandler {
	DefaultDispatchHandler.config = config
	return DefaultDispatchHandler
}
