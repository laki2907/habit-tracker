package main

import (
	"habit-tracker/config"
	"habit-tracker/controllers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	//connection to db
	config.ConnectDB()

	app := fiber.New() //we are creating a HTTP server
	app.Post("/habits", controllers.AddHabitHandler)
	app.Listen(":3000")

}
