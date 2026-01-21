package correlate

import (
	"fmt"
	"net/http"
	"reflect"
	"runtime"

	"github.com/google/uuid"
)

const (
	INFO         = "INFO"
	EXCEPT       = "EXCEPT"
	SUCCESS      = "SUCCESS"
	FATAL        = "FATAL"
	CORR_DETAILS = " Type: %s | correlation_id: %v |status: %v | desc: %v| client_ip: %v | path: %s | funcName: %v| client_headers: %v"
)

type TraceId string          // trace id to navigate the request
type Status string           // status of the request
type Description *string     // description in the request
type HandlerName interface{} //  handler name of the request
type RequestEventTs string   // timestamp of the request's event
type ClientRequestHeaders map[string]any

type Correlate struct {
	TraceId
	HandlerName
	ClientRequestHeaders
	*http.Request
}

// New CorrelateRequest - Initialize New Reguest Correlation for REST
func NewCorrelateRequest(request *http.Request, fnName interface{}) CorrelateService {

	return &Correlate{
		TraceId:              TraceId(uuid.NewString()),
		HandlerName:          fnName,
		ClientRequestHeaders: collectHeaders(request.Header),
		Request:              request,
	}

}

// CorrelateService - Info, Except, Success, Fatal
type CorrelateService interface {
	Info(st *Status, desc *Description)
	Except(st *Status, desc *Description) error
	Success(st *Status, desc *Description)
	Fatal(st *Status, desc *Description) error
}

func (c *Correlate) Info(st *Status, desc *Description) {
	info := fmt.Sprintf(CORR_DETAILS, INFO, c.TraceId, &st, &desc, c.RemoteAddr, c.URL.Path, fn(c.HandlerName), c.ClientRequestHeaders)
	fmt.Println(info)

}
func (c *Correlate) Except(st *Status, desc *Description) error {
	exception := fmt.Sprintf(CORR_DETAILS, EXCEPT, c.TraceId, &st, &desc, c.RemoteAddr, c.URL.Path, fn(c.HandlerName), c.ClientRequestHeaders)

	return fmt.Errorf("%v", exception)
}
func (c *Correlate) Success(st *Status, desc *Description) {

}
func (c *Correlate) Fatal(st *Status, desc *Description) error {
	fatal := fmt.Sprintf(CORR_DETAILS, FATAL, c.TraceId, &st, &desc, c.RemoteAddr, c.URL.Path, fn(c.HandlerName), c.ClientRequestHeaders)

	return fmt.Errorf("%s", fatal)
}

func fn(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

func collectHeaders(header http.Header) ClientRequestHeaders {
	if header == nil {
		return nil
	}
	headers := make(ClientRequestHeaders)
	for key, values := range header {
		headers[key] = values
	}

	return headers
}
