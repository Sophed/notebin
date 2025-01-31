package views

import (
	"notebin/cmd/internal/components"

	"github.com/gofiber/fiber/v2"
	hx "maragu.dev/gomponents-htmx"
	. "maragu.dev/gomponents/html"
)

func ViewApp(c *fiber.Ctx) error {
	c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
	if isHx(c) {
		if loggedIn(c) {
			return c.SendString(components.Render(components.App()))
		} else {
			return c.SendString(components.Render(components.Home()))
		}
	}
	page := components.View("notebin",
		Div(hx.Get("/"), hx.Trigger("load"), hx.Swap("outerHTML"), hx.Headers(components.HX_TOKEN)),
	)
	return c.SendString(components.Render(page))
}
