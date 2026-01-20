package handlers

import (
	"crm-platform-management-api/internal/models"
	"crm-platform-management-api/internal/pkg"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func (handler *Handler) EmailTemplates(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var emailTemplateRequest models.EmailTemplate
	userId := mux.Vars(r)

	err := pkg.CheckHeaders(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error":"invalid header values"`))
		return
	}

	json.NewDecoder(r.Body).Decode(&emailTemplateRequest)

	err = validateEmailTemplateRequest(&emailTemplateRequest)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error":"Template_Name, Subject, Body must be valid"}`))
		return
	}

	err = handler.Datastore.CreateEmailTemplate(&emailTemplateRequest, userId["user_id"])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf(`{"error":"%s"}`, err.Error())))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"success":true,"message":"email template created for user"}`))
	return
}

func validateEmailTemplateRequest(m *models.EmailTemplate) error {
	if m.TemplateName == "" || m.Subject == "" || m.Body == "" {
		return errors.New("Invalid Values")
	}
	return nil
}
