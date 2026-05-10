package stratus

import (
	"log"
	"net/http"
	"os"

	"go.opentelemetry.io/contrib/instrumentation/github.com/aws/aws-lambda-go/otellambda"

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

func (s *Stratus) Post(path string, handler func(http.ResponseWriter, *http.Request)) {
	s.HandleFunc(path, handler)
}

func (s *Stratus) Put(path string, handler func(http.ResponseWriter, *http.Request)) {
	s.HandleFunc(path, handler)
}

func (s *Stratus) Delete(path string, handler func(http.ResponseWriter, *http.Request)) {
	s.HandleFunc(path, handler)
}

func (s *Stratus) Patch(path string, handler func(http.ResponseWriter, *http.Request)) {
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

func (s *Stratus) StratusAdapter(handler Handler) *httpadapter.HandlerAdapter {
	adapter := httpadapter.New(handler)
	return adapter
}

func (s *Stratus) Start() {
	h := s.buildHandler()
	adapter := s.StratusAdapter(h)

	if s.OtelIntegrationEnabled {
		log.Println("Starting Lambda with Otel Integration Enabled...")
		lambda.Start(otellambda.InstrumentHandler(adapter.Proxy))
		return
	}

	log.Println("Starting Lambda with Otel Integration Disabled...")
	lambda.Start(adapter.ProxyWithContext)

}
