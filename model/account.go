package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Account struct {
	ID         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Account_id int                `json:"account_id,omitempty" bson:"account_id,omitempty"`
	Limit      int                `json:"limit,omitempty" bson:"limit,omitempty"`
	Products   []string           `json:"products,omitempty" bson:"products,omitempty"`
}
