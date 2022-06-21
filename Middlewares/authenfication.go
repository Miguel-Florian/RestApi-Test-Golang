package middleware

import (
	"context"
	"net/http"
	"time"

	config "github.com/Miguel-Florian/Electronic-bookshop-of-Higher-science-computer-school-of-Logbessou/Config"
	controllers "github.com/Miguel-Florian/Electronic-bookshop-of-Higher-science-computer-school-of-Logbessou/Controllers"
	models "github.com/Miguel-Florian/Electronic-bookshop-of-Higher-science-computer-school-of-Logbessou/Models"
	"github.com/Miguel-Florian/Electronic-bookshop-of-Higher-science-computer-school-of-Logbessou/responses"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		controllers.ValidateToken()
		c.Next()
	}
}

var adminCollection *mongo.Collection = config.GetCollection(config.DB, "admin")
var validate = validator.New()

func AuthLogging() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data_receive models.AdminLogin
		var u models.Admin
		if err := c.ShouldBindJSON(&data_receive); err != nil {
			return
		}
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		err := adminCollection.FindOne(ctx, bson.M{"email": data_receive.Email}).Decode(&u)
		if err != nil {
			c.JSON(http.StatusNotFound, responses.UserResponse{Status: http.StatusNotFound, Message: "Unrecaheable Email !", Data: map[string]interface{}{"data": err.Error() + ",Email Introuvable"}})
			return
		}
		if err := bcrypt.CompareHashAndPassword([]byte("migflor04"), []byte(data_receive.Password)); err != nil {
			c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "Incorrect password", Data: map[string]interface{}{"data": err.Error() + ",Code de sécurité incorrect"}})
			return
		}
		if u.Username == "Miguel" {
			c.JSON(http.StatusAccepted, responses.UserResponse{Status: http.StatusAccepted, Message: "Credentials ok", Data: map[string]interface{}{"data": "Connected"}})
			return
		}
		c.Next()
	}
}

/*
import (
	common "github.com/Miguel-Florian/Electronic-bookshop-of-Higher-science-computer-school-of-Logbessou/Common"
	"github.com/gin-gonic/gin"
)
/*api.Use(func(ctx *gin.Context) {
		authHeader := ctx.Request.Header.Get("Authorization")
		if authHeader == ""{
			ctx.AbortWithStatusJSON(http.StatusUnauthorized,gin.H{"message":"No token"})
			return
		}
		if len(authHeader) != 2{
			ctx.AbortWithStatusJSON(http.StatusUnauthorized,gin.H{"message":"Malformed token"})
			return
		}else{
			jwtToken := authHeader[1]
			token, err := jwt.Parse(string(jwtToken), func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
				}
				return []byte(SecretKey), nil
			})
			if claims, ok := token.Claims.(jwt.StandardClaims); ok && token.Valid {
				ctx := ctx.Request.Context().Value(claims.Issuer).(jwt.Claims)
				// Access context values in handlers like this
				// props, _ := r.Context().Value("props").(jwt.MapClaims)
				ctx.ServeHTTP(w, r.WithContext(ctx))
			} else {
				ctx.AbortWithStatusJSON(http.StatusUnauthorized,gin.H{"message":"Unauthorized"})
				return
			}

		}
	})*/

/*func ValidateToken(t string){

}
*/
/*func middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := strings.Split(r.Header.Get("Authorization"), "Bearer ")
		if len(authHeader) != 2 {
			fmt.Println("Malformed token")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Malformed Token"))
		} else {
			jwtToken := authHeader[1]
			token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
				}
				return []byte(SecretKey), nil
			})

			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				ctx := context.WithValue(r.Context(), "props", claims)
				// Access context values in handlers like this
				// props, _ := r.Context().Value("props").(jwt.MapClaims)
				next.ServeHTTP(w, r.WithContext(ctx))
			} else {
				fmt.Println(err)
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("Unauthorized"))
			}
		}
	})
}*/

/*func AuthorizeJWT() gin.HandlerFunc{
	return func (c *gin.Context{
		const BEARER SCHEMA  = "Bearer "
		authHeader := c.Request.Header.Get("Authorization")
		tokenString := authHeader[len(BEARER SCHEMA) :]

		token,err :=
	})
}*/
