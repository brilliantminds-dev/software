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

func TestContactHandler(t *testing.T) {

	handler := Handler{
		Datastore: mockDataStore,
		Context:   context.TODO(),
	}

	tests := []struct {
		name     string
		body     models.Contact
		expected int
	}{
		{
			"create_contact_happyPath",
			models.Contact{

				Name:     "test user",
				Phone:    "804-555-5555",
				Email:    "test@gmail.com",
				Category: "personal",
			},
			http.StatusCreated,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			body, _ := json.Marshal(tt.body)
			req := httptest.NewRequest("POST", "/{user_id}/contacts", bytes.NewBuffer(body))
			req.Header.Set("Accept", "application/json")
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Client-Id", models.CRM_ADMIN)

			w := httptest.NewRecorder()

			handler.Contact(w, req)

			if w.Code != tt.expected {
				t.Error("test failed")
			}
		})
	}
}
