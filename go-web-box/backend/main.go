package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

type List struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

func main() {

	app := fiber.New()

	lists := []List{}

	app.Post("/addlists", func(c *fiber.Ctx) error {

		list := &List{}

		if err := c.BodyParser(list); err != nil {
			return err
		}

		list.ID = len(lists) + 1
		lists = append(lists, *list)

		return c.JSON(list)
	})

	log.Fatal(app.Listen(":4040"))
}
