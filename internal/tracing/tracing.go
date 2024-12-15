package tracing

import (
	jaegerExporter "go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.12.0"
)

var (
	// Store a global trace provider variable to clear it before closing
	Tracer *tracesdk.TracerProvider
)

func NewTracer(url, name string) error {
	// Create the Jaeger exporter
	exp, err := jaegerExporter.New(jaegerExporter.WithCollectorEndpoint(jaegerExporter.WithEndpoint(url)))
	if err != nil {
		return err
	}
	Tracer = tracesdk.NewTracerProvider(
		// Always be sure to batch in production.
		tracesdk.WithBatcher(exp),
		// Record information about this application in a Resource.
		tracesdk.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(name),
		)),
	)
	return nil
}
