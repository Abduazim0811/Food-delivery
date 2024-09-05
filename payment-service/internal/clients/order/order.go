package order

import (
	"context"
	"log"
	"os"
	"payment-service/protos/orderproto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
)

func DialgrpcOrder() orderproto.OrderServiceClient {
	conn, err := grpc.NewClient(os.Getenv("order_url"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("failed to dial grpc client order: ", err)
	}

	return orderproto.NewOrderServiceClient(conn)
}

func Total(ctx context.Context, id string)(float32, error){
	res, err := DialgrpcOrder().GetbyIdOrder(ctx, &orderproto.GetOrderReq{OrderId: id})
	return res.Totalamount, err
}
