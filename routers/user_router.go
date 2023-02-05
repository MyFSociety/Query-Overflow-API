package routers

import (
	"harry/query-overflow/controllers"

	"github.com/gofiber/fiber/v2"
)

func UserRouters(app *fiber.App) {
	// accessing the user controller
	app.Get("/users", controllers.GetAllUsers)
	app.Get("/users/:id", controllers.GetAUser)

	// modify the user controller

	app.Put("/users/:id", controllers.UpdateUser)
	app.Delete("/users/:id", controllers.DeleteAUser)

}
