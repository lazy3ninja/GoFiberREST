package routes

import (
	"example.com/hello/pkg/controllers"
	"github.com/gofiber/fiber/v2"
)

func CardRoute(app *fiber.App) {
	app.Post("/card", controllers.CreateCard)
	app.Get("/card/:id", controllers.GetCard)
	app.Get("/cards", controllers.GetCards)
	app.Delete("/card/:id", controllers.DeleteCard)
	app.Put("/card/:id", controllers.UpdateCard)

}
