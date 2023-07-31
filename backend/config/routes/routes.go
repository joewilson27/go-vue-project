package routes

import (
	"github.com/gofiber/fiber/v2"
	pkg "github.com/joewilson27/go-vue-project/pkg/calc"
)

func SetupRoutes(app *fiber.App) {
	// tester
	app.Get("/hello", func(c *fiber.Ctx) error {
		return c.SendString("hello, Mars!")
	})

	// calculate handler
	CalculateRoute(app)
}

func CalculateRoute(app fiber.Router) { // func CalculateRoute(app fiber.Router) {

	app.Post("/calc", pkg.MainCalc)

}
