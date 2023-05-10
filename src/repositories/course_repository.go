package repositories

import (
	"context"
	"fiber-clean/src/configs"
	"fiber-clean/src/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var courseCollection *mongo.Collection = configs.GetCollection(configs.DB, "courses")

func CreateNewCourse(course models.Course) (*mongo.InsertOneResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := courseCollection.InsertOne(ctx, course)

	return result, err
}

func QueryCourses(
	filter interface{},
	projection interface{},
	sort interface{},
) ([]*bson.M, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var courses []*bson.M

	projectionOpts := options.Find().SetProjection(projection)
	sortOpts := options.Find().SetSort(sort)

	cursor, err := courseCollection.Find(ctx, filter, projectionOpts, sortOpts)
	if err != nil {
		return courses, err
	}

	if err = cursor.All(ctx, &courses); err != nil {
		return courses, err
	}
	cursor.Close(ctx)

	return courses, nil
}

func GetAllCourses() ([]*bson.M, error) {
	emptyFilterOpt := bson.M{"exists": true}
	projectAllOpt := bson.M{}
	defaultSortOpt := bson.M{}
	return QueryCourses(emptyFilterOpt, projectAllOpt, defaultSortOpt)
}

func UpdateCourse(courseId string, course models.Course) (*mongo.UpdateResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objId, err := primitive.ObjectIDFromHex(courseId)
	if err != nil {
		return &mongo.UpdateResult{UpsertedCount: 0}, err
	}

	pByte, _ := bson.Marshal(course)

	var update bson.M
	_ = bson.Unmarshal(pByte, &update)

	result, err := courseCollection.UpdateOne(ctx, bson.M{"_id": objId}, bson.M{"$set": update})
	return result, err
}

func DeleteCourse(courseId string) (*mongo.UpdateResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objId, err := primitive.ObjectIDFromHex(courseId)
	if err != nil {
		return &mongo.UpdateResult{UpsertedCount: 0}, err
	}

	update := bson.M{"exists": false}
	result, err := courseCollection.UpdateOne(ctx, bson.M{"_id": objId}, bson.M{"$set": update})
	return result, err
}
