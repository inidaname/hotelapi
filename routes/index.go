package routes

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
	"github.com/inidaname/hotelapi/controllers"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", controllers.Home)
	app.Post("/login", controllers.UserLogin)
	app.Post("/createuser", controllers.CreateUser)

	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("secret"),
	}))

}
