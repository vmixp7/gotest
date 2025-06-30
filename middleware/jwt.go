package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"strings"
)

var jwtSecret = []byte("secret")

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		if auth == "" || !strings.HasPrefix(auth, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"msg": "Invalid token"})
			c.Abort()
			return
		}
		tokenStr := strings.TrimPrefix(auth, "Bearer ")

		token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"msg": "Token expired or invalid"})
			c.Abort()
			return
		}
		c.Next()
	}
}
