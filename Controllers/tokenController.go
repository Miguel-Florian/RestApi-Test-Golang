package controllers

import (
	"net/http"

	common "github.com/Miguel-Florian/Electronic-bookshop-of-Higher-science-computer-school-of-Logbessou/Common"
	models "github.com/Miguel-Florian/Electronic-bookshop-of-Higher-science-computer-school-of-Logbessou/Models"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
	//"go.mongodb.org/mongo-driver/bson"
)

type TokenRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func GenerateToken(context *gin.Context) {
	var request TokenRequest
	var user models.User
	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	// check if email exists and password is correct
	//a := userCollection.FindOne(context, bson.M{"email": user.Email}).Decode(&user)
	err := userCollection.FindOne(context, bson.M{"email": request.Email}).Decode(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Error ! Unexisting email"})
		context.Abort()
		return
	}
	credentialError := user.CheckPassword(request.Password)
	if credentialError != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		context.Abort()
		return
	}
	tokenString, err := common.GenerateJWT(user.Email, user.Username)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}

	context.SetCookie("jwt-token", tokenString, 3600, "/user/login", "localhost", false, true)
	context.JSON(http.StatusOK, gin.H{"token": tokenString})
}
