package server

import (
	"context"
	_ "idefav-httpserver/auto"
	"idefav-httpserver/cfg"
	"idefav-httpserver/components/shutdown"
	"idefav-httpserver/components/warmup"
	"idefav-httpserver/handler"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Start() {
	var serverConfig = cfg.SetUp()
	mux := http.DefaultServeMux
	mux.Handle("/", handler.SetUpDispatchHandler(serverConfig))
	serv := &http.Server{
		Addr:              serverConfig.Address,
		ReadHeaderTimeout: 1000 * time.Millisecond,
		IdleTimeout:       1800 * 1000 * time.Millisecond,
		ReadTimeout:       1000 * time.Millisecond,
		Handler:           mux,
	}
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if !serverConfig.Warmup {
			log.Println("Warmup not enable!")
		} else {
			log.Println("Warmup now!")
			warmup.RunWarmup()
			log.Println("Warmup completed!")
		}

		log.Println("Server listen at:" + serverConfig.Address)
		err := serv.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatal("start server failed!")
		}
	}()
	log.Print("Server Started")
	<-done
	log.Print("Server Stopped")

	doShutdown(serverConfig, serv)
}

func doShutdown(serverConfig *cfg.ServerConfig, serv *http.Server) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(serverConfig.GracefulShutdownWaitTimeMs)*time.Millisecond)
	defer func() {
		if serverConfig.GracefulShutdown {
			log.Println("Server is shutting down and begin cleaning!")
			shutdown.RunShutdownClean()
			log.Println("Server is down, and clean completed!")
		}
		cancel()
	}()

	if err := serv.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown Failed:%+v", err)
	}
	log.Print("Server Exited Properly")
}
