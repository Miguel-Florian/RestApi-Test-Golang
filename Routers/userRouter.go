package routers

import (
	controllers "github.com/Miguel-Florian/Electronic-bookshop-of-Higher-science-computer-school-of-Logbessou/Controllers"
	middleware "github.com/Miguel-Florian/Electronic-bookshop-of-Higher-science-computer-school-of-Logbessou/Middlewares"
	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	//router.LoadHTMLGlob("Templates/*")
	//All routes related to users comes here
	api := router.Group("/api/user")
	{
		api.POST("/login", controllers.LoginUser())
		api.POST("/register", controllers.RegisterUser())
		api.POST("/logout", controllers.LogoutUser())
	}
	api.Use(middleware.Auth())

}
