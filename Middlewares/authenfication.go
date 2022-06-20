package middleware

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

/*func Auth() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenString := context.GetHeader("Authorization")
		if tokenString == "" {
			context.JSON(401, gin.H{"error": "request does not contain an access token"})
			context.Abort()
			return
		}
		err := ValidateToken(tokenString)
		if err != nil {
			context.JSON(401, gin.H{"error": err.Error()})
			context.Abort()
			return
		}
		context.Next()
	}
}
func ValidateToken(t string){

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
