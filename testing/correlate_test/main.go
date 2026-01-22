package main

import (
	"net/http"

	"github.com/brilliantminds-dev/software/libraries/packages/correlate"
)

func ExampleHandler(w http.ResponseWriter, r *http.Request) {

	log := correlate.NewCorrelateRequest(r, ExampleHandler)
	status := correlate.Status("200 OK")
	desc := correlate.Description("Successful API CAll")

	log.Info(status, desc)

	w.Write([]byte(`{"hello":"world"}`))

}
func main() {
	http.HandleFunc("/", ExampleHandler)
	http.ListenAndServe(":8080", nil)
}
