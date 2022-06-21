package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Admin struct {
	ID       primitive.ObjectID `json:"_id,onitempty" bson:"_id,onitempty"`
	Username string             `json:"username,onitempty" bson:"username,onitempty"`
	Email    string             `json:"email,onitempty" bson:"email,onitempty,unique"`
	Password string             `json:"password" bson:"password,onitempty"`
	//HashPassword []byte        `json:"hashpassword,omitempty "`
}
type AdminLogin struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
