package controllers

import (
	"habit-tracker/config"
	"habit-tracker/models"

	"github.com/gofiber/fiber/v2"
)

func AddHabitHandler(c *fiber.Ctx) error {
	type Request struct {
		Name string `json:"name"`
	}

	var req Request

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid input",
		})
	}
	habit := models.Habit{
		Name:   req.Name,
		Streak: 0,
	}

	if err := config.DB.Create(&habit).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Could not create a habit",
		})
	}
	return c.JSON(habit)
}

func GetAllHabitsHandler(c *fiber.Ctx) error {
	var habits []models.Habit //creating an empty slice

	if err := config.DB.Find(&habits).Error; err != nil { //config.DB.Find() returns a DB object so to acess only the error val we use .Error
		return c.Status(500).JSON(fiber.Map{
			"error": "Could not fetch items",
		})
	}
	return c.JSON(habits) //converts slice into json array and gives it back to the user
}
