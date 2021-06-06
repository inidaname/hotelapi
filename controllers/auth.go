package controllers

import (
	"time"

	"github.com/asaskevich/govalidator"
	jwt "github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/inidaname/hotelapi/models"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

// Login func
func UserLogin(c *fiber.Ctx) error {
	user := &models.User{}

	if err := c.BodyParser(user); err != nil {
		return err
	}

	if user.Email == "" || user.Password == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(&fiber.Map{
			"success": false,
			"message": "Must provide email and password",
			"status":  400,
		})
	}

	coll := mgm.Coll(user).SimpleFind(user, bson.M{"email": user.Email, "password": user.Password})

	return c.Status(fiber.StatusAccepted).JSON(&fiber.Map{
		"success": true,
		"message": "You are welcome to Hotel API",
		"data":    coll,
	})
}

// Create user
func CreateUser(c *fiber.Ctx) error {

	payload := models.NewUser(models.User{})
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"success": false,
			"message": "Something went wrong",
			"status":  fiber.StatusInternalServerError,
		})
	}

	// Validating fields
	if _, err := govalidator.ValidateStruct(payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"success": false,
			"message": err.Error(),
			"status":  fiber.StatusBadRequest,
		})
	}

	// hashing password
	bytes, err := bcrypt.GenerateFromPassword([]byte(payload.Password), 17)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"success": false,
			"message": "Something went wrong",
			"status":  fiber.StatusInternalServerError,
		})
	}

	payload.Password = string(bytes)

	userColl := mgm.Coll(payload)

	if err := userColl.Create(payload); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"success": false,
			"message": "Something went wrong",
			"status":  fiber.StatusInternalServerError,
		})
	}

	// implementing token
	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = payload.FullName
	claims["user"] = true
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"success": false,
			"message": "Something went wrong",
			"status":  fiber.StatusInternalServerError,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(&fiber.Map{
		"success": true,
		"message": "User created successfully",
		"status":  fiber.StatusCreated,
		"data":    payload,
		"token":   t,
	})
}
