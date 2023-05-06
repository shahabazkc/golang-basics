package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserModel struct {
	ID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Username string `json:"username,omitempty"`
	Name string `json:"name,omitempty"`
	Password string `json:"password,omitempty"`
	IsActive bool `json:"isActive" bson:"isActive"`
}

