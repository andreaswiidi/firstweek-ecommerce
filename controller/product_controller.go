package controller

import (
	"one-week-project-ecommerce/data/request"
	"one-week-project-ecommerce/data/response"
	"one-week-project-ecommerce/helper"
	"one-week-project-ecommerce/service"

	"github.com/gofiber/fiber/v2"
)

type ProductController struct {
	productService service.ProductServcice
}

func NewProductController(service service.ProductServcice) *ProductController {
	return &ProductController{productService: service}
}

func (controller *ProductController) Create(ctx *fiber.Ctx) error {
	createProductRequest := request.CreateProductRequest{}
	err := ctx.BodyParser(&createProductRequest)
	helper.ErrorPanic(err)

	controller.productService.Create(createProductRequest)

	webResponse := response.Response{
		Code:    200,
		Message: "Success",
		Data:    createProductRequest,
	}
	return ctx.Status(fiber.StatusCreated).JSON(webResponse)
}

func (controller *ProductController) FindAll(ctx *fiber.Ctx) error {
	productResponse := controller.productService.FindAll()
	webResponse := response.Response{
		Code:    200,
		Message: "Success",
		Data:    productResponse,
	}
	return ctx.Status(fiber.StatusCreated).JSON(webResponse)
}

func (controller *ProductController) FindByUPI(ctx *fiber.Ctx) error {
	productUPI := ctx.Params("upi")

	productResponse := controller.productService.FindbyUPI(productUPI)
	webResponse := response.Response{
		Code:    200,
		Message: "Success",
		Data:    productResponse,
	}
	return ctx.Status(fiber.StatusCreated).JSON(webResponse)
}
