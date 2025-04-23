package middlewares

import (
	"net/http"
	"strings"

	"golang_gin/app/libraries"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(c *gin.Context) {
	// Example: check for API key header
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthenticated",
		})
		return
	}

	token := strings.TrimPrefix(authHeader, "Bearer ")
	if token == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthenticated",
		})
		return
	}

	paseto := libraries.NewPasetoToken()
	output, err := paseto.ParseToken(token)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthenticated",
		})
		return
	}

	c.Set("userID", output.Subject)

	// Proceed to next handler
	c.Next()
}
