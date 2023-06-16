package repository

import (
	"context"
	"fmt"
	"log"

	"github.com/phonebook-project/config"
	"github.com/phonebook-project/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repo struct {
	Database *mongo.Database
	Ctx      context.Context
}

// Contact returns a single contact from the db
func (r *Repo) Contact(name, phone, email *string, id *primitive.ObjectID) (*models.Contact, error) {

	if id != nil {
		if id.IsZero() {
			return nil, ErrInvalidData
		}
		filter := bson.D{
			{"id", id},
		}
		destination := &models.Contact{}
		err := r.Database.Collection(config.DB_COLLECTION_USERS).FindOne(r.Ctx, filter).Decode(destination)
		if err != nil {
			log.Println(err)
			return nil, ErrNotFound
		}
		return destination, nil
	}
	if checkString(email) {
		filter := bson.D{
			{"email", email},
		}
		destination := &models.Contact{}
		err := r.Database.Collection(config.DB_COLLECTION_USERS).FindOne(r.Ctx, filter).Decode(destination)
		if err != nil {
			log.Println(err)
			return nil, ErrNotFound
		}
		return destination, nil
	}
	if checkString(phone) {
		filter := bson.D{
			{"phoneNumbers", bson.D{{"$all", bson.A{phone}}}},
		}
		destination := &models.Contact{}
		err := r.Database.Collection(config.DB_COLLECTION_USERS).FindOne(r.Ctx, filter).Decode(destination)
		if err != nil {
			log.Println(err)

			return nil, ErrNotFound
		}
		return destination, nil
	}
	if checkString(name) {

		filter := bson.D{
			{"name", name},
		}
		destination := &models.Contact{}
		err := r.Database.Collection(config.DB_COLLECTION_USERS).FindOne(r.Ctx, filter).Decode(destination)
		if err != nil {
			log.Println(err)

			return nil, ErrNotFound
		}
		return destination, nil
	}

	return nil, ErrInvalidData

}
func checkString(input *string) bool {
	if input != nil {
		if *input != "" {
			return true
		}
	}
	return false
}

// Contacts returns all records from the db
func (r *Repo) Contacts() ([]*models.Contact, error) {
	var contacts []*models.Contact
	cur, err := r.Database.Collection(config.DB_COLLECTION_USERS).Find(r.Ctx, bson.D{{}})
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
	_, err := r.Contact(nil, nil, nil, &contact.Id)
	if err == nil {
		return nil, ErrAlreadyExists
	}
	contact.ImageID, _ = primitive.ObjectIDFromHex(config.DEFAULT_IMAGE_ID)
	_, err = r.Database.Collection(config.DB_COLLECTION_USERS).InsertOne(r.Ctx, contact)
	if err != nil {
		return nil, err
	}
	return r.Contact(&contact.Name, &contact.Phones[0], &contact.Email, &contact.Id)
}

// UpdateContact updates an existing contact
func (r *Repo) UpdateContact(contact models.Contact) (*models.Contact, error) {
	check, err := r.Contact(&contact.Name, nil, &contact.Email, &contact.Id)
	if check == nil {
		return nil, ErrNotFound
	}

	filter := bson.D{
		{"id", contact.Id},
	}
	update := bson.D{{"$set", contact}}

	_, err = r.Database.Collection(config.DB_COLLECTION_USERS).UpdateOne(r.Ctx, filter, update)
	if err != nil {
		return nil, ErrCantUpdate
	}
	return r.Contact(&contact.Name, nil, &contact.Email, &contact.Id)
}

// DeleteContact deletes an existing contact
func (r *Repo) DeleteContact(id primitive.ObjectID) (bool, error) {
	check, _ := r.Contact(nil, nil, nil, &id)
	if check == nil {
		return false, ErrNotFound
	}
	filter := bson.D{
		{"id", id},
	}
	_, err := r.Database.Collection(config.DB_COLLECTION_USERS).DeleteOne(r.Ctx, filter)
	if err != nil {
		return false, ErrCantDelete
	}
	return true, nil
}

