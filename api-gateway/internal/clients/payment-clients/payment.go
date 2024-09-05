package paymentclients

import (
	"api-gateway/internal/protos/paymentproto"
	"log"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
)

func DialGrpcPayment() paymentproto.PaymentServiceClient {
	conn, err := grpc.NewClient(os.Getenv("payment_url"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("dial grpc payment error:", err)
	}

	return paymentproto.NewPaymentServiceClient(conn)
}
