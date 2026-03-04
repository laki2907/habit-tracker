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
	app.Get("/habits", controllers.GetAllHabitsHandler)
	app.Get("/habits/:id", controllers.GetByIdHandler)
	app.Put("/habits/:id", controllers.UpdateHabitHandler)
	app.Delete("/habits/:id", controllers.DeleteHabitHandler)
	app.Post("/habits/:id/completed", controllers.CompletedHabitController)
	app.Listen(":3000")
}
