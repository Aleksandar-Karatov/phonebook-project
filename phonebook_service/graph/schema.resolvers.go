package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.33

import (
	"context"
	"encoding/base64"
	"errors"
	"log"

	"github.com/phonebook-project/graph/model"
	"github.com/phonebook-project/models"
	"github.com/phonebook-project/repository"
	"github.com/phonebook-project/services"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreateContact is the resolver for the createContact field.
func (r *mutationResolver) CreateContact(ctx context.Context, name string, phones []string, notes *string, email string) (*model.Contact, error) {
	objId := primitive.NewObjectID()
	contact := models.Contact{
		Id:     objId,
		Name:   name,
		Phones: phones,
		Notes:  *notes,
		Email:  email,
	}
	output, err := services.CreateContact(&repository.Repo{Database: r.Database, Ctx: r.Ctx}, contact)
	if err != nil {
		return nil, err
	}
	return &model.Contact{ID: output.Id.Hex(), Name: output.Name, Phones: output.Phones, Notes: &output.Notes, Email: output.Email, ImageID: output.ImageID.Hex()}, nil
}

// MergeContacts is the resolver for the mergeContacts field.
func (r *mutationResolver) MergeContacts(ctx context.Context, id string) (*model.Contact, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	output, err := services.MergeContacts(&repository.Repo{Database: r.Database, Ctx: r.Ctx}, objID)
	if err != nil {
		return nil, err
	}
	return &model.Contact{ID: output.Id.Hex(), Name: output.Name, Phones: output.Phones, Notes: &output.Notes, Email: output.Email, ImageID: output.ImageID.Hex()}, nil
}

// UpdateContact is the resolver for the updateContact field.
func (r *mutationResolver) UpdateContact(ctx context.Context, id string, name string, phones []string, email string, notes string) (*model.Contact, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid data")
	}
	contact := models.Contact{
		Id:     objID,
		Name:   name,
		Phones: phones,
		Notes:  notes,
		Email:  email,
	}
	output, err := services.UpdateContact(&repository.Repo{Database: r.Database, Ctx: r.Ctx}, contact)
	if err != nil {
		return nil, err
	}
	return &model.Contact{ID: output.Id.Hex(), Name: output.Name, Phones: output.Phones, Notes: &output.Notes, Email: output.Email, ImageID: output.ImageID.Hex()}, nil
}

// DeleteContact is the resolver for the deleteContact field.
func (r *mutationResolver) DeleteContact(ctx context.Context, id string) (bool, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return false, errors.New("invalid data")
	}
	return services.DeleteContact(&repository.Repo{Database: r.Database, Ctx: r.Ctx}, objID)
}

// UpdateContactPicture is the resolver for the updateContactPicture field.
func (r *mutationResolver) UpdateContactPicture(ctx context.Context, userID string, imageID string, name string, imageData string, contentType string) (*model.Image, error) {
	data := []byte(imageData)
	image := models.Image{
		Data:        data,
		Name:        name,
		ContentType: contentType,
	}

	userObjID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, errors.New("invalid data")
	}

	if userID == imageID {
		image.Id = primitive.NewObjectID()
	} else {
		imageObjID, err := primitive.ObjectIDFromHex(imageID)
		if err != nil {
			return nil, errors.New("invalid data")
		}
		image.Id = imageObjID
	}

	output, err := services.UpdateContactPicture(&repository.Repo{Database: r.Database, Ctx: r.Ctx}, userObjID, image)
	if err != nil {
		return nil, err
	}
	return &model.Image{ImageID: output.Id.Hex(), Name: output.Name, ContentType: output.ContentType, ImageData: string(output.Data)}, nil
}

// Contacts is the resolver for the contacts field.
func (r *queryResolver) Contacts(ctx context.Context) ([]*model.Contact, error) {
	resp, err := services.Contacts(&repository.Repo{Database: r.Database, Ctx: r.Ctx})
	if err != nil {
		return nil, err
	}
	var out []*model.Contact

	for _, contact := range resp {
		out = append(out, &model.Contact{ID: contact.Id.Hex(), Name: contact.Name, Phones: contact.Phones, Notes: &contact.Notes, Email: contact.Email, ImageID: contact.ImageID.Hex()})
	}
	log.Println("CALLED CONTACTS")
	return out, nil
}

// Contact is the resolver for the contact field.
func (r *queryResolver) Contact(ctx context.Context, name *string, phone *string, email *string, id *string) (*model.Contact, error) {
	objID := &primitive.ObjectID{}
	if id != nil {
		if primitive.IsValidObjectID(*id) {
			*objID, _ = primitive.ObjectIDFromHex(*id)
		}
	} else {
		objID = nil
	}

	output, err := services.Contact(&repository.Repo{Database: r.Database, Ctx: r.Ctx}, name, phone, email, objID)
	if err != nil {
		return nil, err
	}
	return &model.Contact{ID: output.Id.Hex(), Name: output.Name, Phones: output.Phones, Notes: &output.Notes, Email: output.Email, ImageID: output.ImageID.Hex()}, nil
}

// Image is the resolver for the image field.
func (r *queryResolver) Image(ctx context.Context, imageID string) (*model.Image, error) {
	objID, err := primitive.ObjectIDFromHex(imageID)
	if err != nil {
		return nil, errors.New("invalid data")
	}

	output, err := services.Image(&repository.Repo{Database: r.Database, Ctx: r.Ctx}, objID)
	if err != nil {
		return nil, err
	}
	log.Println("CALLED IMAGE")

	return &model.Image{ImageID: output.Id.Hex(), Name: output.Name, ContentType: output.ContentType, ImageData: base64.StdEncoding.EncodeToString(output.Data)}, nil
}

// Search is the resolver for the search field.
func (r *queryResolver) Search(ctx context.Context, filter string) ([]*model.Contact, error) {
	resp, err := services.Search(&repository.Repo{Database: r.Database, Ctx: r.Ctx}, filter)
	if err != nil {
		return nil, err
	}
	var out []*model.Contact

	for _, contact := range resp {
		out = append(out, &model.Contact{ID: contact.Id.Hex(), Name: contact.Name, Phones: contact.Phones, Notes: &contact.Notes, Email: contact.Email, ImageID: contact.ImageID.Hex()})
	}

	return out, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
