package routes

import (
	"github.com/gofiber/fiber/v2"
)

func SetRoutes(a *fiber.App) {

	a.Get("/", func(c *fiber.Ctx) error {

		return c.SendString("Hello, Welcome to fiber , /hey")

	})
	a.Get("/hey", func(c *fiber.Ctx) error {

		return c.SendString("Hello Post , Welcome to fiber , /hi")

	})
	a.Get("/hey/hi", func(c *fiber.Ctx) error {

		return c.SendString("Hello, Welcome to fiber , /hello")

	})

}
