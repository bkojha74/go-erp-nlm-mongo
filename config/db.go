package config

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func ConnectDB() {
	/*url := url.URL{
		Scheme: "mongodb",
		Host:   "localhost:27017",
	}*/
	url := os.Getenv("MONGO_URI")

	client, err := mongo.NewClient(options.Client().ApplyURI(url))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	DB = client.Database("erp")
}
