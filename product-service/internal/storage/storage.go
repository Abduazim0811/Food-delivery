package storage

import (
	"context"
	"log"
	"os"
	"product-service/internal/infrastructura/mongodb"
	"product-service/internal/service"
	productservice "product-service/product_service"
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

	hotelCollection := client.Database("Product").Collection("product")

	return client, hotelCollection, nil
}

func Handler() *productservice.Service{
	client, collection, err := NewMongodb()
	if err != nil {
		log.Println("connection mongodb error:", err)
		return nil
	}

	repo := mongodb.NewProductMongodb(client, collection)
	service := service.NewProductService(repo)
	handler := productservice.NewService(service)

	return handler
}
