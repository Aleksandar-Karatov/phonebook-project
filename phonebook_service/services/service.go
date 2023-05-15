package services

import (
	"errors"

	"github.com/phonebook-project/models"
	"github.com/phonebook-project/repository"
)

// CreateContact creates a new contact
func CreateContact(repo *repository.Repo, contact models.Contact) (*models.Contact, error) {

	if contact.Email != "" && contact.Name != "" && len(contact.Phones) > 0 {
		resp, err := repo.CreateContact(contact)
		if err != nil {
			if err == repository.ErrAlreadyExists {
				return nil, errors.New("Contact already exists")
			}
			return nil, errors.New("Could not create record in db")
		}
		return resp, nil
	}
	return nil, errors.New("Data was not valid")
}

// Contact returns a single contact from the db
func Contact(repo *repository.Repo, name, phone, email, id string) (*models.Contact, error) {
	if name == "" && phone == "" && email == "" {
		return nil, errors.New("Invalid input")
	}
	resp, err := repo.Contact(name, phone, email, id)
	if err != nil {
		if err == repository.ErrNotFound {
			return nil, errors.New("Contact was not found")
		}
		return nil, err
	}
	return resp, nil
}

//Contacts returns all records from the db

func Contacts(repo *repository.Repo) ([]*models.Contact, error) {
	resp, err := repo.Contacts()
	if err != nil {
		if err == repository.ErrNotFound {
			return nil, errors.New("Contacts were not found")
		}
		return nil, err
	}

	return resp, nil
}

// UpdateContact updates an existing contact
func UpdateContact(repo *repository.Repo, contact models.Contact) (*models.Contact, error) {
	if contact.Email != "" && len(contact.Phones) > 0 && contact.Name != "" {
		resp, err := repo.UpdateContact(contact)
		if err != nil {
			if err == repository.ErrNotFound {
				return nil, errors.New("Contact doesn`t exist")
			}
			return nil, errors.New("Could not update record in db")
		}
		return resp, nil
	}
	return nil, errors.New("Data was not valid")

}

// DeleteContact deletes an existing contact
func DeleteContact(repo *repository.Repo, email string) (bool, error) {
	flag, err := repo.DeleteContact(email)
	if err != nil {
		if err == repository.ErrNotFound {
			return false, errors.New("Contact doesn`t exist")
		}
		return false, errors.New("Could not delete record")
	}
	return flag, err
}

// MergeContacts merges two existing contacts with the same name
func MergeContacts(repo *repository.Repo, email_one, email_two string) (models.Contact, error) {
	return models.Contact{}, nil
}
