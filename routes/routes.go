package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/james-kariuki-source/Timetable-Management-API/controllers"
)

func TheRoutes(app *fiber.App) {
	//app.Post("/admin/register", controllers.Register)
	app.Post("/admin/login", controllers.AdminLogin)
	app.Get("/admin/dashboard", controllers.Admin)
	app.Post("/admin/logout", controllers.AdminLogout)

}
