package services

import (
	"errors" //used to create error msgs
	"habit-tracker/config"
	"habit-tracker/models"
	"time"

	"github.com/golang-jwt/jwt/v5" //used to creare and sign JWT tokens
	"golang.org/x/crypto/bcrypt"   //used for hashing password
)

//REGISTER

func RegisterUser(name, email, password string) (models.User, error) {
	//hashing the passwrod
	//bcrypt expects in bytes
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return models.User{}, err
	}

	user := models.User{
		Name:     name,
		Email:    email,
		Password: string(hash), //converting hashed password (bytes) to string and storing
	}
	//Storing the new user to the DB
	if err := config.DB.Create(&user).Error; err != nil {
		return models.User{}, err
	}
	//while returning the userobj we dont want to expose the password even the hashed one
	user.Password = " "
	return user, nil
}

//LOGIN

func LoginUser(email, password string) (string, error) {
	var user models.User
	if err := config.DB.Where("email=?", email).First(&user).Error; err != nil {
		return " ", errors.New("Invalid credentials")
	}
	//comparing and checking the password is right
	//conv to arr of bytes and compare both
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("Invalid credentials")
	}

	//create JWT token
	//claims is the data stored inside the token this is the payload part of the JWT token
	claims := jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"exp":     time.Now().Add(config.JwtExpiry).Unix(), //current time + expiration time in Unix format
	}
	//JWT token : header.payload.signature
	//header is automaticaaly created (header is the HS256 part)
	//jwt.NewWithClaims(...)--> Create a new JWT token using these claims
	//jwt.SigningMethodHS256 --> signing method/algo used to create the token signature
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//actual signing process:
	signed, err := token.SignedString(config.JwtSecret)
	if err != nil {
		return "", err
	}
	return signed, nil
}
