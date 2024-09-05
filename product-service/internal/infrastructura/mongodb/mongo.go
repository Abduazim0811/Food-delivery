package mongodb

import (
	"context"
	"fmt"
	"log"
	"product-service/internal/entity/product"
	"product-service/internal/infrastructura/repository"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductMongodb struct {
	client     *mongo.Client
	collection *mongo.Collection
}

var ctx context.Context

func NewProductMongodb(client *mongo.Client, collection *mongo.Collection) repository.ProductRepository {
	return &ProductMongodb{client: client, collection: collection}
}

func (p *ProductMongodb) AddProduct(req product.CreateReq) error {
	_, err := p.collection.InsertOne(ctx, req)
	if err != nil {
		log.Println("error inserting product: ", err)
		return fmt.Errorf("error inserting product: %v", err)
	}
	return nil
}

func (p *ProductMongodb) GetByIdProduct(req product.ProductResponse) (*product.Product, error) {
	var res product.Product
	id, err := primitive.ObjectIDFromHex(req.ID)
	if err != nil {
		log.Println("objectid error")
		return nil, fmt.Errorf("invalid objectid format: %v", err)
	}
	err = p.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&res)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Println("product not found")
			return nil, fmt.Errorf("product not found: %v", err)
		}
		return nil, fmt.Errorf("error finding product: %v", err)
	}
	return &res, nil

}

func (p *ProductMongodb) GetAll() (*[]product.Product, error) {
	var resproduct []product.Product
	cursor, err := p.collection.Find(ctx, bson.M{})
	if err != nil {
		log.Println("failed to get all products")
		return nil, fmt.Errorf("failed to get all products: %v", err)
	}

	for cursor.Next(ctx) {
		var res product.Product
		if err := cursor.Decode(&res); err != nil {
			log.Println("Failed to decode product:", err)
			return nil, fmt.Errorf("failed to decode product: %v", err)
		}
		resproduct = append(resproduct, res)
	}
	if err := cursor.Err(); err != nil {
		log.Println("Cursor error:", err)
		return nil, fmt.Errorf("cursor error: %v", err)
	}
	return &resproduct, nil
}

func (p *ProductMongodb) Update(req product.Product) error {
	id, err := primitive.ObjectIDFromHex(req.ID)
	if err != nil {
		log.Println("Invalid ObjectID format:", err)
		return fmt.Errorf("invalid objectid format: %v", err)
	}

	filter := bson.M{"_id": id}

	update := bson.M{
		"$set" : bson.M{
			"name" : req.Name,
			"description" : req.Description,
			"price" : req.Price,
		},
	}
	res, err := p.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Println("Error updating product:", err)
		return fmt.Errorf("failed to update product: %v", err)
	}

	if res.MatchedCount == 0 {
		log.Println("No product found with the given ID")
		return fmt.Errorf("no product found with the given ID: %v", mongo.ErrNoDocuments)
	}

	return nil
}

func (p *ProductMongodb) Delete(req product.ProductResponse)error{
	id, err := primitive.ObjectIDFromHex(req.ID)
	if err != nil {
		log.Println("Invalid ObjectID format:", err)
		return fmt.Errorf("invalid objectid format: %v", err)
	}

	res, err := p.collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		log.Println("Error deleting product:", err)
		return fmt.Errorf("failed to delete product: %v", err)
	}

	if res.DeletedCount == 0 {
		log.Println("No product found with the given ID")
		return fmt.Errorf("no product found with the given ID: %v", mongo.ErrNoDocuments)
	}

	return nil
}
