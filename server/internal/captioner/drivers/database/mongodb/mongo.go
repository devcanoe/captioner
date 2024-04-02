package mongodb

import (
	"context"
	"fmt"
	"time"

	"captioner.com.ng/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoStore struct {
	Client *mongo.Database
}

func NewMongoStore(cfg *config.Config) *MongoStore {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.Uri))
	if err != nil {
		fmt.Println("Error connecting to data base", err)
	}
	fmt.Println("Connected Successfully")
	return &MongoStore{
		Client: client.Database(cfg.Database),
	}
}
