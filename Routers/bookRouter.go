package routers

import (
	controllers "github.com/Miguel-Florian/Electronic-bookshop-of-Higher-science-computer-school-of-Logbessou/Controllers"
	"github.com/gin-gonic/gin"
)

func BookRoute(router *gin.Engine) {
	//All routes related to book comes here
	router.POST("/book", controllers.CreateBook)
	router.PUT("/book/:id", controllers.UpdateBook)
	router.GET("/book/:id", controllers.GetBookById)
	router.GET("/books", controllers.GetAllBooks)
	router.DELETE("/book/:id", controllers.DeleteBook)
}