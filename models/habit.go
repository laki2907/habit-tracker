package models

import "gorm.io/gorm"

//Habit represents a habit table 
type Habit struct{
	gorm.Model
	Name string
	Status string
	

}
