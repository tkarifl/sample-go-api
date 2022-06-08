package handler

import "github.com/gofiber/fiber/v2"

// Health handle api status
func Health(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "success", "message": "api is healthy!"})
}
