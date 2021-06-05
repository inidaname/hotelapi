package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/inidaname/hotelapi/models"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
)

func UserLogin(c *fiber.Ctx) error {
	user := &models.User{}

	if err := c.BodyParser(user); err != nil {
		return err
	}

	if user.Email == "" || user.Password == "" {
		return c.Status(400).JSON(&fiber.Map{
			"success": false,
			"message": "Must provide email and password",
			"status":  400,
		})
	}

	coll := mgm.Coll(user).SimpleFind(user, bson.M{"email": user.Email, "password": user.Password})

	return c.Status(200).JSON(&fiber.Map{
		"success": true,
		"message": "You are welcome to Hotel API",
		"data":    coll,
	})
}

func CreateUser(c *fiber.Ctx) error {

	payload := models.NewUser(models.User{})

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"success": false,
			"message": "Something went wrong",
			"status":  fiber.StatusInternalServerError,
		})
	}

	userColl := mgm.Coll(payload)

	if err := userColl.Create(payload); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"success": false,
			"message": "Something went wrong",
			"status":  fiber.StatusInternalServerError,
		})
	}

	return c.Status(fiber.StatusOK).JSON(&fiber.Map{
		"success": false,
		"message": "Must provide email and password",
		"status":  400,
	})
}
