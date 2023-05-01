package main

import (
	"api/configs"
	"api/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	configs.ConnectToDB()

	routes.UserRoute(app)

	app.Listen(":3000")
}
