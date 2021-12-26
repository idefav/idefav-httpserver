package healthz

import (
	"errors"
	"idefav-httpserver/cfg"
	"idefav-httpserver/models"
	"net/http"
)

type Health struct {
	name       string
	path       string
	Indicators map[string]HealthIndicator
}

func (h *Health) GetPath() string {
	return h.path
}

func (h *Health) GetMethod() string {
	return http.MethodGet
}

func (h *Health) GetName() string {
	return h.name
}

func (h *Health) Handler(ctx *models.Context) (interface{}, error) {
	var res = cfg.HEALTH
	if !h.checkHealth() {
		res = cfg.UNHEALTHY
		return nil, errors.New(res)
	}
	return res, nil
}

func (h *Health) checkHealth() bool {
	if h.Indicators != nil {
		for _, indicator := range h.Indicators {
			if !indicator.Health() {
				return false
			}
		}
	}
	return true
}

type HealthIndicator interface {
	Health() bool
}
