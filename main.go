package main

import (
	"harry/query-overflow/routers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// album represents data about a record album.

func main() {
	router := fiber.New()

	// Default config
	router.Use(cors.New())

	// routers
	router.Get("/", defaultRoute)

	// Auth
	routers.LoginRouters(router)

	// Start server
	router.Listen(":8080")

}

func defaultRoute(c *fiber.Ctx) error {
	return c.JSON(&fiber.Map{"message": "This is the default route"})
}
