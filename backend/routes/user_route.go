package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/naphattar/GolangAuthLearning/controllers"
)

func UserRoute(app *fiber.App) {
	app.Post("/user", controllers.CreateUser)
}
