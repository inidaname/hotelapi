package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/inidaname/hotelapi/controllers"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", controllers.Home)
	app.Post("/login", controllers.UserLogin)
}
