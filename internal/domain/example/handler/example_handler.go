package handler

import "github.com/gofiber/fiber/v2"

type ExampleHandler interface {
	SayHello(c *fiber.Ctx) error
}
type example struct {
}

func NewExampleHandler() ExampleHandler {
	return &example{}
}
func (h *example) SayHello(c *fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{"message": "hello world"})
}
