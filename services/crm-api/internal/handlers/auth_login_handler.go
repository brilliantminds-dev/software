package handlers

import (
	"crm-platform-management-api/internal/jwt_token"
	"crm-platform-management-api/internal/models"
	"crm-platform-management-api/internal/pkg"
	"encoding/json"
	"fmt"
	"net/http"
)

func (handler Handler) AuthLoginHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	headerErr := pkg.CheckHeaders(r)
	if headerErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`"error":"invalid header values"`))
		return
	}

	var authLogin models.AuthLogin

	err := struct {
		ErrorDetails []string `json:"errorDetails"`
	}{}

	json.NewDecoder(r.Body).Decode(&authLogin)

	// validate request
	errs := validateLoginRequest(&authLogin)

	if len(errs) != 0 {
		err.ErrorDetails = errs
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&err)
		return
	}

	// check login
	user, authErr := handler.Datastore.Login(&authLogin)
	if authErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error":"login failed. username/password incorrect"`))
		return
	}
	genAuthToken := jwt_token.GenerateToken(user)
	token, tokenErr := genAuthToken.SignedString([]byte(models.SECRET))

	if tokenErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error":"unable to retrieve auth token. contact CRM Support."}"`))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf(`{"token":"%v"}`, token)))

}

func validateLoginRequest(m *models.AuthLogin) []string {
	errs := []string{}
	if &m.Username == nil || m.Username == "" {
		errs = append(errs, "Username must not be blank")
	}
	if &m.Password == nil || m.Password == "" {
		errs = append(errs, "Password must not be blank")
	}
	return errs
}
