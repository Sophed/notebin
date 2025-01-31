package api

import (
	"encoding/json"
	"notebin/sessions"
	"notebin/storage"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/sophed/lg"
	"golang.org/x/crypto/bcrypt"
)

type req struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type resp struct {
	Error string `json:"error,omitempty"`
	Token string `json:"token,omitempty"`
}

func Login(c *fiber.Ctx) error {
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
	user, err := storage.METHOD.FindUserByUsername(username)
	if err != nil {
		resp, _ := json.Marshal(resp{
			Error: "Couldn't find an account with this username",
		})
		return c.SendString(string(resp))
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password))
	if err != nil {
		resp, _ := json.Marshal(resp{
			Error: "Incorrect login details",
		})
		return c.SendString(string(resp))
	}
	resp, _ := json.Marshal(resp{
		Token: sessions.NewSession(user.ID),
	})
	lg.Info("[login] " + user.Username + " - " + user.ID)
	return c.SendString(string(resp))
}
