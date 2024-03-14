package router

import (
	"one-week-project-ecommerce/controller"

	"github.com/gofiber/fiber/v2"
)

func NewRouter(productController *controller.ProductController) *fiber.App {
	router := fiber.New()

	router.Get("/healthchecker", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"status":  "Success",
			"message": "welcome",
		})
	})

	router.Route("/product", func(router fiber.Router) {
		router.Post("/", productController.Create)
		router.Get("", productController.FindAll)
	})

	// router.Route("/product/:upi",func(router fiber.Router) {
	// 	router.Delete("",prod)
	// })

	return router

}
