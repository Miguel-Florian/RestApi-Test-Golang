package controllers

import (
	"github.com/gin-gonic/gin"
)

/*type Book struct {
	gorm.Model
	Titre       string   `json:"titre"`
	Description string   `json:"description"`
	Publier     string   `json:"publier,omitempty"`
	Auteur      []string `json:"auteurs"`
	Catégories  []string `json:"categories"`
}

var DB *gorm.DB
var err error

func InitialBookMigration() {
	DB, err = gorm.Open(mysql.Open(config.DNS), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Impossible de se connecter à la base de donnée")
	}
	DB.AutoMigrate(&Book{})
}*/

func CreateBook(c *gin.Context) {
	/*w.Header().Set("Content-Type", "application/json")
	var book Book
	json.NewDecoder(r.Body).Decode(&book)
	DB.Create(&book)
	json.NewEncoder(w).Encode(book)*/
	return
}
func UpdateBook(c *gin.Context) {
	return

}
func GetAllBooks(c *gin.Context) {
	return

}
func GetBookById(c *gin.Context) {
	return

}
func DeleteBook(c *gin.Context) {
	return

}
