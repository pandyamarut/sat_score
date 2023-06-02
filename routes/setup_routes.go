package routes

import (
	"test/endpoints"

	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
)

// CreateRoutes creates all the routes for the application
func CreateRoutes(app *fiber.App, db *gorm.DB) {
	app.Post("/insert", func(c *fiber.Ctx) error {
		return endpoints.Insert(c, db)
	})
	app.Get("/view_all", func(c *fiber.Ctx) error {
		return endpoints.ViewAll(c, db)
	})
	app.Post("/update_score/:name", func(c *fiber.Ctx) error {
		return endpoints.UpdateScore(c, db)
	})
	app.Delete("/delete/:name", func(c *fiber.Ctx) error {
		return endpoints.Delete(c, db)
	})
	app.Get("/get_rank/:name", func(c *fiber.Ctx) error {
		return endpoints.GetRank(c, db)
	})
}
