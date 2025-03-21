package main

import (
	"log"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/lokesh2201013/Usermanagement/database"
	"github.com/lokesh2201013/Usermanagement/routes"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"time"
)

func main() {
	app := fiber.New()

	// Middleware for logging
	app.Use(logger.New())


	database.ConnectDB()

	// Enable CORS
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	//rate limiting
	app.Use(limiter.New(limiter.Config{
		Max:        10,               
		Expiration: 1 * time.Minute,  
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.IP() 
		},
	}))

	// Set up routes
	routes.AuthRoutes(app)

	port := ":8080" 
	log.Printf("Server is running on http://localhost%s\n", port)
	log.Fatal(app.Listen(port))
}
