package routes

import (
	"cid-10-api/controllers"

	"github.com/gofiber/fiber/v2"
)

func CidRoute(app *fiber.App){
	// * Post
	app.Post("/cid", controllers.CreateCid)

	// * Get
	app.Get("/cid", controllers.GetCid)
	app.Get("/cid/validate", controllers.ValidateCid)
}