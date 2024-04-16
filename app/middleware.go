package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wawew/golang-simple-user/entity"
)

func ErrorHandler(c *fiber.Ctx) error {
	err := c.Next()
	if err != nil {
		if err == entity.ErrNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Entity not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err})
	}
	return nil
}
