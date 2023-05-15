package models

type Contact struct {
	Id     string   `bson:"id"`
	Name   string   `bson:"name"`
	Phones []string `bson:"phoneNumbers"`
	Notes  string   `bson:"notes"`
	Email  string   `bson:"email"`
}
