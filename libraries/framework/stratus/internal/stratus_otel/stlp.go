package stratus_otel

import (
	"context"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
	"log"
	"os"
)

type StratusOtelService interface {
	intTracer() *trace.TracerProvider
}

type StratusOtelProvider struct {
	context.Context
	ServiceName string
	Endpoint    string
}

func NewStratusOtelProvider() *StratusOtelProvider {
	return &StratusOtelProvider{
		context.Background(),
		os.Getenv("OTEL_SERVICE_NAME"),
		os.Getenv("OTEL_EXPORTER_OTLP_ENDPOINT"),
	}
}

func (stlp *StratusOtelProvider) InitTracer() *trace.TracerProvider {
	exporter, err := otlptracehttp.New(stlp.Context, otlptracehttp.WithEndpoint(stlp.Endpoint), otlptracehttp.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	tp := trace.NewTracerProvider(
		trace.WithSampler(trace.AlwaysSample()),
		trace.WithBatcher(exporter),
		trace.WithResource(
			resource.NewWithAttributes(
				semconv.SchemaURL,
				semconv.ServiceNameKey.String("stratus-web-framework-serverless-service"))))
	otel.SetTracerProvider(tp)
	return tp

}
