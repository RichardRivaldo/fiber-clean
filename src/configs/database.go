package configs

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateDBConnection() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI(GetEnv("MONGO_URI")))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to database!")
	return client
}

func GetCollection(client *mongo.Client, coll string) *mongo.Collection {
	collection := client.Database(GetEnv("MONGO_DB")).Collection(coll)
	return collection
}

var DB = CreateDBConnection()
var Validator = validator.New()
