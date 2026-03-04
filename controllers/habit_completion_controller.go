package controllers

import (
	"habit-tracker/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func CompletedHabitController (c *fiber.Ctx) error {
	//from json str to integer
	habitID,err := strconv.Atoi(c.Params("id"))
	if err!=nil{
		return c.Status(404).JSON(fiber.Map{
			"error":"ID invalid",
		})
	}

	//calling the service layer 
	completion , err := services.CompleteHabit(uint(habitID))
	if err!=nil{
		return c.Status(500).JSON(fiber.Map{
			"error":"Could not mark as completed",
		})
	}
	//converts the go struct to JSON format  and if everything works okay c.JSON() automatically returns nil for error return stmnt 
	return c.JSON(completion)
}
