package main

import (
	"idefav-httpserver/cfg"
	"idefav-httpserver/handler"
	"idefav-httpserver/handler/headerz"
	"idefav-httpserver/handler/healthz"
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func main() {
	var serverConfig = cfg.SetUp()
	mux := http.NewServeMux()

	// health
	health := healthz.Health{Indicators: map[string]healthz.HealthIndicator{"Default": &healthz.ServerHealth{}}}

	// headerz
	headerz := headerz.HeaderHandler{}

	mux.Handle("/", handler.NewDespatchHandler(&health, &headerz))
	serv := &http.Server{
		Addr:              serverConfig.Address,
		ReadHeaderTimeout: 1000 * time.Millisecond,
		IdleTimeout:       1800 * 1000 * time.Millisecond,
		ReadTimeout:       1000 * time.Millisecond,
		Handler:           mux,
	}
	err := serv.ListenAndServe()
	if err != nil {
		log.Fatal("start server failed!")
	}
}
