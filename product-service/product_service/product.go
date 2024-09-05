package productservice

import (
	"context"
	"fmt"
	"log"
	"product-service/internal/entity/product"
	"product-service/internal/service"
	"product-service/productproto"
)

type Service struct {
	productproto.UnimplementedProductServiceServer
	service *service.ProductService
}

func NewService(service *service.ProductService) *Service {
	return &Service{service: service}
}

func (s *Service) CreateProduct(ctx context.Context, req *productproto.CreateReq) (*productproto.CreateRes, error){
	err := s.service.Createproduct(product.CreateReq{
		Name: req.Name,
		Description: req.Description,
		Price: req.Price,
	})
	if err != nil {
		log.Println("create error product:", err)
		return nil, fmt.Errorf("create error product: %v", err)
	}

	return &productproto.CreateRes{Message: "product created"}, nil
}

func (s *Service) GetByIdProduct(ctx context.Context,req *productproto.ProductResponse) (*productproto.Product, error){
	res, err := s.service.Getbyidproduct(product.ProductResponse{ID: req.Id})
	if err != nil {
		log.Println("get by id error product:", err)
		return nil, fmt.Errorf("get by id product: %v", err)
	}

	return &productproto.Product{
		Id: res.ID,
		Name: res.Name,
		Description: res.Description,
		Price: res.Price,}, nil
}

func (s *Service) GetAllProducts(ctx context.Context,req *productproto.Empty) (*productproto.ListProduct, error){
	res, err := s.service.Getallproducts()
	if err != nil {
		log.Println("get all product error:", err)
		return nil, fmt.Errorf("get all product error: %v", err)
	}

	var productres []*productproto.Product
	for _, product := range *res{
		productres = append(productres, &productproto.Product{
			Id: product.ID,
			Name: product.Name,
			Description:  product.Description,
			Price: product.Price,
		})
	}

	return &productproto.ListProduct{Product: productres}, nil

}

func (s *Service) UpdateProduct(ctx context.Context, req *productproto.Product) (*productproto.CreateRes, error){
	err := s.service.Updateproduct(product.Product{
		ID: req.Id,
		Name: req.Name,
		Description: req.Description,
		Price: req.Price,
	})
	if err != nil {
		log.Println("update product error:", err)
		return nil, fmt.Errorf("update product error: %v", err)
	}

	return &productproto.CreateRes{Message: "product updated"}, nil
}

func (s *Service) DeleteProduct(ctx context.Context,req *productproto.ProductResponse) (*productproto.CreateRes, error){
	err := s.service.Deleteproduct(product.ProductResponse{ID: req.Id})
	if err != nil {
		log.Println("delete product error:", err)
		return nil, fmt.Errorf("delete product error: %v", err)
	}

	return &productproto.CreateRes{Message: "product deleted"}, nil
}