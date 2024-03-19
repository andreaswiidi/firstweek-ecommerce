package service

import (
	"one-week-project-ecommerce/data/request"
	"one-week-project-ecommerce/data/response"
	"one-week-project-ecommerce/model"
	"one-week-project-ecommerce/repository"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type ProductServcice interface {
	Create(product request.CreateProductRequest) response.Response
	UpdateProductUsingUPI(product request.UpdateProductRequest) response.Response
	DeleteProductUsingUPI(productUPI string) response.Response
	FindbyUPI(productUPI string) response.Response
	FindAll() []response.ProductResponse
}

type ProdctServiceImpl struct {
	ProductRepository repository.ProductRepository
	validate          *validator.Validate
}

func NewProductServiceImpl(productRepository repository.ProductRepository, validate *validator.Validate) ProductServcice {
	return &ProdctServiceImpl{
		ProductRepository: productRepository,
		validate:          validate,
	}
}

// Create implements ProductServcice.
func (p *ProdctServiceImpl) Create(product request.CreateProductRequest) response.Response {
	err := p.validate.Struct(product)

	if err != nil {
		webResponseError := response.Response{
			Code:    400,
			Message: "Error",
		}
		return webResponseError
	}

	produuctModel := model.Product{
		Id:    uuid.New(),
		Title: product.Title,
		UPI:   product.UPI,
		Price: product.Price,
	}
	errService := p.ProductRepository.Save(produuctModel)

	if errService != nil {
		webResponseError := response.Response{
			Code:    400,
			Message: "Error",
		}
		return webResponseError
	}

	webResponse := response.Response{
		Code:    200,
		Message: "Success",
	}

	return webResponse
}

// Delete implements ProductServcice.
func (p *ProdctServiceImpl) DeleteProductUsingUPI(productUPI string) response.Response {
	productData, err := p.ProductRepository.FindByUPI(productUPI)
	if err != nil {
		webResponseError := response.Response{
			Code:    fiber.StatusNotFound,
			Message: "Product Not Found",
		}
		return webResponseError
	}

	productData.UpdatedAt = time.Now()
	productData.IsDeleted = true
	productData.UpdatedBy = "System"

	errRepo := p.ProductRepository.DeleteProductUsingUPI(productData)

	if errRepo != nil {
		webResponseError := response.Response{
			Code:    fiber.StatusServiceUnavailable,
			Message: "Error Deleting Product",
		}
		return webResponseError
	}

	webResponse := response.Response{
		Code:    200,
		Message: "Success",
	}

	return webResponse
}

// FindAll implements ProductServcice.
func (p *ProdctServiceImpl) FindAll() []response.ProductResponse {
	result := p.ProductRepository.FindAll()
	var products []response.ProductResponse

	for _, value := range result {
		product := response.ProductResponse{
			UPI:   value.UPI,
			Title: value.Title,
			Price: value.Price,
		}
		products = append(products, product)
	}

	return products
}

// FindbyID implements ProductServcice.
func (p *ProdctServiceImpl) FindbyUPI(productUPI string) response.Response {
	productData, err := p.ProductRepository.FindByUPI(productUPI)
	if err != nil {
		webResponseError := response.Response{
			Code:    fiber.StatusNotFound,
			Message: "Product Not Found",
		}
		return webResponseError
	}
	productResponse := response.ProductResponse{
		UPI:   productData.UPI,
		Title: productData.Title,
		Price: productData.Price,
	}

	webResponse := response.Response{
		Code:    200,
		Message: "Success",
		Data:    productResponse,
	}

	return webResponse
}

// Update implements ProductServcice.
func (p *ProdctServiceImpl) UpdateProductUsingUPI(product request.UpdateProductRequest) response.Response {
	productData, err := p.ProductRepository.FindByUPI(product.UPI)
	if err != nil {
		webResponseError := response.Response{
			Code:    fiber.StatusNotFound,
			Message: "Product Not Found",
		}
		return webResponseError
	}
	productData.Price = product.Price
	productData.Title = product.Title
	errRepo := p.ProductRepository.UpdateProductUsingUPI(productData)
	if errRepo != nil {
		webResponseError := response.Response{
			Code:    fiber.StatusServiceUnavailable,
			Message: "Error Update Data",
		}
		return webResponseError
	}

	webResponse := response.Response{
		Code:    200,
		Message: "Success",
	}

	return webResponse

}
