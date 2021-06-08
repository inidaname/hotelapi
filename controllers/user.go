package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/inidaname/hotelapi/models"
	"github.com/kamva/mgm/v3"
)

func UserUpdate(c *fiber.Ctx) error {
	payload := models.NewUser(models.User{})
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"success": false,
			"message": "Something went wrong",
			"status":  fiber.StatusInternalServerError,
		})
	}

	//@TODO: fine changed data an set for updates

	userColl, err := mgm.Coll(payload).UpdateByID(mgm.Ctx(), c.Params("id"), payload)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(&fiber.Map{
			"success": false,
			"message": "Something went wrong",
			"status":  fiber.StatusNotFound,
		})
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"success": true,
		"message": "Updated successfully",
		"status":  fiber.StatusOK,
		"data":    userColl,
	})
}
