package handlers

import (
	"bytes"
	"context"
	"crm-platform-management-api/internal/mocks"
	"crm-platform-management-api/internal/models"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

var mockDataStore = mocks.NewMockDatastore()

func TestHandler_AuthLoginHandler(t *testing.T) {
	handler := Handler{
		Datastore: mockDataStore,
		Context:   context.TODO(),
	}
	tests := []struct {
		name     string
		body     models.AuthLogin
		expected int
	}{{
		"login_happyPath",
		models.AuthLogin{
			Username: "test-user",
			Password: "password",
		},
		http.StatusOK,
	},
	}

	for _, tt := range tests {

		body, _ := json.Marshal(tt.body)
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("POST", "/auth/login", bytes.NewBuffer(body))
			req.Header.Set("Accept", "application/json")
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Client-Id", models.CRM_ADMIN)

			w := httptest.NewRecorder()

			handler.AuthLoginHandler(w, req)

			if w.Code != tt.expected {
				t.Error("test failed")
			}
		})
	}
}
