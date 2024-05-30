package middleware

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func Logger() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		fmt.Println("Request recieved")
		return c.Next()
	}
}
