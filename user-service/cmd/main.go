package main

import (
	"log"
	"net"
	"os"
	"user-service/internal/storage"
	"user-service/userproto"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"

	"google.golang.org/grpc"
)

func main() {
	db, err := storage.OpenSql(os.Getenv("driver_name"), os.Getenv("postgres_url"))
	if err != nil {
		log.Fatal("failed to connect database:", err)
	}
	defer db.Close()

	handler := storage.Handler(db)

	server := grpc.NewServer()
	userproto.RegisterUserServiceServer(server, handler)
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
