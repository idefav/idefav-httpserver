package httpbin

import (
	"encoding/json"
	"fmt"
	"idefav-httpserver/handler"
	"idefav-httpserver/models"
	"idefav-httpserver/tracing"
	"io/ioutil"
	"net/http"
)

func init() {
	handler.DefaultDispatchHandler.AddHandler(&handler.SimpleHandler{
		Name:   "Headers",
		Path:   "/httpbin/headers",
		Method: http.MethodGet,
		Proc: func(ctx *models.Context) (interface{}, error) {
			asyncReq, err := http.NewRequest(http.MethodGet, "http://httpbin/headers", nil)
			err = tracing.Inject(ctx.Span, asyncReq)
			if err != nil {
				ctx.Span.SetTag("error", true)
				ctx.Span.LogEvent(fmt.Sprintf("Could not inject span context into header: %v", err))
			}
			resp, err := http.DefaultClient.Do(asyncReq)
			if err != nil {
				ctx.Span.SetTag("error", true)
				ctx.Span.LogEvent(fmt.Sprintf("%s /headers error: %v", http.MethodGet, err))
				return nil, err
			}
			if resp.StatusCode != http.StatusOK {
				return nil, fmt.Errorf("access httpbin failed, %v", resp.StatusCode)
			}
			body, err := ioutil.ReadAll(resp.Body)
			var data = make(map[string]interface{})
			err = json.Unmarshal(body, &data)
			return data, err
		},
	})
}
