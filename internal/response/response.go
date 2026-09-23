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

	if errors.Is(err, service.ErrInvalidMovieData) {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "Invalid movie data",
		})
		return
	}

	if errors.Is(err, service.ErrInvalidMovieDuration) {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "Invalid movie duration",
		})
		return
	}

	if errors.Is(err, service.ErrInvalidMovieReleaseDate) {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "Release date must use YYYY-MM-DD format",
		})
		return
	}

	if errors.Is(err, service.ErrMovieGenresRequired) {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "At least one genre is required",
		})
		return
	}

	if errors.Is(err, service.ErrDuplicateGenreIDs) {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "Duplicate genre IDs are not allowed",
		})
		return
	}

	if errors.Is(err, service.ErrInvalidPosterURL) {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "Invalid poster URL",
		})
		return
	}

	if errors.Is(err, service.ErrInvalidTrailerURL) {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "Invalid trailer URL",
		})
		return
	}

	if errors.Is(err, service.ErrMovieGenreNotFound) {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "One or more genres do not exist",
		})
		return
	}

	if errors.Is(err, service.ErrMovieNotFound) {
		c.JSON(404, gin.H{
			"status":  "error",
			"message": "Movie not found",
		})
		return
	}

	if errors.Is(err, service.ErrShowTimeNotInFuture) {
		c.JSON(400, gin.H{
			"status":  "error",
			"message": "Showtime must be in the future",
		})
		return
	}

	c.JSON(500, gin.H{
		"status":  "error",
		"message": "internal server error",
	})
}
