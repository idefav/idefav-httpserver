package main

import (
	"idefav-httpserver/cfg"
	"idefav-httpserver/handler"
	_ "idefav-httpserver/handler/demo"
	_ "idefav-httpserver/handler/headerz"
	_ "idefav-httpserver/handler/healthz"
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func main() {
	var serverConfig = cfg.SetUp()
	mux := http.DefaultServeMux

	mux.Handle("/", handler.DefaultDispatchHandler)
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
