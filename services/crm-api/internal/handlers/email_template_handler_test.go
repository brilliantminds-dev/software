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

func TestEmailTemplateHandler(t *testing.T) {

	handler := Handler{
		Datastore: mockDataStore,
		Context:   context.TODO(),
	}

	tests := []struct {
		name     string
		body     models.EmailTemplate
		expected int
	}{
		{
			"create_contact_happyPath",
			models.EmailTemplate{
				TemplateName: "test-template",
				Subject:      "follow up email",
				Body:         "Hi {name}, Thank you for contacting us",
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

			handler.EmailTemplates(w, req)

			if w.Code != tt.expected {
				t.Error("test failed")
			}
		})
	}
}
