package routes

import (
	"gofiber-starterkit-mysql/internal/domain/example/handler"

	"github.com/gofiber/fiber/v2"
)

type Init struct {
	ExampleHandler handler.ExampleHandler
}

func (i *Init) Setup(app *fiber.App) {
	app.Get("/", i.ExampleHandler.SayHello)
}
