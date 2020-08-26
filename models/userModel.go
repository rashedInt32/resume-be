// Package model provides db model for each instance
package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// User model
type User struct {
	ID         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name       string             `json:"name,omitempty" bson:"name,omitempty"`
	Username   string             `json:"username,omitempty" bson:"username,omitempty"`
	Email      string             `json:"email,omitempty" bson:"email,omitempty"`
	Profession string             `json:"profession,omitempty" bson:"profession,omitempty"`
	DOB        string             `json:"dob,omitempty" bson:"dob,omitempty"`
	Password   string             `json:"password,omitempty" bson:"password,omitempty"`
}

// Auth model
type Auth struct {
	Email    string `json:"email,omitempty" bson:"email,omitempty"`
	Password string `json:"password,omitempty" bson:"password,omitempty"`
}
