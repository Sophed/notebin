package data

import (
	"github.com/sophed/lg"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID           string   `json:"id" bson:"_id"`
	Username     string   `json:"username" bson:"username"`
	PasswordHash string   `json:"password" bson:"password"`
	Notes        []string `json:"notes" bson:"notes"`
	Timestamp    int64    `json:"timestamp" bson:"timestamp"`
}

func Hash(password []byte) string {
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.MinCost)
	if err != nil {
		lg.Fatl(err)
	}
	return string(hash)
}
