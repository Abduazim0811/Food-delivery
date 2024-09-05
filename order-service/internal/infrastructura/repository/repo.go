package repository

import "order_service/internal/entity/order"

type OrderRepository interface {
	AddOrder(req order.CreateOrderReq)(string, error) 
	GetOrderById(req order.GetOrderReq) (*order.GetOrderRes, error)
	UpdateOrder(req order.UpdateReq)error
	DeleteOrder(req order.GetOrderReq)error
}
