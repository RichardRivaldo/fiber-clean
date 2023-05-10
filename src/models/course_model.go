package models

import (
	"mime/multipart"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Course struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Name     string             `bson:"name,omitempty" validate:"required"`
	Category string             `bson:"category,omitempty" validate:"required"`
	Price    int                `bson:"price,omitempty" validate:"required"`
	Details  string             `bson:"details,omitempty"`
	Image    multipart.File     `bson:"image,omitempty"`
	Exists   bool               `bson:"exists,omitempty"`
}
