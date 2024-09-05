package user

import (
	"context"
	"log"
	"os"
	"payment-service/protos/userproto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
)

func DialgrpcUser() userproto.UserServiceClient {
	conn, err := grpc.NewClient(os.Getenv("user_url"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("failed to dial grpc client user: ", err)
	}

	return userproto.NewUserServiceClient(conn)
}


func UserID(ctx context.Context, id int32)error{
	_, err := DialgrpcUser().GetbyIdUser(ctx, &userproto.UserRes{Id: id})
	return err
}