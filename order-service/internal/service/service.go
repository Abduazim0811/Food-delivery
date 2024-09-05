package service

import (
	"order_service/internal/entity/order"
	"order_service/internal/infrastructura/repository"
)

type OrderService struct{
	repo repository.OrderRepository
}

func NewOrderService(repo repository.OrderRepository) *OrderService{
	return &OrderService{repo: repo}
}

func (o *OrderService) Createorder(req order.CreateOrderReq)(string, error){
	return o.repo.AddOrder(req)
}

func (o *OrderService) Getbyidorder(req order.GetOrderReq)(*order.GetOrderRes, error){
	return o.repo.GetOrderById(req)
}

func (o *OrderService) Updateorder(req order.UpdateReq)error{
	return o.repo.UpdateOrder(req)
}

func (o *OrderService) Deleteorder(req order.GetOrderReq)error{
	return o.repo.DeleteOrder(req)
}