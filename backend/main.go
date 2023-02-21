package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/naphattar/GolangAuthLearning/configs"
	"github.com/naphattar/GolangAuthLearning/routes"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{"data": "Fiber is works"})
	})

	configs.ConnectDB()
	routes.UserRoute(app)
	app.Listen(":4000")
}
