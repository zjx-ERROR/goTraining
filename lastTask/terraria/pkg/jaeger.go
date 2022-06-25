package pkg

import (
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.10.0"
)

func TracerProvider(url string) (*tracesdk.TracerProvider, error) {
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(url))
	if err != nil {
		return nil, err
	}
	tp := tracesdk.NewTracerProvider(
		tracesdk.WithBatcher(exp),
		tracesdk.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String("battle.trace"),
			attribute.String("exporter", "jaeger"),
			attribute.Float64("float", 312.23),
		)),
	)
	return tp, nil
}
