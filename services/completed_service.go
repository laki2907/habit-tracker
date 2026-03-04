package services

import (
	"habit-tracker/config"
	"habit-tracker/models"
	"time"
)

// uint --> unassigned integer -->
func CompleteHabit(habitID uint) (models.HabitCompletion, error) {

	//truncate removes the hour part and keeps only the date part ex: 2026-03-04 00:00:00
	today := time.Now().Truncate(24 * time.Hour)

	//check if aldready completed today -- to avoid duplicates
	var exists models.HabitCompletion
	err := config.DB.
		Where("habit_id = ? AND date = ?", habitID, today).
		First(&exists).Error
	//if err!=nil then the record aldready exists

	//if a prexisting row is found then no errors hence duplicate spotted so return the old one and exit
	if err == nil {
		return exists, nil
	}

	//creating a completion object to store the raw data in struct format to save in DB
	completion := models.HabitCompletion{
		HabitID:   habitID,
		Completed: true,
		Date:      time.Now(),
	}

	if err := config.DB.Create(&completion).Error; err != nil {
		return completion, err
	}

	//get the habit with the particular id
	//we need this for streak updation
	var habit models.Habit
	if err := config.DB.First(&habit, habitID).Error; err != nil {
		return models.HabitCompletion{}, err
	}

	//check yesterdays completion
	yesterday := today.AddDate(0, 0, -1)

	var yesterdayCompletion models.HabitCompletion
	err = config.DB.
		Where("habit_id = ? AND date = ?", habitID, yesterday).
		First(&yesterdayCompletion).Error
	if err == nil {
		habit.Streak += 1
	} else {
		habit.Streak = 1
	}
	//as we have updated new streak values use save and store it back into DB
	config.DB.Save(&habit)

	return completion, nil

}
