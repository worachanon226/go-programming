package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

var jobID = ""

func main() {
	app := fiber.New()

	app.Get("/servercheck", func(c *fiber.Ctx) error{
		return c.SendString("Server is OK.")
	})

	app.Get("/output/cpp", func(c *fiber.Ctx) error{
		return c.SendString(OutputCpp(jobID))
	})

	app.Get("/output/py", func(c *fiber.Ctx) error{
		return c.SendString(OutputPy(jobID))
	})

	app.Post("/run/cpp", func(c *fiber.Ctx) error{
		jobID = GenerateFile("cpp",string(c.Body()))
		return c.SendString("Cpp file is generated")
	})

	app.Post("/run/py", func(c *fiber.Ctx) error{
		jobID = GenerateFile("py",string(c.Body()))
		return c.SendString("Python file is generated")
	})

	fmt.Printf("Server is running on port 8000.\n")
	log.Fatal(app.Listen(":8000"))
}