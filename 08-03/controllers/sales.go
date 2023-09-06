package controllers

import (
	"0803/db"
	"0803/models"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func RouteSales(server fiber.Router) {
	router := server.Group("/sales")
	router.Get("", MiddleWare, GetAllSales)
	router.Post("", CreateSales)
	router.Delete(":id", DeleteSalesByID)
}

func MiddleWare(c *fiber.Ctx) error {
	fmt.Println("coba middlerware")

	return c.Next()
}

func GetAllSales(c *fiber.Ctx) error {
	sales := []models.GetSales{}
	err := db.Con.Table("sales").Order("id DESC").Find(&sales).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err,
		})
	}

	return c.JSON(fiber.Map{
		"data": sales,
	})
}

func CreateSales(c *fiber.Ctx) error {
	in := &models.Sales{}
	if err := c.BodyParser(&in); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err,
		})
	}

	err := db.Con.Create(&in).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err,
		})
	}

	return c.JSON(fiber.Map{
		"message": "created",
	})
}

func DeleteSalesByID(c *fiber.Ctx) error {
	id := c.Params("id")

	err := db.Con.Where("id = ?", id).Delete(&models.Sales{}).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err,
		})
	}

	return c.JSON(fiber.Map{
		"message": "deleted",
	})
}
