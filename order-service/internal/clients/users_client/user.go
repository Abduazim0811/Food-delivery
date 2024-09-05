package usersclient

import (
	"context"
	"log"
	"order_service/protos/userproto"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
)

func DialgrpcUser() userproto.UserServiceClient {
	conn, err := grpc.NewClient(os.Getenv("users_url"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("failed to dial grpc client product: ", err)
	}
	return userproto.NewUserServiceClient(conn)
}

func GetUsers(ctx context.Context, id int32)error{
	_,err := DialgrpcUser().GetbyIdUser(ctx, &userproto.UserRes{Id: id})
	return err
}
