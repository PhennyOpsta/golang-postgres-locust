package api

import (
	"database/sql"
	"go-api/usecase/product"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetProductDetail(c *fiber.Ctx, db *sql.DB) error {

	productID := c.Params("id")
	newProductId, err := strconv.Atoi(productID)
	if err != nil {
		return c.JSON(fiber.Map{"result": "Product id invalid."})
	}

	product, err := product.GetProductName(newProductId, db)
	if err != nil {
		return c.JSON(fiber.Map{"result": "Get Product error"})
	}

	if product.ProductID == 0 {
		return c.JSON(fiber.Map{"result": "Don't have this product id."})
	}
	return c.JSON(fiber.Map{"result": product.Name})

}
