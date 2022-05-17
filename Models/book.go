package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Book struct {
	ID          primitive.ObjectID `json:"_id,onitempty" bson:"_id,onitempty"`
	Title       string             `json:"title,onitempty" bson:"title,onitempty"`
	Description string             `json:"description" bson:"title,descritpion"`
	DateSortie  string             `json:"date sortie,omitempty" bson:"date sortie,omitempty"`
	Auteur      []string           `json:"auteurs,onitempty" bson:"auteurs,onitempty"`
	Catégories  []string           `json:"categories,onitempty" bson:"categories,onitempty"`
}
