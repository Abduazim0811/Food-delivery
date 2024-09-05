package repository

import "product-service/internal/entity/product"

type ProductRepository interface {
	AddProduct(req product.CreateReq) error
	GetByIdProduct(req product.ProductResponse) (*product.Product, error)
	GetAll() (*[]product.Product, error)
	Update(req product.Product) error 
	Delete(req product.ProductResponse) error
}
