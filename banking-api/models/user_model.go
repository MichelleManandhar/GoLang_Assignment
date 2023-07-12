package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	//create new struct to represent a single task in the db.
	Id primitive.ObjectID `json:"id,omitempty"` //primitive pkg to set the type of the id of eah task
	//since MongoDB uses ObjectId for the _id field by defautlt.
	// default behavior of MongoDB is that the lowercase field name is used as the key for each exported field when it is being serialized
	//make it ignore empty fields and make the field required
	Name       string  `json:"name,omitempty" `
	AccountNum int     `json:"accountNum,omitempty" `
	Balance    float64 `json:"balance,omitempty" `
	Withdraw   float64 `json:"withdraw" `
	Deposit    float64 `json:"deposit" `
}
