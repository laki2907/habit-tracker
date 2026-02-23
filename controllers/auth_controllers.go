package controllers

import (
	"fmt"
	"habit-tracker/config"
	"habit-tracker/models"
	"log"
)

func AddHabit(name string) {
	habit := models.Habit{
		Name:   name,
		Status: "pending",
	}

	result := config.DB.Create(&habit)
	if result.Error != nil {
		log.Fatal("Error inserting habit", result.Error)
		return //helps in exiting if error found instead of continuing
	}
	fmt.Println("Habit added successfully ID: ", habit.ID)

}
