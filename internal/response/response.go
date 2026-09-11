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

	if errors.Is(err, service.ErrTicketSoldOut) {
		c.JSON(409, gin.H{
			"status":  "error",
			"message": "Ticket sold out",
		})
		return
	}

	if errors.Is(err, service.ErrInvalidQuantity) {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "Invalid quantity",
		})
		return
	}

	if errors.Is(err, service.ErrTransactionNotFound) {
		c.JSON(404, gin.H{
			"status":  "error",
			"message": "Transaction not found",
		})
		return
	}

	if errors.Is(err, service.ErrTransactionAlreadyCancelled) {
		c.JSON(409, gin.H{
			"status":  "error",
			"message": "Transaction already cancelled",
		})
		return
	}

	if errors.Is(err, service.ErrGenreNameRequired) {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "Genre name is required",
		})
		return
	}

	if errors.Is(err, service.ErrGenreNameTooLong) {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "Genre name is too long",
		})
		return
	}

	if errors.Is(err, service.ErrGenreAlreadyExists) {
		c.JSON(409, gin.H{
			"status":  "error",
			"message": "Genre name already exists",
		})
		return
	}

	c.JSON(500, gin.H{
		"status":  "error",
		"message": "internal server error",
	})
}
