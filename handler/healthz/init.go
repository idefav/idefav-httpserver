package healthz

import "idefav-httpserver/handler"

func init() {
	health := Health{Indicators: map[string]HealthIndicator{"Default": &ServerHealth{}}}
	handler.DefaultDispatchHandler.AddHandler(&health)
}
