package product

type Product struct {
	ID          string  `json:"id" bson:"_id,omitempty"`
	Name        string  `json:"name" bson:"name"`
	Description string  `json:"description" bson:"description"`
	Price       float32 `json:"price" bson:"price"`
}

type CreateRes struct {
	Message string `json:"message" bson:"message"`
}

type CreateReq struct {
	Name        string  `json:"name" bson:"name"`
	Description string  `json:"description" bson:"description"`
	Price       float32 `json:"price" bson:"price"`
}

type ProductResponse struct {
	ID string `json:"id" bson:"_id"`
}

type Empty struct{}

type ListProduct struct {
	Product Product `json:"product" bson:"product"`
}
