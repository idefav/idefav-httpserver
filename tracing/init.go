package tracing

import (
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
	"github.com/uber/jaeger-client-go/zipkin"
	"io"
	"os"
)

// Init returns an instance of Jaeger Tracer.
func Init(service string) (opentracing.Tracer, io.Closer) {
	os.Setenv("JAEGER_SERVICE_NAME", service)
	cfg, err := config.FromEnv()
	if err != nil {
		panic(fmt.Sprintf("ERROR: failed to read config from env vars: %v\n", err))
	}
	propagator := zipkin.NewZipkinB3HTTPHeaderPropagator()
	tracer, closer, err := cfg.NewTracer(
		config.Logger(jaeger.StdLogger),
		config.Gen128Bit(false),
		config.ZipkinSharedRPCSpan(true),
		config.Extractor(opentracing.HTTPHeaders, propagator),
		config.Injector(opentracing.HTTPHeaders, propagator),
	)
	if err != nil {
		panic(fmt.Sprintf("ERROR: cannot init Jaeger: %v\n", err))
	}
	return tracer, closer
}
