package views

import (
	"notebin/cmd/internal/components"

	"github.com/gofiber/fiber/v2"
)

func ViewLogin(c *fiber.Ctx) error {
	c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
	if isHx(c) {
		return c.SendString(components.Render(components.Login()))
	}
	return c.Redirect("/")
}
