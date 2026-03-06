package middlewares

//Before request reaches the actual handler, middleware runs first.
//acts as the middleman b/w the user and the controller

import (
	"habit-tracker/config"

	jwtware "github.com/gofiber/contrib/jwt" //this is the Fiber JWT middleware package
	"github.com/gofiber/fiber/v2"
)

// this function returns another funciton
func Protected() func(c *fiber.Ctx) error {
	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{
			Key: config.JwtSecret,
		},
	})

}
