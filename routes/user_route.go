package routes

import (
	"cid-10-api/controllers"

	"github.com/gofiber/fiber/v2"
)

func UserRoute(app *fiber.App){
	app.Post("/user", controllers.CreateUser)
	app.Post("/user/login", controllers.Login)
}