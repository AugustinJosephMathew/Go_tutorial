package main

import (
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/gofiber/fiber/v2"
)

type item struct {
	ID    int     `json:"id"`
	Title string  `json:"title"`
	Price float64 `json:"price"`
}

var ITEMS = []item{
	{1, "Item1", 10.99},
	{2, "Item2", 15.99},
	{3, "Item3", 19.25},
}

func Getallitems(c *fiber.Ctx) error {
	log.Infof("Recieved a %s request for items", c.Method())
	return c.JSON(ITEMS)
}
func Getitem(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.SendString("invalid Id")
	}
	for _, val := range ITEMS {
		if val.ID == id {
			return c.JSON(val)
		}

	}
	return c.Status(http.StatusBadRequest).SendString("Item not found")
}

func Createitem(c *fiber.Ctx) error {
	var it item
	err := c.BodyParser(&it)
	if err != nil {
		return c.Status(http.StatusBadRequest).SendString("Invalid Json")
	}
	it.ID = len(ITEMS) + 1
	ITEMS = append(ITEMS, it)
	return c.JSON(it)
}

func main() {

	app := fiber.New()
	app.Get("/items", Getallitems)
	app.Get("/items/:id", Getitem)
	app.Post("/items/create", Createitem)
	err := app.Listen(":3000")
	if err != nil {
		log.Fatal(err)
	}
}
