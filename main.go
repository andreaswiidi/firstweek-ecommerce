package main

import (
	"log"
	"one-week-project-ecommerce/config"
	"one-week-project-ecommerce/controller"
	"one-week-project-ecommerce/model"
	"one-week-project-ecommerce/repository"
	"one-week-project-ecommerce/router"
	"one-week-project-ecommerce/service"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func main() {
	loadConfig, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("could not read env variable", err)
	}

	db := config.ConnectionDB(&loadConfig)
	validate := validator.New()

	db.Table("products").AutoMigrate(&model.Product{})

	productRepository := repository.NewProductRepositoryImpl(db)
	productService := service.NewProductServiceImpl(productRepository, validate)
	productController := controller.NewProductController(productService)

	routes := router.NewRouter(productController)

	app := fiber.New()

	app.Mount("/api", routes)

	log.Fatal(app.Listen(":8080"))

}
