package main

import (
	"log"
	"notebin/cmd/internal/views"
	"notebin/config"
	"notebin/storage"
	"notebin/storage/mongodb"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/sophed/lg"
)

func main() {
	err := config.Load("config.kdl")
	if err != nil {
		lg.Fatl(err)
	}
	storage.METHOD = &mongodb.StorageMongo{
		URI:       config.Get().Mongo.URI,
		CollUsers: config.Get().Mongo.CollUsers,
		CollNotes: config.Get().Mongo.CollNotes,
	}
	err = storage.METHOD.Test()
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
