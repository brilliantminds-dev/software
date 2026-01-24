package main

import (
	"context"
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

func (s *Stratus) Get(path string, handler func(http.ResponseWriter, *http.Request)) {
	s.HandleFunc(path, handler)

}

func (s *Stratus) Use(m types.MiddleWare) {
	s.MiddleLayers = append(s.MiddleLayers, m)
}

func (s *Stratus) buildHandler() Handler {
	var h Handler = s.StratusInterface.(Handler)

	// Apply middleware (inside → outside)
	for i := len(s.MiddleLayers) - 1; i >= 0; i-- {
		h = s.MiddleLayers[i](h)
	}

	// Wrap with OTEL last (outermost)
	if s.OtelIntegrationEnabled {
		h = otelhttp.NewHandler(h, os.Getenv("OTEL_SERVICE_NAME"))
	}

	return h
}

func NewStratus() *Stratus {

	return &Stratus{
		StratusInterface: http.NewServeMux(),
	}
}

func (s *Stratus) Start() {
	h := s.buildHandler()

	adapter := httpadapter.New(h)

	lambda.StartWithContext(context.Background(), adapter.ProxyWithContext)

}
