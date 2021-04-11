package database

import (
	"log"

	"github.com/dacharat/wire-gin/utils"
	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var Set = wire.NewSet(NewDatabase)

func NewDatabase() *mongo.Database {
	ctx, cancel := utils.GetContext()
	defer cancel()

	opts := options.Client().ApplyURI("mongodb://localhost:27017")

	c, err := mongo.Connect(ctx, opts)
	if err != nil {
		log.Fatal(err)
	}

	if err := c.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatal(err)
	}

	return c.Database("test")
}
