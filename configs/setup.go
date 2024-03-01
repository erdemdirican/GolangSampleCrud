package configs

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() *mongo.Client {

	ctx, _ := context.WithTimeout(context.Background(), 20*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(EnvMongoURI()))
	if err != nil {
		log.Fatalln((err))
	}
	err = client.Ping(ctx, nil)

	if err != nil {
		log.Fatalln((err))
	}

	return client
}

var DB *mongo.Client = ConnectDB()

func Collection(client *mongo.Client, collectionName string) *mongo.Collection {
	return client.Database("golangAPI").Collection(collectionName)
}