// MergeContacts merges two existing contacts with the same name
func (r *Repo) MergeContacts(userID primitive.ObjectID) (*models.Contact, error) {
	contact, err := r.Contact(nil, nil, nil, &userID)
	if err != nil {
		return nil, ErrNotFound
	}
	filter := bson.D{
		{"email", contact.Email},
		{"name", contact.Name},
	}
	var duplicates []*models.Contact

	cur, err := r.Database.Collection(config.DB_COLLECTION_USERS).Find(r.Ctx, filter)
	for cur.Next(r.Ctx) {
		var contact models.Contact
		err := cur.Decode(&contact)
		if err != nil {
			return nil, err
		}
		duplicates = append(duplicates, &contact)
	}

	if len(duplicates) <= 1 {
		return contact, ErrNoDuplicates
	}
	for _, duplicate := range duplicates {
		if duplicate.Id == userID {
			continue
		}
		contact.Notes = fmt.Sprintf("%s \n %s", contact.Notes, duplicate.Notes)
		contact.Phones = append(contact.Phones, duplicate.Phones...)
		_, err := r.DeleteContact(duplicate.Id)
		if err != nil {
			log.Println(err)
		}
	}

	resp, err := r.UpdateContact(*contact)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// UpdateContactPicture updates an existing image
func (r *Repo) UpdateContactPicture(userID primitive.ObjectID, image models.Image) (*models.Image, error) {
	check, err := r.Image(image.Id)
	if check == nil {
		_, err = r.Database.Collection(config.DB_COLLECTION_IMAGES).InsertOne(r.Ctx, image)
		if err != nil {
			return nil, err
		}
	} else {
		filter := bson.D{
			{"_id", image.Id},
		}
		update := bson.D{{"$set", image}}

		_, err = r.Database.Collection(config.DB_COLLECTION_IMAGES).UpdateOne(r.Ctx, filter, update)
		if err != nil {
			return nil, ErrCantUpdate
		}

	}
	contact, err := r.Contact(nil, nil, nil, &userID)
	if err != nil {
		return nil, ErrNotFound
	}
	contact.ImageID = image.Id
	_, err = r.UpdateContact(*contact)
	if err != nil {
		return nil, err
	}

	return r.Image(image.Id)
}

func (r *Repo) Image(imageID primitive.ObjectID) (*models.Image, error) {
	filter := bson.D{
		{"_id", imageID},
	}
	destination := &models.Image{}
	err := r.Database.Collection(config.DB_COLLECTION_IMAGES).FindOne(r.Ctx, filter).Decode(destination)
	if err != nil {
		log.Println(err)
		return nil, ErrNotFound
	}
	return destination, nil

}

func (r *Repo) Search(filterStr string) ([]*models.Contact, error) {
	contactsByName := r.predictiveSearch(filterStr, "name")
	if len(contactsByName) != 0 {
		return contactsByName, nil
	}
	contactsByEmail := r.predictiveSearch(filterStr, "email")
	if len(contactsByEmail) != 0 {
		return contactsByEmail, nil
	}
	contactsByPhone := r.predictiveSearch(filterStr, "phoneNumbers")
	if len(contactsByPhone) != 0 {
		return contactsByPhone, nil
	}
	return nil, nil
}

func (r *Repo) predictiveSearch(filterStr, fieldName string) []*models.Contact {
	var contacts []*models.Contact
	findOptions := options.Find()

	findOptions.SetSort(bson.D{{fieldName, 1}})
	findOptions.SetLimit(10)
	filter := bson.D{{fieldName, primitive.Regex{Pattern: filterStr, Options: "i"}}}
	if fieldName == "phoneNumbers" {
		filter = bson.D{
			{fieldName, bson.D{{"$all", bson.A{primitive.Regex{Pattern: filterStr, Options: ""}}}}},
		}
	}
	cur, err := r.Database.Collection(config.DB_COLLECTION_USERS).Find(r.Ctx, filter, findOptions)
	if err != nil {
		return nil
	}
	for cur.Next(r.Ctx) {
		var contact models.Contact
		err := cur.Decode(&contact)
		if err != nil {
			return nil
		}
		contacts = append(contacts, &contact)
	}
	return contacts
}
