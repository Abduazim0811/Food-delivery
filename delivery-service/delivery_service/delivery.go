package deliveryservice

import (
	"context"
	orderclient "delivery_service/internal/clients/order_client"
	"delivery_service/internal/entity/delivery"
	"delivery_service/internal/service"
	"delivery_service/protos/deliveryproto"
	"fmt"
	"log"
)

type Service struct {
	deliveryproto.UnimplementedDeliveryServiceServer
	service *service.DeliveryService
}

func NewService(service *service.DeliveryService) *Service {
	return &Service{service: service}
}

func (s *Service) CreateDelivery(ctx context.Context, req *deliveryproto.CreateDeliveryReq) (*deliveryproto.CreateDeliveryRes, error){
	err := orderclient.Order(ctx, req.OrderId)
	if err != nil {
		log.Println("order not found")
		return nil, fmt.Errorf("order not found: %v", err)
	}

	deliveryID, err := s.service.Createdelivery(delivery.CreateDeliveryReq{
		OrderID:  req.OrderId,
		Address: req.Address,
		Status: "Pending",
	})
	if err != nil {
		log.Println("create delivery error:", err)
		return nil, fmt.Errorf("create delivery error: %v", err)
	}

	return &deliveryproto.CreateDeliveryRes{Message: "create delivery", Id: deliveryID}, nil
}

func (s *Service) GetDeliveryStatus(ctx context.Context,req *deliveryproto.GetDeliveryStatusReq) (*deliveryproto.Delivery, error){
	res, err := s.service.GetStatusdelivery(delivery.GetDeliveryStatusReq{DeliveryID: req.DeliveryId})
	if err != nil {
		log.Println("get by status error:", err)
		return nil, fmt.Errorf("get by status: %v", err)
	}

	return &deliveryproto.Delivery{
		Id: 	res.ID,
		OrderId: res.OrderID,
		Address: res.Address,
		Status: res.Status,
	}, nil
}

func (s *Service) UpdateDeliveryStatus(ctx context.Context,req *deliveryproto.UpdateDeliveryStatusReq) (*deliveryproto.UpdateDeliveryStatusRes, error){
	err := s.service.Updatedelivery(delivery.UpdateDeliveryStatusReq{DeliveryID: req.DeliveryId, Status: req.Status})
	if err != nil {
		log.Println("update delivery error:", err)
		return nil, fmt.Errorf("update delivery error: %v", err)
	}

	return &deliveryproto.UpdateDeliveryStatusRes{Message: "updated"}, nil
}
