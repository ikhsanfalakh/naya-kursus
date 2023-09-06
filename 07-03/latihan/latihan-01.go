package latihan

import (
	"github.com/gofiber/fiber/v2"
)

type student struct {
	ID    string
	Name  string
	Grade int
}

var data = []student{
	{"E001", "ethan", 21},
	{"W001", "wick", 22},
	{"B001", "bourne", 23},
	{"B002", "bond", 23},
}

func GetUsers(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"data": data})
}

func GetUser(c *fiber.Ctx) error {
	id := c.Query("id")

	for _, each := range data {
		if each.ID == id {
			return c.JSON(each)
		}
	}

	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"message": "User not found",
	})
}
