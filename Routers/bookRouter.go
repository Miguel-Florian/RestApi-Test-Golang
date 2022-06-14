package routers

import (
	controllers "github.com/Miguel-Florian/Electronic-bookshop-of-Higher-science-computer-school-of-Logbessou/Controllers"
	"github.com/gin-gonic/gin"
)

func BookRoute(router *gin.Engine) {
	//All routes related to book comes here

	api := router.Group("/api/book")
	{
		api.POST("/register", controllers.CreateBook())
		api.PUT("/:id", controllers.UpdateBook())
		api.GET("/:id", controllers.GetBookById())
		api.GET("/books", controllers.GetAllBooks())
		api.DELETE("/:id", controllers.DeleteBook())
	}

}
