package models

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	StatusAvailable   = 1
	StatusUnAvailable = 0
)

type Employee struct {
	ID          primitive.ObjectID `bson:"_id"`
	Name        string             `bson:"name"`
	DateOfBirth string             `bson:"date_of_birth"`
	Status      int                `bson:"status"`
}
