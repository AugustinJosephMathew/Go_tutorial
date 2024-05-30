package main

import (
	"fiber1/config"
	"fiber1/middleware"
	routes "fiber1/router"

	"github.com/gofiber/fiber/v2"
)

// type todo struct {
// 	id    int
// 	name  string
// 	marks float64
// }

// var dos = []todo{
// 	{id: 1, name: "ravi", marks: 34.5},
// 	{id: 2, name: "shiva ", marks: 77.99},
// }

func main() {
	app := fiber.New()
	// app.Use(middleware.Logger())
	// app.Get("/", func(c *fiber.Ctx) error {
	// 	return c.SendString("Hello World")
	// })

	app.Use(middleware.Logger())

	routes.SetRoutes(app)
	dataurl := config.Getpath()

	// app.Get("/testapi", func(c *fiber.Ctx) error {

	// 	return c.Status(200).JSON(fiber.Map{
	// 		"success": true,
	// 		"message": " Api created successfully",
	// 	})
	// })

	app.Put("/hey/hi/hello", func(c *fiber.Ctx) error {

		return c.SendString("Hello, Welcome to fiber , the data url is" + c.Params(dataurl))

	})

	app.Listen(":4002")

}
