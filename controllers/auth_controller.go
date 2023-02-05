package controllers

import (
	"context"
	"harry/query-overflow/models"
	"harry/query-overflow/utils"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

// Login is the controller for the login route

func Login(c *fiber.Ctx) error {

	var ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	var user models.LoginUser
	var foundUser models.User

	defer cancel()

	if err := c.BodyParser(&user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(models.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	//use the validator library to validate required fields
	if validationErr := validate.Struct(&user); validationErr != nil {
		return c.Status(http.StatusBadRequest).JSON(models.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": validationErr.Error()}})
	}

	//check if user exists
	err := userCollection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&foundUser)
	defer cancel()

	if err != nil {
		return c.Status(http.StatusNotFound).JSON(models.UserResponse{Status: http.StatusNotFound, Message: "user doesn't exit", Data: &fiber.Map{"data": err.Error()}})
	}

	msg := bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(user.Password))

	if msg != nil {
		return c.Status(http.StatusForbidden).JSON(models.UserResponse{Status: http.StatusForbidden, Message: "credentials not valid", Data: &fiber.Map{"data": msg.Error()}})
	}

	token, _ := utils.GenerateAllTokens(foundUser.Email, foundUser.Name, foundUser.Id.String())

	return c.JSON(&fiber.Map{"token": token, "tokenType": "bearer"})

}
