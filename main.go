package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	db "github.com/james-kariuki-source/Timetable-Management-API/connection"
	"github.com/james-kariuki-source/Timetable-Management-API/routes"
)

func main() {

	db.Connect()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	routes.TheRoutes(app)

	app.Listen(":8080")
}
