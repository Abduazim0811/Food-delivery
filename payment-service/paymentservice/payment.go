package paymentservice

import (
	"context"
	"fmt"
	"log"
	"payment-service/internal/clients/order"
	"payment-service/internal/clients/user"
	"payment-service/internal/entity/payment"
	"payment-service/internal/service"
	"payment-service/protos/paymentproto"
)

type Service struct {
	paymentproto.UnimplementedPaymentServiceServer
	service *service.PaymentService
}

func NewService(service *service.PaymentService) *Service {
	return &Service{service: service}
}

func (s *Service) ProcessPayment(ctx context.Context, req *paymentproto.ProcessPaymentRequest) (*paymentproto.ProcessPaymentResponse, error){
	err := user.UserID(ctx, req.UserId)
	if err != nil {
		log.Println("user not found")
		return nil, fmt.Errorf("user not found: %v", err)
	}

	totalamount , err := order.Total(ctx, req.OrderId)
	if err != nil {
		log.Println("order not found")
		return nil, fmt.Errorf("order not found: %v", err)
	}
	res, err := s.service.Processpayment(ctx, payment.ProcessPaymentRequest{
		UserID: req.UserId,
		OrderID: req.OrderId,
		TotalAmount: totalamount,
		PaymentMethod: req.PaymentMethod,
		PaymentDetails: req.PaymentDetails,
	})
	if err != nil {
		log.Println("process payment error:", err)
		return nil, fmt.Errorf("process payment error: %v", err)
	}

	return &paymentproto.ProcessPaymentResponse{
		Success: res.Success,
		Message: res.Message,
		TransactionId: res.TransactionID,
	}, nil
}


func (s *Service) RefundPayment(ctx context.Context,req *paymentproto.RefundPaymentRequest) (*paymentproto.RefundPaymentResponse, error){
	res, err := s.service.Refundpayment(ctx, payment.RefundPaymentRequest{
		TransactionID: req.TransactionId,
		RefundAmount: req.RefundAmount,
		Reason: req.Reason,
	})
	if err != nil {
		log.Println("RefundPayment error: ", err)
		return nil, fmt.Errorf("RefundPayment error: %v", err)
	}

	return &paymentproto.RefundPaymentResponse{
		Success: res.Success,
		Message: res.Message,
	}, nil
}