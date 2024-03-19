package repository

import (
	"one-week-project-ecommerce/data/request"
	"one-week-project-ecommerce/helper"
	"one-week-project-ecommerce/model"

	"gorm.io/gorm"
)

type ProductRepository interface {
	Save(product model.Product) error
	UpdateProductUsingUPI(product model.Product) error
	DeleteProductUsingUPI(product model.Product) error
	FindByUPI(productUPI string) (model.Product, error)
	FindAll() []model.Product
}

type ProductRepositoryImpl struct {
	Db *gorm.DB
}

func NewProductRepositoryImpl(Db *gorm.DB) ProductRepository {
	return &ProductRepositoryImpl{Db: Db}
}

// Delete implements ProductRepository.
func (p *ProductRepositoryImpl) DeleteProductUsingUPI(product model.Product) error {
	result := p.Db.Model(&product).Updates(product)
	return result.Error
}

// FindAll implements ProductRepository.
func (p *ProductRepositoryImpl) FindAll() []model.Product {
	var product []model.Product
	result := p.Db.Where(`"isdeleted" = ?`, false).Find(&product)
	helper.ErrorPanic(result.Error)
	return product
}

// FindByUPI implements ProductRepository.
func (p *ProductRepositoryImpl) FindByUPI(productUPI string) (model.Product, error) {
	var product model.Product

	err := p.Db.Where(`"UPI" = ? and "isdeleted" = ?`, productUPI, false).First(&product).Error
	if err != nil {
		return product, err
	} else {
		return product, nil
	}
}

// Save implements ProductRepository.
func (p *ProductRepositoryImpl) Save(product model.Product) error {
	result := p.Db.Create(&product)

	return result.Error
}

// Update implements ProductRepository.
func (p *ProductRepositoryImpl) UpdateProductUsingUPI(product model.Product) error {
	var updateProduct = request.UpdateProductRequest{
		Title: product.Title,
		Price: product.Price,
	}
	result := p.Db.Model(&product).Updates(updateProduct)
	return result.Error
}
