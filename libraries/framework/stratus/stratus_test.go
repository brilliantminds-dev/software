package stratus

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestNewStratus(t *testing.T) {
	stratus := NewStratus()
	compare := &Stratus{}
	assert.NotEmpty(t, stratus)
	assert.True(t, reflect.TypeOf(compare) == reflect.TypeOf(stratus))
}

func TestStratus_Get(t *testing.T) {

	tdt := []struct {
		name               string
		handler            func(w http.ResponseWriter, r *http.Request)
		expectedStatusCode int
		actualStatusCode   int
	}{
		{
			"impact_get_status_ok",
			func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(200)
			},
			200,
			0,
		},
	}
	for _, tt := range tdt {

		r := httptest.NewRequest("GET", "/impact/test", nil)
		w := httptest.NewRecorder()
		handler := http.HandlerFunc(tt.handler)
		handler(w, r)
		tt.actualStatusCode = w.Code

		if tt.actualStatusCode != tt.expectedStatusCode {
			t.Logf("want %v,  got %v test failed", tt.actualStatusCode, tt.expectedStatusCode)
			t.Fail()
		}

	}
}
