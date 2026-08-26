package response

import (
	"errors"
	"online-ticketing/internal/service"

	"github.com/gin-gonic/gin"
)

func ResponseSuccess(c *gin.Context, httpStatus int, message string, data any) {
	c.JSON(httpStatus, gin.H{
		"status":  "success",
		"message": message,
		"data":    data,
	})
}

func ResponseError(c *gin.Context, err error) {
	if errors.Is(err, service.ErrInvalidCredentials) {
		c.JSON(401, gin.H{
			"status":  "error",
			"message": "Invalid email or password",
		})
		return
	}

	c.JSON(500, gin.H{
		"status":  "error",
		"message": "internal server error",
	})
}
