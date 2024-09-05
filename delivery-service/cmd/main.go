package main

import (
	"delivery_service/internal/storage"
	"delivery_service/protos/deliveryproto"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
)

func main() {
	handler := storage.Handler()

	server := grpc.NewServer()
	deliveryproto.RegisterDeliveryServiceServer(server, handler)
	lis, err := net.Listen("tcp", os.Getenv("server_url"))
	if err != nil {
		log.Fatal("Unable to listen :", err)
	}
	defer lis.Close()

	log.Println("Server is listening on port ", os.Getenv("server_url"))
	if err = server.Serve(lis); err != nil {
		log.Fatal("Unable to serve :", err)
	}
}
