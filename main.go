package main

import (
	"log"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/lokesh2201013/Usermanagement/database"
	"github.com/lokesh2201013/Usermanagement/routes"
)

func main() {
	// Initialize Fiber app
	app := fiber.New()

	// Middleware for logging
	app.Use(logger.New())

	// Connect to the database
	database.ConnectDB()

	// Enable CORS
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	// Set up routes
	routes.AuthRoutes(app)

	// Start the server
	port := ":8080" // Specify the port to run the server on
	log.Printf("Server is running on http://localhost%s\n", port)
	log.Fatal(app.Listen(port))
}
