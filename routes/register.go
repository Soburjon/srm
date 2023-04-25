package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/work/SRM/api/V1/controllers"
)

func RegisterRoutes(app *fiber.App, c *controllers.Api) {
	routes := app.Group("/register")
	routes.Post("/login/", c.Login)
}
