package database

import (
	"project/configs"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"context"
	"log"
	"time"
)

var Client *mongo.Client
var DB *mongo.Database

func CreateClient() {
	Client, err := mongo.NewClient(options.Client().ApplyURI(configs.MONGO_HOST))
	if err != nil {
		log.Println("[FATAL] could not create client for database")
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)

	err = Client.Connect(ctx)

	defer cancel()

	if err != nil {
		log.Println("[FATAL] could not connect to database")
		panic(err)
	}

	DB = Client.Database(configs.CLIENT_DATABASE)

	return
}
