package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", gin.H{
		"title": "Register",
	})
}
func LoginUser(c *gin.Context) {
	return
}
func GetUserById(c *gin.Context) {
	return
}
func GetAllUsers(c *gin.Context) {
	return
}
func DeleteUser(c *gin.Context) {
	return
}
