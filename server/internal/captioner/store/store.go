package store

import (
	"context"
	"fmt"
	"log"
	"time"

	"captioner.com.ng/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func createIndex(db *mongo.Client) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	index := []mongo.IndexModel{
		{Keys: bson.D{{Key: "email", Value: 1}}, Options: options.Index().SetUnique(true)},
		{Keys: bson.D{{Key: "username", Value: 1}}, Options: options.Index().SetUnique(true)},
	}

	_, err := db.Database("captioner").Collection("users").Indexes().CreateMany(ctx, index)

	if err != nil {
		log.Fatal(err)
	}

}

func ConnectDB() *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.EnvMongoURI()))
	createIndex(client)
	if err != nil {
		log.Fatal("Could not connect to Database", err)
	}
	defer cancel()
	fmt.Println("Connected to Database")
	return client
}
