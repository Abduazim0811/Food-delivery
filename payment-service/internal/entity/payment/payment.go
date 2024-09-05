package payment


type ProcessPaymentRequest struct {
    UserID         int32  `json:"user_id"`
    OrderID        string  `json:"order_id"`
    TotalAmount    float32 `json:"total_amount"`
    PaymentMethod  string  `json:"payment_method"`
    PaymentDetails string  `json:"payment_details"`
}

type ProcessPaymentResponse struct {
    Success       bool   `json:"success"`
    Message       string `json:"message"`
    TransactionID string `json:"transaction_id"`
}

type RefundPaymentRequest struct {
    TransactionID string  `json:"transaction_id"`
    RefundAmount  float32 `json:"refund_amount"`
    Reason        string  `json:"reason"`
}

type RefundPaymentResponse struct {
    Success bool   `json:"success"`
    Message string `json:"message"`
}