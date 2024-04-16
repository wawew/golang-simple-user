package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wawew/golang-simple-user/usecase"
)

func NewController(userUseCase usecase.IUserUseCase) *fiber.App {
	app := fiber.New()
	app.Use(ErrorHandler)

	var DisplayUserRequest struct {
		Userid string `json:"Userid"`
	}

	app.Post("/DisplayUser", func(c *fiber.Ctx) error {
		if err := c.BodyParser(&DisplayUserRequest); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
		}
		user, err := userUseCase.GetUserByID(DisplayUserRequest.Userid)
		if err != nil {
			return err
		}
		return c.JSON(user)
	})

	app.Post("/DisplayAllUsers", func(c *fiber.Ctx) error {
		users, err := userUseCase.GetAllUsers()
		if err != nil {
			return err
		}
		return c.JSON(users)
	})

	return app
}
