package main

import (
	"log"
	"sort"

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

	app.Patch("/update", func(c *fiber.Ctx) error {

		sort.Slice(lists, func(i, j int) bool {
			return lists[i].Title < lists[j].Title
		})

		return c.JSON(lists)
	})

	app.Post("/addlists", func(c *fiber.Ctx) error {

		list := &List{}

		if err := c.BodyParser(list); err != nil {
			return err
		}

		list.ID = len(lists) + 1
		lists = append(lists, *list)

		return c.JSON(list)
	})

	app.Get("/get", func(c *fiber.Ctx) error {
		return c.JSON(lists)
	})

	app.Get("/get/:id", func(c *fiber.Ctx) error {

		id, err := c.ParamsInt("id")

		if err != nil {
			return c.Status(401).SendString("Invalid ID")
		}

		var list List

		for i, t := range lists {
			if t.ID == id {
				list = lists[i]
			}
		}
		return c.JSON(list)
	})

	log.Fatal(app.Listen(":4040"))
}
