package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Product struct {
	ID    primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name  string             `bson:"name" json:"name"`
	Price float64            `bson:"price" json:"price"`
}

/*type Product struct {
	ID    primitive.ObjectID `json:"id"`
	Name  string             `json:"product_name"`
	Price float64            `json:"price"`
}*/
