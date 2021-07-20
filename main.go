package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/mnn59/web-new-back/database"
	"github.com/mnn59/web-new-back/routes"
)

func main() {

	database.Connect()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true, // now front-end can get cookie that we sent from back-end
	}))

	routes.Setup(app)

	app.Listen(":8000")


	//_, err := gorm.Open(mysql.Open("root:@/go_test"), &gorm.Config{})
	//
	//if err != nil {
	//	panic("could not connect to database")
	//}
	//
	//app := fiber.New()
	//
	//app.Get("/", controllers.Hello)
	//
	//app.Listen(":8000")

}
