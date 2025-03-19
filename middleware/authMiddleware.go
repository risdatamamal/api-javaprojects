package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// TODO: ERROR
func validateToken(token string) (bool, string) {
	// For the purpose of this example, any token "valid-token" is considered valid
	// and it returns "admin" role. In real applications, implement actual validation logic.
	if token == "valid-token" {
		return true, "Admin"
	}

	return false, ""
}

func AuthMiddleware(requiredRole string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")

		if token == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token required"})
			ctx.Abort()
			return
		}

		isValid, role := validateToken(token)
		if !isValid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization token"})
			ctx.Abort()
			return
		}

		if role != requiredRole {
			ctx.JSON(http.StatusForbidden, gin.H{"error": "You don't have permission to access this resource"})
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
