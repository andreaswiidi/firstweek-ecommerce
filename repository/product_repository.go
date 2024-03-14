package repository

import (
	"errors"
	"one-week-project-ecommerce/data/request"
	"one-week-project-ecommerce/helper"
	"one-week-project-ecommerce/model"

	"gorm.io/gorm"
)

type ProductRepository interface {
	Save(product model.Product)
	Update(product model.Product)
	Delete(productUPI string)
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
func (p *ProductRepositoryImpl) Delete(productUPI string) {
	// var product model.Product

	// result := p.Db.Where("UPI = ? ",productUPI)
	panic("unimplemented")
}

// FindAll implements ProductRepository.
func (p *ProductRepositoryImpl) FindAll() []model.Product {
	var product []model.Product
	result := p.Db.Find(&product)
	helper.ErrorPanic(result.Error)
	return product
}

// FindByUPI implements ProductRepository.
func (p *ProductRepositoryImpl) FindByUPI(productUPI string) (model.Product, error) {
	var product model.Product

	result := p.Db.Find(&product, productUPI)
	if result != nil {
		return product, nil
	} else {
		return product, errors.New("product is not found")
	}
}

// Save implements ProductRepository.
func (p *ProductRepositoryImpl) Save(product model.Product) {
	result := p.Db.Create(&product)
	helper.ErrorPanic(result.Error)
}

// Update implements ProductRepository.
func (p *ProductRepositoryImpl) Update(product model.Product) {
	var updateProduct = request.UpdateProductRequest{
		UPI:   product.UPI,
		Title: product.Title,
		Price: product.Price,
	}
	result := p.Db.Model(&product).Updates(updateProduct)
	helper.ErrorPanic(result.Error)
}
