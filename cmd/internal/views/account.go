package views

import (
	"notebin/cmd/internal/components"

	"github.com/gofiber/fiber/v2"
)

func ViewAccount(c *fiber.Ctx) error {
	c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
	if isHx(c) {
		return c.SendString(components.Render(components.Account()))
	}
	return c.Redirect("/")
}
