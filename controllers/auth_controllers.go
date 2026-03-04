package controllers

import (
	"habit-tracker/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// POST
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

	habit, err := services.CreateHabit(req.Name)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Could not create habit",
		})
	}

	return c.JSON(habit)
}

// GET
func GetAllHabitsHandler(c *fiber.Ctx) error {

	habits, err := services.GetAllHabits()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Not ablle to fetch habits",
		})
	}
	return c.JSON(habits) //converts slice into json array and gives it back to the user
}

// GetBYID
func GetByIdHandler(c *fiber.Ctx) error {
	//read the id from the url
	idParam := c.Params("id")

	//conv the id to int
	idInt, err := strconv.Atoi(idParam)
	if err != nil || idInt < 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid habit id",
		})
	}

	habit, err := services.GetHabitByID(idInt)

	//returning the habit of desired id
	return c.JSON(habit)

}

// PUT
func UpdateHabitHandler(c *fiber.Ctx) error {

	//get id from the url & conversion to integer
	idParam := c.Params("id")
	idInt, err := strconv.Atoi(idParam)
	if err != nil || idInt < 0 {
		return c.Status(404).JSON(fiber.Map{
			"error": "Invalid Id",
		})
	}

	//taking the body (from the request)
	type Request struct {
		Name string `json:"name"`
		//can add more fields if we want later
	}
	var req Request
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid  JSON body ",
		})
	}
	habits, err := services.UpdateHabit(idInt, req.Name)
	return c.JSON(habits)
}

// DELETE
func DeleteHabitHandler(c *fiber.Ctx) error {
	//get the id u want to del from param and typecast
	IdParam := c.Params("id")
	IdInt, err := strconv.Atoi(IdParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid habit Id",
		})
	}


	err = services.Deletehabit(IdInt)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(404).JSON(fiber.Map{
				"error": "Habit not found",
			})
		}
		return c.Status(500).JSON(fiber.Map{
			"error": "Databbase error",
		})
	}

	//if error is nil then this return stmnt works
	return c.JSON(fiber.Map{
		"message": "Deleted sucessfully",
	})

}
