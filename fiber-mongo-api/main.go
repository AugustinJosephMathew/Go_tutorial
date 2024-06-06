package main

import (
	"fiber-mongo-api/configs" //add this
	"fiber-mongo-api/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	//run database
	configs.ConnectDB()
	routes.UserRoute(app)

	app.Listen(":3000")
}
