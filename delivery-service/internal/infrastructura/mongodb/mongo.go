package mongodb

import (
	"context"
	"delivery_service/internal/entity/delivery"
	"delivery_service/internal/infrastructura/repository"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type DeliveryMongodb struct {
	client     *mongo.Client
	collection *mongo.Collection
}

var ctx context.Context

func NewDeliveryMongodb(client *mongo.Client, collection *mongo.Collection) repository.DeliveryRepository {
	return &DeliveryMongodb{client: client, collection: collection}
}

func (d *DeliveryMongodb) AddDelivery(req delivery.CreateDeliveryReq) (string, error) {
	result, err := d.collection.InsertOne(ctx, req)
	if err != nil {
		log.Println("create delivery error:", err)
		return "", fmt.Errorf("create delivery error: %v", err)
	}

	deliveryId := result.InsertedID.(primitive.ObjectID).Hex() 
	return deliveryId, nil
}

func (d *DeliveryMongodb) GetDeliveryStatus(req delivery.GetDeliveryStatusReq) (*delivery.Delivery, error) {
	var deliveryres delivery.Delivery

	err := d.collection.FindOne(ctx, bson.M{"_id": req.DeliveryID}).Decode(&deliveryres)
	if err != nil {
		log.Println("error find one delivery:", err)
		return nil, fmt.Errorf("error find one delivery: %v", err)
	}

	return &deliveryres, nil
}

func (d *DeliveryMongodb) Update(req delivery.UpdateDeliveryStatusReq) error {
	_, err := d.collection.UpdateOne(ctx, bson.M{"_id": req.DeliveryID}, bson.M{"$set": bson.M{"status": req.Status}})
	if err != nil {
		log.Println("error update delivery: ", err)
		return fmt.Errorf("error update delivery: %v", err)
	}

	return nil
}
