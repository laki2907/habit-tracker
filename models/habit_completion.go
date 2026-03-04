package models

import "time"

type HabitCompletion struct{
	ID uint  `gorm:"primary key"`
	HabitID uint
	Completed bool
	Date time.Time
}