package utils

import (
	"context"
	"log"
	"time"

	"team-maker-api/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func InitDB() {
	if config.AppConfig == nil {
		log.Fatal("Config not loaded. Call LoadConfig first")
	}

	clientOpts := options.Client().ApplyURI(config.AppConfig.MongoURI)
	client, err := mongo.NewClient(clientOpts)
	if err != nil {
		log.Fatalf("Mongo client error: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := client.Connect(ctx); err != nil {
		log.Fatalf("Mongo connect error: %v", err)
	}

	DB = client.Database(config.AppConfig.MongoDB)
	log.Println("MongoDB connected")
}
