package controllers

import (
	"habit-tracker/services"

	"github.com/gofiber/fiber/v2" //fiber framework --> c *Ctx : context of the current request
)

func RegisterHandler(c *fiber.Ctx) error {
	type Req struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	//fetching the JSON body sent by the client and storing it in req
	var req Req
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid Input!",
		})
	}
	//next call the service layer
	//if there is any problem in the service layer then an internal server error will be sent
	user, err := services.RegisterUser(req.Name, req.Email, req.Password)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Could not register",
		})
	}
	return c.Status(201).JSON(user)

}
func LoginHandler(c *fiber.Ctx) error {
	type Req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var req Req
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{ //400--> bad request
			"error": "Invalid input",
		})

	}
	token, err := services.LoginUser(req.Email, req.Password)
	if err != nil { //401-->Unauthourized
		return c.Status(401).JSON(fiber.Map{
			"error": "Invalid credentials",
		})

	}
	return c.JSON(fiber.Map{"token": token})

}
