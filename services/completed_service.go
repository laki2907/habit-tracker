package services

import(
	"habit-tracker/models"
	"habit-tracker/config"
	"time"
)

func CompleteHabit(habitID uint)(models.HabitCompletion,error){
	//creating a completion object to store the raw data in struct format to save in DB
	completion := models.HabitCompletion{
		HabitID : habitID,
		Completed: true,
		Date: time.Now(),
	}

	if err:=config.DB.Create(&completion).Error;err!=nil{
		return models.HabitCompletion{},err
	}

	return completion,nil
	
}