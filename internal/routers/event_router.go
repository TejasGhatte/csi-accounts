package routers

import (
	"csi-accounts/internal/controllers"

	"github.com/gofiber/fiber/v2"
)

func EventRouter(app *fiber.App) {
	eventRouter := app.Group("/events")

	eventRouter.Get("/", controllers.GetEvents)
	eventRouter.Get("/:eventID", controllers.GetEvent)
	eventRouter.Post("/", controllers.CreateEvent)
	eventRouter.Patch("/:eventID", controllers.UpdateEvent)
}