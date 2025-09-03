package repository

import (
	"context"
	"go-api/model"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductRepository struct {
	collection *mongo.Collection
}

func NewProductRepository(client *mongo.Client) *ProductRepository {
	collection := client.Database("mydb").Collection("products")
	return &ProductRepository{collection: collection}
}

func (pr *ProductRepository) GetProducts() ([]model.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := pr.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var products []model.Product
	for cursor.Next(ctx) {
		var product model.Product
		if err := cursor.Decode(&product); err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	return products, nil
}

func (pr *ProductRepository) CreateProduct(product model.Product) (model.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := pr.collection.InsertOne(ctx, product)

	if err != nil {
		return model.Product{}, err
	}

	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		product.ID = oid
	}

	return product, nil

}

func (pr *ProductRepository) GetProductByID(id primitive.ObjectID) (model.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var product model.Product
	err := pr.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&product)
	if err != nil {
		return model.Product{}, err
	}

	return product, nil
}

func (pr *ProductRepository) UpdateProduct(id primitive.ObjectID, updatedData model.Product) (model.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	update := bson.M{
		"$set": bson.M{
			"name":  updatedData.Name,
			"price": updatedData.Price,
		},
	}

	_, err := pr.collection.UpdateOne(ctx, bson.M{"_id": id}, update)
	if err != nil {
		return model.Product{}, err
	}

	updatedData.ID = id
	return updatedData, nil
}

func (pr *ProductRepository) DeleteProduct(id primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := pr.collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}

/*
import (
	"database/sql"
	"fmt"
	"go-api/model"
)

type ProductRepository struct {
	connection *sql.DB
}

func NewProductRepository(connection *sql.DB) ProductRepository {

	return ProductRepository{
		connection: connection,
	}

}

func (pr *ProductRepository) GetProducts() ([]model.Product, error) {

	query := "SELECT id, product_name, price FROM product"
	rows, err := pr.connection.Query(query)
	if err != nil {
		fmt.Println("Error fetching products:", err)
		return []model.Product{}, err
	}

	var productList []model.Product
	var productObj model.Product

	for rows.Next() {
		err = rows.Scan(
			&productObj.ID,
			&productObj.Name,
			&productObj.Price)

		if err != nil {
			fmt.Println("Error scanning product:", err)
			return []model.Product{}, err
		}

		productList = append(productList, productObj)
	}

	rows.Close()

	return productList, nil

}*/
