package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Course struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Name      string             `bson:"name,omitempty" validate:"email,required"`
	Category  string             `bson:"category,omitempty" validate:"required"`
	Price     bool               `bson:"price,omitempty" validate:"required"`
	Details   string             `bson:"details,omitempty"`
	ImageLink string             `bson:"image_link,omitempty"`
	IsDeleted bool               `bson:"is_deleted,omitempty" validate:"required"`
}
