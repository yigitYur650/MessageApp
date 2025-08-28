// cmd/main.go

package main

import (
	"hatirlagpt/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()

	app.Post("/register", handlers.Register)
	app.Post("/login", handlers.Login)
	app.Post("/messages/send", handlers.SendMessage)

	app.Use(cors.New())
	app.Listen(":3000")
}
