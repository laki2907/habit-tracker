package controllers

import (
	"habit-tracker/config"
	"habit-tracker/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
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

	//query DB
	var habit models.Habit
	if err := config.DB.First(&habit, idInt).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(404).JSON(fiber.Map{
				"error": "Habit not found",
			})
		}
		return c.Status(500).JSON(fiber.Map{
			"error": "Database error",
		})
	}

	//returning the habit of desired id
	return c.JSON(habit)

}
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

	//Fetching the existing habit in that id
	var habit models.Habit
	if err := config.DB.First(&habit, idInt).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(404).JSON(fiber.Map{
				"error": "Habit not found",
			})
		}
		return c.Status(500).JSON(fiber.Map{
			"error": "Database error",
		})
	}
	//update fields in habit object
	habit.Name = req.Name

	//save to DB
	if err := config.DB.Save(&habit).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Not able to update in DB",
		})
	}

	return c.JSON(habit)
}
func DeleteHabitHandler(c *fiber.Ctx) error {
	//get the id u want to del from param and typecast
	IdParam := c.Params("id")
	IdInt, err := strconv.Atoi(IdParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid habit Id",
		})
	}
	//Get the habit frm the particular id
	var habit models.Habit
	if err := config.DB.First(&habit, IdInt).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Habit not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Database error",
		})
	}

	//del from DB
	if err := config.DB.Delete(&habit, IdInt).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Was not able to delete habit",
		})
	}
	return c.JSON(fiber.Map{
		"message": "Deleted sucessfully",
		"id":      habit.ID,
	})

}
