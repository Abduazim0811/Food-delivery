package productclient

import (
	"context"
	"fmt"
	"log"
	"order_service/protos/productproto"
	"os"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func DialProductGrpc() productproto.ProductServiceClient {
	conn, err := grpc.NewClient(os.Getenv("product_url"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("failed to dial grpc client product: ", err)
	}

	return productproto.NewProductServiceClient(conn)
}

func Products(ctx context.Context, id string) (float32, error) {
	res, err := DialProductGrpc().GetByIdProduct(ctx, &productproto.ProductResponse{Id: id})
	if err != nil {
		return 0, fmt.Errorf("product not found: %v", err)
	}
	return res.Price, err
}
