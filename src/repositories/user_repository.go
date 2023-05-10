package repositories

import (
	"context"
	"fiber-clean/src/configs"
	"fiber-clean/src/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "users")

func RegisterUser(user models.User) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := userCollection.InsertOne(ctx, user)
	return result, err
}

func DeleteUser(userId string) (*mongo.UpdateResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return &mongo.UpdateResult{UpsertedCount: 0}, err
	}

	update := bson.M{"exists": false}
	result, err := userCollection.UpdateOne(ctx, bson.M{"_id": objId}, bson.M{"$set": update})
	return result, err
}
