package products

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	ID    primitive.ObjectID `bson:"_id"`
	Code  string             `bson:"code"`
	Price int                `bson:"price"`
}
