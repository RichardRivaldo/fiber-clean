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

func getCollectionCount(coll string, filter interface{}) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	count, err := configs.GetCollection(configs.DB, coll).CountDocuments(ctx, filter)
	return count, err
}

func GetStatistics() (models.Statistics, error) {
	statistics := models.Statistics{}
	userCount, err := getCollectionCount("users", bson.M{"exists": true})
	if err != nil {
		return statistics, err
	}

	courseCount, err := getCollectionCount("courses", bson.M{"exists": true})
	if err != nil {
		return statistics, err
	}

	freeCourseFilter := bson.D{
		{
			Key:   "price",
			Value: bson.D{{Key: "$eq", Value: 0}}},
		{
			Key:   "exists",
			Value: true,
		},
	}
	freeCourseCount, err := getCollectionCount("courses", freeCourseFilter)
	if err != nil {
		return statistics, err
	}

	statistics.UserCount = userCount
	statistics.CourseCount = courseCount
	statistics.FreeCourseCount = freeCourseCount

	return statistics, nil
}
