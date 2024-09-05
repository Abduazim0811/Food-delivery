package storage

import (
	"context"
	deliveryservice "delivery_service/delivery_service"
	"delivery_service/internal/infrastructura/mongodb"
	"delivery_service/internal/service"
	"log"
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

	hotelCollection := client.Database("Delivery").Collection("delivery")

	return client, hotelCollection, nil
}

func Handler() *deliveryservice.Service{
	clien, collection, err := NewMongodb()
	if err != nil {
		log.Println("connection mongodb error:", err)
		return nil
	}

	repo := mongodb.NewDeliveryMongodb(clien, collection)
	service := service.NewDeliveryService(repo)
	handler := deliveryservice.NewService(service)
	
	return handler
}
