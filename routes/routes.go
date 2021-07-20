package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mnn59/web-new-back/controllers"
)

func Setup(app *fiber.App)  {
	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)
	app.Get("/api/user", controllers.User)
}