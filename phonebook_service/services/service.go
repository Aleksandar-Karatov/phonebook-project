package services

import (
	"errors"

	"github.com/phonebook-project/models"
	"github.com/phonebook-project/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
func Contact(repo *repository.Repo, name, phone, email *string, id *primitive.ObjectID) (*models.Contact, error) {
	resp, err := repo.Contact(name, phone, email, id)
	if err != nil {
		if err == repository.ErrNotFound {
			return nil, errors.New("Contact was not found")
		}
		if err == repository.ErrInvalidData {
			return nil, errors.New("Invalid input data")
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
func DeleteContact(repo *repository.Repo, id primitive.ObjectID) (bool, error) {
	flag, err := repo.DeleteContact(id)
	if err != nil {
		if err == repository.ErrNotFound {
			return false, errors.New("Contact doesn`t exist")
		}
		return false, errors.New("Could not delete record")
	}
	return flag, err
}

// MergeContacts merges two existing contacts with the same name
func MergeContacts(repo *repository.Repo, userID primitive.ObjectID) (*models.Contact, error) {
	if userID.IsZero() {
		return nil, errors.New("User id is empty")
	}

	resp, err := repo.MergeContacts(userID)
	if err == repository.ErrNotFound {
		return nil, errors.New("Can not find records")
	}
	if err == repository.ErrNoDuplicates {
		return nil, errors.New("No contacts with this name and email")
	}
	if err != nil {
		return nil, errors.New("Can not merge contacts")
	}

	return resp, nil
}

func UpdateContactPicture(repo *repository.Repo, userID primitive.ObjectID, image models.Image) (*models.Image, error) {
	if userID.IsZero() {
		return nil, errors.New("User not specified")
	}
	if image.ContentType != "" && len(image.Data) == 0 && image.Name != "" && !image.Id.IsZero() {
		resp, err := repo.UpdateContactPicture(userID, image)
		if err != nil {
			return nil, err
		}
		return resp, nil
	}

	return nil, errors.New("Invalid image data")
}

func Image(repo *repository.Repo, imageID primitive.ObjectID) (*models.Image, error) {
	if imageID.IsZero() {
		return nil, errors.New("Invalid image data")
	}

	resp, err := repo.Image(imageID)
	if err == repository.ErrNotFound {
		return nil, errors.New("Image could not be found")
	}
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func Search(repo *repository.Repo, filter string) ([]*models.Contact, error) {
	if filter == "" {
		return nil, nil
	}
	resp, err := repo.Search(filter)
	if err != nil {
		if err == repository.ErrNotFound {
			return nil, nil
		}
		return nil, err
	}

	return resp, nil

}
