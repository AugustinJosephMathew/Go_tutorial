package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Define a dynamic route that captures a user's ID
	app.Get("/users/:id", func(c *fiber.Ctx) error {	
		// Get the user ID from the route parameters
		userID := c.Params("id")

		return c.SendString("User ID: " + userID)
	})

	app.Get("/json", func(c *fiber.Ctx) error {
		// Create a JSON response
		response := fiber.Map{
			"message": "Hello, Fiber!",
		}

		// Send the JSON response
		return c.JSON(response)
	})
	app.Get("/error", func(c *fiber.Ctx) error {
		// Simulate an error
		err := ()
		if err != nil {
			// Trigger a panic with the error
			panic(err)
		}
		return c.SendString("No error occurred")
	})

	// Define an error handler
	app.Use(func(c *fiber.Ctx) error {
		if err := recover(); err != nil {
			// Handle the error and respond with an error message
			return c.Status(500).SendString("Internal Server Error")
		}
		return c.Next()
	})

	app.Listen(":3000")
}
