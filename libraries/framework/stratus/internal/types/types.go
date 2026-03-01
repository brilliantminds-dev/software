package types

import (
	"net/http"
)

type Layer http.Handler
type MiddleLayers []MiddleWare
type OtelIntegrationEnabled bool
type MiddleWare func(http.Handler) http.Handler

const (
	GET    = "GET"
	POST   = "POST"
	PUT    = "PUT"
	PATCH  = "PATCH"
	DELETE = "DELETE"
)
