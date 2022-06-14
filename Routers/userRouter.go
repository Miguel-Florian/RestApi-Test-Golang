package routers

import (
	controllers "github.com/Miguel-Florian/Electronic-bookshop-of-Higher-science-computer-school-of-Logbessou/Controllers"
	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	//router.LoadHTMLGlob("Templates/*")
	//All routes related to users comes here
	api := router.Group("/api/user")
	{
		api.POST("/save", controllers.CreateUser())
		api.POST("/login", controllers.LoginUser())
		api.POST("/register", controllers.RegisterUser())
		api.POST("/logout", controllers.LogoutUser())
		api.GET("/:id", controllers.GetUserById())
		api.PUT("/update/:id", controllers.UpdateUserById())
		api.GET("/users", controllers.GetAllUsers())
		api.DELETE("/delete/:id", controllers.DeleteUser())
	}

}
