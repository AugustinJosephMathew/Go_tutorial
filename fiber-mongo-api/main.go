package main

import (
	"fiber-mongo-api/configs" //add this
	"fiber-mongo-api/controllers"

	"github.com/gofiber/fiber/v2"
)

func UserRoute(app *fiber.App) {
	//All routes related to users comes here

	app.Post("/user", controllers.CreateUser)
	app.Get("/user/:userId", controllers.GetAUser)
	app.Put("/user/:userId", controllers.EditAUser)
	app.Delete("/user/:userId", controllers.DeleteAUser)
	app.Get("/users", controllers.GetAllUsers)
}

func main() {
	app := fiber.New()

	//run database
	configs.ConnectDB()
	UserRoute(app)

	app.Listen(":6000")
}