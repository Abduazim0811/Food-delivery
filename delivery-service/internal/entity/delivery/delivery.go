package delivery

type Delivery struct {
    ID      	string `json:"id" bson:"_id,omitempty"`
    OrderID 	string `json:"order_id" bson:"order_id"`
	CurierID 	string `json:"curier_id" bson:"curier_id"`
    Address 	string `json:"address" bson:"address"`
    Status  	string `json:"status" bson:"status"`
}

type CreateDeliveryReq struct {
    OrderID 	string `json:"order_id" bson:"order_id"`
	CurierID 	string `json:"curier_id" bson:"curier_id"`
	Status  	string `json:"status" bson:"status"`
    Address 	string `json:"address" bson:"address"`
}

type CreateDeliveryRes struct {
    Message    string `json:"message" bson:"message"`
}

type GetDeliveryStatusReq struct {
    DeliveryID string `json:"delivery_id" bson:"_id"`
}

type GetDeliveryStatusRes struct {
    DeliveryID string `json:"delivery_id" bson:"_id"`
    Status     string `json:"status" bson:"status"`
}

type UpdateDeliveryStatusReq struct {
    DeliveryID string `json:"delivery_id" bson:"_id"`
    Status     string `json:"status" bson:"status"`
}

type UpdateDeliveryStatusRes struct {
    Message string `json:"message" bson:"message"`
}
