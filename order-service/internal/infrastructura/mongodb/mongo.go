package mongodb

import (
	"context"
	"fmt"
	"log"
	"order_service/internal/entity/order"
	"order_service/internal/infrastructura/repository"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type OrderMongodb struct {
	client     *mongo.Client
	collection *mongo.Collection
}

var ctx context.Context

func NewOrderMongodb(client *mongo.Client, collection *mongo.Collection) repository.OrderRepository {
	return &OrderMongodb{client: client, collection: collection}
}

func (o *OrderMongodb) AddOrder(req order.CreateOrderReq) (string, error) {
	res, err := o.collection.InsertOne(ctx, req)
	if err != nil {
		log.Println("Error inserting order:", err)
		return "", fmt.Errorf("error inserting order: %v", err)
	}
	insertedId := res.InsertedID.(primitive.ObjectID).Hex()
	return insertedId, nil
}

func (o *OrderMongodb) GetOrderById(req order.GetOrderReq) (*order.GetOrderRes, error) {
	var res order.GetOrderRes
	id, err := primitive.ObjectIDFromHex(req.OrderID)
	if err != nil {
		log.Println("Invalid ObjectID format:", err)
		return nil, fmt.Errorf("invalid objectid format: %v", err)
	}

	err = o.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&res)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Println("Order not found")
			return nil, fmt.Errorf("order not found: %v", err)
		}
		log.Println("Error finding order:", err)
		return nil, fmt.Errorf("error finding order: %v", err)
	}

	return &res, nil
}

func (o *OrderMongodb) UpdateOrder(req order.UpdateReq) error {
	id, err := primitive.ObjectIDFromHex(req.OrderID)
	if err != nil {
		log.Println("Invalid ObjectID format:", err)
		return fmt.Errorf("invalid objectid format: %v", err)
	}

	filter := bson.M{"_id": id}

	update := bson.M{
		"$set": bson.M{
			"product": req.Product,
			"address": req.Address,
			"totalamount" : req.TotalAmount,
		},
	}
	res, err := o.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Println("Error updating order:", err)
		return fmt.Errorf("failed to update order: %v", err)
	}

	if res.MatchedCount == 0 {
		log.Println("No order found with the given ID")
		return fmt.Errorf("no order found with the given ID")
	}

	return nil
}

func (o *OrderMongodb) DeleteOrder(req order.GetOrderReq) error {
	id, err := primitive.ObjectIDFromHex(req.OrderID)
	if err != nil {
		log.Println("Invalid ObjectID format:", err)
		return fmt.Errorf("invalid objectid format: %v", err)
	}

	res, err := o.collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		log.Println("Error deleting order:", err)
		return fmt.Errorf("failed to delete order: %v", err)
	}

	if res.DeletedCount == 0 {
		log.Println("No order found with the given ID")
		return fmt.Errorf("no order found with the given ID")
	}
	return nil
}
