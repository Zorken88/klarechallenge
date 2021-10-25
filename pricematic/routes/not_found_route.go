package routes

import "github.com/gofiber/fiber/v2"

// Generate a not found route
func NotFoundRoute(a *fiber.App) {
	a.Use(
		func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": true,
				"msg":   "The url that you trying doesn't exists",
			})
		},
	)
}
