package main

import (
	"context"
	"crm-platform-management-api/internal/config"
	"crm-platform-management-api/internal/datastore"
	"crm-platform-management-api/internal/handlers"
	"crm-platform-management-api/internal/interfaces"
	"crm-platform-management-api/internal/middleware"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func newHandler(db interfaces.Datastore, ctx context.Context) *handlers.Handler {
	return &handlers.Handler{db, ctx}
}



func main() {

	ctx := context.TODO()

	appConfig := config.GetAppConfig()

	// invoke db datastore
	db := datastore.NewCRMDb(appConfig)

	// using handler to use Endpoint interfaces
	handler := newHandler(db, ctx)

	// set up router
	mux := mux.NewRouter()

	// setup logger middleware

	mux.HandleFunc("/crm-platform-management-api/admin/health", handler.GetHealth).Methods("GET")
	mux.HandleFunc("/crm-platform-management-api/create-user", handler.CreateCRMUser).Methods("POST")
	mux.HandleFunc("/crm-platform-management-api/auth/login", handler.AuthLoginHandler).Methods("POST")
	mux.HandleFunc("/crm-platform-management-api/{user_id}/contacts", handler.Contact).Methods("POST", "GET")
	mux.HandleFunc("/crm-platform-management-api/{user_id}/email/email-templates", handler.EmailTemplates).Methods("POST")

	mux.Use(middleware.LoggingMiddleware)

	log.Fatal(http.ListenAndServeTLS(":8080", "certs/server.crt", "certs/server.key", mux))

}
