package userclient

import (
	"api-gateway/internal/protos/userproto"
	"log"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
)


func DialGrpcUser()userproto.UserServiceClient{
	conn, err := grpc.NewClient(os.Getenv("user_url"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("dial grpc user error:", err)
	}

	return userproto.NewUserServiceClient(conn)
}