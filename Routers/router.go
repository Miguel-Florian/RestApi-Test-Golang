package routers

import (
	"net/http"

	controllers "github.com/Miguel-Florian/E-School/Controllers"
	"github.com/gin-gonic/gin"
)

func InitAllRoutes() {
	router := gin.Default()
	router.SetTrustedProxies(nil)
	router.LoadHTMLGlob("Templates/*")

	router.GET("/user/register", func(c *gin.Context) {
		c.HTML(http.StatusOK, "register.html", gin.H{
			"title": "Register",
		})
	})
	router.GET("/user/login", controllers.LoginUser)
	router.POST("/user/:id", controllers.GetUserById)
	router.POST("/user/", controllers.GetAllUsers)
	router.DELETE("/users/:id", controllers.DeleteUser)

	router.POST("/book", controllers.CreateBook)
	router.PUT("/book/:id", controllers.UpdateBook)
	router.GET("/book/:id", controllers.GetBookById)
	router.GET("/books", controllers.GetAllBooks)
	router.DELETE("/book/:id", controllers.DeleteBook)

	router.Run("localhost:8080")
}
