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

	userColl, err := mgm.Coll(payload).UpdateByID(mgm.Ctx(), c.Params("id"), payload)
}
