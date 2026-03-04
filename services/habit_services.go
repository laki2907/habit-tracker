package services

import (
	"habit-tracker/config"
	"habit-tracker/models"
)

//service layer handles DB operations only no fiber is written here

func CreateHabit(name string) (models.Habit, error) {

	habit := models.Habit{
		Name:   name,
		Streak: 0,
	}

	if err := config.DB.Create(&habit).Error; err != nil {
		return models.Habit{}, err
	}

	return habit, nil

}

func GetAllHabits(page int, limit int) ([]models.Habit, error) {
	
	//incase of negative or 0 page or limit values
	if page < 1 {
		page = 1
	}

	if limit < 1 {
		limit = 10
	}
	//offset --> how many records we want to skip / from which habit u want to fetch
	offset := (page - 1) * limit

	//creating a slice which can contain several habits
	var habits []models.Habit

	err := config.DB.
		Limit(limit).
		Offset(offset).
		Find(&habits).Error
	if err != nil {
		return nil, err
	}

	return habits, nil

}

func GetHabitByID(id int) (models.Habit, error) {
	//query DB
	var habit models.Habit
	if err := config.DB.First(&habit, id).Error; err != nil {
		return models.Habit{}, err
	}

	return habit, nil
}

func UpdateHabit(id int, name string) (models.Habit, error) {
	var habit models.Habit
	if err := config.DB.First(&habit, id).Error; err != nil {
		return models.Habit{}, err
	}

	habit.Name = name

	if err := config.DB.Save(&habit).Error; err != nil {
		return models.Habit{}, err
	}

	return habit, nil

}

func Deletehabit(id int) error {

	var habit models.Habit

	if err := config.DB.First(&habit, id).Error; err != nil {
		return err
	}

	if err := config.DB.Delete(&habit).Error; err != nil {
		return err
	}

	return nil

}
