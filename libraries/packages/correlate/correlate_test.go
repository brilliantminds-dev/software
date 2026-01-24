package correlate_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	c "github.com/brilliantminds-dev/software/libraries/packages/correlate"
)

func TestCorrelate(t *testing.T) {

	tdt := []struct {
		name       string
		handler    http.HandlerFunc // interface for http handler functions
		statusCode int
	}{{
		name:       "Info Handler Test",
		handler:    CorrelateInfoHandler,
		statusCode: 200,
	}, {
		name:       "Except Handler Test",
		handler:    CorrelateExceptHandler,
		statusCode: 400,
	}, {
		name:       "Success Handler Test",
		handler:    CorrelateSuccessHandler,
		statusCode: 200,
	}, {
		name:       "Fatal Handler Test",
		handler:    CorrelateFatalHandler,
		statusCode: 500,
	}}

	for _, tt := range tdt {
		t.Run(tt.name, func(t *testing.T) {

			r := httptest.NewRequest("GET", "/correlate", nil)
			w := httptest.NewRecorder()

			r.Header.Set("User-Agent", "Go-http-client/1.1")

			tt.handler.ServeHTTP(w, r)
			fmt.Println(w.Result().StatusCode)
			if w.Result().StatusCode != tt.statusCode {
				t.Errorf("Expected status code %d, got %d", tt.statusCode, w.Result().StatusCode)
			}
		})
	}
}

func CorrelateInfoHandler(w http.ResponseWriter, r *http.Request) {
	correlate := c.NewCorrelateRequest(r, CorrelateInfoHandler)
	correlate.Info(c.Status("Warning"), c.Description("Successful with warnings"))
	w.WriteHeader(200)
	return
}

func CorrelateExceptHandler(w http.ResponseWriter, r *http.Request) {
	correlate := c.NewCorrelateRequest(r, CorrelateExceptHandler)
	correlate.Fatal(c.Status("Baq request from client"), c.Description("Client missing required fields"))
	w.WriteHeader(400)
	return
}

func CorrelateFatalHandler(w http.ResponseWriter, r *http.Request) {
	correlate := c.NewCorrelateRequest(r, CorrelateFatalHandler)
	correlate.Fatal(c.Status("Error with function [name]"), c.Description("Description of the error"))
	w.WriteHeader(500)
	return
}

func CorrelateSuccessHandler(w http.ResponseWriter, r *http.Request) {
	correlate := c.NewCorrelateRequest(r, CorrelateSuccessHandler)
	correlate.Success(c.Status("Success func [name] call"), c.Description("Successful call of Service"))
	w.WriteHeader(200)
	return
}
