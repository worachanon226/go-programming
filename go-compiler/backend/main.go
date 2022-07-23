package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/servercheck", func(c *fiber.Ctx) error{
		return c.SendString("Server is OK.")
	})

	app.Post("/run/cpp", func(c *fiber.Ctx) error{

		GenerateFile("cpp",string(c.Body()))

		return c.SendString("file is generated")
	})

	fmt.Printf("Server is running on port 8000.")
	log.Fatal(app.Listen(":8000"))
}