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
		panic("Config not loaded. Call LoadConfig first")
	}

	clientOpts := options.Client().ApplyURI(config.AppConfig.MongoURI)
	client, err := mongo.NewClient(clientOpts)
	if err != nil {
		panic("Mongo client error: " + err.Error())
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := client.Connect(ctx); err != nil {
		panic("Mongo connect error: " + err.Error())
	}

	// Tambahkan ping ke database untuk cek koneksi
	if err := client.Ping(ctx, nil); err != nil {
		panic("Mongo ping error: " + err.Error())
	}

	DB = client.Database(config.AppConfig.MongoDB)
	log.Println("MongoDB connected")
}
