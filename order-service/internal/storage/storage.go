package storage

import (
	"context"
	"log"
	"order_service/internal/infrastructura/mongodb"
	"order_service/internal/service"
	orderservce "order_service/order_servce"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
)

func NewMongodb() (*mongo.Client, *mongo.Collection, error) {
	clientOptions := options.Client().ApplyURI(os.Getenv("mongo_url"))
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, nil, err
	}

	hotelCollection := client.Database("Order").Collection("orders")

	return client, hotelCollection, nil
}

func Handler() *orderservce.Service{
	client, collection, err := NewMongodb()
	if err != nil {
		log.Println("connection mongodb error:", err)
		return nil
	}

	repo := mongodb.NewOrderMongodb(client, collection)
	service := service.NewOrderService(repo)
	handler := orderservce.NewService(service)
	return handler
}
