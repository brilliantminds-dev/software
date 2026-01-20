package handlers

import (
	"bytes"
	"context"
	"crm-platform-management-api/internal/models"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateCRMUserHandler(t *testing.T) {

	handler := Handler{
		Datastore: mockDataStore,
		Context:   context.TODO(),
	}

	tests := []struct {
		name     string
		body     models.CRMUser
		expected int
	}{
		{
			"create_crm_user_happyPath",
			models.CRMUser{
				Username: "test-user",
				Password: "pass!@#$",
			},
			http.StatusCreated,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			body, _ := json.Marshal(tt.body)
			req := httptest.NewRequest("POST", "/create-user", bytes.NewBuffer(body))
			req.Header.Set("Accept", "application/json")
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Client-Id", models.CRM_ADMIN)

			w := httptest.NewRecorder()

			handler.CreateCRMUser(w, req)

			if w.Code != tt.expected {
				t.Error("test failed")
			}
		})
	}
}
