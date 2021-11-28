package components

import (
	"idefav-httpserver/components/warmup"
	"log"
)

func init() {
	warmup.Add("Demo", func() error {
		log.Println("demo warmup!")
		return nil
	})
}
