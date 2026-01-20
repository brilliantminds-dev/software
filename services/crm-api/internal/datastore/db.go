package datastore

import (
	"crm-platform-management-api/internal/config"
	"crm-platform-management-api/internal/interfaces"
	"crm-platform-management-api/internal/models"
	"crm-platform-management-api/internal/pkg"
	"database/sql"
	"errors"
	"fmt"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"log"
	"time"
)

type DB struct {
	*sql.DB
}

func NewCRMDb(config *config.AppConfig) interfaces.Datastore {

	// set up datastore
	conn, err := sql.Open("postgres", fmt.Sprintf("host=%s  dbname=%s user=%s port=%d sslmode=disable", config.Host, config.DBName, config.User, 5432))
	if err != nil {
		log.Fatal("Error connecting to database: " + err.Error())
	}

	return &DB{conn}

}

func (d *DB) CreateCRMUser(user *models.CRMUser) error {

	// Stores CRM User to the database

	date := time.Now().Format("2006-01-02")

	_, err := d.Query(fmt.Sprintf(`insert into crm_pltfrm_mgmt_users values('%s','%s','%s','%s','%s')`, uuid.New(), user.Username, pkg.Hash(user.Password), date, date))
	if err != nil {
		log.Fatal("error saving new CRM User" + err.Error())
		return errors.New("error when inserting new CRM User" + err.Error())
	}

	return nil
}

func (d *DB) Login(auth *models.AuthLogin) (*models.CRMUser, error) {
	var crmUser models.CRMUser

	query, err := d.Query(fmt.Sprintf(`select * from crm_pltfrm_mgmt_users where username='%s'`, auth.Username))

	if err != nil {
		log.Fatal("error getting credentials from CRM Database " + err.Error())
	}

	for query.Next() {
		query.Scan(&crmUser.ID, &crmUser.Username, &crmUser.Password, &crmUser.CreatedAt, &crmUser.UpdatedAt)

	}

	// checking to see if passwords match based on hash and password from user
	isMatched := pkg.CheckHash(auth.Password, crmUser.Password)

	if isMatched {
		return &crmUser, nil
	} else {
		return nil, errors.New("password does not match")
	}

}

func (d *DB) CreateContact(m *models.Contact, userid string) error {
	contactId := uuid.New()
	contactCategoryID := d.getContactCategoryId(m)

	// first lets insert into contact_category table
	_, err := d.Query(fmt.Sprintf(`insert into crm_pltfrm_mgmt_contact_category values('%s','%s')`, contactCategoryID, m.Category))
	if err != nil {
		log.Println("category already exists.. adding to contact to exisitng category.")
	}

	// create the contact to specific category
	_, err = d.Query(fmt.Sprintf(`insert into crm_pltfrm_mgmt_contacts values('%s','%s','%s','%s','%s','%s')`, contactId, m.Name, m.Phone, m.Email, userid, contactCategoryID))
	if err != nil {
		return errors.New(fmt.Sprintf("error creating new contact: %s", err.Error()))
	}
	return nil
}

func (d *DB) GetContacts(id *string, userId string) (*[]models.Contact, error) {

	var contact models.Contact
	var contactList []models.Contact

	var q string

	if &id == nil || *id == "" { // contact_category id is null, lets get all the contacts for this user
		q = fmt.Sprintf(`select name, phone, email from 
                              crm_pltfrm_mgmt_contacts as crmc
                              where crmc.crm_id='%s'`, userId)
	} else { // if contact_category is populated, filter specific list of contacts based on contact category id
		q = fmt.Sprintf(`select name, phone, email from 
                crm_pltfrm_mgmt_contacts as crmc join crm_pltfrm_mgmt_contact_category as crmcc
                on crmc.contact_category_id=crmcc.id
                where crmc.crm_id = '%s'
                and crmcc.id='%s'`, userId, *id)
	}

	query, err := d.Query(q)
	if err != nil {
		return nil, err
	}

	for query.Next() {
		query.Scan(&contact.Name, &contact.Phone, &contact.Email)
		contactList = append(contactList, contact)
	}

	return &contactList, nil
}

func (d *DB) UserExists(user *models.CRMUser) bool {
	query, _ := d.Query(fmt.Sprintf(`select username from crm_pltfrm_mgmt_users where username='%s'`, user.Username))

	return query.Next()
}

func (d *DB) CreateEmailTemplate(m *models.EmailTemplate, userId string) error {

	var templateID = uuid.New()

	createContactExec := fmt.Sprintf(`insert into crm_pltfrm_mgmt_email_templates values('%s','%s','%s','%s','%s')`, templateID, m.TemplateName, m.Subject, m.Body, userId)
	_, err := d.Query(createContactExec)
	if err != nil {
		return err
	}

	return nil

}

func (d *DB) getContactCategoryId(c *models.Contact) string {
	// find the existing contact category
	var existingCategoryId string
	query, _ := d.Query(fmt.Sprintf("select id from crm_pltfrm_mgmt_contact_category where contact_category = '%s'", c.Category))

	// if there are results, lets return the existing contact category id
	if query.Next() {
		query.Scan(&existingCategoryId)
		return existingCategoryId
	}

	// if this is a new contact category, lets create a new id to store it
	return uuid.New().String()

}
