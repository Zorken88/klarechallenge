package routes

import (
	"github.com/KlareTeam/interview-challenges/go/pricematic/handlers"
	"github.com/gofiber/fiber/v2"
)

func PublicRoutes(r fiber.Router) {

	// * GET `/api/v1/products` <-- Lista todos los productos
	// * GET `/api/v1/products/:id` <-- Obtiene el producto :id
	// * POST `/api/vi/products` <-- Inserta un producto en la base de datos
	// * PATCH `/api/vi/products/:id` <-- Actualiza uno o mas campos del producto :id
	// * DELETE `/api/vi/products/:id` <-- Elimina el producto :id
	// * GET `/api/v1/products/:id/prices` <-- Devuelve lista de la historia de los precios del producto :id

	v1 := r.Group("/v1")

	// GET method
	v1.Get("/products", handlers.GetProducts)
	v1.Get("/products/:id", handlers.GetProduct)
	v1.Get("/products/:id/prices", handlers.GetHistoryPrices)

	// POST method
	v1.Post("/products", handlers.CreateProduct)

	// PATCH method
	v1.Patch("/products/:id", handlers.UpdateProduct)

	// DELETE method
	v1.Delete("/products/:id", handlers.DeleteProduct)

}
