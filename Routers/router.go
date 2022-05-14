package routers

import (
	controllers "github.com/Miguel-Florian/E-School/Controllers"
	"github.com/gin-gonic/gin"
)

func InitAllRoutes() {
	router := gin.Default()
	router.SetTrustedProxies(nil)
	router.LoadHTMLGlob("Templates/*")

	//Routes for User entity
	router.GET("/user/register", controllers.CreateUser)
	router.GET("/user/login", controllers.LoginUser)
	router.POST("/user/:id", controllers.GetUserById)
	router.POST("/user/", controllers.GetAllUsers)
	router.DELETE("/users/:id", controllers.DeleteUser)

	// Routes for Book entity
	router.POST("/book", controllers.CreateBook)
	router.PUT("/book/:id", controllers.UpdateBook)
	router.GET("/book/:id", controllers.GetBookById)
	router.GET("/books", controllers.GetAllBooks)
	router.DELETE("/book/:id", controllers.DeleteBook)

	router.Run("localhost:8080")
}
