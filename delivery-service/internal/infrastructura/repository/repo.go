package repository

import "delivery_service/internal/entity/delivery"

type DeliveryRepository interface {
	AddDelivery(req delivery.CreateDeliveryReq) (string, error)
	GetDeliveryStatus(req delivery.GetDeliveryStatusReq)(*delivery.Delivery, error)
	Update(req delivery.UpdateDeliveryStatusReq)error
}
