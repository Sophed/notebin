package storage

import (
	"notebin/data"
)

var METHOD StorageMethod

type StorageMethod interface {
	AddUser(*data.User) error
	RemoveUser(*data.User) error
	FindUser(string) (*data.User, error)
	AddNote(*data.Note) error
	RemoveNote(*data.Note) error
	FindNote(string) (*data.Note, error)
	GetNotes(*data.User) (*[]data.Note, error)
	Test() error
}
