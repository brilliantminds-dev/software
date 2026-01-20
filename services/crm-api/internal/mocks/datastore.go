package mocks

import (
	"crm-platform-management-api/internal/models"
	"time"
)

// mocks setup for unit testing

type mockDatastore struct {
}

func NewMockDatastore() *mockDatastore {
	return &mockDatastore{}
}

func (m *mockDatastore) CreateCRMUser(user *models.CRMUser) error {
	//TODO implement me
	return nil
}

func (m *mockDatastore) UserExists(user *models.CRMUser) bool {
	//TODO implement me
	return false
}

func (m *mockDatastore) CreateContact(c *models.Contact, s string) error {
	//TODO implement me
	return nil
}

func (m *mockDatastore) GetContacts(id *string, params string) (*[]models.Contact, error) {
	//TODO implement me

	return &[]models.Contact{}, nil
}

func (m *mockDatastore) Login(auth *models.AuthLogin) (*models.CRMUser, error) {
	return &models.CRMUser{
		ID:        "abc-123-456-dfgh",
		Username:  "test-user",
		Password:  "password",
		CreatedAt: time.Now().Format(time.RFC3339),
		UpdatedAt: time.Now().Format(time.RFC3339),
	}, nil
}

func (m *mockDatastore) CreateEmailTemplate(e *models.EmailTemplate, s string) error {
	//TODO implement me
	return nil
}
