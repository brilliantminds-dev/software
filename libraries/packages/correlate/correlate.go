package correlate

import (
	"fmt"
	"net/http"
	"reflect"
	"runtime"

	"github.com/google/uuid"
)

// Constants INFO, EXCEPT, SUCCESS, FATAL
const (
	Info        = "INFO"
	Except      = "EXCEPT"
	Success     = "SUCCESS"
	Fatal       = "FATAL"
	CorrDetails = " Type: %s | correlation_id: %v | status: %v | desc: %v| host: %v | path: %s | funcName: %v| client_headers: %v"
)

// TraceID, Status, Description, HandlerName, RequestEventTs, ClientRequestHeaders
type TraceID string          // trace id to navigate the request
type Status string           // status of the request
type Description string      // description in the request
type HandlerName interface{} //  handler name of the request
type RequestEventTs string   // timestamp of the request's event
type ClientRequestHeaders map[string]any

type Correlate struct {
	TraceID
	HandlerName
	ClientRequestHeaders
	*http.Request
}

// New CorrelateRequest - Initialize New Reguest Correlation for REST
func NewCorrelateRequest(request *http.Request, fnName interface{}) CorrelateService {

	return &Correlate{
		TraceID:              TraceID(uuid.NewString()),
		HandlerName:          fnName,
		ClientRequestHeaders: collectHeaders(request.Header),
		Request:              request,
	}

}

// CorrelateService - Info, Except, Success, Fatal
type CorrelateService interface {
	Info(st Status, desc Description)
	Except(st Status, desc Description)
	Success(st Status, desc Description)
	Fatal(st Status, desc Description)
}

func (c *Correlate) Info(st Status, desc Description) {
	info := fmt.Sprintf(CorrDetails, Info, c.TraceID, st, desc, c.Host, c.URL.Path, fn(c.HandlerName), c.ClientRequestHeaders)

	fmt.Println(info)

}
func (c *Correlate) Except(st Status, desc Description) {
	exception := fmt.Sprintf(CorrDetails, Except, c.TraceID, st, desc, c.Host, c.URL.Path, fn(c.HandlerName), c.ClientRequestHeaders)

	fmt.Println(exception)
}
func (c *Correlate) Success(st Status, desc Description) {
	success := fmt.Sprintf(CorrDetails, Success, c.TraceID, st, desc, c.Host, c.URL.Path, fn(c.HandlerName), c.ClientRequestHeaders)

	fmt.Println(success)

}
func (c *Correlate) Fatal(st Status, desc Description) {
	fatal := fmt.Sprintf(CorrDetails, Fatal, c.TraceID, st, desc, c.Host, c.URL.Path, fn(c.HandlerName), c.ClientRequestHeaders)

	fmt.Println(fatal)
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
