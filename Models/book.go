package models

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Publier     time.Time `json:"publier,omitempty"`
	Auteur      []string  `json:"auteurs"`
	Catégories  []string  `json:"categories"`
}

echo "# E-School" >> README.md
git init
git add README.md
git commit -m "first commit"
git branch -M main
git remote add origin https://github.com/Miguel-Florian/E-School.git
git push -u origin main
