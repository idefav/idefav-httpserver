package headerz

import "idefav-httpserver/handler"

func init() {
	headerz := HeaderHandler{}
	handler.DefaultDispatchHandler.AddHandler(&headerz)
}
