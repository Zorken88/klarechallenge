package routes

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
)

func SwaggerRoute(a *fiber.App) {
	// Create a group for swagger
	route := a.Group("/swagger")

	// Routes for Get Method
	route.Get("*", swagger.Handler)
}
