package main

import (
	"0803/controllers"
	"0803/db"

	"github.com/gofiber/fiber/v2"
)

func main() {
	//const layoutDateTime = "2006-01-02 15:04:05"

	//saleDate, _ := time.Parse(layoutDateTime, "2023-08-31 00:00:00")
	//saleTime, _ := time.Parse(layoutDateTime, "2023-08-31 10:50:39")

	db.Connect()

	server := fiber.New()
	controllers.RouteSales(server)

	server.Listen(":8080")
}
