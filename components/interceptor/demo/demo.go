package demo

import (
	"idefav-httpserver/components/interceptor"
	"math/rand"
	"net/http"
	"time"
)

func init() {
	interceptor.Add("RandomRt", 1, func(writer http.ResponseWriter, request *http.Request) error {
		n := rand.Intn(1000)
		time.Sleep(time.Duration(n) * time.Millisecond)
		return nil
	})
}
