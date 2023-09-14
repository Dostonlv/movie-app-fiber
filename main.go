package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"log"
	"movie/database"
	"movie/router"
)

func main() {
	app := fiber.New()

	api := app.Group("/api", logger.New())
	user := api.Group("/user")
	user.Post("/create", router.User)
	database.InitDB()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	log.Fatal(app.Listen(":3000"))
}
