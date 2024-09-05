package service

import (
	"context"
	"payment-service/internal/entity/payment"
	"payment-service/internal/infrastructura/repository"
)

type PaymentService struct{
	repo repository.PaymentRepository
}

func NewPaymentPostgres(repo repository.PaymentRepository) *PaymentService{
	return &PaymentService{repo: repo}
}

func (p *PaymentService) Processpayment(ctx context.Context, req payment.ProcessPaymentRequest)(*payment.ProcessPaymentResponse, error){
	return p.repo.ProcessPayment(ctx, req)
}

func (p *PaymentService) Refundpayment(ctx context.Context, req payment.RefundPaymentRequest)(*payment.RefundPaymentResponse, error){
	return p.repo.RefundPayment(ctx, req)
}