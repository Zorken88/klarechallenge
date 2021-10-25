package handlers

import (
	"errors"
	"strconv"

	"github.com/KlareTeam/interview-challenges/go/pricematic/database"
	"github.com/KlareTeam/interview-challenges/go/pricematic/models"
	"github.com/gofiber/fiber/v2"
)

// * GET `/api/v1/products` <-- Lista todos los productos
// * GET `/api/v1/products/:id` <-- Obtiene el producto :id
// * POST `/api/vi/products` <-- Inserta un producto en la base de datos
// * PATCH `/api/vi/products/:id` <-- Actualiza uno o mas campos del producto :id
// * DELETE `/api/vi/products/:id` <-- Elimina el producto :id
// * GET `/api/v1/products/:id/prices` <-- Devuelve lista de la historia de los precios del producto :id

// GetProducts func gets all products
// @Description Get all products.
// @Summary get all products from database.
// @Tags Products
// @Accept json
// @Produce json
// @Success 200 {array} models.Product
// @Router /api/v1/products [get]
func GetProducts(c *fiber.Ctx) error {
	products := []models.Product{}

	database.Database.Db.Find(&products)

	return c.JSON(fiber.Map{
		"error":    false,
		"msg":      nil,
		"count":    len(products),
		"products": products,
	})
}

// GetProduct func gets produt by given id
// @Description Get product by given id.
// @Summary get product from database by given id or 404.
// @Tags Product
// @Accept json
// @Produce json
// @Param id path number true "Product ID"
// @Success 200 {object} models.Product
// @Router /api/v1/products/{id} [get]
func GetProduct(c *fiber.Ctx) error {

	// Catch product id from url
	id, err := strconv.ParseUint(c.Params("id"), 10, 0)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	product := models.Product{}

	if err := findProduct(id, &product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":    true,
			"msg":      err.Error(),
			"products": nil,
		})
	}

	return c.JSON(fiber.Map{
		"error":    false,
		"msg":      nil,
		"products": product,
	})
}

// Find in database the product by id
func findProduct(id uint64, product *models.Product) error {
	database.Database.Db.Find(&product, "id = ?", id)

	if product.ID == 0 {
		return errors.New("product does not exist")
	}

	return nil
}

// GetHistoryPrices func gets the history of prices by product id
// @Description Get the price history by given product id.
// @Summary get the price history by given product id
// @Tags Product
// @Accept json
// @Produce json
// @Param id path number true "Product ID"
// @Success 200 {array} models.Product
// @Router /api/v1/products/{id}/prices [get]
func GetHistoryPrices(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 0)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	product := models.Product{}

	if err := database.Database.Db.Preload("Prices").Find(&product, "products.id = ?", id).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"error":    false,
		"msg":      nil,
		"products": product,
	})
}

// CreateProduct func creates a new product
// @Description Create a new product.
// @Summary create a new product
// @Tags Product
// @Accept json
// @Produce json
// @Param product body models.Product true "Product"
// @Success 201 {object} models.Product
// @Router /api/v1/products [post]
func CreateProduct(c *fiber.Ctx) error {
	product := models.Product{}
	price := models.Price{}

	if err := c.BodyParser(&product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":    true,
			"msg":      err.Error(),
			"products": nil,
		})
	}

	price.Value = product.ActualPrice

	product.Prices = append(product.Prices, price)

	if err := database.Database.Db.Create(&product).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":    true,
			"msg":      err.Error(),
			"products": nil,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"error":    false,
		"msg":      "Successfully created product",
		"products": product,
	})
}

// UpdateProduct func update values for product
// @Description Update values for product by id.
// @Summary update values for product by id
// @Tags Product
// @Accept json
// @Produce json
// @Param id path number true "Product ID"
// @Param product body models.Product false "Product"
// @Success 204 {string} status "ok"
// @Router /api/v1/products/{id} [patch]
func UpdateProduct(c *fiber.Ctx) error {
	product := models.Product{}
	updateProduct := models.Product{}
	price := models.Price{}

	if err := c.BodyParser(&updateProduct); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":    true,
			"msg":      err.Error(),
			"products": nil,
		})
	}

	id, err := strconv.ParseUint(c.Params("id"), 10, 0)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	if err = findProduct(id, &product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":    true,
			"msg":      err.Error(),
			"products": nil,
		})
	}

	// TODO Implement a real PATCH solution with reflection
	if updateProduct.ActualPrice >= 0 {
		price.ProductRefer = product.ID
		price.Value = updateProduct.ActualPrice

		if err := database.Database.Db.Create(&price).Error; err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error":    true,
				"msg":      err.Error(),
				"products": nil,
			})
		}

	}
	product.Name = updateProduct.Name
	product.ActualPrice = updateProduct.ActualPrice

	if err := database.Database.Db.Updates(&product).Error; err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":    true,
			"msg":      err.Error(),
			"products": nil,
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

// Delete func for deletes product by id
// @Description Delete product by given id.
// @Summary delete product by given id
// @Tags Product
// @Accept json
// @Produce json
// @Param id path number true "Product ID"
// @Success 204 {string} status "ok"
// @Router /api/v1/products/{id} [delete]
func DeleteProduct(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 0)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	product := models.Product{}

	if err = findProduct(id, &product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":    true,
			"msg":      err.Error(),
			"products": nil,
		})
	}

	if err = database.Database.Db.Delete(&product).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":    true,
			"msg":      err.Error(),
			"products": nil,
		})
	}
	return c.SendStatus(fiber.StatusNoContent)
}
