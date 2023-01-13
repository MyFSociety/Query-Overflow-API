package controllers

import (
	"github.com/gofiber/fiber/v2"
)

func CreateUser(c *fiber.Ctx) error {
	return c.JSON(&fiber.Map{"message": "This is the create user route"})
}

func UpdateUser(c *fiber.Ctx) error {
	return c.JSON(&fiber.Map{"message": "This is the update user route"})
}

func DeleteAUser(c *fiber.Ctx) error {
	return c.JSON(&fiber.Map{"message": "This is the delete a user route"})
}
