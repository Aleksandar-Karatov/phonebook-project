package repository

import (
	"context"
	"log"

	"github.com/phonebook-project/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repo struct {
	Collection *mongo.Collection
	Ctx        context.Context
}

// Contact returns a single contact from the db
func (r *Repo) Contact(name, phone, email, id string) (*models.Contact, error) {
	if id != "" {
		filter := bson.D{
			{"id", id},
		}
		destination := &models.Contact{}
		err := r.Collection.FindOne(r.Ctx, filter).Decode(destination)
		if err != nil {
			log.Println(err)
			return nil, ErrNotFound
		}
		return destination, nil
	}
	if email != "" {
		filter := bson.D{
			{"email", email},
		}
		destination := &models.Contact{}
		err := r.Collection.FindOne(r.Ctx, filter).Decode(destination)
		if err != nil {
			log.Println(err)
			return nil, ErrNotFound
		}
		return destination, nil
	}
	if phone != "" {
		filter := bson.D{
			{"phoneNumbers", bson.D{{"$all", bson.A{phone}}}},
		}
		destination := &models.Contact{}
		err := r.Collection.FindOne(r.Ctx, filter).Decode(destination)
		if err != nil {
			log.Println(err)

			return nil, ErrNotFound
		}
		return destination, nil
	}
	if name != "" {
		filter := bson.D{
			{"name", name},
		}
		destination := &models.Contact{}
		err := r.Collection.FindOne(r.Ctx, filter).Decode(destination)
		if err != nil {
			log.Println(err)

			return nil, ErrNotFound
		}
		return destination, nil
	}

	return nil, ErrNotFound

}

// Contacts returns all records from the db
func (r *Repo) Contacts() ([]*models.Contact, error) {
	var contacts []*models.Contact
	cur, err := r.Collection.Find(r.Ctx, bson.D{{}})
	if err != nil {
		return nil, err
	}
	for cur.Next(r.Ctx) {
		var contact models.Contact
		err := cur.Decode(&contact)
		if err != nil {
			return nil, err
		}
		contacts = append(contacts, &contact)
	}
	return contacts, nil
}

// CreateContact creates a new contact
func (r *Repo) CreateContact(contact models.Contact) (*models.Contact, error) {
	_, err := r.Contact(contact.Name, contact.Phones[0], contact.Email, contact.Id)
	if err == nil {
		return nil, ErrAlreadyExists
	}
	_, err = r.Collection.InsertOne(r.Ctx, contact)
	if err != nil {
		return nil, err
	}
	return r.Contact(contact.Name, contact.Phones[0], contact.Email, contact.Id)
}

// UpdateContact updates an existing contact
func (r *Repo) UpdateContact(contact models.Contact) (*models.Contact, error) {
	check, err := r.Contact(contact.Name, contact.Phones[0], contact.Email, contact.Id)
	if check == nil {
		return nil, ErrNotFound
	}

	filter := bson.D{
		{"id", contact.Id},
	}
	update := bson.D{{"$set", contact}}

	_, err = r.Collection.UpdateOne(r.Ctx, filter, update)
	if err != nil {
		return nil, err
	}
	return r.Contact(contact.Name, contact.Phones[0], contact.Email, contact.Id)
}

// DeleteContact deletes an existing contact
func (r *Repo) DeleteContact(email string) (bool, error) {
	check, _ := r.Contact("", "", email, "")
	if check == nil {
		return false, ErrNotFound
	}
	filter := bson.D{
		{"email", email},
	}
	_, err := r.Collection.DeleteOne(r.Ctx, filter)
	if err != nil {
		return false, err
	}
	return true, nil
}

// MergeContacts merges two existing contacts with the same name
func (r *Repo) MergeContacts(email_one, email_two string) (*models.Contact, error) {
	return &models.Contact{}, nil
}
