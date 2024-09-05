package service

import (
	"product-service/internal/entity/product"
	"product-service/internal/infrastructura/repository"
)

type ProductService struct{
	repo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) *ProductService{
	return &ProductService{repo: repo}
}

func (p *ProductService) Createproduct(req product.CreateReq) error{
	return p.repo.AddProduct(req)
}

func (p *ProductService) Getbyidproduct(req product.ProductResponse) (*product.Product, error){
	return p.repo.GetByIdProduct(req)
}

func (p *ProductService) Getallproducts()(*[]product.Product, error){
	return p.repo.GetAll()
}

func (p *ProductService) Updateproduct(req product.Product) error{
	return p.repo.Update(req)
}

func (p *ProductService) Deleteproduct(req product.ProductResponse) error{
	return p.repo.Delete(req)
}