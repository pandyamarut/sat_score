package main

import (
	"log"

	"test/routes"
	"test/types"

	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	// Connect to the PostgreSQL database
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=marutpandya dbname=test sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Auto-migrate the database schema
	db.AutoMigrate(&types.SATResult{})

	// Initialize Fiber app
	app := fiber.New()

	// Setup routes
	routes.CreateRoutes(app, db)

	// Start the server
	log.Fatal(app.Listen(":8080"))
}
