package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

var jobID = ""

func main() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173",
		AllowHeaders: "Origin, Context-Type, Accept",
	}))

	app.Get("/servercheck", func(c *fiber.Ctx) error{
		return c.SendString("Server is OK.")
	})

	app.Get("/output/cpp", func(c *fiber.Ctx) error{
		return c.SendString(ExecuteCpp(jobID))
	})

	app.Get("/output/py", func(c *fiber.Ctx) error{
		return c.SendString(ExecutePy(jobID))
	})

	app.Post("/run/cpp", func(c *fiber.Ctx) error{
		jobID = GenerateFile("cpp",string(c.Body()))
		return c.SendString("Cpp file is generated")
	})

	app.Post("/run/py", func(c *fiber.Ctx) error{
		jobID = GenerateFile("py",string(c.Body()))
		return c.SendString("Python file is generated")
	})

	fmt.Printf("Server is running on port 8080.\n")
	log.Fatal(app.Listen(":8080"))
}