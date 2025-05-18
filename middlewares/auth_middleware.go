package middlewares

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)


func JWTAuthMiddleware() gin.HandlerFunc{
	return func(c *gin.Context){
		tokenString := c.GetHeader("Authorization")
		if tokenString == ""{
			c.AbortWithStatusJSON(http.StatusUnauthorized , gin.H{"error" : "Missing Authorization header"})
			return
		}
		tokenString = strings.TrimPrefix(tokenString,"Bearer ")
		token , err := jwt.Parse(tokenString,func(token *jwt.Token) (interface{},error){
			return []byte(os.Getenv("JWT_SECRET")),nil
		})

		if err != nil || !token.Valid {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
		}
		c.Next()
	}

}