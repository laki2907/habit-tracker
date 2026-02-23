package main

import (
	"habit-tracker/config"
	"habit-tracker/controllers"
)

func main() {
	//connection to db
	config.ConnectDB()

	//testing 
	controllers.AddHabit("Drink water")
	controllers.AddHabit("sleep")

}
