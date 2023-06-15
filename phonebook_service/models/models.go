package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Contact struct {
	Id      primitive.ObjectID `bson:"id"`
	Name    string             `bson:"name"`
	Phones  []string           `bson:"phoneNumbers"`
	Notes   string             `bson:"notes"`
	Email   string             `bson:"email"`
	ImageID primitive.ObjectID `bson:"image_id"`
}

type Image struct {
	Id          primitive.ObjectID `bson:"_id"`
	Name        string             `bson:"name"`
	Data        []byte             `bson:"data"`
	ContentType string             `bson:"contentType"`
}
