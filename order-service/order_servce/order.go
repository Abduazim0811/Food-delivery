package orderservce

import (
	"context"
	"fmt"
	"log"
	productclient "order_service/internal/clients/product_client"
	usersclient "order_service/internal/clients/users_client"
	"order_service/internal/entity/order"
	"order_service/internal/service"
	"order_service/protos/orderproto"
)

type Service struct {
	orderproto.UnimplementedOrderServiceServer
	service *service.OrderService
}

func NewService(service *service.OrderService) *Service {
	return &Service{service: service}
}

func (s *Service) CreateOrder(ctx context.Context,req *orderproto.CreateOrderReq) (*orderproto.CreateOrderRes, error){
	err := usersclient.GetUsers(ctx, req.UserId)
	if err != nil {
		log.Println("user not found")
		return nil, fmt.Errorf("user not found")
	}
	var Product []order.Product
	var totalamount float32

	for _, product := range req.Product {
		summ, err := productclient.Products(ctx, product.ProductId)
		if err != nil {
			log.Println("product not found")
			return nil, fmt.Errorf("product not found")
		}
		totalamount += summ*float32(product.Quantity)
		Product = append(Product, order.Product{
			ProductID: product.ProductId,
			Quantity:  product.Quantity,
		})
	}

	id, err := s.service.Createorder(order.CreateOrderReq{
		Product: Product,
		UserID:  req.UserId,
		Address: req.Address,
		TotalAmount: totalamount,
	})

	if err != nil {
		log.Println("create order error:", err)
		return nil, fmt.Errorf("create order error: %v", err)
	}

	return &orderproto.CreateOrderRes{OrderId: id}, nil
}

func (s *Service) GetbyIdOrder(ctx context.Context,req *orderproto.GetOrderReq) (*orderproto.GetOrderRes, error){
	res, err := s.service.Getbyidorder(order.GetOrderReq{OrderID: req.OrderId})
	if err != nil {
		log.Println("get by id order error")
		return nil, fmt.Errorf("get by id order error: %v", err)
	}
	var productres []*orderproto.ProductRes
	for _, product := range res.Product{
		productres = append(productres, &orderproto.ProductRes{
			ProductId: product.ProductID,
			Quantity:  product.Quantity,
		})
	}

	return &orderproto.GetOrderRes{
		OrderId:     res.OrderID,
		UserId:      res.UserID,
		Status:      res.Status,
		Product:     productres,
		Address:     res.Address,
		Totalamount: res.TotalAmount,
		}, nil
}

func (s *Service) UpdateOrder(ctx context.Context,req *orderproto.UpdateReq) (*orderproto.UpdateOrderRes, error){
	var Product []order.Product
	var totalamount float32
	for _, product := range req.Product {
		summ, err := productclient.Products(ctx, product.ProductId)
		if err != nil {
			log.Println("product not found")
			return nil, fmt.Errorf("product not found")
		}

		totalamount += summ*float32(product.Quantity)
		Product = append(Product, order.Product{
			ProductID: product.ProductId,
			Quantity:  product.Quantity,
		})
	}

	err := s.service.Updateorder(order.UpdateReq{
		OrderID: req.OrderId,
		Product: Product,
		Address: req.Address,
		TotalAmount: totalamount,
	})

	if err != nil {
		log.Println("update error: ", err)
		return nil, fmt.Errorf("update error: %v", err)
	}

	return &orderproto.UpdateOrderRes{Message: "updated"}, nil
}

func (s *Service) DeleteOrder(ctx context.Context, req *orderproto.GetOrderReq) (*orderproto.UpdateOrderRes, error){
	err := s.service.Deleteorder(order.GetOrderReq{OrderID: req.OrderId})
	if err != nil {
		log.Println("delete error: ",err)
		return nil, fmt.Errorf("delete error: %v", err)
	}
	return &orderproto.UpdateOrderRes{Message: "deleted"}, nil
}