package controllers

import (
	"context"
	"net/http"
	"time"

	config "github.com/Miguel-Florian/Electronic-bookshop-of-Higher-science-computer-school-of-Logbessou/Config"
	models "github.com/Miguel-Florian/Electronic-bookshop-of-Higher-science-computer-school-of-Logbessou/Models"
	"github.com/Miguel-Florian/Electronic-bookshop-of-Higher-science-computer-school-of-Logbessou/responses"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	//"go.mongodb.org/mongo-driver/bson/primitive"
)

var bookCollection *mongo.Collection = config.GetCollection(config.DB, "books")
var validated = validator.New()

func CreateBook() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var book models.Book
		defer cancel()
		// validatind request body
		if err := c.BindJSON(&book); err != nil {
			c.JSON(http.StatusBadRequest, responses.BookResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		//using the library validation to validate required fields
		if validationErr := validated.Struct(&book); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.BookResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}
		newBook := models.Book{
			ID:          primitive.NewObjectID(),
			Title:       book.Title,
			Description: book.Description,
			DateSortie:  book.DateSortie,
			Auteur:      book.Auteur,
			Categories:  book.Categories,
		}
		result, err := bookCollection.InsertOne(ctx, newBook)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.BookResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		c.JSON(http.StatusCreated, responses.BookResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})
	}
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
