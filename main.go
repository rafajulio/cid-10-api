package main

import (
	"cid-10-api/configs"
	"cid-10-api/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	//run db
	configs.ConnectDB()

	//routes
	routes.CidRoute(app)
	routes.UserRoute(app)
	
	app.Listen(":6000")
}
