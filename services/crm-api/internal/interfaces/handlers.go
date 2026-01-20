package interfaces

import "net/http"

type HandlerInterface interface {
	GetHealth(w http.ResponseWriter, r *http.Request)
	CreateCRMUser(w http.ResponseWriter, r *http.Request)
	AuthLoginHandler(w http.ResponseWriter, r *http.Request)
	Contact(w http.ResponseWriter, r *http.Request)
	EmailTemplates(writer http.ResponseWriter, request *http.Request)
}
