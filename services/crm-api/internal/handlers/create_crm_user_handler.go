package handlers

import (
	"crm-platform-management-api/internal/models"
	"crm-platform-management-api/internal/pkg"
	"encoding/json"
	"net/http"
	"strings"
)

func (handler *Handler) CreateCRMUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	//validate headers
	headerErr := pkg.CheckHeaders(r)
	if headerErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`"error":"invalid header values"`))
		return
	}

	var newCRMUser models.CRMUser

	err := struct {
		Error []string `json:"error"`
	}{}

	// unmarshall body
	json.NewDecoder(r.Body).Decode(&newCRMUser)

	// validating request
	errs := validateRequest(&newCRMUser)
	err.Error = errs
	// return bad request if there are validation errors
	if len(validateRequest(&newCRMUser)) != 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&err)
		return

	}

	// checking to see if user already exits
	if exists := handler.Datastore.UserExists(&newCRMUser); exists {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`"error":"User already exists"`))
		return
	}

	// create user
	createCRMErr := handler.Datastore.CreateCRMUser(&newCRMUser)
	if createCRMErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(createCRMErr.Error()))
		return
	}

	w.WriteHeader(http.StatusCreated)
	return

}

func validateRequest(m *models.CRMUser) []string {
	errorDetails := []string{}
	if &m.Username == nil || m.Username == "" {
		errorDetails = append(errorDetails, "Username cannot be null")
	}
	if &m.Password == nil || m.Password == "" {
		errorDetails = append(errorDetails, "Password cannot be null")
	}
	if len(m.Password) < 7 {
		errorDetails = append(errorDetails, "Password must be at least 8 characters")
	}

	if checkSpecialCharacters(m.Password) < 2 {
		errorDetails = append(errorDetails, "Password must have two special characters")

	}
	return errorDetails

}

func checkSpecialCharacters(password string) int {
	// check for special characters
	sc := "!@#$%^&*"
	count := 0
	p := strings.Split(password, "")
	for _, pass := range p {
		if strings.ContainsAny(pass, sc) {
			count++
		}
	}
	return count
}
