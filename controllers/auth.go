package controllers

import (
	"log"

	"github.com/asaskevich/govalidator"
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
	_, err := govalidator.ValidateStruct(payload)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"success": false,
			"message": err.Error(),
			"status":  fiber.StatusBadRequest,
		})
	}
	log.Println(payload)

	userColl := mgm.Coll(payload)

	if err := userColl.Create(payload); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"success": false,
			"message": "Something went wrong",
			"status":  fiber.StatusInternalServerError,
		})
	}

	if err := payload.DefaultModel.Creating(); err != nil {
		log.Println("This happened")
	}

	return c.Status(fiber.StatusCreated).JSON(&fiber.Map{
		"success": true,
		"message": "User created successfully",
		"status":  fiber.StatusCreated,
		"data":    payload,
	})
}
