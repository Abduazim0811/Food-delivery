package orderclient

import (
	"log"
	"notification-service/internal/protos/orderproto"
	"os"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func DialGrpcOrder() orderproto.OrderServiceClient {
	conn, err := grpc.NewClient(os.Getenv("order_url"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("dial grpc user error:", err)
	}

	return orderproto.NewOrderServiceClient(conn)
}
