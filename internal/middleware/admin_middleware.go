package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AdminMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if ctx.GetString("role") != "admin" {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error": "Forbidden: admin access required",
			})
			return
		}

		ctx.Next()
	}
}
