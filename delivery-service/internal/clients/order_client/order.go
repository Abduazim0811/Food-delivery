package orderclient

import (
	"context"
	"delivery_service/protos/orderproto"
	"fmt"
	"log"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
)

func DialgrpcOrder() orderproto.OrderServiceClient {
	conn, err := grpc.NewClient(os.Getenv("order_url"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("failed to dial grpc client product: ", err)
	}

	return orderproto.NewOrderServiceClient(conn)
}


func Order(ctx context.Context, id string)error{
	_, err := DialgrpcOrder().GetbyIdOrder(ctx, &orderproto.GetOrderReq{OrderId: id})
	if err != nil {
		log.Println("order not found")
		return fmt.Errorf("order not found: %v", err)
	}

	return nil
}