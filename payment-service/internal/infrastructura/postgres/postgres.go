package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"payment-service/internal/entity/payment"
	"payment-service/internal/infrastructura/repository"

	"github.com/Masterminds/squirrel"
)

type PaymentPostgres struct{
	db *sql.DB
}

func NewPaymentPostgres(db *sql.DB) repository.PaymentRepository{
	return &PaymentPostgres{db: db}
}

func (p *PaymentPostgres) ProcessPayment(ctx context.Context, req payment.ProcessPaymentRequest)(*payment.ProcessPaymentResponse, error){
	tx, err := p.db.Begin()
	if err != nil {
		log.Println("error starting transaction: ", err)
		return nil, fmt.Errorf("error starting transaction: %v", err)
	}

	sql, args, err := squirrel.
		Insert("payments").
		Columns("user_id", "order_id", "total_amount", "payment_method", "payment_details").
        Values(req.UserID, req.OrderID, req.TotalAmount, req.PaymentMethod, req.PaymentDetails).
        PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		tx.Rollback()
		log.Println("insert payment error", err)
		return nil, fmt.Errorf("insert payment error: %v", err)
	}

	result, err := tx.ExecContext(ctx, sql, args...)
	if err != nil {
		tx.Rollback()
		log.Println("payment exex error:", err)
		return nil, fmt.Errorf("payment exec error: %v", err)
	}

	transactionID, err := result.LastInsertId()
    if err != nil {
        tx.Rollback()
		log.Println("error transaction error: ", err)
        return nil, fmt.Errorf("error transaction id: %v", err)
    }

	if err := tx.Commit(); err != nil {
		log.Println("transaction confirmation error:", err)
		return nil, fmt.Errorf("transaction confirmation error: %v", err)
	}

	return &payment.ProcessPaymentResponse{
		Success: true,
		Message: "Payment processed successfully",
		TransactionID: fmt.Sprintf("%d", transactionID),
	}, nil
}

func (p *PaymentPostgres) RefundPayment(ctx context.Context, req payment.RefundPaymentRequest)(*payment.RefundPaymentResponse,error){
	tx, err := p.db.Begin()
	if err != nil {
		log.Println("error starting transaction: ", err)
		return nil, fmt.Errorf("error starting transaction: %v", err)
	}

	sql, args, err := squirrel.
		Insert("refunds").
		Columns("transaction_id", "refund_amount", "reason").
        Values(req.TransactionID, req.RefundAmount, req.Reason).
        PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		tx.Rollback()
		log.Println("error creating request:", err)
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	_, err = tx.ExecContext(ctx, sql, args...)
	if err != nil {
		tx.Rollback()
		log.Print("exec contex error: ", err)
		return nil, fmt.Errorf("exec context error: %v", err)
	}

	if err := tx.Commit(); err != nil {
		log.Println("transaction confirmation error", err)
        return nil, fmt.Errorf("transaction confirmation error: %v", err)
    }

	return &payment.RefundPaymentResponse{
		Success: true,
		Message: "Refund processed successfully",
	}, nil
}