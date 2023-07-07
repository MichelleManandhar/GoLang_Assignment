package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id primitive.ObjectID `json:"id,omitempty"`
	//make it ignore empty fields and make the field required
	Name       string  `json:"name,omitempty" validate:"required"`
	AccountNum int     `json:"accountNum,omitempty" validate:"required"`
	Balance    float64 `json:"balance,omitempty" validate:"required"`
}
