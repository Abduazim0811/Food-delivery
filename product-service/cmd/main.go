package main

import (
	"log"
	"net"
	"os"
	"product-service/internal/storage"
	"product-service/productproto"

	"google.golang.org/grpc"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
)

func main() {
	handler := storage.Handler()
	server := grpc.NewServer()
	productproto.RegisterProductServiceServer(server, handler)
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
