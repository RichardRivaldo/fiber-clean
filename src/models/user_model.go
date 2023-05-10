package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Email    string             `bson:"email,omitempty" validate:"email,required"`
	Password string             `bson:"password,omitempty" validate:"required"`
	Exists   bool               `bson:"exists,omitempty"`
}
