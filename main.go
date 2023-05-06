package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yi-lin-1234/purchase-review-backend/database"
	"github.com/yi-lin-1234/purchase-review-backend/router"
	"log"
	"os"
)

func main() {

	// Fiber instance
	app := fiber.New()

	// Connect to the Database
	database.ConnectDB()

	// Set up the router
	router.SetupRoutes(app)

	// Start server(local development)
	//log.Fatal(app.Listen(":8000"))

	// Start server(pro development)
	port := os.Getenv("PORT")
	log.Fatal(app.Listen("0.0.0.0:" + port))
}
