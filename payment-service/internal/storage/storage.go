package storage

import (
	"database/sql"
	"log"
	"payment-service/internal/infrastructura/postgres"
	"payment-service/internal/service"
	"payment-service/paymentservice"
)

func OpenSql(driverName, url string) (*sql.DB, error) {
	db, err := sql.Open(driverName, url)
	if err != nil {
		log.Println("failed to open database:", err)
		return nil, err
	}

	if err = db.Ping(); err != nil {
		log.Println("Unable to connect to database:", err)
		return nil, err
	}
	return db, err
}

func Handler(db *sql.DB) *paymentservice.Service {
	repo := postgres.NewPaymentPostgres(db)
	service := service.NewPaymentPostgres(repo)
	handler := paymentservice.NewService(service)
	return handler
}
