package router

import (
	"manualVuln/handler"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// SetupRoutes setup router api
func SetupRoutes(app *fiber.App) {
	// Middleware
	api := app.Group("/api", logger.New())
	api.Get("/", handler.Health)

	// Users
	user := api.Group("/user")
	user.Get("/", handler.GetAllUsers)
	user.Get("/:id", handler.GetUser)
	user.Post("/",  handler.CreateUser)
	user.Put("/:id", handler.UpdateUser)
	user.Delete("/:id", handler.DeleteUser)
}
