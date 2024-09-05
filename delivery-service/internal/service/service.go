package service

import (
	"delivery_service/internal/entity/delivery"
	"delivery_service/internal/infrastructura/repository"
)

type DeliveryService struct{
	repo repository.DeliveryRepository
}

func NewDeliveryService(repo repository.DeliveryRepository) *DeliveryService{
	return &DeliveryService{repo: repo}
}

func (d *DeliveryService) Createdelivery(req delivery.CreateDeliveryReq)(string,error){
	return d.repo.AddDelivery(req)
}

func (d *DeliveryService) GetStatusdelivery(req delivery.GetDeliveryStatusReq)(*delivery.Delivery, error){
	return d.repo.GetDeliveryStatus(req)
}

func (d *DeliveryService) Updatedelivery(req delivery.UpdateDeliveryStatusReq)error{
	return d.repo.Update(req)
}