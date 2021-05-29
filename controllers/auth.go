package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/inidaname/hotelapi/models"
)

func UserLogin(c *fiber.Ctx) error {
	user := models.User{}

	if err := c.BodyParser(&user); err != nil {
		return err
	}

	return c.Status(200).JSON(&fiber.Map{
		"success": true,
		"message": "You are welcome to Hotel API",
	})
}
