package api

import (
	"encoding/json"
	"notebin/data"
	"notebin/sessions"
	"notebin/storage"
	"strings"
	"time"
	"unicode"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/sophed/lg"
)

func Register(c *fiber.Ctx) error {
	var req req
	err := json.Unmarshal(c.Body(), &req)
	if err != nil {
		resp, _ := json.Marshal(resp{
			Error: "Malformed request",
		})
		return c.SendString(string(resp))
	}
	if req.Username == "" || req.Password == "" {
		resp, _ := json.Marshal(resp{
			Error: "One or more fields are empty",
		})
		return c.SendString(string(resp))
	}
	username := strings.ToLower(req.Username)
	if !validUsername(username) {
		resp, _ := json.Marshal(resp{
			Error: "Invalid username",
		})
		return c.SendString(string(resp))
	}
	if !validPassword(req.Password) {
		resp, _ := json.Marshal(resp{
			Error: "Invalid password",
		})
		return c.SendString(string(resp))
	}
	user := data.User{
		ID:           newUUID(),
		Username:     username,
		PasswordHash: data.Hash([]byte(req.Password)),
		Notes:        []string{},
		Timestamp:    time.Now().UnixMilli(),
	}
	err = storage.METHOD.AddUser(&user)
	if err != nil {
		var msg string
		if err == storage.ErrConflict {
			msg = "An account with this username already exists"
		} else {
			msg = "Failed to create your account"
		}
		resp, _ := json.Marshal(resp{
			Error: msg,
		})
		return c.SendString(string(resp))
	}
	resp, _ := json.Marshal(resp{
		Token: sessions.NewSession(user.ID),
	})
	lg.Info("[register] " + user.Username + " - " + user.ID)
	return c.SendString(string(resp))
}

func newUUID() string {
	return uuid.NewString() // TODO: check for conflicts
}

func validUsername(username string) bool {
	if len(username) > 128 {
		return false
	}
	for _, c := range username {
		if !unicode.IsLetter(c) && !unicode.IsNumber(c) && c != '-' {
			return false
		}
	}
	return true
}

func validPassword(password string) bool {
	return !(len(password) > 128)
}
