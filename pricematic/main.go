package main

import (
	"log"

	"github.com/KlareTeam/interview-challenges/go/pricematic/database"
	_ "github.com/KlareTeam/interview-challenges/go/pricematic/docs"
	"github.com/KlareTeam/interview-challenges/go/pricematic/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// @title API
// @version 1.0
// @description Swagger docs
// @contact.name Miguel Muñoz
// @contact.email mmunozl88@gmail.com
// BasePath /
func main() {

	// Connected with Database
	database.ConnectDb()

	// Create fiber app
	app := fiber.New()

	app.Use(
		// add CORS for api calls
		cors.New(),
		// Add logger for api calls
		logger.New(),
	)

	app.Static("/", "public")

	// Routes
	api := app.Group("/api")
	routes.SwaggerRoute(app)
	routes.PublicRoutes(api)
	routes.NotFoundRoute(app)

	// Listen on port 8080
	log.Fatal(app.Listen(":8080"))

}
