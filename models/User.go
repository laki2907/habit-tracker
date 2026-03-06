package models

import "time"

type User struct {
	//the orange ones are called struct tags --> they give extra info to lib

	ID        uint      `gorm:"primary key" json:"id"` //gorm-->sets it as primary  key in the DB && json-->when converting frm struct to JSON use id as a field name
	Name      string    `json:"name"`
	Email     string    `gorm:"uniqueIndex" json:"email"`
	Password  string    `json:"_"` //_ means do not include this in the json response body
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// gorm : "..." → database rules
// json : "..." → API response format
