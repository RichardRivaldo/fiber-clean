package repositories

import (
	"context"
	"fiber-clean/src/configs"
	"fiber-clean/src/models"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "users")

func RegisterUser(user models.User) (*mongo.InsertOneResult, error) {
	newUser := models.User{
		ID:        primitive.NewObjectID(),
		Email:     user.Email,
		Password:  user.Password,
		IsDeleted: false,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := userCollection.InsertOne(ctx, newUser)

	return result, err
}
