package routers

import (
	"csi-accounts/internal/routes"

	"github.com/gofiber/fiber/v2"
	// Assuming routes are defined in the "internal/routes" package
)

// welcome is a handler for the /api endpoint
func welcome(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Welcome to the API!",
	})
}

// SetUp initializes the routes for the application
func SetUp(app *fiber.App) {
	app.Get("/api", welcome)                // Welcome route
	app.Post("/api/users", routes.CreateUser) // CreateUser route from routes package
}
