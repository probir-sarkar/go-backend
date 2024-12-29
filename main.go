package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/probir-sarkar/go-backend/database"
	"github.com/probir-sarkar/go-backend/handlers"
)

func main() {
	// Connect to the database
	if err := database.ConnectDatabase(); err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	// Migrate the schema
	// database.Migrate()

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Post("/contact", handlers.SubmitContactForm)

	app.Listen(":3000")
}
