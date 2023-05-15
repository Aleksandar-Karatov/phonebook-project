package repository

import "github.com/pkg/errors"

var (
	ErrAlreadyExists = errors.New("this record already exists in the database")
	ErrNotFound      = errors.New("Record for this contact was not found")
)
