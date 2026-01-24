## correlate

##### Logging library for your  REST API in Go
<hr>

###### Owner / Engineer: <a href="https://www.linkedin.com/in/javar-harris-a0167118a/"> Javar Harris </a>
###### Contact: brilliantmindsdev@gmail.com
###### Profile: https://javar-harris.carrd.co
###### Released Versions: [CHANGELOG.md](CHANGELOG.md)

<hr>

#### Install


```go

go get github.com/brilliantminds-dev/software/commits/libraries/packages/correlate/v.1.0.0

```

<hr>

#### Types

INFO

SUCCESS

EXCEPT 

FATAL

<hr>

#### Input Variables

type Status string

type Description string

<hr>

#### Setup

```go

correlate := correlate.NewCorrelateRequest(request, function)

```

<hr>

#### Definitions


~~~

Function:

    NewCorrelateRequest // initalizes CorrelateService

Attributes:

   request *http.Request - request from the client

   func interface{} - the handler  i.e TestHander - Note: do not include call operator "()"

Interface:

   CorrelationService - implements the abstract functions

     Info()
     Success()
     Except()
     Fatal()

~~~

<hr>

#### Usage

```go


func CorrelateInfoHandler(w http.ResponseWriter, r *http.Request) {
	correlate := c.NewCorrelateRequest(r, CorrelateInfoHandler)
	correlate.Info(c.Status("Warning"), c.Description("Successful with warnings"))
	w.WriteHeader(200)
	return
}

func CorrelateExceptHandler(w http.ResponseWriter, r *http.Request) {
	correlate := c.NewCorrelateRequest(r, CorrelateExceptHandler)
	correlate.Except(c.Status("Baq request from client"), c.Description("Client missing required fields"))
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

```

<hr>

#### Example Logs in CloudWatch, Docker Logs, Terminal, etc

```bash

 # INFO Logs

Type: INFO | correlation_id: baec0d3b-f82b-4165-a54c-cfb6fc05c330 | status: Warning | desc: Successful with warnings| host: example.com | path: /correlate | funcName: github.com/brilliantminds-dev/software/libraries/packages/correlate_test.CorrelateInfoHandler| client_headers: map[User-Agent:[Go-http-client/1.1]]

 # Exception Logs

 Type: EXCEPT | correlation_id: f08c10a4-a268-42ce-aa9c-4f462e0ce246 | status: Baq request from client | desc: Client missing required fields| host: example.com | path: /correlate | funcName: github.com/brilliantminds-dev/software/libraries/packages/correlate_test.CorrelateExceptHandler| client_headers: map[User-Agent:[Go-http-client/1.1]]

 # Success Logs

 Type: SUCCESS | correlation_id: b9f5d93f-a295-4546-b0b5-168d09109fc9 | status: Success func [name] call | desc: Successful call of Service| host: example.com | path: /correlate | funcName: github.com/brilliantminds-dev/software/libraries/packages/correlate_test.CorrelateSuccessHandler| client_headers: map[User-Agent:[Go-http-client/1.1]]

 # Fatal Logs

 Type: FATAL | correlation_id: 735ccf1f-3403-4dd3-b408-b64fd3f5b66b | status: Error with function [name] | desc: Description of the error| host: example.com | path: /correlate | funcName: github.com/brilliantminds-dev/software/libraries/packages/correlate_test.CorrelateFatalHandler| client_headers: map[User-Agent:[Go-http-client/1.1]]



