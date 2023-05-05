package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	purchaseRoutes "github.com/yi-lin-1234/purchase-review-backend/internal/routes/purchase"
)

func SetupRoutes(app *fiber.App) {
	// Middleware
	app.Use(logger.New())
	app.Use(cors.New())

	// Create a /api/v1 endpoint
	v1 := app.Group("/api/v1")

	// Setup purchase routes, can use same syntax to add routes for more models
	purchaseRoutes.SetupPurchaseRoutes(v1)
}
