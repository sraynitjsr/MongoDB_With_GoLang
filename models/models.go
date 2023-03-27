package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	Name   string             `bson:"name,omitempty"`
	Gender string             `bson:"gender,omitempty"`
	Age    int                `bson:"age,omitempty"`
}
