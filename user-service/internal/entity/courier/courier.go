package courier

import "time"

type Courier struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	Email          string    `json:"email"`
	PasswordHash   string    `json:"password_hash"`
	PhoneNumber    string    `json:"phone_number"`
	DeliveryArea   string    `json:"delivery_area"`
	AssignedOrders []int     `json:"assigned_orders"`
	CreatedAt      time.Time `json:"created_at"`
}
