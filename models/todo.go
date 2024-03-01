package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Todo struct {
	Id      primitive.ObjectID
	Title   string
	Content string
}
