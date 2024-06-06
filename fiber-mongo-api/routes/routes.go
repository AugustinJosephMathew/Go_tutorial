package routes

import (
	"fiber-mongo-api/controllers"

	"github.com/gofiber/fiber/v2"
)

func UserRoute(app *fiber.App) {
	//All routes related to users comes here

	app.Post("/user", controllers.CreateUser)
	// app.Get("/user/:userId", controls.GetAUser)
	// app.Put("/user/:userId", controls.EditAUser)
	// app.Delete("/user/:userId", controls.DeleteAUser)
	// app.Get("/users", controls.GetAllUsers)
}
