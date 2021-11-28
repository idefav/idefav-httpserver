package healthz

import "idefav-httpserver/handler"

var healthHandler = &Health{
	name:       "HEALTH",
	path:       "/healthz",
	Indicators: map[string]HealthIndicator{},
}

func AddHealthIndicator(name string, indicator HealthIndicator) {
	healthHandler.Indicators[name] = indicator
}

func init() {
	handler.DefaultDispatchHandler.AddHandler(healthHandler)
}
