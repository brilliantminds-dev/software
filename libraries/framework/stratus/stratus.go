package stratus

import (
	"context"
	"fmt"
	"github.com/brilliantminds-dev/software/libraries/framework/stratus/internal/stratus_otel"
	"go.opentelemetry.io/contrib/instrumentation/github.com/aws/aws-lambda-go/otellambda"
	"go.opentelemetry.io/otel"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter" // Added missing import
	"github.com/brilliantminds-dev/software/libraries/framework/stratus/internal/types"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

type Stratus struct {
	StratusInterface
	types.MiddleLayers
	types.OtelIntegrationEnabled
}

type Router interface {
	HandleFunc(string, func(http.ResponseWriter, *http.Request))
}
type Handler interface {
	ServeHTTP(http.ResponseWriter, *http.Request)
}

type StratusInterface interface {
	Router
}

func (s *Stratus) StratusRouter() StratusInterface {
	return s.StratusInterface.(*http.ServeMux)
}

func (s *Stratus) StratusResource(methods []string, path string, handler func(http.ResponseWriter, *http.Request)) {
	tracer := otel.Tracer("my-lambda-tracer")
	// Start a span for Stratus resource
	_, span := tracer.Start(context.Background(), "handler-span")
	event := fmt.Sprintf("StratusResource called for path : %s", path)
	span.AddEvent(event)
	defer span.End()

	s.HandleFunc(path, handler)

}

func (s *Stratus) Use(m types.MiddleWare) {
	s.MiddleLayers = append(s.MiddleLayers, m)
}

func (s *Stratus) buildHandler() Handler {
	var h = s.StratusInterface.(Handler)

	// Apply middleware (inside → outside)
	for i := len(s.MiddleLayers) - 1; i >= 0; i-- {
		h = s.MiddleLayers[i](h)
	}

	// Wrap with OTEL last (outermost)
	if s.OtelIntegrationEnabled {
		svc := os.Getenv("OTEL_SERVICE_NAME")
		if svc == "" || &svc == nil {
			log.Fatal("OTEL_SERVICE_NAME env var is not set. please set export or set variable")

		}

		h = otelhttp.NewHandler(h, svc)

	}

	return h
}

func NewStratus() *Stratus {

	return &Stratus{
		StratusInterface: &http.ServeMux{},
	}
}

func (s *Stratus) Start() {
	h := s.buildHandler()

	adapter := httpadapter.New(h)
	if s.OtelIntegrationEnabled {
		log.Println("Starting Lambda with Otel Integration Enabled...")
		stlp := stratus_otel.NewStratusOtelProvider(os.Getenv("OTEL_SERVICE_NAME"), "0.0.0.0:4318") // will use live one from monoscope soon
		stp := stlp.InitTracer()
		defer func() {
			if err := stp.Shutdown(context.Background()); err != nil {
				log.Printf("Error shutting down Stratus Web Framework Serverless Service OTEL tracer: %v", err)
			}
		}()

		lambda.Start(otellambda.InstrumentHandler(adapter.Proxy,
			otellambda.WithTracerProvider(stp)))
		return
	} else {
		log.Println("Starting Lambda with Otel Integration Disabled...")
		lambda.StartWithContext(context.Background(), adapter.ProxyWithContext)

	}

}
