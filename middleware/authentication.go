package middleware

import (
	"net/http"

	"github.com/risdatamamal/api-javaprojects/helpers"

	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		verifyToken, err := helpers.VerifyToken(ctx)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"error":   "Unauthorized",
				"message": err.Error(),
			})
			return
		}

		ctx.Set("userData", verifyToken)
		ctx.Next()
	}
}
