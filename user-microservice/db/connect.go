package db

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoInstance struct {
	Client *mongo.Client
	Db     *mongo.Database
}

var mg MongoInstance

// Connect - function to connect to mongodb in docker or atlas
func Connect(mongoURI string, dbName string) (*MongoInstance, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	if err != nil {
		return &MongoInstance{}, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	db := client.Database(dbName)

	if err != nil {
		return &MongoInstance{}, err
	}

	mg = MongoInstance{
		Client: client,
		Db:     db,
	}
	fmt.Println("Connected to db")
	return &mg, nil
}
