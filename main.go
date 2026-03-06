package main

import (
	"habit-tracker/config"
	"habit-tracker/controllers"
	"habit-tracker/middlewares"

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
	//jWT
	app.Post("/auth/register", controllers.RegisterHandler)
	app.Post("/auth/login", controllers.LoginHandler)
	//middleware --> does not give direct acess routes to the middleware
	//grouping--> groups the routes that share the same base path & the same middleware
	protected := app.Group("/api", middlewares.Protected())

	protected.Get("/habits", controllers.GetAllHabitsHandler)
	protected.Post("/habits", controllers.AddHabitHandler)
	protected.Post("/habits/:id/complete", controllers.CompletedHabitController)
}
