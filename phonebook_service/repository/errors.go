package repository

import "github.com/pkg/errors"

var (
	ErrAlreadyExists = errors.New("this record already exists in the database")
	ErrNotFound      = errors.New("Record for this contact was not found")
	ErrNoDuplicates  = errors.New("No other contact with this name and email")
	ErrCantDelete    = errors.New("Can't delete object")
	ErrCantUpdate    = errors.New("Cant update object")
	ErrInvalidData   = errors.New("Invalid input data")
)
