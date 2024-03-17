package service

import (
	"one-week-project-ecommerce/data/request"
	"one-week-project-ecommerce/data/response"
	"one-week-project-ecommerce/helper"
	"one-week-project-ecommerce/model"
	"one-week-project-ecommerce/repository"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type ProductServcice interface {
	Create(product request.CreateProductRequest)
	Update(product request.UpdateProductRequest)
	Delete(productUPI string)
	FindbyUPI(productUPI string) response.ProductResponse
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
func (p *ProdctServiceImpl) Create(product request.CreateProductRequest) {
	err := p.validate.Struct(product)
	helper.ErrorPanic(err)
	produuctModel := model.Product{
		Id:    uuid.New(),
		Title: product.Title,
		UPI:   product.UPI,
		Price: product.Price,
	}
	p.ProductRepository.Save(produuctModel)
}

// Delete implements ProductServcice.
func (p *ProdctServiceImpl) Delete(productUPI string) {
	productData, err := p.ProductRepository.FindByUPI(productUPI)
	helper.ErrorPanic(err)

	productData.UpdatedAt = time.Now()
	productData.IsDeleted = true
	productData.UpdatedBy = "System"

	p.ProductRepository.Delete(productData)
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
func (p *ProdctServiceImpl) FindbyUPI(productUPI string) response.ProductResponse {
	productData, err := p.ProductRepository.FindByUPI(productUPI)
	if err != nil {
		panic(helper.NewNotFoundError(err.Error()))
	}
	productResponse := response.ProductResponse{
		UPI:   productData.UPI,
		Title: productData.Title,
		Price: productData.Price,
	}
	return productResponse
}

// Update implements ProductServcice.
func (p *ProdctServiceImpl) Update(product request.UpdateProductRequest) {
	productData, err := p.ProductRepository.FindByUPI(product.UPI)
	helper.ErrorPanic(err)
	productData.Price = product.Price
	productData.Title = product.Title
	p.ProductRepository.Update(productData)

}
