package types

type Layer http.Handler
type MiddleLayers []MiddleWare

type Get func(string, http.ResponseWriter, *http.Request)
