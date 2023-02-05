package main

import (
	"harry/query-overflow/routers"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	router := fiber.New()

	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	// Default config
	router.Use(cors.New())

	// routers
	router.Get("/", defaultRoute)

	// Auth
	routers.LoginRouters(router)

	// Signup
	routers.SignUpRouters(router)

	// User
	routers.UserRouters(router)

	// Start server
	router.Listen(":" + os.Getenv("PORT"))

}

func defaultRoute(c *fiber.Ctx) error {
	return c.JSON(&fiber.Map{"message": "This is the default route"})
}
