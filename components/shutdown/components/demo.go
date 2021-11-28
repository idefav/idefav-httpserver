package components

import (
	"idefav-httpserver/components/shutdown"
	"log"
	"math"
	"time"
)

func init() {
	shutdown.Add(&shutdown.DefaultGracefulShutdownComponent{
		Name:  "Demo",
		Order: math.MinInt,
		Proc: func() {
			log.Println("Cleaning...")
			time.Sleep(2 * time.Second)
			log.Println("Clean Done!")
		},
	})
}
