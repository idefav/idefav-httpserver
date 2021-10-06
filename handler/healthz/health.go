package healthz

import (
	"idefav-httpserver/cfg"
	"idefav-httpserver/handler"
	"io"
	"net/http"
)

const Name = "HEALTH"
const Path = "/healthz"

type Health struct {
	Indicators map[string]HealthIndicator
}

func (h *Health) Path() string {
	return Path
}

func (h *Health) Method() string {
	return http.MethodGet
}

func (h *Health) Name() string {
	return Name
}

func (h *Health) Handler(writer http.ResponseWriter, request *http.Request) (int, *handler.Response) {
	var res = cfg.HEALTH
	var code = http.StatusOK
	if !h.checkHealth() {
		res = cfg.UNHEALTHY
		code = http.StatusServiceUnavailable
	}
	return code, &handler.Response{
		Code:    code,
		Message: res,
	}
}

func (h *Health) checkHealth() bool {
	if h.Indicators != nil {
		for _, indicator := range h.Indicators {
			if !indicator.health() {
				return false
			}
		}
	}
	return true
}

type HealthIndicator interface {
	health() bool
}

func (h *Health) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	var res = cfg.HEALTH
	var code = http.StatusOK
	if !h.checkHealth() {
		res = cfg.UNHEALTHY
		code = http.StatusServiceUnavailable
	}
	writer.WriteHeader(code)
	_, _ = io.WriteString(writer, res)
}
