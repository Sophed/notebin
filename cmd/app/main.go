package main

import (
	"log"
	"notebin/cmd/internal/views"
	"notebin/config"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/sophed/lg"
)

func main() {
	err := config.Load("config.kdl")
	if err != nil {
		lg.Fatl(err)
	}
	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})

	app.Get("/", views.ViewApp)
	app.Get("/account", views.ViewAccount)

	app.Static("/static", config.Get().StaticDir)

	port := strconv.Itoa(config.Get().Port)
	lg.Info("listening on http://127.0.0.1:" + port + "/")
	log.Fatal(app.Listen(":" + port))
}
