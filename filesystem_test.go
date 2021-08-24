package main

import (
	"testing"

	"github.com/gofiber/fiber/v2"
)

func Test_FileSystem(t *testing.T) {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
}
