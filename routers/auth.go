package routers

import (
	"harry/query-overflow/controllers"

	"github.com/gofiber/fiber/v2"
)

func LoginRouters(app *fiber.App) {
	app.Post("/login", controllers.Login)
}
