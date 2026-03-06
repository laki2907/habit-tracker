package config

import "time"

var JwtSecret = []byte("super_secret_change_me")//move to env file
var JwtExpiry = time.Hour * 24 //validity of the token 
