package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       uint
	Username string
	Email    string
	Password string
}

// HashPassword hashes the user's password
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func main() {
	U := User{}
	app := fiber.New()
	fmt.Println("Enter the id")
	fmt.Scan(&U.ID)
	fmt.Println("Enter the username ")
	fmt.Scan(&U.Username)
	fmt.Println("Enter the password")
	fmt.Scan(U.Password)
	fmt.Printf("The password is ")
	// Define your routes and handle database operations
	// ...
	app.Get("/login", func(c *fiber.Ctx) error {
		hash, err := HashPassword(U.Password)
		fmt.Printf("The password is  %v", hash)
		if err != nil {
			panic(err)
		}

		return c.SendString("The password is" + hash)

	})

	app.Listen(":3006")
}
