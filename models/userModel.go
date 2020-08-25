// Package model provides db model for each instance
package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// User model
type User struct {
	ID         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	FirstName  string             `json:"firstname,omitempty" bson:"firstname,omitempty"`
	LastName   string             `json:"lastname,omitempty" bson:"lastname,omitempty"`
	Email      string             `json:"email,omitempty" bson:"email,omitempty"`
	Profession string             `json:"profession,omitempty" bson:"profession,omitempty"`
	DOB        string             `json:"dob,omitempty" bson:"dob,omitempty"`
}
