package stratus_otel

import (
	"context"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
	"log"
)

type StratusOtelService interface {
	intTracer() *trace.TracerProvider
}

type StratusOtelProvider struct {
	context.Context
	ServiceName string
	Endpoint    string
}

func NewStratusOtelProvider(svc string, endpoint string) *StratusOtelProvider {
	return &StratusOtelProvider{
		context.Background(),
		svc,
		endpoint,
	}
}

func (stlp *StratusOtelProvider) InitTracer() *trace.TracerProvider {
	exporter, err := otlptracehttp.New(context.Background(), otlptracehttp.WithEndpoint("127.0.0.1:4318"), otlptracehttp.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	tp := trace.NewTracerProvider(
		trace.WithBatcher(exporter), trace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL, semconv.ServiceNameKey.String("stratus-web-framework-serveless-service"))))
	otel.SetTracerProvider(tp)
	return tp

}
