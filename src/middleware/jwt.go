package middleware

import (
	"application/src/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.Request.Header.Get("Authorization")
		// Check if Authorization header exists and is formatted correctly
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "authorization header not provided or improperly formatted"})
			c.Abort()
			return
		}

		claims, err := usecase.ParseToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "error parsing token"})
			c.Abort()
			return
		}

		// Store userID and other details in the context
		c.Set("userID", claims.Subject)
		c.Set("username", claims.Username)

		c.Next()
	}
}
