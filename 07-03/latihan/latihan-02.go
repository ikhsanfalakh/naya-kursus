package latihan

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func DataUsers(url string) ([]User, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var users []User
	err = json.NewDecoder(resp.Body).Decode(&users)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func GetUsersBaru(url string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		users, err := DataUsers(url)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"message": "Server Internal Error",
			})
		}

		return c.JSON(fiber.Map{"data": users})
	}
}

func GetUsersNew(url string, c *fiber.Ctx) error {
	users, err := DataUsers(url)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Server Internal Error",
		})
	}

	return c.JSON(fiber.Map{"data": users})
}

func GetUserNew(url string, c *fiber.Ctx) error {
	id := c.Query("id")
	users, err := DataUsers(url)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error getting users",
		})
	}

	for _, user := range users {
		if fmt.Sprintf("%d", user.ID) == id {
			return c.JSON(fiber.Map{"data": user})
		}
	}

	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"message": "User not found",
	})
}
