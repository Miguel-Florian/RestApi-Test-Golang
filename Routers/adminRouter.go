package routers

import (
	controllers "github.com/Miguel-Florian/Electronic-bookshop-of-Higher-science-computer-school-of-Logbessou/Controllers"
	"github.com/gin-gonic/gin"
)

func AdminRoute(router *gin.Engine) {
	//router.LoadHTMLGlob("Templates/*")
	//All routes related to users comes here
	api := router.Group("/api/admin", gin.BasicAuth(gin.Accounts{
		"Miguel": "migflor04",
	}))

	api.POST("user/save", controllers.CreateUser())
	api.GET("user/:id", controllers.GetUserById())
	api.PUT("user/update/:id", controllers.UpdateUserById())
	api.GET("user/users", controllers.GetAllUsers())
	api.DELETE("user/delete/:id", controllers.DeleteUser())

}
