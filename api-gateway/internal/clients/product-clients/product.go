package productclients

import (
	"api-gateway/internal/protos/productproto"
	"log"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
)

func DialGrpcProduct() productproto.ProductServiceClient {
	conn, err := grpc.NewClient(os.Getenv("product_url"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("dial grpc product error:", err)
	}

	return productproto.NewProductServiceClient(conn)
}
