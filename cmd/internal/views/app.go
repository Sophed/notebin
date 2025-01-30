package views

import (
	"notebin/cmd/internal/components"

	"github.com/gofiber/fiber/v2"
)

func ViewApp(c *fiber.Ctx) error {
	c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
	if isHx(c) {
		return c.SendString(components.Render(components.Home()))
	}
	page := components.View("notebin",
		components.Home(),
	)
	return c.SendString(components.Render(page))
}
