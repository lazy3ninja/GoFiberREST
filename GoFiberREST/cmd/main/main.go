package main

import (
	"example.com/hello/pkg/routes"
	"github.com/gofiber/fiber/v2"
)

//context is used to manage cancellations timeouts and deadlines :)

func main() {
	app := fiber.New()
	// config.ConnectDb()
	routes.CardRoute(app)
	app.Listen(":8000")

}
