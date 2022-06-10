package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        primitive.ObjectID `json:"_id,onitempty" bson:"_id,onitempty"`
	Username  string             `json:"username,onitempty" bson:"username,onitempty"`
	FirstName string             `json:"firstname,onitempty" bson:"firstname,onitempty"`
	LastName  string             `json:"lastname,onitempty" bson:"lastname,onitempty"`
	Email     string             `json:"email,onitempty" bson:"email,onitempty,unique"`
	Password  string             `json:"password,onitempty" bson:"password,onitempty"`
	//HashPassword []byte        `json:"hashpassword,omitempty "`
}

func (user *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}
func (user *User) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword))
	if err != nil {
		return err
	}
	return nil
}
