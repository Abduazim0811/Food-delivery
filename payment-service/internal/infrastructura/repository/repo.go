package repository

import (
	"context"
	"payment-service/internal/entity/payment"
)

type PaymentRepository interface {
	ProcessPayment(ctx context.Context, req payment.ProcessPaymentRequest) (*payment.ProcessPaymentResponse, error)
	RefundPayment(ctx context.Context, req payment.RefundPaymentRequest)(*payment.RefundPaymentResponse,error)
}
