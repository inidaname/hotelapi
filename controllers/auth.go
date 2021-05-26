package controllers

import "github.com/gofiber/fiber/v2"

func UserLogin(c *fiber.Ctx) error {
	return c.Status(200).JSON(&fiber.Map{
		"success": true,
		"message": "You are welcome to Hotel API",
	})
}
