package deliveryclient

import (
	"api-gateway/internal/protos/deliveryproto"
	"log"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
)

func DialGrpcDelivery() deliveryproto.DeliveryServiceClient {
	conn, err := grpc.NewClient(os.Getenv("delivery_url"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("dial grpc delivery error:", err)
	}

	return deliveryproto.NewDeliveryServiceClient(conn)
}
