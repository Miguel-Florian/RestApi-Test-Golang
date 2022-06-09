package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Book struct {
	ID          primitive.ObjectID `json:"_id,onitempty" bson:"_id,onitempty"`
	Title       string             `json:"title,onitempty" bson:"title,onitempty"`
	Description string             `json:"description" bson:"description,descritpion"`
	DateSortie  string             `json:"date sortie,omitempty" bson:"date sortie,omitempty"`
	Auteur      []string           `json:"auteurs,onitempty" bson:"auteurs,onitempty"`
	Categories  []string           `json:"categories,onitempty" bson:"categories,onitempty"`
}
