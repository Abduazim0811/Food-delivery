package order

type Product struct{
    ProductID string `json:"product_id" bson:"product_id"`
    Quantity  int32  `json:"quantity" bson:"quantity"`
}

type Order struct {
    ID        string `json:"id" bson:"_id,omitempty"`
    Product   []Product `json:"product" bson:"product"`
    UserID    int32 `json:"user_id" bson:"user_id"`
    Status    string `json:"status" bson:"status"`
}

type CreateOrderReq struct {
    Product []Product `json:"product" bson:"product"`
    UserID    int32 `json:"user_id" bson:"user_id"`
    Address string    `json:"address" bson:"address"`
    TotalAmount float32   `json:"totalamount" bson:"totalamount"`
}

type CreateOrderRes struct {
    Message string `json:"message" bson:"message"`
}

type GetOrderReq struct {
    OrderID string `json:"id" bson:"_id"`
}

type GetOrderRes struct {
	OrderID     string    `json:"order_id" bson:"order_id"`
	Product     []Product `json:"product" bson:"product"`
	Address     string    `json:"address" bson:"address"`
	UserID      int32     `json:"user_id" bson:"user_id"`
	Status      string    `json:"status" bson:"status"`
	TotalAmount float32   `json:"totalamount" bson:"totalamount"`
}

type UpdateReq struct {
	OrderID string    `json:"order_id" bson:"order_id"`
	Product []Product `json:"product" bson:"product"`
	Address string    `json:"address" bson:"address"`
    TotalAmount float32   `json:"totalamount" bson:"totalamount"`
}
