package handlers

import (
	"context"
	"crm-platform-management-api/internal/interfaces"
	"net/http"
)

type Handler struct {
	interfaces.Datastore
	context.Context
}

func (handler *Handler) GetHealth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"UP"}`))
}
