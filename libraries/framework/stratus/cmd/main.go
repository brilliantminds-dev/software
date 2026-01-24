package main

import (
	"net/http"

	s "github.com/brilliantminds-dev/software/libraries/framework/stratus"
)

func SampleHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, Stratus!"))
}

func main() {

	stratus := s.NewStratus()
	stratus.OtelIntegrationEnabled = false

	stratus.Get("/hello", SampleHandler)

	stratus.Start()

}
