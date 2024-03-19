package controller

import (
	"one-week-project-ecommerce/data/request"
	"one-week-project-ecommerce/data/response"
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
	if err != nil {
		webResponseError := response.Response{
			Code:    400,
			Message: "Error",
		}

		return ctx.Status(fiber.StatusBadRequest).JSON(webResponseError)
	}

	result := controller.productService.Create(createProductRequest)

	return ctx.Status(fiber.StatusCreated).JSON(result)
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

	result := controller.productService.FindbyUPI(productUPI)

	return ctx.Status(result.Code).JSON(result)
}

func (controller *ProductController) UpdateProductUsingUPI(ctx *fiber.Ctx) error {
	updateProductRequest := request.UpdateProductRequest{}

	err := ctx.BodyParser(&updateProductRequest)

	if err != nil {
		webResponseError := response.Response{
			Code:    400,
			Message: "Error",
		}

		return ctx.Status(fiber.StatusBadRequest).JSON(webResponseError)
	}

	productUPI := ctx.Params("upi")
	updateProductRequest.UPI = productUPI

	result := controller.productService.UpdateProductUsingUPI(updateProductRequest)

	return ctx.Status(result.Code).JSON(result)

}

func (controller *ProductController) DeleteProductUsingUPI(ctx *fiber.Ctx) error {

	productUPI := ctx.Params("upi")
	result := controller.productService.DeleteProductUsingUPI(productUPI)
	return ctx.Status(result.Code).JSON(result)

}
