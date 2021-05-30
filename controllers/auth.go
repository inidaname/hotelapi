package controllers

import (
	"fmt"

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

	user := &models.User{}

	usercoll := mgm.Coll(user).Create(user)

	fmt.Println(usercoll)
	return c.Status(200).JSON(&fiber.Map{
		"success": false,
		"message": "Must provide email and password",
		"status":  400,
	})
}
