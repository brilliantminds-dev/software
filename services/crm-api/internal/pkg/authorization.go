package pkg

import (
	"crm-platform-management-api/internal/models"
	"errors"
	"net/http"
	"slices"
)

func CheckHeaders(r *http.Request) error {
	// validate headers

	accept := r.Header.Get("Accept")
	contentType := r.Header.Get("Content-Type")
	clientID := r.Header.Get("Client-Id")
	headerValue := "application/json"

	validClientIDs := []string{models.CRM_ADMIN, models.CRM_USER, models.CRM_READ_ONLY}

	if accept != headerValue || contentType != headerValue || !slices.Contains(validClientIDs, clientID) {
		return errors.New("invalid header values")
	}

	return nil

}
