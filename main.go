package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yi-lin-1234/purchase-review-backend/database"
	"github.com/yi-lin-1234/purchase-review-backend/router"
	"log"
)

func main() {

	// Fiber instance
	app := fiber.New()

	// Connect to the Database
	database.ConnectDB()

	// Set up the router
	router.SetupRoutes(app)

	// Start server
	log.Fatal(app.Listen(":8000"))
}
