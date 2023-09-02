package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Cid struct {
	Id		primitive.ObjectID `json:"_id,omitempty"`
	Code 	string			   `json:"code,omitempty" validate:"required"`
	Title	string			   `json:"title,omitempty" validate:"required"`
}	