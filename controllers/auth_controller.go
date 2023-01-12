package controllers

import (
	"github.com/gofiber/fiber/v2"
)

// Login is the controller for the login route

func Login(c *fiber.Ctx) error {
	return c.JSON(&fiber.Map{"message": "This is the login route"})
}
