package main

import (
	"gofiber-starterkit-mysql/config"
	"log"
)

func main() {
	config.NewEnvirontment()
	logger := config.NewLogger()
	validation := config.NewValidator()
	db := config.NewMysql()
	app := config.NewFiber()

	config.Initialize(&config.Configuration{
		Logger:     logger,
		Validation: validation,
		DB:         db,
		App:        app,
	})

	if err := app.Listen(":3000"); err != nil {
		log.Fatalf("gofiber error: %s", err.Error())
	}
}
