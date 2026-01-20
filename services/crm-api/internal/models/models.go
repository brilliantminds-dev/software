package models

type CRMUser struct {
	ID        string `json:"id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type AuthLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Contact struct {
	ID                string `json:"id,omitempty"`
	Name              string `json:"name"`
	Phone             string `json:"phone"`
	Email             string `json:"email"`
	Category          string `json:"contact_type,omitempty"`
	CRMID             string `json:",omitempty"` // below comment
	ContactCategoryID string `json:",omitempty"` //dont show sensitive data in response
}

type EmailTemplate struct {
	ID           string `json:"id"`
	TemplateName string `json:"template_name"`
	Subject      string `json:"subject"`
	Body         string `json:"body"`
}

// valid client id values
const (
	CRM_ADMIN     = "CRM-ADMIN"
	CRM_USER      = "CRM USER"
	CRM_READ_ONLY = "CRM_READ_ONLY"
)

const SECRET = "jwt-auth-test-token"
