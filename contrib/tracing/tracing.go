package tracing

import (
	"os"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

type TracingConf interface{
	GetEndpoint() string
	GetCustomName() string
}

type tracingConfig struct {
	endpoint string
	serviceName string
	fraction float64
}

type Option func(tr *tracingConfig)



func InitTracer(trConf TracingConf, opts ...Option) {
	var (
		tc = &tracingConfig{
			serviceName: trConf.GetCustomName(),
			endpoint: trConf.GetEndpoint(),
			// rate based on the parent span to 100%
			fraction: 1.0,
		}
		host, _ = os.Hostname()
	)

	for _, opt := range opts {
		opt(tc)
	}

	if tc.endpoint == "" {
		return
	}
	if tc.serviceName == "" {
		tc.serviceName = host
	}

	// Create the Jaeger exporter
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(trConf.GetEndpoint())))
	if err != nil {
		return 
	}
	tp := tracesdk.NewTracerProvider(
		// Set the sampling rate based on the `tracingConfig`
		tracesdk.WithSampler(tracesdk.ParentBased(tracesdk.TraceIDRatioBased(tc.fraction))),
		// Always be sure to batch in production.
		tracesdk.WithBatcher(exp),
		// Record information about this application in an Resource.
		tracesdk.WithResource(resource.NewSchemaless(
				semconv.ServiceNameKey.String(tc.serviceName),
				attribute.String("host", host),
		)),
	)
	otel.SetTracerProvider(tp)
}

func WithEndpoint(endpoint string) Option {
	return func(tr *tracingConfig) {
		tr.endpoint = endpoint
	}
}

func WithServiceName(name string) Option {
	return func(tr *tracingConfig) {
		tr.serviceName = name
	}
}

func WithRatioBased(fraction float64) Option {
	return func(tr *tracingConfig) {
		tr.fraction = fraction
	}
}
