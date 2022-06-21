package controllers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	config "github.com/Miguel-Florian/Electronic-bookshop-of-Higher-science-computer-school-of-Logbessou/Config"
	models "github.com/Miguel-Florian/Electronic-bookshop-of-Higher-science-computer-school-of-Logbessou/Models"
	"github.com/Miguel-Florian/Electronic-bookshop-of-Higher-science-computer-school-of-Logbessou/responses"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
)

var userCollection *mongo.Collection = config.GetCollection(config.DB, "users")
var validate = validator.New()

const SecretKey = "SECRETKEY"

func LoginUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data_receive models.UserLogin
		var u models.User
		if err := c.ShouldBindJSON(&data_receive); err != nil {
			return
		}
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		err := userCollection.FindOne(ctx, bson.M{"email": data_receive.Email}).Decode(&u)
		if err != nil {
			c.JSON(http.StatusNotFound, responses.UserResponse{Status: http.StatusNotFound, Message: "Unrecaheable Email !", Data: map[string]interface{}{"data": err.Error() + ",Email Introuvable"}})
			return
		}
		if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(data_receive.Password)); err != nil {
			c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "Incorrect password", Data: map[string]interface{}{"data": err.Error() + ",Code de sécurité incorrect"}})
			return
		}
		expTime := time.Now().Add(48 * time.Hour)
		claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
			ExpiresAt: expTime.Unix(),
			Id:        u.ID.String(),
			Issuer:    u.Email,
		})
		token, err := claims.SignedString([]byte(SecretKey))
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "Couldn't login", Data: map[string]interface{}{"data": "Impossible de vous connecter"}})
			return
		}
		cookie := http.Cookie{
			Name:     "Jwt-Token",
			Value:    token,
			Path:     "/api/user/login",
			Domain:   "Localhost",
			Expires:  expTime,
			MaxAge:   3600 * 48,
			Secure:   false,
			HttpOnly: true,
		}
		c.SetCookie(cookie.Name, cookie.Value, cookie.MaxAge, cookie.Path, cookie.Domain, cookie.Secure, cookie.HttpOnly)
		//c.Writer.Header().Add("Authorisation", cookie.Value)
		c.JSON(http.StatusAccepted, responses.UserResponse{Status: http.StatusAccepted, Message: "success", Data: map[string]interface{}{"data": u.Username}})

		c.Redirect(http.StatusFound, "/api/user/register")
	}
}

func RegisterUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var u models.User
		if err := c.ShouldBindJSON(&u); err != nil {
			return
		}
		if u.Password == "" {
			c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": "Invalid field, check it !"}})
			return
		}
		pass, _ := bcrypt.GenerateFromPassword([]byte(u.Password), 14)
		password := string(pass[:])
		user := models.User{
			ID:        primitive.NewObjectID(),
			Username:  u.Username,
			FirstName: u.FirstName,
			LastName:  u.LastName,
			Email:     u.Email,
			Password:  password,
		}
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if user.Username == "" || user.FirstName == "" || user.LastName == "" || user.Email == "" {
			c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": "Invalid field, check it !"}})
			return
		}
		result, err := userCollection.InsertOne(ctx, user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		c.JSON(http.StatusCreated, responses.UserResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})
		c.Redirect(http.StatusFound, "/api/book/books")
	}
}

func LogoutUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie := http.Cookie{
			Name:     "jwt-token",
			Value:    "",
			Path:     "/user/login",
			Domain:   "localhost",
			Expires:  time.Now().Add(-time.Hour),
			MaxAge:   0,
			Secure:   false,
			HttpOnly: true,
		}
		c.SetCookie(cookie.Name, cookie.Value, cookie.MaxAge, cookie.Path, cookie.Domain, cookie.Secure, cookie.HttpOnly)
		c.JSON(http.StatusOK, responses.UserResponse{Status: http.StatusOK, Message: "successfully disconnect", Data: map[string]interface{}{"data": "successfully disconnect"}})
		c.Redirect(http.StatusFound, "/api/user/login")
	}
}

func ValidateToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		const BEARER = "Bearer "
		authHeader := c.Request.Header.Get("Authorization")
		tokenString := authHeader[len(BEARER):]
		if tokenString == "" || authHeader == "" {
			c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "Token inexistant", Data: map[string]interface{}{"data": "Token inexistant ou invalide"}})
			return
		}
		token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(SecretKey), nil
		})
		claims := token.Claims.(jwt.MapClaims)
		var user models.User
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		err := userCollection.FindOne(ctx, bson.M{"email": claims["iss"]}).Decode(&user)
		if err != nil {
			c.JSON(http.StatusNotFound, responses.UserResponse{Status: http.StatusNotFound, Message: "Unexisting Email ", Data: map[string]interface{}{"data": "Email inexistant"}})
			return
		}
		c.JSON(http.StatusAccepted, responses.UserResponse{
			Status:  http.StatusAccepted,
			Message: "Token valid, User connected",
			Data:    map[string]interface{}{"data": user.FirstName + " " + user.LastName},
		})
		return
	}
}
