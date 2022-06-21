package main

import (
	"fmt"

	config "github.com/Miguel-Florian/Electronic-bookshop-of-Higher-science-computer-school-of-Logbessou/Config"
	routers "github.com/Miguel-Florian/Electronic-bookshop-of-Higher-science-computer-school-of-Logbessou/Routers"
	"github.com/gin-gonic/gin"
	//models "github.com/Miguel-Florian/Electronic-bookshop-of-Higher-science-computer-school-of-Logbessou/Models"
)

func main() {
	fmt.Println("Starting server ...")
	r := gin.Default()
	r.RedirectTrailingSlash = true
    r.RedirectFixedPath = true
	//run database
	config.ConnectDB()

	// Initialize all routes
	routers.UserRoute(r)
	routers.BookRoute(r)
	routers.AdminRoute(r)

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"data": "Hello from Gin-gonic & mongoDB",
		})
	})
	r.SetTrustedProxies(nil)

	r.Run("localhost:8080")
}
