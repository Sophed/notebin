package views

import "github.com/gofiber/fiber/v2"

func isHx(c *fiber.Ctx) bool {
	return c.Get("HX-Request") == "true"
}
