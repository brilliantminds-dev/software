package interfaces

import "crm-platform-management-api/internal/models"

type Datastore interface {
	CreateCRMUser(user *models.CRMUser) error
	Login(auth *models.AuthLogin) (*models.CRMUser, error)
	UserExists(user *models.CRMUser) bool
	CreateContact(m *models.Contact, s string) error
	GetContacts(id *string, params string) (*[]models.Contact, error)
	CreateEmailTemplate(m *models.EmailTemplate, s string) error
}
