package views

import (
	"notebin/cmd/internal/components"
	"notebin/sessions"

	"github.com/gofiber/fiber/v2"
	"github.com/sophed/lg"
)

func loggedIn(c *fiber.Ctx) bool {
	return sessions.ValidToken(c.Get("token"))
}

func ViewAccount(c *fiber.Ctx) error {
	c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
	if !isHx(c) {
		return c.Redirect("/")
	}
	if loggedIn(c) {
		lg.Dbug("logged in")
		return c.SendString(components.Render(components.Account()))
	} else {
		return c.SendString(components.Render(components.Login()))
	}
}
