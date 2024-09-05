package jwt

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func TokenMiddleware(expectedType string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
			c.Abort()
			return
		}

		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
			c.Abort()
			return
		}

		token := tokenParts[1]
		if expectedType == "user" && strings.HasPrefix(token, "user_") {
			c.Set("user_type", "user")
		} else if expectedType == "courier" && strings.HasPrefix(token, "courier_") {
			c.Set("user_type", "courier")
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		c.Next()
	}
}
