package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/inidaname/hotelapi/config"
	"github.com/inidaname/hotelapi/routes"
	"github.com/joho/godotenv"
)

func main() {
	app := fiber.New()

	err := godotenv.Load()

	if err != nil {
		panic(err)
	}

	config.ConnectDB()

	app.Use(cors.New())

	app.Use(func(c *fiber.Ctx) error {
		if c.Is("json") {
			return c.Next()
		}
		return c.SendString("Only JSON allowed!")
	})

	routes.SetupRoutes(app)

	app.Listen(":3000")
}
