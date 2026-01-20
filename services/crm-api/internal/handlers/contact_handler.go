package handlers

import (
	"crm-platform-management-api/internal/models"
	"crm-platform-management-api/internal/pkg"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func (handler Handler) Contact(w http.ResponseWriter, r *http.Request) {
	// if its a post method, create new contact
	params := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "POST" {

		var newContact models.Contact

		headerErr := pkg.CheckHeaders(r)
		if headerErr != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`"error":"invalid header values"`))
			return
		}
		json.NewDecoder(r.Body).Decode(&newContact)

		errs := validateContactRequest(&newContact)

		if len(errs) != 0 {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"error":"Name, Phone, Email is required"}`))
			return
		}

		err := handler.Datastore.CreateContact(&newContact, params["user_id"])
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf(`{"error":"Error creating contact %s"`, err.Error())))
			return
		}
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"success":true, "message":"contact created"}`))
		return

	}

	// GET Endpoint
	categoryId := r.URL.Query().Get("category_id")

	contacts, err := handler.Datastore.GetContacts(&categoryId, params["user_id"])
	if err != nil {
		w.WriteHeader(404)
		w.Write([]byte(`no contacts found for this specific user`))
	}

	json.NewEncoder(w).Encode(&contacts)

}

func validateContactRequest(c *models.Contact) []string {
	errs := []string{}
	if c.Name == "" || c.Phone == "" || c.Email == "" || c.Category == "" {
		errs = append(errs, "Name, Phone, Email, Contact Type is required.")
	}
	return errs
}
