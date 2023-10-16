package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/theritikchoure/shorten/pkg/middleware"
	"github.com/theritikchoure/shorten/pkg/routes"
	"github.com/theritikchoure/shorten/pkg/utils"
)

func main() {

	// Define a fiber app
	app := fiber.New()

	// Middlewares
	middleware.FiberMiddleware(app) // Register fiber's middleware for app.

	// Routes
	routes.PublicRoutes(app) // Register public routes for app

	// Start server (with or without graceful shutdown)
	if os.Getenv("STAGE_STATUS") == "dev" {
		utils.StartServer(app)
	} else {
		utils.StartServerWithGracefulShutdown(app)
	}
}
