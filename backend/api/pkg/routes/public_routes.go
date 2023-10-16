package routes

import "github.com/gofiber/fiber/v2"

func PublicRoutes(a *fiber.App) {
	//Create routes group
	route := a.Group("/api/v1")

	route.Get("/books")

}
