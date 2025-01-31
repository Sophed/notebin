package storage

import (
	"errors"
	"notebin/data"
)

var METHOD StorageMethod
var ErrConflict = errors.New("account already exists")

type StorageMethod interface {
	AddUser(*data.User) error
	RemoveUser(*data.User) error
	FindUserByID(string) (*data.User, error)
	FindUserByUsername(string) (*data.User, error)
	AddNote(*data.Note) error
	RemoveNote(*data.Note) error
	FindNote(string) (*data.Note, error)
	GetNotes(*data.User) (*[]data.Note, error)
	Test() error
}
