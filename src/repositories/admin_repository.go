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

var adminCollection *mongo.Collection = configs.GetCollection(configs.DB, "admins")

func RegisterAdmin(admin models.Admin) (*mongo.InsertOneResult, error) {
	newAdmin := models.Admin{
		ID:       primitive.NewObjectID(),
		Email:    admin.Email,
		Password: admin.Password,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := adminCollection.InsertOne(ctx, newAdmin)

	return result, err
}

func DeleteUser(userId string) (*mongo.DeleteResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return &mongo.DeleteResult{DeletedCount: 0}, err
	}

	result, err := userCollection.DeleteOne(ctx, bson.M{"_id": objId})
	return result, err
}
